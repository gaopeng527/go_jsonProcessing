// JSON Processing project main.go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	//	parseJson()
	generateJson()
}

func parseJson() {
	// 1.知道结构采用官方提供的方式
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	err := json.Unmarshal([]byte(str), &s)
	checkErr(err)
	fmt.Println(s)
	// 2.不知道结构官方解析方式
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b, &f)
	checkErr(err)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case bool:
			fmt.Println(k, "is bool", vv)
		case float64:
			fmt.Println(k, "is numbers", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	// 3.采用Bity公司开源的simplejson包，在处理未知结构的JSON时很方便
	js, err := simplejson.NewJson([]byte(`{
    "test": {
        "array": [1, "2", 3],
        "int": 10,
        "float": 5.150,
        "bignum": 9223372036854775807,
        "string": "simplejson",
        "bool": true
    }
	}`))
	arr, _ := js.Get("test").Get("array").Array()
	fmt.Println(arr)
	i, _ := js.Get("test").Get("int").Int()
	fmt.Println(i)
	ms := js.Get("test").Get("string").MustString()
	fmt.Println(ms)
}

func generateJson() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(&s)
	checkErr(err)
	fmt.Println(string(b))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
