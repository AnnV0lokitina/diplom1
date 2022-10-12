package handler

import (
	pb "github.com/AnnV0lokitina/diplom1/proto"
	log "github.com/sirupsen/logrus"
	"io"
)

func (h *Handler) RestoreFile(in *pb.RestoreFileRequest, stream pb.SecureStorage_RestoreFileServer) error {
	session := in.GetSession()
	//httpShutdownCh := make(chan struct{})

	r, w := io.Pipe()
	// fileinfo !!!
	buf := make([]byte, 16)
	for {
		log.Println("read start")
		n, err := r.Read(buf)
		log.Printf("n = %v err = %v b = %v\n", n, err, string(buf))
		log.Printf("b[:n] = %q\n", buf[:n])
		if err == io.EOF || n == 0 {
			log.Println("eof")
			break
		}
		//n, err = w.Write(buf[:n])
		resp := &pb.RestoreFileResponse{
			Content: buf[:n],
			//Info: { !!!!
			//	Name: ,
			//	Type: ,
			//	Time:
			//},
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
		//log.Println(n)
		//if err != nil {
		//	return err
		//}
	}

	h.service.RestoreFile(session, fileType, fileName, w, time)

	//resp := &pb.RestoreFileResponse{
	//	Content: []byte("123"),
	//}
	//if err := stream.Send(resp); err != nil {
	//	return err
	//}
	return nil
	//return status.Errorf(codes.Unimplemented, "method SendFile not implemented")
}
