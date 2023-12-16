package v1

import (
	"fmt"
	"github.com/wuntsong-org/vxwk-go-sdk/utils"
	"net/url"
)

type Options func(any)

func Status(s ...int64) Options {
	return func(a any) {
		if len(s) == 0 {
			return
		}

		q, ok := a.(*url.Values)
		if !ok {
			return
		}

		if q == nil {
			return
		}

		q.Set("status", fmt.Sprintf("[%s]", utils.JoinIntToString(s, ",", true)))
	}
}

func MediaType(s ...string) Options {
	return func(a any) {
		if len(s) == 0 {
			return
		}

		q, ok := a.(*url.Values)
		if !ok {
			return
		}

		if q == nil {
			return
		}

		q.Set("mediaType", fmt.Sprintf("[%s]", utils.JoinStringToString(s, ",", true)))
	}
}

func Page(page, pageSize int64) Options {
	return func(a any) {
		if page <= 0 {
			return
		}

		if pageSize <= 0 {
			return
		}

		q, ok := a.(*url.Values)
		if !ok {
			return
		}

		if q == nil {
			return
		}

		q.Set("page", fmt.Sprintf("%d", page))
		q.Set("pagesize", fmt.Sprintf("%d", pageSize))
	}
}

func Src(src string) Options {
	return func(a any) {
		if len(src) == 0 {
			return
		}

		q, ok := a.(*url.Values)
		if !ok {
			return
		}

		if q == nil {
			return
		}

		q.Set("src", src)
	}
}

func Name(name string) Options {
	return func(a any) {
		if len(name) == 0 {
			return
		}

		q, ok := a.(*url.Values)
		if !ok {
			return
		}

		if q == nil {
			return
		}

		q.Set("name", name)
	}
}
