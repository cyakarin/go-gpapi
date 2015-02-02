package gpapi

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/iij/go-gpapi/protocol"
	"hash"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	HmacSHA1           = "HmacSHA1"
	HmacSHA256         = "HmacSHA256"
	SignatureVersion2  = "2"
	APIVersion20130901 = "20130901"
	EndpointJSON       = "https://gp.api.iij.jp/json"
	EndpointXML        = "https://gp.api.iij.jp/xml"
	TimeLayout         = "2006-01-02T15:04:05Z"
	PostContentType    = "application/x-www-form-urlencoded"
)

type API struct {
	AccessKey  string
	SecretKey  string
	Endpoint   string
	SignMethod string
	Expires    time.Duration
}

func NewAPI(accesskey, secretkey string) *API {
	dur, _ := time.ParseDuration("1h")
	return &API{AccessKey: accesskey,
		SecretKey:  secretkey,
		Endpoint:   EndpointJSON,
		SignMethod: HmacSHA256,
		Expires:    dur,
	}
}

// add common parameters from API struct
func (a API) FixParam(param url.Values) url.Values {
	expir := time.Now().Add(a.Expires).UTC()
	base := protocol.CommonArg{APIVersion: APIVersion20130901,
		AccessKeyId:      a.AccessKey,
		Action:           "dummy",
		Signature:        "dummy",
		Expire:           expir.Format(TimeLayout),
		SignatureMethod:  a.SignMethod,
		SignatureVersion: SignatureVersion2}
	baseparm, err := ValidateRequest(base)
	if err != nil {
		log.Fatalln("baseparam error", base, err)
	}
	for k, v := range param {
		baseparm.Set(k, v[0])
	}
	baseparm.Del("Signature")
	return baseparm
}

// write parameter string to io.Writer
func (a API) WriteParam(w io.Writer, param url.Values) {
	io.WriteString(w, param.Encode())
}

func convert1(r byte) string {
	passchar := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_~.-"
	if strings.ContainsRune(passchar, rune(r)) {
		return string(r)
	}
	return fmt.Sprintf("%%%02X", r)
}

func CustomEscape(v string) string {
	res := ""
	for _, c := range []byte(v) {
		res += convert1(c)
	}
	return res
}

// calculate HMAC Signature
func (a API) GetSignature(method string, param url.Values) string {
	param = a.FixParam(param)
	delete(param, "Signature")
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var qs []string
	for _, k := range keys {
		s := CustomEscape(k) + "=" + CustomEscape(param.Get(k))
		qs = append(qs, s)
	}
	purl, _ := url.Parse(a.Endpoint)
	tgtstr := method + "\n"
	tgtstr += purl.Host + "\n"
	tgtstr += purl.Path + "\n"
	tgtstr += strings.Join(qs, "&")
	var hfn func() hash.Hash
	switch a.SignMethod {
	case HmacSHA1:
		hfn = sha1.New
	case HmacSHA256:
		hfn = sha256.New
	}
	mac := hmac.New(hfn, []byte(a.SecretKey))
	io.WriteString(mac, tgtstr)
	macstr := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(macstr)
}

// add signature to request
func (a API) AddSignature(method string, param url.Values) url.Values {
	param.Set("Signature", a.GetSignature(method, param))
	return param
}

func (a API) Call(req interface{}, res interface{}) (err error) {
	val, err := ValidateRequest(req)
	if err != nil {
		log.Println("invalid request", req, err)
		return
	}
	resp, err := a.Post(val)
	if err != nil {
		log.Println("request error", val, err)
		return
	}
	err = ParseResponse(resp.Body, res)
	return
}

// low-level Post
func (a API) Post(param url.Values) (resp *http.Response, err error) {
	param = a.FixParam(param)
	param = a.AddSignature("POST", param)
	cl := http.Client{}
	resp, err = cl.PostForm(a.Endpoint, param)
	return
}

// low-level Get
func (a API) Get(param url.Values) (resp *http.Response, err error) {
	param = a.FixParam(param)
	param = a.AddSignature("GET", param)
	cl := http.Client{}
	resp, err = cl.Get(a.Endpoint + "?" + param.Encode())
	return
}
