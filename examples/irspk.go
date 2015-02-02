package main

import (
	"fmt"
	"github.com/iij/gpapi"
	"github.com/iij/gpapi/protocol"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	api := gpapi.NewAPI(os.Getenv("IIJAPI_ACCESS_KEY"), os.Getenv("IIJAPI_SECRET_KEY"))
	gpsc := os.Args[1]
	gcsc := os.Args[2]
	pubkey := os.Args[3]
	log.Println("gp", gpsc, "gc", gcsc, "key", pubkey)
	if !(strings.HasPrefix(gpsc, "gp") && strings.HasPrefix(gcsc, "gc")) {
		log.Fatal("invalid service code", gpsc, gcsc)
	}
	pubkeyb, err := ioutil.ReadFile(pubkey)
	if err != nil {
		log.Fatal("invalid pubkey", pubkey)
	}
	pubkeys := string(pubkeyb)
	param := protocol.ImportRootSshPublicKeyArg{GpServiceCode: gpsc, GcServiceCode: gcsc, PublicKey: pubkeys}
	v, err := gpapi.ValidateRequest(param)
	log.Println("val", v, err)
	res, err := api.Post(v)
	log.Println("post", res, err)
	prd, pwr := io.Pipe()
	wr := io.MultiWriter(os.Stdout, pwr)
	go io.Copy(wr, res.Body)
	rsp := protocol.ImportRootSshPublicKeyResponse{}
	err = gpapi.ParseResponse(prd, &rsp)
	fmt.Println("")
	log.Printf("resp: err=%v, rsp=%+v\n", err, rsp)
}
