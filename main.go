package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

// VERSION is added by goxc during the build phase.
var VERSION string

func serve(ctx *cli.Context) error {
	port := ctx.Int("port")
	if port == 0 {
		freePort, err := getPort()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		port = freePort
	}

	dir := ctx.String("dir")
	if len(dir) > 0 {
		cleanedDir, err := getDirectory(dir)
		if err != nil {
			return err
		}

		dir = cleanedDir
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		dir = pwd
	}

	log.Printf("Now serving %s at http://localhost:%d/", dir, port)

	http.Handle("/", logWrapper(http.FileServer(http.Dir(dir))))

	if err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil); err != nil {
		return err
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = "http-server"
	app.Usage = "A simple no frills http file server"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Usage: "specify a port to bind to (a open port will be chosen if not provided or 0)",
		},
		cli.StringFlag{
			Name:  "dir, d",
			Usage: "specify a folder to serve (the current directory will be chosen if not provided)",
		},
	}
	app.Action = serve

	app.Run(os.Args)
}
