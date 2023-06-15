package helper

import (
	"alta/air-bnb/app/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func InitMidtrans() *coreapi.Client {
	appConfig := config.ReadEnv()
	client := coreapi.Client{}
	client.New(appConfig.MIDTRANS_SERVER_KEY, midtrans.Sandbox)
	return &client
} 