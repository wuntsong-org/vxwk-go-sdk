package v1

import (
	"github.com/wuntsong-org/vxwk-go-sdk/client"
	"github.com/wuntsong-org/vxwk-go-sdk/model"
	errors "github.com/wuntsong-org/wterrors"
	"net/url"
)

type LiveCode struct {
	c *client.Client
}

func NewLiveCode(c *client.Client) *LiveCode {
	return &LiveCode{
		c: c,
	}
}

func (s *LiveCode) GetList(opt ...Options) (resp *model.GetLiveCodeListResp, err errors.WTError) {
	val := url.Values{}
	for _, o := range opt {
		o(&val)
	}

	resp = new(model.GetLiveCodeListResp)
	_, err = s.c.SendGetRequests("/api/v1/user/livecode/list", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LiveCode) GetInfo(id string, opt ...Options) (resp *model.GetLiveCodeResp, err errors.WTError) {
	val := url.Values{}

	for _, o := range opt {
		o(&val)
	}

	val.Set("id", id)

	resp = new(model.GetLiveCodeResp)
	_, err = s.c.SendGetRequests("/api/v1/user/livecode", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LiveCode) Create(req model.CreateLiveCodeReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/livecode/create", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LiveCode) Update(req model.UpdateLiveCodeReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/livecode/update", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LiveCode) Delete(id string, opt ...Options) (resp *model.Resp, err errors.WTError) {
	req := model.DeleteID{
		ID: id,
	}

	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/livecode/delete", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
