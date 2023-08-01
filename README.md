# Managing GitHub Teams w/ Pulumi IaC

- one single YAML file as input (`config.yaml`)
- Pulumi w/ Golang
- GitHub Actions

## Initialize

Create a stack, configure remote backend, and set up GitHub owner/token:

```bash
pulumi new github-go
aws s3api create-bucket --bucket tiexin-pulumi-test-state --region us-east-1
pulumi login s3://tiexin-pulumi-test-state
export PULUMI_CONFIG_PASSPHRASE=""
pulumi config set github:owner ironcoreworks
pulumi config set github:token  --secret
```

## Capture the Initial State

Start from an empty state, with no code in the stack, first, we need to import existing teams.

For GitHub, existing team membership doesn't need to be explicitly imported, as they are merely establishing cross references between GitHub Users and Teams.

To import a team, run:

```bash
pulumi import github:index/team:Team manually-created-team 8324430
```

Add the generated code into the stack, then unprotect the state:

```bash
pulumi state unprotect urn:pulumi:dev::pulumi-go-gh-iac::github:index/team:Team::manually-created-team
```

Last, run a `pulumi preview` which should show no changes because now the code should reflect the existing infrastructure.

## Adding Teams and Membership

Change the `config.yaml` file accordingly, no need to change the Pulumi code.
