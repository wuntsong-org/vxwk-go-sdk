package v1

import (
	"github.com/wuntsong-org/vxwk-go-sdk/client"
	"github.com/wuntsong-org/vxwk-go-sdk/model"
	errors "github.com/wuntsong-org/wterrors"
	"net/url"
)

type CardDY struct {
	c *client.Client
}

func NewCardDY(c *client.Client) *CardDY {
	return &CardDY{
		c: c,
	}
}

func (s *CardDY) GetImgUrl(id string, opt ...Options) (fileURL string, err errors.WTError) {
	val := url.Values{}

	for _, o := range opt {
		o(&val)
	}

	val.Set("projectid", id)

	resp, err := s.c.SendGetRequests("/api/v1/admin/carddy/img", val, nil)
	if err != nil {
		return "", err
	}

	fileURL = resp.Header.Get("Location")
	if len(fileURL) == 0 {
		return "", errors.Errorf("get url fail")
	}

	return fileURL, nil
}

func (s *CardDY) GetList(opt ...Options) (resp *model.GetCardDYListResp, err errors.WTError) {
	val := url.Values{}
	for _, o := range opt {
		o(&val)
	}

	resp = new(model.GetCardDYListResp)
	_, err = s.c.SendGetRequests("/api/v1/user/carddy/list", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CardDY) GetInfo(id string, opt ...Options) (resp *model.GetCardDYResp, err errors.WTError) {
	val := url.Values{}

	for _, o := range opt {
		o(&val)
	}

	val.Set("id", id)

	resp = new(model.GetCardDYResp)
	_, err = s.c.SendGetRequests("/api/v1/user/carddy", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CardDY) Create(req model.CreateCardDYReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/carddy/create", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CardDY) Update(req model.UpdateCardDYReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/carddy/update", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CardDY) Delete(id string, opt ...Options) (resp *model.Resp, err errors.WTError) {
	req := model.DeleteID{
		ID: id,
	}

	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/carddy/delete", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
