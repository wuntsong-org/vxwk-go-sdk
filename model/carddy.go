package model

type CardDY struct {
	ID           string `json:"id"`                 // 短链ID
	DomainID     string `json:"domainID,omitempty"` // 短链域名ID
	Domain       string `json:"domain"`
	DomainSuffix string `json:"domainSuffix"`
	Hash         string `json:"hash"` // 短链哈希
	Type         int64  `json:"type"`
	Title        string `json:"title"` // 短链title
	ImgKey       string `json:"imgKey"`
	TargetUrl    string `json:"targetUrl"`
	Url          string `json:"url"`
	StopAt       int64  `json:"stopAt"` // 短链结束时间 (可为NULL)
	StartAt      int64  `json:"startAt"`
	Status       int64  `json:"status"`
	TestMode     bool   `json:"testMode"`
	Style        string `json:"style"`
}

type GetCardDYListData struct {
	Card  []CardDY `json:"card"`
	Count int64    `json:"count"`
}

type GetCardDYListResp struct {
	Resp
	Data GetCardDYListData `json:"data"`
}

type GetCardDYData struct {
	Card CardDY `json:"card"`
}

type GetCardDYResp struct {
	Resp
	Data GetCardDYData `json:"data"`
}

type CreateCardDYReq struct {
	DomainID  string `json:"domainID" cs-checker:"UUIDString"` // 短链域名ID
	Img       string `json:"img" c-ignore:"true"`
	TargetURL string `json:"targetUrl" cs-max:"200" cs-checker:"HttpString"`
	Title     string `json:"title" cs-max:"20"` // 短链title
	Type      int64  `json:"type" ci-checker:"ShortLinkType"`
	StartAt   int64  `json:"startAt,optional" ci-zero:"ignore" ci-min:"0"`
	StopAt    int64  `json:"stopAt,optional" ci-zero:"ignore" ci-min:"0"`
	Hash      string `json:"hash,optional" cs-max:"100"`
	TestMode  bool   `json:"testMode"`
	Style     string `json:"style" cs-checker:"Style"`
}

type UpdateCardDYReq struct {
	ID        string `json:"id" cs-checker:"UUIDString"` // 短链ID
	Title     string `json:"title" cs-max:"20"`          // 短链title
	Img       string `json:"img" c-ignore:"true"`
	TargetURL string `json:"targetUrl" cs-max:"200" cs-checker:"HttpString"`
	Type      int64  `json:"type"`
	StartAt   int64  `json:"startAt,optional" ci-zero:"ignore" ci-min:"0"`
	StopAt    int64  `json:"stopAt,optional" ci-zero:"ignore" ci-min:"0"`
	TestMode  bool   `json:"testMode"`
	Style     string `json:"style" cs-checker:"Style"`
	DeleteImg bool   `json:"deleteImg"`
}
