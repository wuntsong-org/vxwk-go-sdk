package v1

import (
	"github.com/wuntsong-org/vxwk-go-sdk/client"
	"github.com/wuntsong-org/vxwk-go-sdk/model"
	errors "github.com/wuntsong-org/wterrors"
	"net/url"
)

type External struct {
	c *client.Client
}

func NewExternal(c *client.Client) *External {
	return &External{
		c: c,
	}
}

func (s *LiveCodeFile) GetLogoUrl(id string, opt ...Options) (fileURL string, err errors.WTError) {
	val := url.Values{}

	for _, o := range opt {
		o(&val)
	}

	val.Set("projectid", id)

	resp, err := s.c.SendGetRequests("/api/v1/admin/external/img", val, nil)
	if err != nil {
		return "", err
	}

	fileURL = resp.Header.Get("Location")
	if len(fileURL) == 0 {
		return "", errors.Errorf("get url fail")
	}

	return fileURL, nil
}

func (s *External) GetList(opt ...Options) (resp *model.GetExternalListResp, err errors.WTError) {
	val := url.Values{}
	for _, o := range opt {
		o(&val)
	}

	resp = new(model.GetExternalListResp)
	_, err = s.c.SendGetRequests("/api/v1/user/external/list", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *External) GetInfo(id string, opt ...Options) (resp *model.GetExternalResp, err errors.WTError) {
	val := url.Values{}

	for _, o := range opt {
		o(&val)
	}

	val.Set("id", id)

	resp = new(model.GetExternalResp)
	_, err = s.c.SendGetRequests("/api/v1/user/external", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *External) Create(req model.CreateExternalReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/external/create", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *External) Update(req model.UpdateExternalReq, opt ...Options) (resp *model.Resp, err errors.WTError) {
	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/external/update", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *External) Delete(id string, opt ...Options) (resp *model.Resp, err errors.WTError) {
	req := model.DeleteID{
		ID: id,
	}

	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/external/delete", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
