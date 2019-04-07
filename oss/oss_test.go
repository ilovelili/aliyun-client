package oss

import (
	"testing"
)

// TestUpload test object upload
func TestUpload(t *testing.T) {
	oss := NewService("<<api key>>", "<<api secret>>")
	oss.SetEndPoint("oss-cn-shanghai.aliyuncs.com")
	oss.SetBucket("dong-feng")

	opts := &UploadOptions{
		FileName: "./test.png",
		Public:   true,
	}

	resp := oss.Upload(opts)
	if resp.Error != nil {
		t.Error(resp.Error)
	}
}
