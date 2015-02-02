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
	param := protocol.GetContractInformationArg{GpServiceCode: os.Args[1]}
	if len(os.Args) > 2 {
		scd := os.Args[2]
		switch {
		case strings.HasPrefix(scd, "gc"):
			param.GcServiceCode = scd
		case strings.HasPrefix(scd, "gl"):
			param.GlServiceCode = scd
		case strings.HasPrefix(scd, "gx"):
			param.GxServiceCode = scd
		}
	}
	v, err := gpapi.ValidateRequest(param)
	log.Println("val", v, err)
	res, err := api.Post(v)
	log.Println("post", res, err)
	prd, pwr := io.Pipe()
	wr := io.MultiWriter(os.Stdout, pwr)
	go io.Copy(wr, res.Body)
	rsp := protocol.GetContractInformationResponse{}
	err = gpapi.ParseResponse(prd, &rsp)
	fmt.Println("")
	log.Printf("resp: err=%v, rsp=%+v\n", err, rsp)
}
