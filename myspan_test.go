package myspan

import (
	"testing"
	"fmt"
	"bytes"
	"encoding/binary"
	"time"
	"io/ioutil"
	"net/http"

	"go.opencensus.io/trace"
)

func TestCreateSpan(t *testing.T){
	t.Run("testCreateSpan",func(t *testing.T){
		//name, endpoint, traceSample, prob := ParamInit()
		// fmt.Println(name, endpoint, traceSample, prob)
		span := CreateSpan("test2", "127.0.0.1:9411", "always", 0)
		defer func(){
			span.End()
			time.Sleep(30000*time.Millisecond)
		}()
		dowork()
	})
}

func dowork(){
	response, _ := http.Get("http://localhost:3220/hello")
	if response != nil {
		fmt.Println(response)
	} else {
		fmt.Println("response nil !!")
	}
	defer response.Body.Close()
	num, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(num))
	// byte 转换为 int
	buf := bytes.NewBuffer([]byte{0x00})
	_, err := binary.ReadVarint(buf)
	if err != nil {
		fmt.Println(trace.StatusCodeUnknown)
	}
}
