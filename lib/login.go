package lib

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"iptv-helper/util"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var (
	lt = ""
	viewgood = ""
)

func Login(username, password string) *Iptv {
	if videoLoginPwd, exists := checkCache(); exists {
		return &Iptv{
			iptvUsername: username,
			iptvPassword: videoLoginPwd,
		}
	}

	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		req.Host = util.RedirectHost
		req.Header.Set("Host", util.RedirectHost)
		req.URL.Host = util.RedirectHost
		return nil
	}
	beforeLogin(&client)
	fmt.Println("publicKey:", util.PwdEncoderInstance.PublicKey)
	fmt.Println("lt:", lt)

	viewgood = sendLoginRequest(&client, username, password)
	fmt.Println(viewgood)

	loginUserPassword := util.GetLoginUsernamePassword(&client)

	iptvWatcher := Iptv{
		iptvUsername: username,
		iptvPassword: loginUserPassword,
	}
	afterLogin(loginUserPassword)

	return &iptvWatcher
}

func afterLogin(videoUserPwd string) {
	setCache(videoUserPwd)
}

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
	getRequest, _ := http.NewRequest("GET", util.LoginUrl, nil)
	util.SetRequestHeader(getRequest)
	resp, err := client.Do(getRequest)
	if err != nil {
		return nil, err
	}
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

	postReq, _ := http.NewRequest("POST", util.LoginUrl, strings.NewReader(values.Encode()))
	util.SetRequestHeader(postReq)

	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	return resp.Cookies()[0].Value
}