package gateway

import (
	"google.golang.org/grpc"
)

type Implementation struct {
	yandexFoodConn *grpc.ClientConn
}

func NewGateway(yandexFoodConn *grpc.ClientConn) *Implementation {
	return &Implementation{
		yandexFoodConn: yandexFoodConn,
	}
}
