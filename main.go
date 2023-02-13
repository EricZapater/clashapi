package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/EricZapater/clashapi/environment"
	"github.com/EricZapater/clashapi/model"
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

func saveToDb(env environment.Environment, player model.Runaway, year int, week int, data time.Time) error {
	sqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.DbHost, env.DbPort, env.DbUser, env.DbPass, env.DbName)
	db, err := sql.Open("postgres", sqlString)
	if err != nil {
		return fmt.Errorf("error trying to connect to db: %v", err)
	}
	defer db.Close()
	query := `INSERT INTO historicalData(year, week, day, player, battlesDone, comments)VALUES($1,$2,$3,$4,$5,$6)`
	_, err = db.Exec(query, year, week, data, player.Name, player.DecksUsedToday, "")
	if err != nil {
		return (err)
	}
	return nil
}

func main() {
	ip := getip2()
	fmt.Println(ip)
	for {
		env := environment.LoadEnvironment()
		zona, _ := time.Now().Zone()
		loc, _ := time.LoadLocation(zona)

		iniTime := time.Now().In(loc) //UTC().Add(time.Duration(offset) * time.Second)
		if int(iniTime.Weekday()) == 5 || int(iniTime.Weekday()) == 6 || int(iniTime.Weekday()) == 7 || int(iniTime.Weekday()) == 0 || int(iniTime.Weekday()) == 1 {
			fmt.Println(iniTime)
			if (iniTime.Hour() == env.HoraFinal && iniTime.Minute() == env.MinutFinal) || (iniTime.Hour() == env.HoraAvis && iniTime.Minute() == env.MinutAvis) {
				fmt.Printf("Send: %v\n", time.Now().In(loc))
				runaways := service.GetRunaways(env)
				err := service.SendRunaways(env, runaways)

				if err != nil {
					log.Printf("error sending runaways: %v\n", err)
				}
				players := runaways
				for _, player := range players {
					var iweek int
					_, iweek = iniTime.ISOWeek()
					if int(iniTime.Weekday()) == 1 {
						iweek = iweek - 1
					}
					err := saveToDb(env, player, iniTime.Year(), iweek, iniTime)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
		time.Sleep(1 * time.Minute)
	}

}
