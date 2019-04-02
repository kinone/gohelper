package gohelper

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
)

func Trim(str string) string {
	return strings.TrimSpace(str)
}

func StrToLower(str string) string {
	return strings.ToLower(str)
}

func StrToUpper(str string) string {
	return strings.ToUpper(str)
}

func IsIMEI(imei string) bool {
	match, _ := regexp.MatchString(`^\d{15}$`, imei)

	return match
}

func MD5(str string) string {
	bs := md5.Sum([]byte(str))
	return hex.EncodeToString(bs[:])
}

func LocalIP() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if nil != err {
		return nil, err
	}

	var ips []string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}

func HttpBuildQuery(params map[string]string) string {
	seg := make([]string, len(params))

	i := 0
	for k, v := range params {
		k = url.QueryEscape(k)
		v = url.QueryEscape(v)
		seg[i] = strings.Join([]string{k, v}, "=")
		i++
	}

	return strings.Join(seg, "&")
}

func HttpGet(url string) (string, error) {
	http.DefaultClient.Timeout = 3 * time.Second
	res, err := http.Get(url)
	if nil != err {
		return "", err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if nil != err {
		return "", err
	}

	return string(resBody), nil
}

func NewStopSignal() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	return ch
}
