package main

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
	_ "github.com/coredns/coredns/plugin/debug"
	_ "github.com/coredns/coredns/plugin/log"

	"gophers.dev/cmds/donutdns/plugins/donutdns"
)

var directives = []string{
	"donutdns",
	"debug",
	"log",
	"startup",
	"shutdown",
}

func init() {
	dnsserver.Port = "1053"
	dnsserver.Directives = directives
	caddy.SetDefaultCaddyfileLoader(donutdns.PluginName, caddy.LoaderFunc(func(serverType string) (caddy.Input, error) {
		return caddy.CaddyfileInput{
			Filepath:       donutdns.PluginName,
			Contents:       []byte(corefile),
			ServerTypeName: "dns",
		}, nil
	}))
}

func main() {
	coremain.Run()
}

const corefile = `.:1053 {
  debug
  log
  donutdns {
	allow example.com
	block facebook.com
	block instagram.com
  }
}`
