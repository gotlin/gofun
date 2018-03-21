package gorotinepool

import (
	"testing"
	"time"
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"github.com/hu17889/go_spider/core/common/mlog"
	"log"
)

func TestGpool_NewPool(t *testing.T) {

	p := NewPool(3)

	for {
		p.Get()
		go func() {
			defer p.Free()
			time.Sleep(2 * time.Second)
			log.Println("end")
		}()
	}

}

func TestHttp(t *testing.T){
	client := &http.Client{}

	httpreq, err := http.NewRequest("GET", "http://www.baidu.com", strings.NewReader("&a=1"))

	var resp *http.Response
	if resp, err = client.Do(httpreq); err != nil {
		fmt.Println(err)
	}


	var sorbody []byte
	if sorbody, err = ioutil.ReadAll(resp.Body); err != nil {
		mlog.LogInst().LogError(err.Error())
		// For gb2312, an error will be returned.
		// Error like: simplifiedchinese: invalid GBK encoding
		// return ""
	}
	//e,name,certain := charset.DetermineEncoding(sorbody,contentTypeStr)
	bodystr := string(sorbody)

	resp.Body.Close()
	fmt.Println(bodystr)

}