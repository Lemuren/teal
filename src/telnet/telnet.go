package telnet

import (
	"net"
	"time"
)

// Connect establishes a TCP connection to the Telnet server.
func Connect(url string, timeout time.Duration) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", url, timeout)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// SendAndListen sends a message string to the server and returns a response
// if successful.
func SendAndListen(conn net.Conn, message string) (string, error) {
	err := Send(conn, message)
	if err != nil {
		return "", err
	}

	response, err := Listen(conn)
	if err != nil {
		return "", err
	}

	return response, nil
}

// Send sends a message string to the server.
func Send(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message + "\r\n"))
	if err != nil {
		return err
	}
	return nil
}

// Listen listens for a response from the server.
func Listen(conn net.Conn) (string, error) {
	response := make([]byte, 1024)
	_, err := conn.Read(response)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
