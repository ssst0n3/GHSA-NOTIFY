#!/bin/sh -l
#set -ex
#env
#mkdir -p repo
#git clone ${GITHUB_SERVER_URL}/${GITHUB_REPOSITORY} repo
#cd repo

/notifier
ls -lah output

#git config --global user.email "ssst0n3@gmail.com"
#git config --global user.name ${GITHUB_REPOSITORY_OWNER}
#git config --global user.password ${GHTOKEN}
#git add .
#git commit -m "container.xml updated by github action"
#git push