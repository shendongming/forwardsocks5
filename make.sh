(cd src/cmd/forwardsocks5server/;GOOS=linux go build -o ../../../bin/forwardsocks5server.linux main.go )
(cd src/cmd/forwardsocks5work/;GOOS=windows go build -o ../../../bin/forwardsocks5work.exe  main.go )
(cd src/cmd/forwardsocks5work/;GOOS=linux go build -o ../../../bin/forwardsocks5work.linux  main.go )
(cd src/cmd/forwardsocks5work/;GOOS=darwin go build -o ../../../bin/forwardsocks5work.mac  main.go )
