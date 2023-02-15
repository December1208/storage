package service

import (
	"context"
	"net/url"
	"path/filepath"
	pb "storage/api/v1"
	"storage/config"
	"storage/pkg"
)

func (s StorageService) UploadFile(c context.Context, req *pb.UploadFileRequest) (*pb.UploadFileReply, error) {

	dir := filepath.Join(config.Instance.Server.MediaRoot, req.PathPrefix)
	identity, err := s.storage.Save(c, dir, req.File)
	if err != nil {
		return &pb.UploadFileReply{}, pkg.Errorf(499, "UPLOAD_FILE_ERROR", err.Error())
	}
	domain := c.Value("HOST").(string)
	fileUrl, _ := url.Parse(domain)
	fileUrl.Path = filepath.Join(config.Instance.Server.MediaUrlPrefix, identity)
	result := pb.UploadFileReply{
		Identity: identity,
		Url:      fileUrl.String(),
		Domain:   domain,
		Size:     0,
		Type:     "",
		Name:     req.File.Filename,
	}

	return &result, nil
}
