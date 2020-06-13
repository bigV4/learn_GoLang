//解析网页
package main
import (
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
)
func main() {
	fmt.Println("Hello, world")
	url := "http://www.baidu.com/"
	download(url)
}
func download(url string) {
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
	}
}

/*
go常见的解析器xpath、jquery、正则都有，直接搜索即可，我这里偷懒，直接用别人写好的轮子collectlinks，可以提取网页中所有的链接.
下载方法go get -u github.com/jackdanger/collectlinks
*/