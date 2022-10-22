package grpctest

import (
	pb "github.com/AnnV0lokitina/diplom1/proto"
	"time"
)

const ExistingUser = "ExistingUser"
const NewUser = "NewUser"
const WithErrorSession = "WithErrorSession"
const CorrectSession = "CorrectSession"

var ChunkSize = 16
var TestFileContent string
var TestFileDate time.Time

type Handler struct {
	pb.UnimplementedSecureStorageServer
}
