package main

import (
	"time"
)

type Config struct {
	Endpoint string `json:"endpoint"`
	Username string `json:"username"`
	Password string `json:"password"`
	SrcDir   string `json:"src_dir"`
	DstDri   string `json:"dst_dir"`
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type AuthLoginResponse struct {
	BaseResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}
type FSListResponse struct {
	BaseResponse
	Data struct {
		Content  []*FSListContentItem `json:"content"`
		Total    int                  `json:"total"`
		Readme   string               `json:"readme"`
		Header   string               `json:"header"`
		Write    bool                 `json:"write"`
		Provider string               `json:"provider"`
	} `json:"data"`
}

type FSListContentItem struct {
	Name     string    `json:"name"`
	Size     int       `json:"size"`
	IsDir    bool      `json:"is_dir"`
	Modified time.Time `json:"modified"`
	Created  time.Time `json:"created"`
	Sign     string    `json:"sign"`
	Type     int       `json:"type"`
}
