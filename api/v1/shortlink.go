package v1

import (
	errors "github.com/wuntsong-org/wterrors"
	"go-sdk/client"
	"go-sdk/model"
	"net/url"
)

type ShortLink struct {
	c *client.Client
}

func NewShortLink(c *client.Client) *ShortLink {
	return &ShortLink{
		c: c,
	}
}

func (s *ShortLink) GetList(opt ...Options) (resp *model.GetShortLinkListResp, err errors.WTError) {
	val := url.Values{}
	for _, o := range opt {
		o(&val)
	}

	resp = new(model.GetShortLinkListResp)
	_, err = s.c.SendGetRequests("/api/v1/user/shortlink/list", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ShortLink) GetInfo(id string, opt ...Options) (resp *model.GetShortLinkResp, err errors.WTError) {
	val := url.Values{}

	for _, o := range opt {
		o(&val)
	}

	val.Set("id", id)

	resp = new(model.GetShortLinkResp)
	_, err = s.c.SendGetRequests("/api/v1/user/shortlink", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ShortLink) Create(req model.CreateShortLinkReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/shortlink/create", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ShortLink) Update(req model.UpdateShortLinkReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/shortlink/update", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ShortLink) Delete(id string, opt ...Options) (resp *model.Resp, err errors.WTError) {
	req := model.DeleteID{
		ID: id,
	}

	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/shortlink/delete", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
