package github

import (
	"fmt"
	"strconv"

	"github.com/ironcoreworks/pulumi-go-gh-iac/internal/pkg/config"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func SetupTeams(ctx *pulumi.Context) error {
	// imported teams
	importedTeams(ctx)

	org, err := config.LoadOrganizationConfig()
	if err != nil {
		return err
	}

	for _, team := range org.Teams {
		if err := setupParentTeam(ctx, &team); err != nil {
			return err
		}
	}

	return nil
}

func setupParentTeam(ctx *pulumi.Context, parentTeam *config.Team) error {
	// setup parent team
	ghParentTeam, err := github.NewTeam(ctx, parentTeam.Slug, &github.TeamArgs{
		Description: pulumi.String(parentTeam.Description),
		Name:        pulumi.String(parentTeam.Name),
		Privacy:     pulumi.String("closed"),
	}, pulumi.Protect(false))
	if err != nil {
		fmt.Println("encountered error creating new Pulumi github parent team: ", parentTeam.Name)
		return err
	}

	// add members that belong to the parent team
	addMembers(ctx, parentTeam.Members, ghParentTeam, parentTeam.Name)

	// set each child team's parent team ID to the current team ID
	if err := setupChildTeams(ctx, parentTeam.Teams, ghParentTeam); err != nil {
		return err
	}

	return nil
}

func setupChildTeams(ctx *pulumi.Context, childTeams []config.Team, ghParentTeam *github.Team) error {
	for _, childTeam := range childTeams {
		// create child team
		ghChildTeam, err := github.NewTeam(ctx, childTeam.Slug, &github.TeamArgs{
			Description: pulumi.String(childTeam.Description),
			Name:        pulumi.String(childTeam.Name),
			Privacy:     pulumi.String("closed"),
			ParentTeamId: ghParentTeam.ID().ApplyT(func(id interface{}) int {
				// we need to re-cast id as an int so we can then transform it into a pulumi.IntOutput, which can be used to set the ParentTeamId.
				x := fmt.Sprintf("%v", id)
				y, _ := strconv.Atoi(x)
				return y
			}).(pulumi.IntOutput),
		}, pulumi.Protect(false), pulumi.Parent(ghParentTeam))
		if err != nil {
			fmt.Println("encountered error creating new Pulumi github child team: ", childTeam.Name)
			return err
		}

		// add members that belong to the child team
		if err := addMembers(ctx, childTeam.Members, ghChildTeam, childTeam.Name); err != nil {
			return err
		}

		// handle nested child teams recursively
		if len(childTeam.Teams) > 0 {
			if err := setupChildTeams(ctx, childTeam.Teams, ghChildTeam); err != nil {
				return err
			}
		}
	}

	return nil
}
