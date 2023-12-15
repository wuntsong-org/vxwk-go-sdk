package v1

import (
	"encoding/base64"
	"github.com/wuntsong-org/vxwk-go-sdk/client"
	"github.com/wuntsong-org/vxwk-go-sdk/model"
	errors "github.com/wuntsong-org/wterrors"
	"net/url"
)

type LiveCodeFile struct {
	c *client.Client
}

func NewLiveCodeFile(c *client.Client) *LiveCodeFile {
	return &LiveCodeFile{
		c: c,
	}
}

func (s *LiveCodeFile) GetList(opt ...Options) (resp *model.GetLiveCodeListResp, err errors.WTError) {
	val := url.Values{}
	for _, o := range opt {
		o(&val)
	}

	resp = new(model.GetLiveCodeListResp)
	_, err = s.c.SendGetRequests("/api/v1/user/livecode/file/list", val, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LiveCodeFile) GetURL(id string, opt ...Options) (fileURL string, err errors.WTError) {
	val := url.Values{}

	for _, o := range opt {
		o(&val)
	}

	val.Set("id", id)

	resp, err := s.c.SendGetRequests("/api/v1/user/livecode/file", val, nil)
	if err != nil {
		return "", err
	}

	fileURL = resp.Header.Get("Location")
	if len(fileURL) == 0 {
		return "", errors.Errorf("get url fail")
	}

	return fileURL, nil
}

func (s *LiveCodeFile) Create(file []byte, name string, opt ...Options) (resp *model.UploadFileResp, err errors.WTError) {
	fileBase64 := base64.StdEncoding.EncodeToString(file)
	req := model.UpdateFileReq{
		File: "base64:" + fileBase64,
		Name: name,
	}

	for _, o := range opt {
		o(&req)
	}

	resp = new(model.UploadFileResp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/livecode/file/update", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LiveCodeFile) UpdateName(id string, name string, opt ...Options) (resp *model.Resp, err errors.WTError) {
	req := model.UpdateFileNameReq{
		Fid:  id,
		Name: name,
	}

	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/livecode/file/name/update", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LiveCodeFile) Delete(id string, opt ...Options) (resp *model.Resp, err errors.WTError) {
	req := model.DeleteID{
		ID: id,
	}

	for _, o := range opt {
		o(&req)
	}

	resp = new(model.Resp)
	_, err = s.c.SendPostRequests(req, "/api/v1/user/livecode/file/delete", resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
