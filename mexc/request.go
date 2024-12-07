package mexc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/msw-x/moon/refl"
	"github.com/msw-x/moon/ustring"

	"github.com/msw-x/moon/uhttp"
)

func GetPub[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodGet, path, req, transform, false)
}

func Get[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodGet, path, req, transform, true)
}

func Post[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodPost, path, req, transform, true)
}

func Delete[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodDelete, path, req, transform, true)
}

func request[R, T any](c *Client, method string, path string, request any, transform func(R) (T, error), sign bool) (r Response[T]) {
	var attempt int
	for {
		r = req(c, method, path, request, transform, sign)
		if r.StatusCode != http.StatusOK && c.onTransportError != nil {
			if c.onTransportError(r.Error, r.StatusCode, attempt) {
				attempt++
				continue
			}
		}
		break
	}
	return
}

func req[R, T any](c *Client, method string, path string, request any, transform func(R) (T, error), sign bool) (r Response[T]) {
	var perf *uhttp.Performer
	switch method {
	case http.MethodGet:
		perf = c.c.Get(path).Params(request)
	case http.MethodPost:
		//perf = c.c.Post(path).Params(request)
		perf = ParamsCustom(c.c.Post(path), request)
	case http.MethodDelete:
		perf = c.c.Delete(path).Params(request)
	default:
		r.Error = fmt.Errorf("forbidden method: %s", method)
		return
	}
	if sign {
		if c.s == nil {
			r.Error = errors.New("api key is empty")
			r.NetError = true
			return
		}

		if perf.Request.Header == nil {
			perf.Request.Header = make(http.Header)
		}
		switch method {
		case http.MethodGet, http.MethodPost, http.MethodDelete:
			c.s.Sign(perf)
		}
	}
	h := perf.Do()
	if h.Error == nil {
		r.StatusCode = h.StatusCode
		if h.StatusCode != http.StatusOK {
			r.Error = errors.New(h.BodyString())
		}
		if h.BodyExists() {
			raw := new(R)
			err := h.Json(raw)
			if err != nil {
				r.Error = fmt.Errorf("json decode error: %s - %s", err, h.BodyString())
				return
			}
			if r.Ok() {
				r.Data, r.Error = transform(*raw)
			}
		}
		if sign {
			//r.SetErrorIfNil(h.HeaderTo(&r.Limit))
		}
	} else {
		r.Error = h.Error
		r.NetError = true
	}
	return
}

func ParamsCustom(o *uhttp.Performer, s any) *uhttp.Performer {
	refl.WalkOnTagsAny(s, "url", func(v any, name string, flags []string) {
		param(o, name, v, OmitEmpty(flags))
	})
	return o
}

func param(o *uhttp.Performer, name string, value any, omitempty bool) {
	if o.Request.Params == nil {
		o.Request.Params = make(url.Values)
	}
	if v, omit := Marshal(value, omitempty); !omit {
		name = ustring.TitleLowerCase(name)
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			elems, _ := json.Marshal(value)
			o.Request.Params.Set(name, string(elems))
			return
		}
		o.Request.Params.Set(name, v)
	}
}

func OmitEmpty(flags []string) bool {
	for _, flag := range flags {
		if flag == "omitempty" {
			return true
		}
	}
	return false
}
func IsEmpty(s string) bool {
	return s == "" || s == "0"
}

func Marshal(v any, omitempty bool) (s string, omit bool) {
	s = fmt.Sprint(v)
	omit = omitempty && IsEmpty(s)
	return
}
