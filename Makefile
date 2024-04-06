all src/main.go src/telnet/*.go src/cmd/*.go src/cli/*.go:
	cd src; go build; mv teal ../bin

