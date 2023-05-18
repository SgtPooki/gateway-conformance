package tests

import (
	"testing"

	"github.com/ipfs/gateway-conformance/tooling/test"
	. "github.com/ipfs/gateway-conformance/tooling/test"
)

func TestCors(t *testing.T) {
	cidHello := "bafkqabtimvwgy3yk" // hello

	tests := SugarTests{
		{
			Name: "GET Response from Gateway should contain CORS headers",
			Request: Request().
				Path("ipfs/{{CID}}/", cidHello),
			Response: Expect().
				Headers(
					Header("Access-Control-Allow-Origin").Equals("*"),
					Header("Access-Control-Allow-Methods").Equals("GET"),
					Header("Access-Control-Allow-Headers").Has("Range"),
					Header("Access-Control-Expose-Headers").Has(
						"Content-Range",
						"Content-Length",
						"X-Ipfs-Path",
						"X-Ipfs-Roots",
					),
				),
		},
		{
			Name: "OPTIONS to Gateway succeeds",
			Request: Request().
				Method("OPTIONS").
				Path("ipfs/{{CID}}/", cidHello),
			Response: Expect().
				Headers(
					Header("Access-Control-Allow-Origin").Equals("*"),
					Header("Access-Control-Allow-Methods").Equals("GET"),
					Header("Access-Control-Allow-Headers").Has("Range"),
					Header("Access-Control-Expose-Headers").Has(
						"Content-Range",
						"Content-Length",
						"X-Ipfs-Path",
						"X-Ipfs-Roots",
					),
				),
		},
	}

	test.Run(t, tests)
}
