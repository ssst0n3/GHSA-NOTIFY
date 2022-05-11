#!/bin/sh -l
set -ex
env
mkdir -p repo
git clone ${GITHUB_SERVER_URL}/${GITHUB_REPOSITORY} repo
cd repo

/notifier
ls -lah output
