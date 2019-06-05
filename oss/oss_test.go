package oss

import (
	"testing"
)

// TestUpload test object upload
func TestUpload(t *testing.T) {
	oss := NewService("<<api key>>", "<<secret key>>")
	oss.SetEndPoint("oss-cn-shanghai.aliyuncs.com")
	oss.SetBucket("dong-feng")

	opts1 := &UploadOptions{
		ObjectName: "test",
		Public:     true,
		IsFolder:   true,
	}

	resp := oss.Upload(opts1)
	if resp.Error != nil {
		t.Error(resp.Error)
	}

	opts2 := &UploadOptions{
		ObjectName:   "../test/index.html",
		Public:       true,
		ParentFolder: "test",
	}

	resp = oss.Upload(opts2)
	if resp.Error != nil {
		t.Error(resp.Error)
	}
}
