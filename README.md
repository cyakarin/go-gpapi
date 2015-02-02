# Golang binding for GPAPI

# Usage

## get your gpapi access key

- open https://help.api.iij.jp and login with IIJ SA Code
- click `+ アクセスキーを追加する` button
- `アクセスキー` and `シークレットキー` are your access key id / secret key

## install gpapi command

- go get github.com/iij/go-gpapi/cli
- go build -o gpapi github.com/iij/go-gpapi/cli
- install -c gpapi /usr/bin/gpapi

## gpapi command

```
# gpapi
NAME:
   gpapi - IIJ GP API Client CLI

USAGE:
   gpapi [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   AddCloneVirtualMachines, ACVM
   AddDiskOption, ADO
   AddFwLbOption, AFLO
   AddGlobalAddressOption, AGAO
   AddLbNode, ALN
   AddLbPool, ALP
   AddLbVirtualServer, ALVS
   AddVirtualIpAddress, AVIA
   AddVirtualMachines, AVM
   AddVlanOption, AVO
   AttachFwLb, AFL
   AttachVlan, AV
   ChangeFwLbOptionType, CFLOT
   ChangeVirtualMachineType, CVMT
   CloneVirtualMachine, CVM
   DeleteDiskOption, DDO
   DeleteFwLbOption, DFLO
   DeleteGlobalAddressOption, DGAO
   DeleteLbNode, DLN
   DeleteLbPool, DLP
   DeleteLbVirtualServer, DLVS
   DeleteVirtualIpAddress, DVIA
   DeleteVirtualMachine, DVM
   DeleteVlanOption, DVO
   DescribeFw, DF
   DescribeLb, DL
   DescribeVirtualMachine, DVM2
   DetachFwLb, DFL
   DetachVlan, DV
   DisableIPv6, DIP
   Echo, E
   EnableIPv6, EIP
   GetContractInformation, GCI
   GetContractStatus, GCS
   GetDomainNameByIpAddress, GDNBIA
   GetFilteringRuleList, GFRL
   GetGpServiceCodeList, GGSCL
   GetLbHostList, GLHL
   GetLbPool, GLP
   GetLbStatus, GLS
   GetLbVirtualServer, GLVS
   GetServiceCodeList, GSCL
   GetSnatRuleList, GSRL
   GetVirtualMachineStatus, GVMS
   GetVirtualMachineStatusList, GVMSL
   ImportRootSshPublicKey, IRSPK
   InitializeVirtualMachine, IVM
   RebootVirtualMachine, RVM
   ResetDomainNameByIpAddress, RDNBIA
   ResetIptables, RI
   ResetNetworkConfiguration, RNC
   ResetSshdConfiguration, RSC
   SetDomainNameByIpAddress, SDNBIA
   SetFilteringRuleList, SFRL
   SetLabel, SL
   SetLbNode, SLN
   SetLbPool, SLP
   SetLbPoolName, SLPN
   SetLbVirtualServer, SLVS
   SetLbVirtualServerName, SLVSN
   SetSnatRuleList, SSRL
   SetVirtualIpAddressName, SVIAN
   StartVirtualMachine, SVM
   StopVirtualMachine, SVM2
   watch, w
   help, h				Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --AccessKeyId 	Access Key [$IIJAPI_ACCESS_KEY]
   --SecretKey 		Secret Key [$IIJAPI_SECRET_KEY]
   --help, -h		show help
   --version, -v	print the version
```

- export IIJAPI_ACCESS_KEY=`your api key id`
- export IIJAPI_SECRET_KEY=`your api secret key`
- gpapi GetGpServiceCodeList | jq .

```js
{
  "GetGpServiceCodeListResponse": {
    "RequestId": "<uuid>",
    "GpServiceCodeList": [
      "gp<your gpcode>"
    ]
  }
}
```

- gpapi GetServiceCodeList --GpServiceCode `gp....` | jq .

```js
{
  "GetServiceCodeListResponse": {
    "RequestId": "<uuid>",
    "CustomerCode": "SGA....",
    "GpServiceCode": "gp.....",
    "GcServiceCodeList": [
      "gc....."
    ],
    "GnbServiceCodeList": [],
    "GxServiceCodeList": [],
    "GlServiceCodeList": [],
    "GvmServiceCodeList": [],
    "GvsServiceCodeList": [],
    "GomServiceCodeList": [
      "gom....."
    ]
  }
}
```

- API call
    - gpapi DescribeVirtualMachine --GpServiceCode `gp....` --GcServiceCode `gc....`
    - gpapi GetVirtualMachineStatus --GpServiceCode `gp....` --GcServiceCode `gc....`
    - gpapi StartVirtualMachine --GpServiceCode `gp....` --GcServiceCode `gc....`
- Other Commands
    - gpapi watch --GpServiceCode `gp....` --GcServiceCode `gc....`

# Usage for golang users

```go
package main

import (
	"github.com/iij/go-gpapi"
	"github.com/iij/go-gpapi/protocol"
	"log"
	"os"
)

func main() {
	api := gpapi.NewAPI(os.Getenv("IIJAPI_ACCESS_KEY"), os.Getenv("IIJAPI_SECRET_KEY"))
	param := protocol.EchoArg{Param: "Hello, World"}
	res := &protocol.EchoResponse{}
	if err := api.Call(param, res); err!=nil{
		log.Println("error", err)
	}else{
		log.Printf("res: param=%+v, res=%+v", param, res)
	}
}
```

Enjoy.
