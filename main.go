package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
)

func checkNetgear(gatewayIP string, gatewayPassword string) string {
	c := colly.NewCollector(
		colly.AllowedDomains(gatewayIP),
		colly.Debugger(&debug.LogDebugger{}),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("################ Visiting URL:", r.URL, r.Method, "################ Headers:", r.Headers, "################ Body: ", r.Body)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("################ Response:", r.StatusCode, "################ Headers:", r.Headers, "################ Body:", string(r.Body))
	})

	c.OnHTML("span[class]", func(e *colly.HTMLElement) {
		itemprop := e.Attr("class")
		if itemprop == "m_datausage_dataTransferred" {
			fmt.Println("################ Gateway mobile data usage:", e.Text)
		}
	})

	err := c.Post("http://"+gatewayIP+"/Forms/config", map[string]string{
		"session.password": gatewayPassword,
	})
	if err != nil {
		log.Fatal(err)
	}

	c.Visit("http://" + gatewayIP + "/index.html")

	return "################ Finished checking mobile usage data on Netgear: " + gatewayIP
}

func checkNju(njuLogin string, njuPassword string) string {
	c := colly.NewCollector(
		colly.AllowedDomains("www.njumobile.pl"),
		colly.Debugger(&debug.LogDebugger{}),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("################ Visiting URL:", r.URL, r.Method, "################ Headers:", r.Headers, "################ Body: ", r.Body)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("################ Response:", r.StatusCode, "################ Headers:", r.Headers, "################ Body:", string(r.Body))
	})

	c.OnHTML("div[class]", func(e *colly.HTMLElement) {
		itemprop := e.Attr("class")
		if itemprop == "box-slider-info" {
			fmt.Println("################ Nju mobile data usage:", e.Text)
		}
	})

	err := c.Post("https://www.njumobile.pl/logowanie?_DARGS=/profile-processes/login/login.jsp.portal-login-form", map[string]string{
		"login-form":    njuLogin,
		"password-form": njuPassword,
	})
	if err != nil {
		log.Fatal(err)
	}

	siteCookies := c.Cookies("https://www.njumobile.pl/logowanie?_DARGS=/profile-processes/login/login.jsp.portal-login-form")
	strCookies := ""
	fmt.Println("################ Cookies:", siteCookies)
	for _, element := range siteCookies {
		fmt.Println("################ Set-Cookie:", element)
		strCookies += element.Name + "=" + element.Value + ";"
	}
	fmt.Println("################ Str-Cookies:", strCookies)

	err = c.Request("POST",
		// "https://www.njumobile.pl/logowanie",
		// "https://www.njumobile.pl/logowanie?_DARGS=/profile-processes/login/login.jsp.portal-login-form",
		"https://www.njumobile.pl/logowanie?_DARGS=/profile-processes/login/login.jsp.portal-login-form",
		strings.NewReader("login-form="+njuLogin+"&password-form="+njuPassword+"&login-submit=zaloguj się"),
		nil,
		http.Header{
			"Content-Type": []string{"application/x-www-form-urlencoded"},
			"Origin":       []string{"https://www.njumobile.pl"},
			"Cookie":       []string{strCookies},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	c.Visit("https://www.njumobile.pl/mojekonto/stan-konta")

	return "################ Finished checking mobile usage data on Nju"
}

func main() {
	argsWithoutProg := os.Args[1:]

	fmt.Println(checkNetgear(argsWithoutProg[0], argsWithoutProg[1]))
	fmt.Println(checkNju(argsWithoutProg[2], argsWithoutProg[3]))
}
