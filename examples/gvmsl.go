package main

import (
	"fmt"
	"github.com/iij/gpapi"
	"github.com/iij/gpapi/protocol"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	api := gpapi.NewAPI(os.Getenv("IIJAPI_ACCESS_KEY"), os.Getenv("IIJAPI_SECRET_KEY"))
	param := protocol.GetVirtualMachineStatusListArg{}
	for _, v := range os.Args[1:] {
		switch {
		case strings.HasPrefix(v, "gp"):
			param.GpServiceCode = v
		case strings.HasPrefix(v, "gc"):
			param.GcServiceCode = append(param.GcServiceCode, v)
		}
	}
	v, err := gpapi.ValidateRequest(param)
	log.Println("val", v, err)
	res, err := api.Post(v)
	log.Println("post", res, err)
	prd, pwr := io.Pipe()
	wr := io.MultiWriter(os.Stdout, pwr)
	go io.Copy(wr, res.Body)
	rsp := protocol.GetVirtualMachineStatusListResponse{}
	err = gpapi.ParseResponse(prd, &rsp)
	fmt.Println("")
	log.Printf("resp: err=%v, rsp=%+v\n", err, rsp)
}
