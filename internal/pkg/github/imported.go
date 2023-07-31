package github

import (
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func importedTeams(ctx *pulumi.Context) error {
	_, err := github.NewTeam(ctx, "manually-created-team", &github.TeamArgs{
		Name:    pulumi.String("Manually Created Team"),
		Privacy: pulumi.String("closed"),
	}, pulumi.Protect(true))
	if err != nil {
		return err
	}
	return nil
}
