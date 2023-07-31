# Managing GitHub Teams w/ Pulumi IaC

## Steps

```bash
pulumi new github-go
aws s3api create-bucket --bucket tiexin-pulumi-test-state --region us-east-1
pulumi login s3://tiexin-pulumi-test-state

export PULUMI_CONFIG_PASSPHRASE=""
pulumi config set github:owner ironcoreworks
pulumi config set github:token  --secret

pulumi preview
pulumi up

pulumi import github:index/team:Team manually-created-team 8324430
pulumi state unprotect urn:pulumi:dev::pulumi-go-gh-iac::github:index/team:Team::manually-created-team
```

