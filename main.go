package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/EricZapater/clashapi/environment"
	"github.com/EricZapater/clashapi/service"
)

type IP struct {
	Query string
}

func getip2() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func main() {
	ip := getip2()
	fmt.Println(ip)
	for {
		env := environment.LoadEnvironment()
		zona, _ := time.Now().Zone()
		loc, _ := time.LoadLocation(zona)
		fmt.Println(zona)

		iniTime := time.Now().In(loc) //UTC().Add(time.Duration(offset) * time.Second)
		fmt.Println(iniTime)
		if (iniTime.Hour() == env.HoraFinal && iniTime.Minute() == env.MinutFinal) || (iniTime.Hour() == env.HoraAvis && iniTime.Minute() == env.MinutAvis) {
			fmt.Printf("Send: %v\n", time.Now().In(loc))

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
