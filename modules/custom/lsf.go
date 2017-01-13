package custom

import (
	"net/http"
	"errors"
	"time"
	"net/url"
	"strings"
)

const uri = "https://www.lsf.hs-weingarten.de/qisserver/servlet/de.his.servlet.RequestDispatcherServlet?state=user&type=1&category=auth.login&startpage=portal.vm"
var lsf_err = errors.New("lsf_error")

var clientRed = &http.Client{
	Timeout:       5 * time.Second,
	CheckRedirect: Redirect,
	Transport: &http.Transport{
		Proxy:             http.ProxyFromEnvironment,
		DisableKeepAlives: true,
	},
}

//Redirect is used to prevent the http client from following the Location Header
//if the http statuscode is 302
func Redirect(req *http.Request, via []*http.Request) error {
	return lsf_err
}

func CheckValidUser(username, password string) (bool, error) {

	param := url.Values{}
	param.Set("username", username)
	param.Add("password", password)
	param.Add("submit", "Anmelden")

	req, err := http.NewRequest("POST", uri, strings.NewReader(param.Encode()))
	if err != nil {
		return false, err
	}

	req.Header.Add("Host", "www.lsf.hs-weingarten.de")
	req.Header.Add("Origin", "https://www.lsf.hs-weingarten.de")
	req.Header.Add("Referer", "https://www.lsf.hs-weingarten.de/qisserver/rds?state=user&type=0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Language", "de-DE,de;q=0.8,en-US;q=0.6,en;q=0.4")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Pragma", "no-cache")

	resp, err := clientRed.Do(req)
	//check if an error occured and if the error isnt the redirect prevention
	if err != nil && !strings.Contains(err.Error(), "lsf_error") {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 302 {
		return true, nil
	} else if resp.StatusCode == 200 {
		return false, nil
	} else {
		return false, errors.New("lsf nicht erreichbar")
	}
}
