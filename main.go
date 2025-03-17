package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/pagopa/terraform-provider-azure-dx/internal/provider"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	// providerserver.Serve(context.Background(), provider.New(version), providerserver.ServeOpts{
	// 	Address: "registry.terraform.io/pagopa/azure-dx",
	// })

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/pagopa/azure-dx",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}

}
