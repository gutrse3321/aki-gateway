package reverseProxy

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/19 14:22
 * @Title:
 * --- --- ---
 * @Desc:
 */

type Resp struct {
	*http.Response
}

func (r *Resp) PatchHeaders() {
	r.Header.Set("X-Tomonori-Remote", "aki")
	return
}

func (r *Resp) Decompress() ([]byte, error) {
	var err error
	body := r.Body
	defer body.Close()

	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func (r *Resp) Compress(buf []byte) {
	body := ioutil.NopCloser(bytes.NewReader(buf))
	r.Body = body
	r.ContentLength = int64(len(buf))
	r.Header.Set("Content-Length", strconv.Itoa(len(buf)))

	err := r.Body.Close()
	if err != nil {
		log.Printf("response compress body close error:%s\n", err.Error())
	}
}
