package main

import (
	"fmt"
	"log"
	"time"

	"github.com/EricZapater/clashapi/environment"
	"github.com/EricZapater/clashapi/service"
)

func main() {
	for {
		zona, _ := time.Now().Zone()
		loc, _ := time.LoadLocation(zona)
		//fmt.Println(zona, offset)

		iniTime := time.Now().In(loc) //UTC().Add(time.Duration(offset) * time.Second)
		fmt.Println(iniTime)
		if (iniTime.Hour() == 10 && iniTime.Minute() == 35) || (iniTime.Hour() == 9 && iniTime.Minute() == 0) {
			fmt.Printf("Send: %v\n", time.Now().In(loc))
			env := environment.LoadEnvironment()
			runaways := service.GetRunaways(env)
			err := service.SendRunaways(env, runaways)
			if err != nil {
				log.Printf("error sending runaways: %v\n", err)
			}
		}
		//fmt.Println(time.Now())
		time.Sleep(1 * time.Minute)
	}

}
