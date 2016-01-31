#!/usr/bin/env bash
export GOPATH=$GOPATH:`pwd`
mkdir -p bin/mac
mkdir -p bin/windows
mkdir -p bin/linux
(cd src/cmd/forwardsocks5server/;GOOS=linux go build -o ../../../bin/linux/forwardsocks5server.linux main.go )
(cd src/cmd/forwardsocks5server/;GOOS=darwin go build -o ../../../bin/mac/forwardsocks5server.mac main.go )
(cd src/cmd/forwardsocks5work/;GOOS=windows go build -o ../../../bin/windows/forwardsocks5work.exe  main.go )
(cd src/cmd/forwardsocks5work/;GOOS=linux go build -o ../../../bin/linux/forwardsocks5work.linux  main.go )
(cd src/cmd/forwardsocks5work/;GOOS=darwin go build -o ../../../bin/mac/forwardsocks5work.mac  main.go )

