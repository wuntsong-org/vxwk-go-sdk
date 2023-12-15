package model

type External struct {
	ID           string `json:"id"`                 // 短链ID
	DomainID     string `json:"domainID,omitempty"` // 短链域名ID
	Domain       string `json:"domain"`
	DomainSuffix string `json:"domainSuffix"`
	Hash         string `json:"hash"` // 短链哈希
	Type         int64  `json:"type"`
	Title        string `json:"title"` // 短链title
	Describe     string `json:"describe"`
	Tips         string `json:"tips"`
	ImgKey       string `json:"imgKey"`
	Url          string `json:"url"`
	StopAt       int64  `json:"stopAt"` // 短链结束时间 (可为NULL)
	StartAt      int64  `json:"startAt"`
	Status       int64  `json:"status"`
	TestMode     bool   `json:"testMode"`
}

type GetExternalListData struct {
	External []External `json:"external"`
	Count    int64      `json:"count"`
}

type GetExternalListResp struct {
	Resp
	Data GetExternalListData `json:"data"`
}

type ExternalWithLink struct {
	External
	Link  []Link `json:"link"`
	Style string `json:"style"`
}

type GetExternalData struct {
	External ExternalWithLink `json:"external"`
}

type GetExternalResp struct {
	Resp
	Data GetExternalData `json:"data"`
}

type CreateExternalReq struct {
	DomainID string `json:"domainID" cs-checker:"UUIDString"` // 短链域名ID
	Title    string `json:"title" cs-max:"20"`                // 短链title
	Describe string `json:"describe" cs-max:"100"`
	Tips     string `json:"tips,optional" cs-max:"100"`
	Img      string `json:"img,optional" c-ignore:"true"`
	Type     int64  `json:"type" ci-checker:"ShortLinkType"`
	StartAt  int64  `json:"startAt,optional" ci-zero:"ignore" ci-min:"0"`
	StopAt   int64  `json:"stopAt,optional" ci-zero:"ignore" ci-min"0"`
	Link     []Link `json:"link" c-ignore:"true"`
	Hash     string `json:"hash,optional" cs-max:"100"`
	TestMode bool   `json:"testMode"`
	Style    string `json:"style" cs-checker:"Style"`
}

type UpdateExternalReq struct {
	ID        string `json:"id" cs-checker:"UUIDString"` // 短链ID
	DeleteImg bool   `json:"deleteImg"`
	Title     string `json:"title" cs-max:"20"` // 短链title
	Describe  string `json:"describe" cs-max:"100"`
	Tips      string `json:"tips,optional" cs-max:"100"`
	Img       string `json:"img,optional" c-ignore:"true"`
	StartAt   int64  `json:"startAt,optional" ci-zero:"ignore" ci-min:"0"`
	StopAt    int64  `json:"stopAt,optional" ci-zero:"ignore" ci-min"0"`
	Link      []Link `json:"link" c-ignore:"true"`
	TestMode  bool   `json:"testMode"`
	Style     string `json:"style" cs-checker:"Style"`
}
