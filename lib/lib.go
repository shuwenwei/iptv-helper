package lib

import (
	"errors"
	"fmt"
	"iptv-helper/util"
	"net/http"
	"net/http/cookiejar"
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

	loginUserPassword := util.GetLoginUsernamePassword(&client)
	fmt.Println(loginUserPassword)

	iptvWatcher := Iptv{
		iptvUsername: username,
		iptvPassword: loginUserPassword,
	}

	iptvWatcher.GetBaseVideoUrl()
	fmt.Println("baseURL:", iptvWatcher.baseUrl)
}