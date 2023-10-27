package main

import (
	"log"
	"time"
	webdriver "github.com/fedesog/webdriver"
	"os/signal"
	"os"
	"syscall"
)

func StartChromeDriver () {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	chromeDriver := webdriver.NewChromeDriver("/bin/chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}
	err = session.Url("https://visa.vfsglobal.com/blr/en/pol/login")
	if err != nil {
		log.Println(err)
	}
	if session !=nil {
		log.Println(session)
	}
	time.Sleep(10 * time.Second)
	go func() {
		<-c
		session.Delete()
		chromeDriver.Stop()
		os.Exit(1)
		}()
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

}