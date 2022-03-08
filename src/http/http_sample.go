package _http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SampleMain()  {

	// 实际上也是使用client
	resp,_ := http.Get("http://baidu.com/")
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))

	// 直接使用client
	client := http.Client{}
	resp2,_ := client.Get("http://baidu.com/")
	content2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println(string(content2))

	// 最本质的方式
	req, _ := http.NewRequest("GET", "http://baidu.com/", nil)
	resp3,_ := http.DefaultClient.Do(req)
	content3,_ := ioutil.ReadAll(resp3.Body)
	fmt.Println(string(content3))

}
