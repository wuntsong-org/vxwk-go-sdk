package model

type LiveCodeFile struct {
	Fid       string `json:"fid"`
	Name      string `json:"name"`
	MediaType string `json:"mediaType"`
	Status    int64  `json:"status"`
}

type LCGetFileListData struct {
	Count int64          `json:"count"`
	File  []LiveCodeFile `json:"file"`
}

type LCGetFileListResp struct {
	Resp
	Data LCGetFileListData `json:"data"`
}

type UploadFileData struct {
	File string `json:"file"`
}

type UploadFileResp struct {
	Resp
	Data UploadFileData `json:"data"`
}

type UpdateFileReq struct {
	File string `json:"file" c-ignore:"true"`
	Name string `json:"name" cs-max:"50"`
}

type UpdateFileNameReq struct {
	Fid  string `json:"fid" cs-max:"100"`
	Name string `json:"name" cs-max:"50"`
}
