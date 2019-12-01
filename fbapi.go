package fbapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	API_VERSION   = "5.0"
	API_TOKEN_URL = "https://graph.facebook.com/v5.0/oauth/access_token"
	API_ENDPOINT  = "https://graph.facebook.com/"
)

var (
	httpTr *http.Transport
)

func init() {
	httpTr = &http.Transport{
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     10 * time.Minute,
	}
}

// Получение токена
func (fb *Api) GetToken(d TokenData) (ans GetTokenAns, err error) {

	url := fmt.Sprintf("%s?code=%s&client_id=%d&client_secret=%s&redirect_uri=%s&v=%s",
		API_TOKEN_URL, d.Code, d.ClientId, d.ClientSecret, d.RedirectUri, API_VERSION)

	client := &http.Client{}
	resp, err := client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println("[error]", err, string(content))
		return
	}

	err = json.Unmarshal(content, &ans)
	if err != nil {
		log.Println("[error]", err, string(content))
		return
	}

	return
}

// Получаем инфу об аккаунте
func (fb *Api) GetAccount(id string, fields string) (ans GetAccountAns, err error) {
	if fields != "" {
		fields = "&fields=" + fields
	}

	url := fmt.Sprintf("%s%s/?access_token=%s%s",
		API_ENDPOINT, id, fb.AccessToken, fields)

	client := &http.Client{Transport: httpTr}
	resp, err := client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println("[error]", err, string(content))
		return
	}

	err = json.Unmarshal(content, &ans)
	if err != nil {
		log.Println("[error]", err, string(content))
		return
	}

	return
}

// Получаем инфу о рекламах
func (fb *Api) Ads(id string, params map[string]string) (ans AdsAns, err error) {

	urlParams := []string{}
	for k, v := range params {
		urlParams = append(urlParams, fmt.Sprintf("%s=%s", k, v))
	}

	url := fmt.Sprintf("%sv%s/%s/ads?access_token=%s&%s", API_ENDPOINT, API_VERSION, id,
		fb.AccessToken, strings.Join(urlParams, "&"))

	client := &http.Client{Transport: httpTr}
	resp, err := client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println("[error]", err, string(content))
		return
	}

	err = json.Unmarshal(content, &ans)
	if err != nil {
		log.Println("[error]", err, string(content))
		return
	}

	return

}

// Получаем инфу о рекламах
func (fb *Api) Insights(id string, params map[string]string) (ans InsightsAns, err error) {

	urlParams := []string{}
	for k, v := range params {
		urlParams = append(urlParams, fmt.Sprintf("%s=%s", k, v))
	}

	url := fmt.Sprintf("%sv%s/%s/insights?access_token=%s&%s", API_ENDPOINT, API_VERSION, id,
		fb.AccessToken, strings.Join(urlParams, "&"))

	client := &http.Client{Transport: httpTr}
	resp, err := client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println("[error]", err, string(content))
		return
	}

	err = json.Unmarshal(content, &ans)
	if err != nil {
		log.Println("[error]", err, string(content))
		return
	}

	return

}
