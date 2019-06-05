package oss

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"

	sdk "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// UploadOptions upload response
type UploadOptions struct {
	ObjectName   string
	Public       bool
	Expire       time.Duration
	IsFolder     bool   // Is file or folder
	ParentFolder string // Parent folder
	Meta         map[string]string
}

// UploadResponse upload response
type UploadResponse struct {
	// Location location of uploaded file
	Location string
	Error    error
}

// Context context includes endPoint and bucket info
type Context struct {
	EndPoint string
	Bucket   string
}

// Service service defines context
type Service struct {
	Context      *Context
	AccessKey    string
	AccessSecret string
}

// NewService service initializer
func NewService(key, secret string) *Service {
	return &Service{
		Context:      new(Context),
		AccessKey:    key,
		AccessSecret: secret,
	}
}

// SetEndPoint set endpoint
func (s *Service) SetEndPoint(endpoint string) {
	s.Context.check()
	s.Context.EndPoint = endpoint
}

// GetEndPoint get endpoint
func (s *Service) GetEndPoint() string {
	return s.Context.EndPoint
}

// SetBucket set bucket
func (s *Service) SetBucket(bucket string) {
	s.Context.check()
	s.Context.Bucket = bucket
}

// GetBucket get bucket
func (s *Service) GetBucket() string {
	return s.Context.Bucket
}

// GetLocation get location
func (s *Service) GetLocation(objname string) string {
	return fmt.Sprintf("https://%s.%s/%s", s.GetBucket(), s.GetEndPoint(), objname)
}

// Upload upload file
func (s *Service) Upload(opts *UploadOptions) (resp *UploadResponse) {
	resp = new(UploadResponse)
	client, err := sdk.New(s.GetEndPoint(), s.AccessKey, s.AccessSecret)

	bucket, err := client.Bucket(s.GetBucket())
	if err != nil {
		resp.Error = err
		return
	}

	options := []sdk.Option{}
	if opts.Public {
		options = append(options, sdk.ObjectACL(sdk.ACLPublicRead))
	}

	for k, v := range opts.Meta {
		options = append(options, sdk.Meta(k, v))
	}

	filename := opts.ObjectName
	objname := resolveObjName(filename, opts.ParentFolder)

	if opts.IsFolder {
		err = bucket.PutObject(objname+"/", strings.NewReader(""), options...)
	} else {
		err = bucket.PutObjectFromFile(objname, filename, options...)
	}

	resp.Error = err
	if err != nil {
		return
	}

	resp.Location = s.GetLocation(objname)
	return
}

// AsyncUpload async upload
func (s *Service) AsyncUpload(opts *UploadOptions) (respchan chan *UploadResponse) {
	respchan = make(chan *UploadResponse)
	go func() {
		respchan <- s.Upload(opts)
	}()
	return
}

func resolveObjName(fullfilename, parent string) string {
	return path.Join(parent, filepath.Base(fullfilename))
}

func (ctx *Context) check() {
	if ctx == nil {
		log.Fatal("invalid context")
	}
}
