package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/EricZapater/clashapi/environment"
	"github.com/EricZapater/clashapi/model"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unkown fromServer")
		}
	}
	return nil, nil
}

func GetPlayers(env environment.Environment) []model.Runaway {
	url := env.Endpoint
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + env.Bearer
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating new request.\n", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	var clan model.Resp
	json.NewDecoder(resp.Body).Decode(&clan)
	var runaways []model.Runaway
	var runaway model.Runaway
	for _, v := range clan.Clan.Participants {
		isFromClan, err := IsFromClan(v.Tag, env)
		if err != nil {
			log.Println(err)
			return nil
		}
		if isFromClan {
			runaway.Tag = v.Tag
			runaway.Name = v.Name
			runaway.DecksUsed = v.DecksUsed
			runaway.DecksUsedToday = v.DecksUsedToday
			runaways = append(runaways, runaway)
		}
	}
	return runaways
}

func GetRunaways(env environment.Environment) []model.Runaway {
	url := env.Endpoint
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + env.Bearer

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating new request.\n", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	var clan model.Resp
	json.NewDecoder(resp.Body).Decode(&clan)
	var runaways []model.Runaway
	var runaway model.Runaway
	for _, v := range clan.Clan.Participants {
		isFromClan, err := IsFromClan(v.Tag, env)
		if err != nil {
			log.Println(err)
			return nil
		}
		if v.DecksUsedToday < 4 && isFromClan {
			runaway.Tag = v.Tag
			runaway.Name = v.Name
			runaway.DecksUsed = v.DecksUsed
			runaway.DecksUsedToday = v.DecksUsedToday
			runaways = append(runaways, runaway)
		}
	}
	return runaways
}

func send(env environment.Environment, stringMessage string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", env.T_Token)

	var message model.Message

	message.Chat_id = env.T_ChatID
	message.Text = stringMessage

	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("failed to close response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send successful request. Status was %q", response.Status)
	}
	return nil
}

func SendRunaways(env environment.Environment, runaways []model.Runaway) error {
	var s_runaways string
	var l_runaways string
	for _, v := range runaways {
		s_runaways = fmt.Sprintf("%s\nNom: %s\nBaralles totals: %v\nBaralles avui:%v\n", s_runaways, v.Name, v.DecksUsed, v.DecksUsedToday)
		l_runaways = fmt.Sprintf("%s,%s", l_runaways, v.Name)
	}
	err := send(env, s_runaways)
	if err != nil {
		return err
	}
	err = send(env, l_runaways)
	if err != nil {
		return err
	}

	return nil
}

func IsFromClan(usertag string, env environment.Environment) (bool, error) {
	url := fmt.Sprintf("https://api.clashroyale.com/v1/players/%s", strings.Replace(usertag, "#", "%23", 1))
	var bearer = "Bearer " + env.Bearer
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating new request.\n", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Cache-Control", "no-cache")
	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	var clan model.Resp
	json.NewDecoder(resp.Body).Decode(&clan)
	if clan.Clan.Name == "CATALUNYA" {
		return true, nil
	}
	return false, nil
}
