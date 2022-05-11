#!/bin/sh -l

echo "Hello $env"
time=$(date)
echo "::set-output name=time::$time"

/notifier
