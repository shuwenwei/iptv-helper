package util

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func GetLt(page *goquery.Document) string {
	ltSection, _ := page.Find("#cryptoType").Next().Attr("value")
	return ltSection
}

func GetPublicKey(page *goquery.Document) string {
	toSearch := page.Find("body").Text()
	startPos := strings.Index(toSearch, "var n= \"")
	return toSearch[startPos+8:startPos+256+8]
}

func GetLoginUsernamePassword(client *http.Client) string {
	getRequest, _ := http.NewRequest("GET", "http://wyjx.ncu.edu.cn/SPM/sso/Default.aspx?action=pc&host=wyjx.ncu.edu.cn", nil)
	getRequest.Header.Add("Host", "wyjx.ncu.edu.cn")
	resp, err := client.Do(getRequest)
	if err != nil {
		log.Fatal("get loginUserPassword error")
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
}