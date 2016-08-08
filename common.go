package common

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
	"time"
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

func TimeNowString() string {
	return fmt.Sprintf("%v", time.Unix(0, time.Now().UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond))*int64(time.Millisecond)))[:23]
}

func EscapeLatex(s string) string {
	s2 := strings.Replace(s, "\\", "\\textbackslash", -1)
	s2 = strings.Replace(s2, "&", "\\&", -1)
	s2 = strings.Replace(s2, "%", "\\%", -1)
	s2 = strings.Replace(s2, "$", "\\$", -1)
	s2 = strings.Replace(s2, "#", "\\#", -1)
	s2 = strings.Replace(s2, "_", "\\_", -1)
	s2 = strings.Replace(s2, "{", "\\{", -1)
	s2 = strings.Replace(s2, "}", "\\}", -1)
	s2 = strings.Replace(s2, "~", "\\textasciitilde", -1)
	return strings.Replace(s2, "^", "\\textasciicircum", -1)

}

func PadZero(s string) string {
	if s[0] == '.' {
		return "0" + s
	} else {
		return s
	}
}
