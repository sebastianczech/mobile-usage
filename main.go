package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func checkNetgear(gatewayIP string, gatewayPassword string) string {
	c := colly.NewCollector(
		colly.AllowedDomains(gatewayIP),
		// colly.Debugger(&debug.LogDebugger{}),
	)

	err := c.Post("http://"+gatewayIP+"/Forms/config", map[string]string{
		"session.password": gatewayPassword,
	})
	if err != nil {
		// log.Fatal(err)
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

func main() {
	argsWithoutProg := os.Args[1:]

	fmt.Println(checkNetgear(argsWithoutProg[0], argsWithoutProg[1]))
}
