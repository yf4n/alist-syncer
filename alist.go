package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AlistClient struct {
	url      string
	username string
	password string
	token    string
	ctx      context.Context
}

func NewAlistClient(ctx context.Context, url, username, password string) *AlistClient {
	return &AlistClient{
		username: username,
		password: password,
		url:      url,
		ctx:      ctx,
	}
}

func (ac *AlistClient) generateRequest(method, api string, body io.Reader) (*http.Request, error) {
	requrl, err := url.JoinPath(ac.url, api)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ac.ctx, method, requrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", ac.token)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (ac *AlistClient) Login() error {
	// 创建请求体
	reqBody := strings.NewReader("{\"username\":\"" + ac.username + "\",\"password\":\"" + ac.password + "\"}")
	// 发送请求
	req, err := ac.generateRequest("POST", "/api/auth/login", reqBody)
	if err != nil {
		return fmt.Errorf("创建登录请求失败: %v", err)
	}

	// 发送请求
	cli := &http.Client{Timeout: 10 * time.Second}
	resp, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("发送登录请求失败: %v", err)
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取登录响应失败: %v", err)
	}
	// 解析响应
	var loginResp AuthLoginResponse
	if err := json.Unmarshal(body, &loginResp); err != nil {
		return fmt.Errorf("解析登录响应失败: %v", err)
	}
	if loginResp.Code != 200 {
		return fmt.Errorf("登录失败: %s", loginResp.Message)
	}

	// 存储 token
	ac.token = loginResp.Data.Token
	return nil
}

func (ac *AlistClient) FSList(path string) (*FSListResponse, error) {
	// 创建请求体
	reqBody := strings.NewReader(fmt.Sprintf(`{"path":"%s","page":%d,"per_page":%d,"refresh":true}`, path, 1, 0))

	// 发送请求
	req, err := ac.generateRequest("POST", "/api/fs/list", reqBody)
	if err != nil {
		return nil, err
	}

	// 发送请求
	cli := &http.Client{Timeout: 10 * time.Second}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 解析响应
	var fsListResp FSListResponse
	if err := json.Unmarshal(body, &fsListResp); err != nil {
		return nil, err
	}
	if fsListResp.Code != 200 {
		return nil, fmt.Errorf("response error: %s", fsListResp.Message)
	}

	return &fsListResp, nil
}

func (ac *AlistClient) FSRemove(dir string, names []string) error {
	data := map[string]interface{}{
		"names": names,
		"dir":   dir,
	}
	bs, _ := json.Marshal(data)
	// 创建请求体
	reqBody := strings.NewReader(string(bs))

	// 发送请求
	req, err := ac.generateRequest("POST", "/api/fs/remove", reqBody)
	if err != nil {
		return fmt.Errorf("创建删除文件请求失败: %v", err)
	}

	// 发送请求
	cli := &http.Client{Timeout: 10 * time.Second}
	resp, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("发送删除文件请求失败: %v", err)
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取删除文件响应失败: %v", err)
	}
	// 解析响应
	var baseResp BaseResponse
	if err := json.Unmarshal(body, &baseResp); err != nil {
		return fmt.Errorf("解析删除文件响应失败: %v", err)
	}
	if baseResp.Code != 200 {
		return fmt.Errorf("删除文件失败: %s", baseResp.Message)
	}

	return nil
}

func (ac *AlistClient) FSCopy(src, dst string, names []string) error {
	data := map[string]interface{}{
		"src_dir": src,
		"dst_dir": dst,
		"names":   names,
	}
	bs, _ := json.Marshal(data)
	// 创建请求体
	reqBody := strings.NewReader(string(bs))

	// 发送请求
	req, err := ac.generateRequest("POST", "/api/fs/copy", reqBody)
	if err != nil {
		return fmt.Errorf("创建复制文件请求失败: %v", err)
	}

	// 发送请求
	cli := &http.Client{Timeout: 10 * time.Second}
	resp, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("发送复制文件请求失败: %v", err)
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取复制文件响应失败: %v", err)
	}
	// 解析响应
	var baseResp BaseResponse
	if err := json.Unmarshal(body, &baseResp); err != nil {
		return fmt.Errorf("解析复制文件响应失败: %v", err)
	}
	if baseResp.Code != 200 {
		return fmt.Errorf("复制文件失败: %s", baseResp.Message)
	}

	return nil
}

func (ac *AlistClient) FSMkdir(path string) error {
	// 创建请求体
	reqBody := strings.NewReader(fmt.Sprintf(`{"path":"%s"}`, path))

	// 发送请求
	req, err := ac.generateRequest("POST", "/api/fs/mkdir", reqBody)
	if err != nil {
		return err
	}

	// 发送请求
	cli := &http.Client{Timeout: 10 * time.Second}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// 解析响应
	var fsListResp BaseResponse
	if err := json.Unmarshal(body, &fsListResp); err != nil {
		return err
	}
	if fsListResp.Code != 200 {
		return fmt.Errorf("response error: %s", fsListResp.Message)
	}

	return nil
}
