package mexc

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/msw-x/moon/uhttp"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type Sign struct {
	Key    string
	Secret string
}

func NewSign(key, secret string) *Sign {
	o := new(Sign)
	o.Key = key
	o.Secret = secret
	return o
}

func (o *Sign) Sign(perf *uhttp.Performer) {
	if perf.Request.Params == nil {
		perf.Request.Params = make(url.Values)
	}
	perf.Request.Params.Add("timestamp", strconv.FormatInt(time.Now().UnixMilli(), 10))

	encodedParams := encodeSortParams(perf.Request.Params)

	signature := o.signHmac(encodedParams, o.Secret)
	perf.Request.Params.Set("signature", signature)
	o.header(perf.Request.Header)
}

func encodeSortParams(src url.Values) (s string) {
	if len(src) == 0 {
		return
	}
	keys := make([]string, len(src))
	i := 0
	for k := range src {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		s += encodeParam(k, src.Get(k)) + "&"
	}
	s = s[0 : len(s)-1]
	return
}

func encodeParam(name, value string) string {
	params := url.Values{}
	params.Add(name, value)
	return params.Encode()
}

func encodeSortBody(u *url.Values, body []byte) {
	mapData := map[string]any{}
	err := json.Unmarshal(body, &mapData)
	if err != nil {
		panic(err)
	}
	for key, value := range mapData {
		u.Add(key, fmt.Sprintf("%v", value))
	}
	return
}

func (o *Sign) timestamp() string {

	// Get the current time in UTC
	currentTime := time.Now().UTC().UnixMilli()
	return fmt.Sprintf("&timestamp=%d", currentTime)
}

func (o *Sign) header(h http.Header) {
	h.Set("x-mexc-apikey", o.Key)
}

//signValue is requestBody or queryString or both
func (o *Sign) signHmac(preSignedString, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, preSignedString)
	return hex.EncodeToString(h.Sum(nil))
}

func signHmac(preSignedString, secret string) string {

	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, preSignedString)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func nowUtcMs() int {
	return int(time.Now().UTC().Unix())
}
