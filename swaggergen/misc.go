package swaggergen

import (
	"regexp"
	"strings"
)

var basicTypes = map[string]string{
	"bool":              "boolean::",
	"uint":              "integer:int32:",
	"uint8":             "integer:int32:",
	"uint16":            "integer:int32:",
	"uint32":            "integer:int32:",
	"uint64":            "integer:int64:",
	"int":               "integer:int64:",
	"int8":              "integer:int32:",
	"int16":             "integer:int32:",
	"int32":             "integer:int32:",
	"int64":             "integer:int64:",
	"uintptr":           "integer:int64:",
	"float32":           "number:float:",
	"float64":           "number:double:",
	"string":            "string::",
	"complex64":         "number:float:",
	"complex128":        "number:double:",
	"byte":              "string:byte:",
	"rune":              "string:byte:",
	"time.Time":         "string::2006-01-02T15:04:05+08:00",
	"time.Duration":     "integer:int64:",
	"statuscode.Status": "integer:int64:",
	"Money":             "integer:int64:",
}

// 解析出有意义的一行
var parseAtRows = regexp.MustCompile(`@(\S+)\s+(.+)`)

// 解析出路由对应的方法
var parseRouter = regexp.MustCompile(`(\S+)\s+\[(\w+)\]`)

// 解析出参数
var parseParam = regexp.MustCompile(`(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(.+)`)

// 解析出返回
var parseResp = regexp.MustCompile(`(\S+)\s+(.+)`)

// 正则出 每一行的数据
func ParseAtRows(data string) map[string][]string {

	ds := strings.Split(data, "\n")

	for i := 0; i != len(ds); i++ {
		v := ds[i]
		if len(v) != 0 && v[len(v)-1] == '\\' && i+1 != len(ds) {
			v2 := ds[i+1]
			ds[i] = v + " " + v2
			ds = append(ds[:i+1], ds[i+2:]...)
		}

	}

	ret := map[string][]string{}
	for _, v := range ds {
		d := parseAtRows.FindStringSubmatch(v)
		if len(d) >= 3 {
			k := strings.ToLower(d[1])
			ret[k] = append(ret[k], d[2])
		}
	}
	//ffmt.Puts(ret)
	return ret
}
