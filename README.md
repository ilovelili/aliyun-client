# Aliyun Client

Aliyun client designed for easy use. It's used by my own [DongFeng](https://github.com/ilovelili/dongfeng-core) project.

## Dependencies

- [OSS-SDK](https://github.com/aliyun/aliyun-oss-go-sdk) Aliyun OSS SDK

## Object Upload

example:

```Go
import github.com/ilovelili/aliyun-client/oss

func main() {
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
```

## Contact

<route666@live.cn>