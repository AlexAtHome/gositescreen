package main

import (
	"github.com/go-auxiliaries/selenium"
	"io/fs"
	"log"
	"net/url"
	"os"
)

const geckoPath = "geckodriver"

func getDriver() (webdriver selenium.WebDriver) {
	var err error
	caps := selenium.Capabilities(map[string]interface{}{
		"browserName": "firefox",
	})
	if webdriver, err = selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub"); err != nil {
		log.Fatalf("Failed to open session: %s\n", err)
	}
	return webdriver
}

func main() {
	uri, err := url.ParseRequestURI(os.Args[1])
	if err != nil {
		panic(err)
	}
	log.Printf("Testing the website: %v\n", uri.String())

	driver := getDriver()
	defer driver.Quit()

	handleError := func(msg string, err error) {
		driver.Quit()
		log.Fatalf(msg, err)
	}

	err = driver.Get(uri.String())
	if err != nil {
		handleError("Failed to load page: %s\n", err)
	}

	img, err := driver.Screenshot()
	if err != nil {
		handleError("Failed to get a screenshot: %s\n", err)
	}
	if err = os.WriteFile("./screenshot.png", img, fs.ModePerm); err != nil {
		handleError("Failed to save a screenshot: %s\n", err)
	}
	log.Printf("The screenshot has been saved into a file screenshot.png\n")
}

// vim: tabstop=4 shiftwidth=4 noexpandtab
