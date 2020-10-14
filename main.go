package main

import (
	"os"
	"strings"

	"github.com/byungsujeong/jobscrapper/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	/* e.Use(middleware.Logger())
	e.Use(middleware.Recover()) */

	// Routes
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func handleHome(c echo.Context) error {
	return c.File("home.html")
	//return c.String(http.StatusOK, "Hello, World!")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}
