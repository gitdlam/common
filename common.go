package common

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

func ValidBarcode8(s string) bool {
	re := regexp.MustCompile("^[A-Z0-9]{8,8}$")
	return re.MatchString(s)
}

func ValidBarcode7(s string) bool {
	re := regexp.MustCompile("^[A-Z0-9]{7,7}$")
	return re.MatchString(s)
}

func ValidCarton20(s string) bool {
	re := regexp.MustCompile("^[0-9]{20,20}$")
	return re.MatchString(s)
}

func ReverseProxy(proxyPort string, pathMap map[string]string) {

	for urlPath, targetPort := range pathMap {
		u, _ := url.Parse("http://127.0.0.1:" + targetPort)
		http.Handle(urlPath, httputil.NewSingleHostReverseProxy(u))
	}

	http.ListenAndServe(":"+proxyPort, nil)

}
