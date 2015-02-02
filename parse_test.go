package gpapi

import (
	"bytes"
	protocol "github.com/iij/gpapi/protocol"
	"testing"
)

func TestValidateEcho(t *testing.T) {
	ea := protocol.EchoArg{Param: "test"}
	res, err := ValidateRequest(ea)
	if err != nil || res.Get("Param") != "test" {
		t.Errorf("res:%+v, err=%s", res, err)
	}
	t.Log(res, err)
}

func TestValidateSetLabel(t *testing.T) {
	sl := protocol.SetLabelArg{}
	res, err := ValidateRequest(sl)
	if err == nil {
		t.Errorf("res:%+v, err=%s", res, err)
	}
	t.Log(res, err)
	sl = protocol.SetLabelArg{GpServiceCode: "gpcode", Label: "testlabel", ServiceCode: "svcode"}
	res, err = ValidateRequest(sl)
	if err != nil {
		t.Errorf("res:%+v, err=%s", res, err)
	}
}

func TestEncodeEcho(t *testing.T) {
	echostr := "Hello, World"
	sl := protocol.EchoArg{Param: echostr}
	res, err := ValidateRequest(sl)
	t.Log("validate", res.Encode(), err)
	a := NewAPI("testkey", "testsecret")
	res = a.FixParam(res)
	if res.Get("APIVersion") != APIVersion20130901 {
		t.Error("invalid api version", res)
	}
	if res.Get("SignatureMethod") != HmacSHA256 {
		t.Error("invalid signature method", res)
	}
	if res.Get("SignatureVersion") != SignatureVersion2 {
		t.Error("invalid signature version", res)
	}
	if res.Get("Param") != echostr {
		t.Error("invalid param", res)
	}
	if res.Get("AccessKeyId") != "testkey" {
		t.Error("invalid access key", res)
	}
	if res.Get("Action") != "Echo" {
		t.Error("invalid action", res)
	}
	t.Log("fix", res.Encode())
	res = a.AddSignature("POST", res)
	t.Log("post", res.Encode())
}

func TestDecodeEcho(t *testing.T) {
	s := `{ "RequestId": "testid", "Param": "Hello, World" }`
	echores := new(protocol.EchoResponse)
	ParseResponse(bytes.NewBufferString(s), echores)
	if echores.RequestId != "testid" {
		t.Error("invalid reqid", echores)
	}
	if echores.Param != "Hello, World" {
		t.Error("invalid param", echores)
	}
	t.Logf("echores: %+v\n", echores)
}

func TestDecodeGSCL(t *testing.T) {
	s := `{"GetServiceCodeListResponse":{"RequestId":"testid","CustomerCode":"testcc","GpServiceCode":"gp-test","GcServiceCodeList":["gc-test"],"GnbServiceCodeList":[],"GxServiceCodeList":[],"GlServiceCodeList":[],"GvmServiceCodeList":[],"GvsServiceCodeList":[],"GomServiceCodeList":["gom-test"]}}`
	gsclres := new(protocol.GetServiceCodeListResponse)
	ParseResponse(bytes.NewBufferString(s), gsclres)
	if gsclres.RequestId != "testid" {
		t.Error("invalid reqid", gsclres)
	}
	if gsclres.CustomerCode != "testcc" {
		t.Error("invalid customercode", gsclres)
	}
	if gsclres.GpServiceCode != "gp-test" {
		t.Error("invalid gp servicecode", gsclres)
	}
	if len(gsclres.GcServiceCodeList) != 1 || gsclres.GcServiceCodeList[0] != "gc-test" {
		t.Error("invalid gc servicecode", gsclres)
	}
}
