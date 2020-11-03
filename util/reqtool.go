package util

import "net/http"

func SetRequestHeader(req *http.Request) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", "https://cas.ncu.edu.cn:8443/cas/login?service=http%3a%2f%2fwyjx.ncu.edu.cn%2fSPM%2fsso%2fDefault.aspx%3faction%3dpc%26host%3dwyjx.ncu.edu.cn")
	req.Header.Add("Host", "cas.ncu.edu.cn:8443")
	req.Header.Add("Origin", "https://cas.ncu.edu.cn:8443")
	req.Header.Add("Connection", "keep-alive")
}