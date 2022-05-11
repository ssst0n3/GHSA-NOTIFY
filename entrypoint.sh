#!/bin/sh -l
set -ex
env
mkdir -p repo
git clone ${GITHUB_SERVER_URL}/${GITHUB_REPOSITORY} repo
cd repo

/notifier
ls -lah output

git config --global user.email "ssst0n3@gmail.com"
git config --global user.name "ssst0n3"
git add .
git commit -m "container.xml updated by github action"
git push