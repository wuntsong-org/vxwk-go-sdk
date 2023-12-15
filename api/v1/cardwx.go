package v1

import (
	"github.com/wuntsong-org/vxwk-go-sdk/client"
	"github.com/wuntsong-org/vxwk-go-sdk/model"
	errors "github.com/wuntsong-org/wterrors"
	"net/url"
)

type CardWX struct {
	c *client.Client
}

func NewCardWX(c *client.Client) *CardWX {
	return &CardWX{
		c: c,
	}
}

func (s *CardWX) GetImgUrl(id string, opt ...Options) (fileURL string, err errors.WTError) {
	val := url.Values{}

	for _, o := range opt {
		o(&val)
	}

	val.Set("projectid", id)

	resp, err := s.c.SendGetRequests("/api/v1/admin/cardwx/img", val, nil)
	if err != nil {
		return "", err
	}

	fileURL = resp.Header.Get("Location")
	if len(fileURL) == 0 {
		return "", errors.Errorf("get url fail")
	}

	return fileURL, nil
}

func (s *CardWX) GetList(opt ...Options) (resp *model.GetCardWXListResp, err errors.WTError) {
	val := url.Values{}
	for _, o := range opt {
		o(&val)
	}

	resp = new(model.GetCardWXListResp)
	_, err = s.c.SendGetRequests("/api/v1/user/cardwx/list", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CardWX) GetInfo(id string, opt ...Options) (resp *model.GetCardWXResp, err errors.WTError) {
	val := url.Values{}

	for _, o := range opt {
		o(&val)
	}

	val.Set("id", id)

	resp = new(model.GetCardWXResp)
	_, err = s.c.SendGetRequests("/api/v1/user/cardwx", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CardWX) Create(req model.CreateCardWXReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/cardwx/create", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CardWX) Update(req model.UpdateCardWXReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/cardwx/update", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CardWX) Delete(id string, opt ...Options) (resp *model.Resp, err errors.WTError) {
	req := model.DeleteID{
		ID: id,
	}

	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/cardwx/delete", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
