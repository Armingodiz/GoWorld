package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	//	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

var api string
var text string
var ngrokApi string

func main() {
	// TODO cleaning code
	// TODO pacakging app
	setInfo()
	makeTunelForServer()
	fmt.Println("connecting to telegram api ... ")
	tunelURL := getTunelURL()
	fmt.Println("boooooooooooooooooo")
	connectTunelToTelegramAPI(tunelURL)
	startSever()
}

func setInfo() {
	fmt.Println("pleas enter your api : ")
	fmt.Scanln(&api)
	fmt.Println("ENTER THE text you want to show to user after sending /start :")
	fmt.Scanln(&text)
	var numRun string
	fmt.Println("it is your first run ? ")
	fmt.Scan(&numRun)
	if numRun == "yes" {
		fmt.Println("ENTER YOUR TOKEN : ")
		var token string
		fmt.Scan(&token)
		setNgrokApi(token)
	}
}

func connectTunelToTelegramAPI(url string) {
	app2 := "curl"
	arg1 := "-F"
	arg2 := "url=" + url + "/"
	arg3 := "https://api.telegram.org/bot" + api + "/setWebhook"
	cmd := exec.Command(app2, arg1, arg2, arg3)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tss := string(stdout)
	fmt.Println(tss)
}

func getTunelURL() string {
	time.Sleep(5 * time.Second)
	app2 := "curl"
	arg1 := "localhost:4040/api/tunnels"

	cmd := exec.Command(app2, arg1)
	stdout, err := cmd.Output()
	content := NgrokDetail{}
	json.Unmarshal(stdout, &content)
	if err != nil {
		fmt.Println("fucl ofdsadf")
		fmt.Println(content.Tunnels[0].PublicURL, "https")
		return ""
	}
	if strings.Contains(content.Tunnels[0].PublicURL, "https") {
		fmt.Println(content.Tunnels[0].PublicURL, "https")
		return content.Tunnels[0].PublicURL
	}
	fmt.Println(content.Tunnels[1].PublicURL)
	return content.Tunnels[1].PublicURL
}

type NgrokDetail struct {
	Tunnels []struct {
		Name      string `json:"name"`
		URI       string `json:"uri"`
		PublicURL string `json:"public_url"`
		Proto     string `json:"proto"`
		Config    struct {
			Addr    string `json:"addr"`
			Inspect bool   `json:"inspect"`
		} `json:"config"`
		Metrics struct {
			Conns struct {
				Count  int `json:"count"`
				Gauge  int `json:"gauge"`
				Rate1  int `json:"rate1"`
				Rate5  int `json:"rate5"`
				Rate15 int `json:"rate15"`
				P50    int `json:"p50"`
				P90    int `json:"p90"`
				P95    int `json:"p95"`
				P99    int `json:"p99"`
			} `json:"conns"`
			HTTP struct {
				Count  int `json:"count"`
				Rate1  int `json:"rate1"`
				Rate5  int `json:"rate5"`
				Rate15 int `json:"rate15"`
				P50    int `json:"p50"`
				P90    int `json:"p90"`
				P95    int `json:"p95"`
				P99    int `json:"p99"`
			} `json:"http"`
		} `json:"metrics"`
	} `json:"tunnels"`
	URI string `json:"uri"`
}

func setNgrokApi(token string) {
	app2 := "ngrok"
	arg1 := "authtoken"

	cmd := exec.Command(app2, arg1, token)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tss := string(stdout)
	fmt.Println(tss)
}

func startSever() {
	fmt.Println("starting server ...")
	http.ListenAndServe(":8585", http.HandlerFunc(Handler))
}

func makeTunelForServer() {
	app2 := "ngrok"
	arg1 := "http"
	arg2 := "8585"
	cmd := exec.Command(app2, arg1, arg2)
	cmd.Start()
}

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// Check if the message contains the word "marco"
	// if not, return without doing anything
	if !strings.Contains(strings.ToLower(body.Message.Text), "/start") {
		return
	}

	// If the text contains marco, call the `sayPolo` function, which
	// is defined below
	if err := showText(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func showText(chatID int64) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   text,
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	url := "https://api.telegram.org/bot" + api + "/sendMessage"
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
