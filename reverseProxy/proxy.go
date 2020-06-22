package reverseProxy

import (
	"bytes"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/19 14:17
 * @Title:
 * --- --- ---
 * @Desc:
 */

type MsgExchange interface {
	PatchHeaders()
}

type ReverseProxy struct {
	TargetUrl string
	target    *url.URL
	Proxy     *httputil.ReverseProxy
}

func NewReverseProxy() *ReverseProxy {
	r := &ReverseProxy{TargetUrl: "http://aki.red"}
	r.Proxy = r.InitProxy()
	return r
}

func (r *ReverseProxy) SetTarget(targetUrl string) {
	r.TargetUrl = targetUrl
	r.target, _ = url.Parse(targetUrl)
}

func (r *ReverseProxy) InitProxy() *httputil.ReverseProxy {
	proxy := &httputil.ReverseProxy{}
	trans := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			Renegotiation:      tls.RenegotiateFreelyAsClient,
		},
		DialContext: (&net.Dialer{
			Timeout:   300 * time.Second,
			KeepAlive: 300 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   300 * time.Second,
		ResponseHeaderTimeout: 300 * time.Second,
		ExpectContinueTimeout: 150 * time.Second,
		IdleConnTimeout:       300 * time.Second,
		Proxy:                 http.ProxyFromEnvironment,
	}
	proxy.Transport = trans

	proxy.Director = func(req *http.Request) {
		err := r.rewireRequest(req)
		if err != nil {
			log.Printf("proxy director rewite request error: %s\n", err.Error())
		}

		req.URL.Scheme = r.target.Scheme
		req.URL.Host = r.target.Host
		req.URL.Path = r.singleJoiningSlash(r.target.Path, req.URL.Path)
		if r.target.RawQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = r.target.RawQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = r.target.RawQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			req.Header.Set("User-Agent", "")
		}
	}

	proxy.ModifyResponse = r.modifyResponse

	return proxy
}

func (r *ReverseProxy) ServeProxy(c *gin.Context) {
	r.Proxy.ServeHTTP(c.Writer, c.Request)
}

func (r *ReverseProxy) modifyResponse(response *http.Response) (err error) {
	resp := Resp{response}
	resp.PatchHeaders()

	buffer, err := resp.Decompress()
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	buffer = r.cleanURL(buffer)
	resp.Compress(buffer)

	return
}

func (r *ReverseProxy) rewireRequest(request *http.Request) (err error) {
	req := Req{request}
	req.PatchHeaders()

	if request.Body != nil {
		reader := request.Body
		var buffer []byte
		buffer, err = ioutil.ReadAll(reader)
		if err != nil {
			return
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(buffer))
		req.ContentLength = int64(len(buffer))
		req.Header.Set("Content-Length", strconv.Itoa(len(buffer)))

		err = reader.Close()
		if err != nil {
			return
		}

		err = request.Body.Close()
		if err != nil {
			return
		}
	}
	return
}

func (r *ReverseProxy) singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func (r *ReverseProxy) cleanURL(buffer []byte) (buf []byte) {
	buf = bytes.Replace(buffer, []byte("http://"), []byte("http://"), -1)
	return
}
