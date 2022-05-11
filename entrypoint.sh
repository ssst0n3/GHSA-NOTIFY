#!/bin/sh -l
set -ex
mkdir -p repo
git clone -o repo ${GITHUB_REPOSITORY}
cd repo

/notifier
ls -lah output
