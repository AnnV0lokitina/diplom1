package handler

import (
	pb "github.com/AnnV0lokitina/diplom1/proto"
)

// Handler structure holds dependencies for server handlers.
type Handler struct {
	pb.UnimplementedSecureStorageServer

	service Service
}

// NewHandler Create new handler struct.
func NewHandler(service Service) *Handler {
	h := &Handler{}
	h.service = service
	return h
}

//func getUserID(id uint32) (uint32, error) {
//	if id > 0 {
//		return id, nil
//	}
//	return userid.GenerateUserID()
//}
