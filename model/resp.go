package model

type Resp struct {
	Code       string `json:"code"`
	SubCode    string `json:"subCode"`
	NumCode    int64  `json:"_code"`
	NumSubCode int64  `json:"_subCode"`
	Msg        string `json:"msg,omitempty"`
}

func (r Resp) GetCode() string {
	return r.Code
}

func (r Resp) GetSubCode() string {
	return r.SubCode
}

func (r Resp) GetMsg() string {
	return r.Msg
}
