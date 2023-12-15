package model

type LiveCode struct {
	ID           string `json:"id"`                 // 短链ID
	DomainID     string `json:"domainID,omitempty"` // 短链域名ID
	Domain       string `json:"domain"`
	DomainSuffix string `json:"domainSuffix"`
	Hash         string `json:"hash"` // 短链哈希
	Type         int64  `json:"type"`
	Title        string `json:"title"` // 短链title
	Url          string `json:"url"`
	StopAt       int64  `json:"stopAt"` // 短链结束时间 (可为NULL)
	StartAt      int64  `json:"startAt"`
	Status       int64  `json:"status"`
	TestMode     bool   `json:"testMode"`
}

type GetLiveCodeListData struct {
	LiveCode []LiveCode `json:"livecode"`
	Count    int64      `json:"count"`
}

type GetLiveCodeListResp struct {
	Resp
	Data GetLiveCodeListData `json:"data"`
}

type LiveCodeWithLink struct {
	LiveCode
	Link  []Link `json:"link"`
	Style string `json:"style"`
}

type GetLiveCodeData struct {
	LiveCode LiveCodeWithLink `json:"livecode"`
}

type GetLiveCodeResp struct {
	Resp
	Data GetLiveCodeData `json:"data"`
}

type CreateLiveCodeReq struct {
	DomainID string `json:"domainID" cs-checker:"UUIDString"` // 短链域名ID
	Title    string `json:"title" cs-max:"20"`                // 短链title
	Type     int64  `json:"type" ci-checker:"ShortLinkType"`
	StartAt  int64  `json:"startAt,optional" ci-zero:"ignore" ci-min:"0"`
	StopAt   int64  `json:"stopAt,optional" ci-zero:"ignore" ci-min"0"`
	Link     []Link `json:"link" c-ignore:"true"`
	Hash     string `json:"hash,optional" cs-max:"100"`
	TestMode bool   `json:"testMode"`
	Style    string `json:"style" cs-checker:"Style"`
}

type UpdateLiveCodeReq struct {
	ID       string `json:"id" cs-checker:"UUIDString"` // 短链ID
	Title    string `json:"title" cs-max:"20"`          // 短链title
	StartAt  int64  `json:"startAt,optional" ci-zero:"ignore" ci-min:"0"`
	StopAt   int64  `json:"stopAt,optional" ci-zero:"ignore" ci-min"0"`
	Link     []Link `json:"link" c-ignore:"true"`
	TestMode bool   `json:"testMode"`
	Style    string `json:"style" cs-checker:"Style"`
}
