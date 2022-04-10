package cron

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

var WeatherLocations = []string{"tuzla", "sarajevo", "banja-luka", "zenica"}

func ScrapingWeather() {

	for _, city := range WeatherLocations {

		file, err := os.Create("weather-data/" + city + ".csv")
		if err != nil {
			log.Fatalf("ERROR: Could not create file %q: %s\n", city+".csv", err)
			return
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write([]string{"Date", "Temp", "Desc", "Feel"})

		c := colly.NewCollector()

		c.OnHTML(`#wt-ext tbody tr`, func(e *colly.HTMLElement) {
			date := e.ChildText("th .soft")
			temp := e.ChildText("td:nth-child(3)")
			desc := e.ChildText("td:nth-child(4)")
			feel := e.ChildText("td:nth-child(5)")

			writer.Write([]string{
				date,
				temp,
				desc,
				feel,
			})
		})

		c.Visit("https://www.timeanddate.com/weather/bosnia-herzegovina/" + city + "/ext")
		fmt.Println("End of scraping task for " + city + "!")

	}

}
