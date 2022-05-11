# GHSA-NOTIFY

## Usage
Add this file to your rss reader: 

https://github.com/ssst0n3/GHSA-NOTIFY/blob/main/feed/container.xml

## For Developer

### Project Structure Introduction

* .github/workflows: github action dir
  * container-feed.yml: the action to pull and push the latest security advisories, executed by every 5 mins
  * ghcr-push.yml: the action to build and push a container image for quering github api
* cmd/notifier: the go main package
