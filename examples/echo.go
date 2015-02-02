package main

import (
	"github.com/iij/gpapi"
	"github.com/iij/gpapi/protocol"
	"log"
	"os"
)

func main() {
	api := gpapi.NewAPI(os.Getenv("IIJAPI_ACCESS_KEY"), os.Getenv("IIJAPI_SECRET_KEY"))
	param := protocol.EchoArg{Param: "Hello, World"}
	res := &protocol.EchoResponse{}
	if err := api.Call(param, res); err != nil {
		log.Println("error", err)
	} else {
		log.Printf("res: param=%+v, res=%+v", param, res)
	}
}
