#!/bin/bash
export PATH=$PATH:$(go env GOPATH)/bin
str=$(ls -R | grep ./T2 | cut -d':' -f1)
names=`echo $str | tr " " "\n"`
touch grepres.txt
for line in $names; do
 go vet "$line" &>> grepres.txt
done 
golangci-lint run -c .golangci-lint.yaml ./... 