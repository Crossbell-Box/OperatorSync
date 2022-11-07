# xSync Server

Backend server for xSync

## Deploy

We recommend deploying xSync server in docker or kubernetes.

### Requirements

- IPFS-Upload-Relay: Please deploy your own `https://github.com/NaturalSelectionLabs/IPFS-Upload-Relay` server.

### Docker

1. Copy environment files from `deploy/env/.example` to `deploy/env`
2. Edit environment files to match your requirements, like:
   - Set `ETHEREUM_PRIVATE_KEY` to your operator's private key in `worker.env`.
   - Set `IPFS_ENDPOINT` to your own IPFS-Upload-Relay server's URI (like `https://upload-relay.example.ltd`).
   - Edit other fields in different files (like rsshub.stateful.env) for rsshub with platforms logged in
3. Build services by `make build-docker`
4. Start services by `make prod-start`

For more details, please refer to Makefile.

### Kubernetes

> Refer to [.github/workflows/docker-build-push.yml](https://github.com/Crossbell-Box/OperatorSync/blob/develop/.github/workflows/docker-build-push.yml) 

## Develop

This server parse users' RSS feeds to structured data, and then post them on chain.

### Add a new platform

1. Edit `SUPPORTED_PLATFORM` in `common/consts/platform.go`, add target platform's basic information.
2. Create a new directory under `app/worker/platforms/`.
3. We need to implement 2 functions, one for account validate and the other for feed collect. For implementing details, please refer to other platforms.
4. Call account validate func in `app/worker/rpc/jobs/validate_account.go`, and call feed collect func in `app/worker/mq/jobs/dispatch/collect_feeds.go`.
5. Time to test :tada:

If there's any further questions, please open an issue.
