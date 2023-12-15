package client

import (
	"bytes"
	"fmt"
	"github.com/wuntsong-org/vxwk-go-sdk/utils"
	errors "github.com/wuntsong-org/wterrors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type RespInterface interface {
	GetCode() string
	GetSubCode() string
	GetMsg() string
}

type Config struct {
	AccessKey    string
	AccessSecret string
	Endpoint     string
}

func NewClient(c Config) *Client {
	return &Client{
		accessKey:    c.AccessKey,
		accessSecret: c.AccessSecret,
		endpoint:     c.Endpoint,
	}
}

type Client struct {
	accessKey    string
	accessSecret string
	endpoint     string
}

func (c *Client) SendGetRequests(path string, query url.Values, r RespInterface) (*http.Response, errors.WTError) {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	n := utils.GenerateUniqueNumber(18)

	query.Add("xaccesskey", c.accessKey)
	query.Add("xn", n)
	query.Add("xtimestamp", timestamp)
	query.Add("xrunmode", "release")

	signText := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", http.MethodGet, c.accessKey, timestamp, n, path, query.Encode())
	sign := utils.StringCalculateHMAC(c.accessSecret, signText)

	query.Add("xsign", sign)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?%s", c.endpoint, path, query.Encode()), nil)
	if err != nil {
		return nil, errors.WarpQuick(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WarpQuick(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WarpQuick(err)
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("get bad status code: %s", body)
	}

	if r == nil {
		return resp, nil
	}

	err = utils.JsonUnmarshal(body, r)
	if err != nil {
		return nil, errors.WarpQuick(err)
	}

	code := r.GetCode()
	if code == "SUCCESS" {
		return resp, nil
	} else if code == "LOGIC_DENY" {
		return nil, errors.Errorf("get logic resp: %s", r.GetMsg()).SetCode(fmt.Sprintf(r.GetSubCode()))
	}

	return nil, errors.Errorf("requests fail: %s", r.GetMsg()).SetCode(fmt.Sprintf(r.GetCode()))
}

func (c *Client) SendPostRequests(data any, path string, r RespInterface) (*http.Response, errors.WTError) {
	dataByte, jsonErr := utils.JsonMarshal(data)
	if jsonErr != nil {
		return nil, jsonErr
	}

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	n := utils.GenerateUniqueNumber(18)

	signText := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", http.MethodPost, c.accessKey, timestamp, n, path, string(dataByte))
	sign := utils.StringCalculateHMAC(c.accessSecret, signText)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.endpoint, path), bytes.NewBuffer(dataByte))
	if err != nil {
		return nil, errors.WarpQuick(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-AccessKey", c.accessKey)
	req.Header.Set("X-Timestamp", timestamp)
	req.Header.Set("X-N", n)
	req.Header.Set("X-Sign", sign)
	req.Header.Set("X-RunMode", "release")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WarpQuick(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WarpQuick(err)
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("get bad status code: %s", body)
	}

	if r == nil {
		return resp, nil
	}

	err = utils.JsonUnmarshal(body, r)
	if err != nil {
		return nil, errors.WarpQuick(err)
	}

	code := r.GetCode()
	if code == "SUCCESS" {
		return resp, nil
	} else if code == "LOGIC_DENY" {
		return nil, errors.Errorf("get logic resp: %s", r.GetMsg()).SetCode(fmt.Sprintf(r.GetSubCode()))
	}

	return nil, errors.Errorf("requests fail: %s", r.GetMsg()).SetCode(fmt.Sprintf(r.GetCode()))
}
