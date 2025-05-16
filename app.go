package main

import (
	"bili-start-live-helper/model"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	cookiejar "github.com/juju/persistent-cookiejar"

	"golang.org/x/net/publicsuffix"
)

// App struct
type App struct {
	ctx context.Context

	jar    *cookiejar.Jar
	client *http.Client

	csrf       string
	session_id string
}

// NewApp creates a new App application struct
func NewApp() *App {
	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
		Filename:         COOKIES_SAVE_PATH,
	})
	client := &http.Client{
		Jar: jar,
	}
	csrf, session_id := "", ""
	if _, err := os.Stat(COOKIES_SAVE_PATH); err == nil {
		ck := jar.AllCookies()
		for _, cookie := range ck {
			if cookie.Name == "bili_jct" {
				csrf = cookie.Value
			} else if cookie.Name == "SESSDATA" {
				session_id = cookie.Value
			}
		}
	}
	return &App{
		jar:        jar,
		client:     client,
		csrf:       csrf,
		session_id: session_id,
	}
}

func (a *App) GetQRCode() (model.QRCode, error) {
	var qr model.QRCode
	req, _ := http.NewRequest(http.MethodGet, QRCODE_URL, nil)
	resp, err := a.client.Do(req)
	if err != nil {
		return qr, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &qr); err != nil {
		return qr, err
	}
	return qr, nil
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
