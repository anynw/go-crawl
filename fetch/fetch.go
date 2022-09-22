package fetch

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var rateLimit = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	// 延迟10毫秒执行
	<-rateLimit

	client := &http.Client{}
	// url := "https://book.douban.com/"
	reqest, err := http.NewRequest("GET", url, nil)
	// reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36 OPR/66.0.3515.115")
	reqest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	if err != nil {
		// fmt.Println("http get err:",err)
		panic(err)
	}
	resp, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error status code :", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
