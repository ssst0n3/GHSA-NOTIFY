# GHSA-NOTIFY

## Usage
Add these files to your rss reader: 

### Container Software Security Advisories
[![Container Security Feed](https://github.com/ssst0n3/GHSA-NOTIFY/actions/workflows/container-feed.yml/badge.svg?branch=main)](https://github.com/ssst0n3/GHSA-NOTIFY/actions/workflows/container-feed.yml)

https://github.com/ssst0n3/GHSA-NOTIFY/blob/main/feed/container.xml

* packages listened now:
  * github.com/docker/docker
  * github.com/containerd/containerd
  * github.com/opencontainers/runc
  * github.com/containerd/imgcrypt
  * CRI-O
  * if you find more packages, please pull a request
### for other softwares, please pull a request

## For Developer

### Project Structure Introduction

* .github/workflows: github action dir
  * container-feed.yml: the action to pull and push the latest security advisories, executed by every 5 mins
  * ghcr-push.yml: the action to build and push a container image for quering github api
* cmd/notifier: the go main package

### Container Image
[![Push to GHCR](https://github.com/ssst0n3/GHSA-NOTIFY/actions/workflows/ghcr-push.yaml/badge.svg?event=push)](https://github.com/ssst0n3/GHSA-NOTIFY/actions/workflows/ghcr-push.yaml)
