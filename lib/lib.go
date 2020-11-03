package lib

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"iptv-helper/util"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

const (
	loginUrl = "https://cas.ncu.edu.cn:8443/cas/login?service=http://wyjx.ncu.edu.cn/SPM/sso/Default.aspx?action=pc&host=wyjx.ncu.edu.cn"
)

var (
	lt = ""
	viewgood = ""
)

func Run(username, password string) {
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		req.Host = "wyjx.ncu.edu.cn"
		req.Header.Set("Host", "wyjx.ncu.edu.cn")
		req.URL.Host = "wyjx.ncu.edu.cn"
		return nil
	}
	beforeLogin(&client)
	fmt.Println("publicKey:", util.PwdEncoderInstance.PublicKey)
	fmt.Println("lt:", lt)

	viewgood = sendLoginRequest(&client, username, password)
	fmt.Println(viewgood)

	loginUserPassword := getLoginUsernamePassword(&client)
	fmt.Println(loginUserPassword)

	toBaseVideoUrl := getVideoUrl(username, loginUserPassword)
	fmt.Println(toBaseVideoUrl)
	getBaseVideoUrl(toBaseVideoUrl)
}

func getBaseVideoUrl(baseUrl string) string {
	resp, err := http.Get(baseUrl)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	respPlain, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respPlain))
	return util.ParseXML(&respPlain)
}

func getVideoUrl(username, password string) string {
	//baseUrl := "http://wyjx.ncu.edu.cn/VIEWGOOD/adi/portal/load.ashx?ModeType=PlayVOD,StreamType=HTTP_MP4,Ver=8.0.0.2,"
	testUrl := fmt.Sprint("http://wyjx.ncu.edu.cn/VIEWGOOD/adi/portal/load.ashx?" +
		"ModeType=PlayVOD," +
		"StreamType=HTTP_MP4," +
		"Ver=8.0.0.2," +
		"StreamID=7037," +
		"ClassID=53," +
		"ClassName=%e5%bd%b1%e8%a7%86%e6%ac%a3%e8%b5%8f," +
		"assetID=193," +
		"assetName=%e7%88%86%e8%a3%82%e9%bc%93%e6%89%8b," +
		"Episode_ID=1," +
		"Username="+ username + "," +
		"Password=" + password + "," +
		"Redirect=false," +
		"Random="+ fmt.Sprintf("%v", time.Now().Unix()) + "000")
	fmt.Println(testUrl)
	return testUrl
}

func getLoginUsernamePassword(client *http.Client) string {
	getRequest, _ := http.NewRequest("GET", "http://wyjx.ncu.edu.cn/SPM/sso/Default.aspx?action=pc&host=wyjx.ncu.edu.cn", nil)
	getRequest.Header.Add("Host", "wyjx.ncu.edu.cn")
	resp, err := client.Do(getRequest)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	pageDoc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	pageText := pageDoc.Find("body").Text()
	start := strings.Index(pageText, "loginuserpassword=") + 19
	return pageText[start:start+16]
	//pageBody, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(pageBody))
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

	postReq, _ := http.NewRequest("POST", loginUrl, strings.NewReader(values.Encode()))
	reqBody, _ := ioutil.ReadAll(postReq.Body)
	fmt.Println("reqBody:", string(reqBody))
	util.SetRequestHeader(postReq)

	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	//respText, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.Cookies())
	//fmt.Println(string(respText))
	return resp.Cookies()[0].Value
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
	lt = getLt(page)
	util.PwdEncoderInstance.PublicKey = getPublicKey(page)
}

func getPage(client *http.Client) (io.ReadCloser, error) {
	getRequest, _ := http.NewRequest("GET", loginUrl, nil)
	util.SetRequestHeader(getRequest)
	resp, err := client.Do(getRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println("Header:", resp.Header)
	return resp.Body, nil
}

func getLt(page *goquery.Document) string {
	ltSection, _ := page.Find("#cryptoType").Next().Attr("value")
	return ltSection
}

func getPublicKey(page *goquery.Document) string {
	toSearch := page.Find("body").Text()
	startPos := strings.Index(toSearch, "var n= \"")
	return toSearch[startPos+8:startPos+256+8]
}