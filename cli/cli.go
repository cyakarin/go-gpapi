package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/iij/go-gpapi"
	p "github.com/iij/go-gpapi/protocol"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

func getflag(vtyp interface{}) []cli.Flag {
	res := []cli.Flag{}
	val, _ := gpapi.Interface2Values(vtyp)
	val.Del("Action")
	for k, v := range val {
		fl := cli.StringFlag{
			Name:   k,
			Value:  v[0],
			Usage:  v[0],
			EnvVar: strings.ToUpper(k),
		}
		res = append(res, fl)
	}
	return res
}

var shortnames map[string]string = make(map[string]string)

func shortname(n string) (res string) {
	var data []byte
	for _, r := range n {
		if 'A' <= r && r <= 'Z' {
			data = append(data, byte(r))
		}
	}
	name := string(data)
	if _, ok := shortnames[name]; ok {
		name += "2"
	}
	shortnames[name] = n
	return name
}

func add2app(app *cli.App, vtyp interface{}) (err error) {
	val, err := gpapi.Interface2Values(vtyp)
	name := val.Get("Action")
	val.Del("Action")
	cmd := cli.Command{
		Name:      name,
		ShortName: shortname(name),
		Usage:     "Call " + name + " API",
		Action: func(c *cli.Context) {
			api := gpapi.NewAPI(c.GlobalString("AccessKeyId"), c.GlobalString("SecretKey"))
			val := url.Values{}
			val.Set("Action", name)
			for _, v := range c.FlagNames() {
				vs := c.String(v)
				if vs == "required" {
					log.Fatal("missing required option ", v)
				}
				if vs != "" && vs != "optional" {
					if bt, err := ioutil.ReadFile(vs); err == nil {
						val.Set(v, string(bt))
					} else {
						val.Set(v, vs)
					}
				}
			}
			res, err := api.Post(val)
			if err != nil {
				log.Println("post error", err)
			}
			io.Copy(os.Stdout, res.Body)
		},
	}
	flags := getflag(vtyp)
	cmd.Flags = flags
	app.Commands = append(app.Commands, cmd)
	return
}

func statuswatch(c *cli.Context) {
	var prestate string
	pretime := time.Now()
	for {
		api := gpapi.NewAPI(c.GlobalString("AccessKeyId"), c.GlobalString("SecretKey"))
		arg := p.GetVirtualMachineStatusArg{}
		arg.GpServiceCode = c.String("GpServiceCode")
		arg.GcServiceCode = c.String("GcServiceCode")
		rsp := &p.GetVirtualMachineStatusResponse{}
		if err := api.Call(arg, rsp); err != nil {
			log.Println("parse", err)
		}
		if rsp.Status != prestate {
			log.Println(rsp.Status, time.Since(pretime))
			pretime = time.Now()
			prestate = rsp.Status
		}
		time.Sleep(time.Second * 3)
	}
}

func lscode(c *cli.Context) {
	api := gpapi.NewAPI(c.GlobalString("AccessKeyId"), c.GlobalString("SecretKey"))
	// get gp service code
	rootarg := p.GetGpServiceCodeListArg{}
	rootrsp := &p.GetGpServiceCodeListResponse{}
	if err := api.Call(rootarg, rootrsp); err != nil {
		log.Println("Error", err)
	}
	for _, gpcode := range rootrsp.GpServiceCodeList {
		fmt.Printf("%s\n", gpcode)
		scarg := p.GetServiceCodeListArg{GpServiceCode: gpcode}
		scrsp := &p.GetServiceCodeListResponse{}
		if err := api.Call(scarg, scrsp); err != nil {
			log.Println("Error", err)
		}
		sclist := []string{}
		sclist = append(sclist, scrsp.GnbServiceCodeList...)
		sclist = append(sclist, scrsp.GxServiceCodeList...)
		sclist = append(sclist, scrsp.GvsServiceCodeList...)
		sclist = append(sclist, scrsp.GlServiceCodeList...)
		sclist = append(sclist, scrsp.GcServiceCodeList...)
		sclist = append(sclist, scrsp.GvmServiceCodeList...)
		sclist = append(sclist, scrsp.GomServiceCodeList...)
		for _, sc := range sclist {
			fmt.Printf("  %s\n", sc)
		}
	}
}

func appendflag(flag []cli.Flag, name, desc, env string) []cli.Flag {
	fl := cli.StringFlag{
		Name:   name,
		Value:  "",
		Usage:  desc,
		EnvVar: env,
	}
	return append(flag, fl)
}

func main() {
	app := cli.NewApp()
	app.Name = "gpapi"
	app.Usage = "IIJ GP API Client CLI"
	app.Author = ""
	app.Email = ""
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{}
	app.Flags = appendflag(app.Flags, "AccessKeyId", "Access Key", "IIJAPI_ACCESS_KEY")
	app.Flags = appendflag(app.Flags, "SecretKey", "Secret Key", "IIJAPI_SECRET_KEY")
	// cd ../protocol; ls -1 *.go | grep -v Common | awk -F. '{print "\tadd2app(app, p."$1"Arg{})";}'

	add2app(app, p.AddCloneVirtualMachinesArg{})
	add2app(app, p.AddDiskOptionArg{})
	add2app(app, p.AddFwLbOptionArg{})
	add2app(app, p.AddGlobalAddressOptionArg{})
	add2app(app, p.AddLbNodeArg{})
	add2app(app, p.AddLbPoolArg{})
	add2app(app, p.AddLbVirtualServerArg{})
	add2app(app, p.AddVirtualIpAddressArg{})
	add2app(app, p.AddVirtualMachinesArg{})
	add2app(app, p.AddVlanOptionArg{})
	add2app(app, p.AttachFwLbArg{})
	add2app(app, p.AttachVlanArg{})
	add2app(app, p.ChangeFwLbOptionTypeArg{})
	add2app(app, p.ChangeVirtualMachineTypeArg{})
	add2app(app, p.CloneVirtualMachineArg{})
	add2app(app, p.DeleteDiskOptionArg{})
	add2app(app, p.DeleteFwLbOptionArg{})
	add2app(app, p.DeleteGlobalAddressOptionArg{})
	add2app(app, p.DeleteLbNodeArg{})
	add2app(app, p.DeleteLbPoolArg{})
	add2app(app, p.DeleteLbVirtualServerArg{})
	add2app(app, p.DeleteVirtualIpAddressArg{})
	add2app(app, p.DeleteVirtualMachineArg{})
	add2app(app, p.DeleteVlanOptionArg{})
	add2app(app, p.DescribeFwArg{})
	add2app(app, p.DescribeLbArg{})
	add2app(app, p.DescribeVirtualMachineArg{})
	add2app(app, p.DetachFwLbArg{})
	add2app(app, p.DetachVlanArg{})
	add2app(app, p.DisableIPv6Arg{})
	add2app(app, p.EchoArg{})
	add2app(app, p.EnableIPv6Arg{})
	add2app(app, p.GetContractInformationArg{})
	add2app(app, p.GetContractStatusArg{})
	add2app(app, p.GetDomainNameByIpAddressArg{})
	add2app(app, p.GetFilteringRuleListArg{})
	add2app(app, p.GetGpServiceCodeListArg{})
	add2app(app, p.GetLbHostListArg{})
	add2app(app, p.GetLbPoolArg{})
	add2app(app, p.GetLbStatusArg{})
	add2app(app, p.GetLbVirtualServerArg{})
	add2app(app, p.GetServiceCodeListArg{})
	add2app(app, p.GetSnatRuleListArg{})
	add2app(app, p.GetVirtualMachineStatusArg{})
	add2app(app, p.GetVirtualMachineStatusListArg{})
	add2app(app, p.ImportRootSshPublicKeyArg{})
	add2app(app, p.InitializeVirtualMachineArg{})
	add2app(app, p.RebootVirtualMachineArg{})
	add2app(app, p.ResetDomainNameByIpAddressArg{})
	add2app(app, p.ResetIptablesArg{})
	add2app(app, p.ResetNetworkConfigurationArg{})
	add2app(app, p.ResetSshdConfigurationArg{})
	add2app(app, p.SetDomainNameByIpAddressArg{})
	add2app(app, p.SetFilteringRuleListArg{})
	add2app(app, p.SetLabelArg{})
	add2app(app, p.SetLbNodeArg{})
	add2app(app, p.SetLbPoolArg{})
	add2app(app, p.SetLbPoolNameArg{})
	add2app(app, p.SetLbVirtualServerArg{})
	add2app(app, p.SetLbVirtualServerNameArg{})
	add2app(app, p.SetSnatRuleListArg{})
	add2app(app, p.SetVirtualIpAddressNameArg{})
	add2app(app, p.StartVirtualMachineArg{})
	add2app(app, p.StopVirtualMachineArg{})

	app.Commands = append(app.Commands, cli.Command{
		Name:      "watch",
		ShortName: "w",
		Action:    statuswatch,
		Usage:     "VM status watcher",
		Flags:     getflag(p.GetVirtualMachineStatusArg{}),
	})
	app.Commands = append(app.Commands, cli.Command{
		Name:      "lscode",
		ShortName: "lc",
		Usage:     "List Service Code",
		Action:    lscode,
	})
	app.Run(os.Args)
}
