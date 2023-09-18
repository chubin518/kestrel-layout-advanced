package utils

import (
	"errors"
	"io"
	"net"
	"net/http"
	"os"
)

// 下载进度回调函数
type ProcessCallback func(percent int64)

// download 文件下载
func Download(url, dst string, callback ProcessCallback) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer file.Close()
	// 使用io.TeeReader将数据同时写入文件和进度报告回调函数
	reader := io.TeeReader(resp.Body, buildProcessHandler(resp.ContentLength, callback))
	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}
	return nil
}

// IsNetworkError 判断是否为网络错误
func IsNetworkError(err error) bool {
	// 根据具体情况判断是否为网络错误，这可以根据err的类型或内容来定制
	// 以下是一个示例，你可以根据实际情况进行扩展
	_, isNetErr := err.(net.Error)
	return isNetErr
}

type processHandler struct {
	total    int64
	current  int64
	callback ProcessCallback
	set      map[int64]bool
}

func buildProcessHandler(total int64, callback ProcessCallback) *processHandler {
	return &processHandler{
		current:  0,
		set:      make(map[int64]bool),
		total:    total,
		callback: callback,
	}
}

func (h *processHandler) Write(buf []byte) (int, error) {
	n := len(buf)
	h.current += int64(n)
	if h.callback != nil {
		percent := int64(float64(h.current) / float64(h.total) * 100)
		if _, ok := h.set[percent]; !ok {
			h.callback(percent)
		}
		h.set[percent] = true
	}
	return n, nil
}
