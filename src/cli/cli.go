package cli

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Lemuren/teal/telnet"
)

// This function handles the interactive cli during a session.
func CliLoop(url string, timeout time.Duration) {
	// Connect to the server and listen for a message.
	conn, err := telnet.Connect(url, timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	defer conn.Close()
	response, err := telnet.Listen(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
	}
	fmt.Println(response)

	// Main interactive cli loop.
	for {
		fmt.Printf(url + ">")
		scanner := bufio.NewScanner(os.Stdin)
		var message string
		if scanner.Scan() {
			message = scanner.Text()
		}
		response, err := telnet.SendAndListen(conn, message)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
		}
		fmt.Println(response)
		if message == "QUIT" {
			os.Exit(0)
		}
	}
}
