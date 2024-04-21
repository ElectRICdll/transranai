package logic

import (
	"net/http"
)

type Crawler struct {
	dst string
}

func (c *Crawler) GET() (res *http.Response, err error) {
	res, err = http.Get(c.dst)
	return
}