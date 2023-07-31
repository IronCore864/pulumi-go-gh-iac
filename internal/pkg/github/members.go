package github

import (
	"fmt"

	"github.com/ironcoreworks/pulumi-go-gh-iac/internal/pkg/config"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func addMembers(ctx *pulumi.Context, members []config.Member, ghTeam *github.Team, teamName string) error {
	for _, member := range members {
		// unique name for TeamMembership
		utmName := teamName + "-" + member.UserName

		role := member.Role
		if role == "" {
			// role defaults to "member"
			role = "member"
		}

		_, err := github.NewTeamMembership(ctx, utmName, &github.TeamMembershipArgs{
			TeamId:   ghTeam.ID(),
			Username: pulumi.String(member.UserName),
			Role:     pulumi.String(role),
		})
		if err != nil {
			fmt.Printf("encountered error adding member %s to github team %s\n", teamName, member.UserName)
			return err
		}
	}
	return nil
}
