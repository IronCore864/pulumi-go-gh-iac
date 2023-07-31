package main

import (
	"github.com/ironcoreworks/pulumi-go-gh-iac/internal/pkg/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(github.SetupTeams)
}
