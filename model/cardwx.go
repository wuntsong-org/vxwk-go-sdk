package model

type CardWX struct {
	ID        string `json:"id"` // 短链ID
	Type      int64  `json:"type"`
	Title     string `json:"title"` // 短链title
	ImgKey    string `json:"imgKey"`
	Describe  string `json:"describe"`
	TargetUrl string `json:"targetUrl"`
	Url       string `json:"url"`
	ShareUrl  string `json:"shareUrl"`
	StopAt    int64  `json:"stopAt"` // 短链结束时间 (可为NULL)
	StartAt   int64  `json:"startAt"`
	Status    int64  `json:"status"`
	TestMode  bool   `json:"testMode"`
	Style     string `json:"style"`
}

type GetCardWXListData struct {
	Card  []CardWX `json:"card"`
	Count int64    `json:"count"`
}

type GetCardWXListResp struct {
	Resp
	Data GetCardWXListData `json:"data"`
}

type GetCardWXData struct {
	Card CardWX `json:"card"`
}

type GetCardWXResp struct {
	Resp
	Data GetCardWXData `json:"data"`
}

type CreateCardWXReq struct {
	Describe  string `json:"describe,optional"`
	Img       string `json:"img" c-ignore:"true"`
	TargetURL string `json:"targetUrl" cs-max:"200" cs-checker:"HttpString"`
	Title     string `json:"title" cs-max:"20"` // 短链title
	Type      int64  `json:"type" ci-checker:"ShortLinkType"`
	StartAt   int64  `json:"startAt,optional" ci-zero:"ignore" ci-min:"0"`
	StopAt    int64  `json:"stopAt,optional" ci-zero:"ignore" ci-min"0"`
	TestMode  bool   `json:"testMode"`
	Style     string `json:"style" cs-checker:"Style"`
}

type UpdateCardWXReq struct {
	ID        string `json:"id" cs-checker:"UUIDString"` // 短链ID
	DeleteImg bool   `json:"deleteImg"`
	Describe  string `json:"describe,optional"`
	Img       string `json:"img" c-ignore:"true"`
	TargetURL string `json:"targetUrl" cs-max:"200" cs-checker:"HttpString"`
	Title     string `json:"title" cs-max:"20"` // 短链title
	Type      int64  `json:"type" ci-checker:"ShortLinkType"`
	StartAt   int64  `json:"startAt,optional" ci-zero:"ignore" ci-min:"0"`
	StopAt    int64  `json:"stopAt,optional" ci-zero:"ignore" ci-min"0"`
	TestMode  bool   `json:"testMode"`
	Style     string `json:"style" cs-checker:"Style"`
}
