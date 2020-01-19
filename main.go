package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func checkNetgear(gatewayIP string, gatewayPassword string) string {
	c := colly.NewCollector(
		colly.AllowedDomains(gatewayIP),
		colly.Debugger(&debug.LogDebugger{}),
	)

	err := c.Post("http://"+gatewayIP+"/Forms/config", map[string]string{
		"session.password": gatewayPassword,
	})
	if err != nil {
		log.Fatal(err)
	}

	c.OnHTML("span[class]", func(e *colly.HTMLElement) {
		itemprop := e.Attr("class")
		if itemprop == "m_datausage_dataTransferred" {
			fmt.Println("Gateway mobile data usage:", e.Text)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL)
	})

	c.Visit("http://" + gatewayIP + "/index.html")

	return "Finished checking mobile usage data on Netgear: " + gatewayIP
}

func checkNju(njuLogin string, njuPassword string) string {
	c := colly.NewCollector(
		colly.AllowedDomains("www.njumobile.pl"),
		colly.Debugger(&debug.LogDebugger{}),
	)

	// err := c.Post("https://www.njumobile.pl/logowanie", map[string]string{
	// 	"login-form":    njuLogin,
	// 	"password-form": njuPassword,
	// })
	err := c.Request("POST",
		"https://www.njumobile.pl/logowanie",
		strings.NewReader("login-form="+njuLogin+"&password-form"+njuPassword),
		nil,
		http.Header{
			"Content-Type": []string{"application/x-www-form-urlencoded"},
			"Origin":       []string{"https://www.njumobile.pl"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	c.OnHTML("div[class]", func(e *colly.HTMLElement) {
		itemprop := e.Attr("class")
		if itemprop == "box-slider-info" {
			fmt.Println("Nju mobile data usage:", e.Text)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL)
	})

	c.Visit("https://www.njumobile.pl/mojekonto/stan-konta")

	return "Finished checking mobile usage data on Nju"
}

func main() {
	argsWithoutProg := os.Args[1:]

	// fmt.Println(checkNetgear(argsWithoutProg[0], argsWithoutProg[1]))
	fmt.Println(checkNju(argsWithoutProg[2], argsWithoutProg[3]))
}
