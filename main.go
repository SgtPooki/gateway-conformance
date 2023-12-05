package main

// import cli from cmd/gateway-conformance/cli.go
import (
	"os"

	cli "github.com/SgtPooki/gateway-conformance/cmd/gateway-conformance"
)

func main() {
	app := cli.Cli()
	app.Run(os.Args)
}
