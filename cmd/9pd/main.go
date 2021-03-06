package main

import (
	"context"
	"flag"
	"log"
	"os"

	_ "net/http/pprof"

	"github.com/altid/server"
)

var factotum = flag.Bool("f", false, "Enable client authentication via a plan9 factotum")
var usetls = flag.Bool("t", false, "Use TLS")
var chatty = flag.Bool("D", false, "Chatty")
var debug = flag.Bool("d", false, "Debug")

var port = flag.String("p", "564", "Port to listen on")
var addr = flag.String("l", "", "Address to listen on")
var dir = flag.String("m", "/tmp/altid", "Path to Altid services")

var cert string
var key string

func main() {
	flag.Parse()

	if flag.Lookup("h") != nil {
		flag.Usage()
		os.Exit(0)
	}

	if *usetls {
		// TODO(halfwit) config.ServerTLS()
	}

	ctx := context.Background()
	svc := &service{
		listen: *dir,
		chatty: *chatty,
		tls:    *usetls,
	}

	srv, err := server.NewServer(ctx, svc, *dir)
	if err != nil {
		log.Fatal(err)
	}

	if *debug {
		srv.Logger = log.Printf
	}

	if e := srv.Listen(); e != nil {
		log.Fatal(e)
	}
}
