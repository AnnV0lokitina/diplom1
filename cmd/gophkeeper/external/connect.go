package external

import (
	// ...
	pb "github.com/AnnV0lokitina/diplom1/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ExtConnection struct {
}

func NewExtConnection() {
	// устанавливаем соединение с сервером
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// получаем переменную интерфейсного типа UsersClient,
	// через которую будем отправлять сообщения
	c := pb.NewSecureStorageClient(conn)
	log.Println(c)

	// функция, в которой будем отправлять сообщения
	//TestUsers(c)
}
