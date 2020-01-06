#!/bin/bash
git pull &&
go get &&
go build -o main &&
mv main /bin/ezservice
echo "done"
