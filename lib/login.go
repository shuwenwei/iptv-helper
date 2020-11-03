package lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"iptv-helper/util"
	"net/http"
	"net/url"
	"strings"
)

func beforeLogin(client *http.Client) {
	getBody, err := getPage(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer getBody.Close()

	page, err := goquery.NewDocumentFromReader(getBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	lt = util.GetLt(page)
	util.PwdEncoderInstance.PublicKey = util.GetPublicKey(page)
}

func getPage(client *http.Client) (io.ReadCloser, error) {
	getRequest, _ := http.NewRequest("GET", util.LOGIN_URL, nil)
	util.SetRequestHeader(getRequest)
	resp, err := client.Do(getRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println("Header:", resp.Header)
	return resp.Body, nil
}

func sendLoginRequest(client *http.Client, username, password string) string {
	password, err := util.PwdEncoderInstance.EncodePassword([]byte(password))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	values := url.Values{
		"username":   {username},
		"password":   {password},
		"errors": {"0"},
		"imageCodeName": {""},
		"cryptoType": {"1"},
		"lt":         {lt},
		"_eventId":   {"submit"},
	}

	postReq, _ := http.NewRequest("POST", util.LOGIN_URL, strings.NewReader(values.Encode()))
	reqBody, _ := ioutil.ReadAll(postReq.Body)
	fmt.Println("reqBody:", string(reqBody))
	util.SetRequestHeader(postReq)

	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	return resp.Cookies()[0].Value
}