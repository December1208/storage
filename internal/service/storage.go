package service

import (
	"context"
	pb "storage/api/v1"
)

func (s StorageService) UploadFile(c context.Context, req *pb.UploadFileRequest) (*pb.UploadFileReply, error) {

	return &pb.UploadFileReply{}, nil
}
