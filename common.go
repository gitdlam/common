package common

import (
	"encoding/base64"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
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
	//return fmt.Sprintf("%v", time.Unix(0, time.Now().UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond))*int64(time.Millisecond)))[:23]
	return time.Now().Format("2006-01-02 15:04:05.000")
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

type TmpRow7 struct {
	Str1   string
	Str2   string
	Str3   string
	Str4   string
	Str5   string
	Str6   string
	Str7   string
	Int1   int64
	Int2   int64
	Int3   int64
	Int4   int64
	Int5   int64
	Int6   int64
	Int7   int64
	Float1 float64
	Float2 float64
	Float3 float64
	Float4 float64
	Float5 float64
	Float6 float64
	Float7 float64
}

func IndexOf(target int, intSlice []int) int {
	for i, v := range intSlice {
		if v == target {
			return i
		}
	}
	return -1
}

func StringToRows7(data string, config []string) []TmpRow7 {
	var translate [21]int
	strPosition := 1
	intPosition := 8
	floatPosition := 15
	var result []TmpRow7

	for j, t := range config {
		if t == "string" {
			translate[j] = strPosition
			strPosition++
		} else if t == "int64" {
			translate[j] = intPosition
			intPosition++

		} else if t == "float64" {
			translate[j] = floatPosition
			floatPosition++
		}

	}

	lines := strings.Split(data, "\n")
	for _, v := range lines {
		if v == "" {
			continue
		}
		columns := strings.Split(v, "\t")
		row := TmpRow7{}
		index := 0
		for j, col := range columns {
			index = translate[j]
			switch {
			case index == 1:
				row.Str1 = col
			case index == 2:
				row.Str2 = col
			case index == 3:
				row.Str3 = col
			case index == 4:
				row.Str4 = col
			case index == 5:
				row.Str5 = col
			case index == 6:
				row.Str6 = col
			case index == 7:
				row.Str7 = col
			case index >= 7 && index <= 14:
				tmpInt, err := strconv.ParseInt(col, 10, 64)
				if err != nil {
					tmpInt = 0
				}
				switch {
				case index == 8:
					row.Int1 = tmpInt
				case index == 9:
					row.Int2 = tmpInt
				case index == 10:
					row.Int3 = tmpInt
				case index == 11:
					row.Int4 = tmpInt
				case index == 12:
					row.Int5 = tmpInt
				case index == 13:
					row.Int6 = tmpInt
				case index == 14:
					row.Int7 = tmpInt
				}
			case index >= 15 && index <= 21:
				tmpFloat, err := strconv.ParseFloat(col, 64)
				if err != nil {
					tmpFloat = 0
				}
				switch {
				case index == 15:
					row.Float1 = tmpFloat
				case index == 16:
					row.Float2 = tmpFloat
				case index == 17:
					row.Float3 = tmpFloat
				case index == 18:
					row.Float4 = tmpFloat
				case index == 19:
					row.Float5 = tmpFloat
				case index == 20:
					row.Float6 = tmpFloat
				case index == 21:
					row.Float7 = tmpFloat
				}

			}
		}

		result = append(result, row)
	}
	return result
}

func Base64Encode(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}

func Base64Decode(encoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return ""
	}
	return string(decoded)

}

func ExcelDateNumber(s string) int {
	var duration time.Duration

	ref, _ := time.Parse("2006-01-02", "1970-01-01")
	then, err := time.Parse("2006-01-02", s)

	if err == nil {
		duration = then.Sub(ref)
	}
	return int(duration.Hours()/24) + 25569

}

func ExcelRefresh(fileName string) {
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("Excel.Application")
	excel, _ := unknown.QueryInterface(ole.IID_IDispatch)
	oleutil.PutProperty(excel, "Visible", true)

	workbooks := oleutil.MustGetProperty(excel, "Workbooks").ToIDispatch()

	workbook, err := oleutil.CallMethod(workbooks, "Open", fileName)

	if err != nil {
		log.Fatalln(err)
	}
	wb := workbook.ToIDispatch()
	oleutil.CallMethod(excel, "CalculateFull")
	oleutil.CallMethod(wb, "Close", true)

	wb.Release()

	workbooks.Release()

	oleutil.CallMethod(excel, "Quit")
	excel.Release()
	ole.CoUninitialize()

}
