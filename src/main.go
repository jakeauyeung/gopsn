/*
* Playstation Network API
* v0.0.1
* @author   jakeauyeung
* @desc information from Sony PSN servers
 */
package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const (
	baseURL          = "https://auth.api.sonyentertainmentnetwork.com"
	redirectURLOauth = "com.scee.psxandroid.scecompcall://redirect"
	clientID         = "b0d0d7ad-bb99-4ab1-b25e-afa0c76577b0"
	scope            = "sceapp"
	scopePSN         = "psn:sceapp,user:account.get,user:account.settings.privacy.get,user:account.settings.privacy.update,user:account.realName.get,user:account.realName.update,kamaji:get_account_hash"
	csrfToken
	authCode
	clientSecret  = "Zo4y8eGIa3oazIEp"
	duid          = "00000005006401283335353338373035333434333134313a433635303220202020202020202020202020202020"
	state         = "1156936032"
	serviceEntity = "urn:service-entity:psn"
	paramString   = "c2VydmljZV9lbnRpdHk9cHNuJnJlcXVlc3RfdGhlbWU9bGlxdWlk"
)

var (
	debug         = false
	npLanguage    = "en"
	region        = "us"
	email         string
	password      string
	userAgent     = "Mozilla/5.0 (Linux; U; Android 4.3; " + npLanguage + "; C6502 Build/10.4.1.B.0.101) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 PlayStation App/2.55.8/" + npLanguage + "/" + npLanguage
	requestedWith = "com.scee.psxandroid"
	signIN        = baseURL + "/2.0/oauth/authorize?response_type=code&service_entity=" + serviceEntity + "&returnAuthCode=true&state=" + state + "&redirect_uri=" + redirectURLOauth + "&client_id=" + clientID + "&scope=" + scopePSN
)

func main() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", signIN, nil)
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("X-Requested-With", requestedWith)

	response, _ := client.Do(request)

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodyStr := string(body)
		r := strings.NewReader(bodyStr)
		d, err := html.Parse(r)
		doc, err := goquery.NewDocument(d)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("#brandingParams")
		fmt.Println(doc)
	}
}
