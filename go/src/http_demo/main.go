package main
// https://blog.xiec.live/archives/48.html
import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)
func get() {
    client := &http.Client{}

    // 地址为api.imjad.cn的一言接口，执行后可以收到一言数据
    res, err := client.Get("https://api.imjad.cn/hitokoto/")
	fmt.Println(res)
    if err != nil {
        fmt.Println(err)
        return
    }

    result, err := ioutil.ReadAll(res.Body)
    defer res.Body.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%s", result)
}

func post(){
	   // 配置需要发送的值
    type Data struct {
        Page  int `json: "page"`
        Count int `json: "count"`
    }

    data := Data{
        Page:  1,
        Count: 5,
    }

    params, err := json.Marshal(data)
    //string(params): {"page":"xx@xx.com","count":"123456"}

    client := &http.Client{}

    // 地址为api.apiopen.top的网易新闻接口
    res, err := client.Post("https://api.apiopen.top/getWangYiNews", "application/json", bytes.NewBuffer(params))

    if err != nil {
        fmt.Println(err)
        return
    }

    result, err := ioutil.ReadAll(res.Body)
    defer res.Body.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%s", result)
}

func main(){
	post()
}