package pagination

import (
	"net/url"
	"strconv"
)

type Page struct {
	size  *int
	start *string
}

func New() *Page {
	return &Page{}
}

func (p *Page) SetSize(size int) *Page {
	p.size = &size
	return p
}

func (p *Page) GetSize() string {
	if p.size == nil {
		return "1000" // return default value
	}
	return strconv.Itoa(*p.size)
}

func (p *Page) SetStart(start string) *Page {
	p.start = &start
	return p
}

func (p *Page) GetStart() string {
	if p.start == nil {
		return ""
	}
	return *p.start
}

func (p *Page) Apply(req *url.URL) {
	q := req.Query()

	if p.size != nil {
		q.Add("page.size", strconv.Itoa(*p.size))
	}

	if p.start != nil && len(*p.start) > 0 {
		q.Add("page.start", *p.start)
	}

	req.RawQuery = q.Encode()
}
