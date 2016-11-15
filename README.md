# http-server

A simple no frills http file server. Simply run it from your directory of choice
to serve it up or run it with the following:

```
NAME:
   http-server - A simple no frills http file server

USAGE:
   http-server [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value  specify a port to bind to (a open port will be chosen if not provided or 0) (default: 0)
   --dir value, -d value   specify a folder to serve (the current directory will be chosen if not provided)
   --help, -h              show help
   --version, -v           print the version

```

## Install

Either run:

```
go get github.com/wyattjoh/http-server
```

Or download the latest release: https://github.com/wyattjoh/http-server/releases/latest
