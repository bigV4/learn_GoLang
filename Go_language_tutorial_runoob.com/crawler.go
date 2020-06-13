package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)
func main() {
	fmt.Println("Hello, world")
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println(string(body)[:100])
	fmt.Println(resp.StatusCode)
	fmt.Println("StatusCode: " + strconv.Itoa(resp.StatusCode))
}