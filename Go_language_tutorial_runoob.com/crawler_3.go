//并发
package main
import (
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
)
func main() {
	fmt.Println("Hello, world")
	url := "http://www.baidu.com/"
	queue := make(chan string)
	go func() {  //Golang使用关键字go即可开启一个新的go程，也叫goroutine
		queue <- url
	}()
	for uri := range queue {
		download(uri, queue)
	}
}
func download(url string, queue chan string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()
	links := collectlinks.All(resp.Body)
	for _, link := range links {
		fmt.Println("parse url", link)
		go func() {
			queue <- link
		}()
	}
}