package main

import (
	"fmt"
	"go.bug.st/serial"
	"log"
)

func main() {

	modeWrite := &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: 1,
	}

	firstPort, err := serial.Open("COM1", modeWrite)
	if err != nil {
		log.Fatal(err)
	}
	defer firstPort.Close()

	secondPort, err := serial.Open("COM2", modeWrite)
	if err != nil {
		log.Fatal(err)
	}
	defer secondPort.Close()

	message := "test"
	n, err := firstPort.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
	if n != len(message) {
		log.Fatal("Write did not write enough bytes")
	}
	fmt.Printf("Отправлено сообщение |%s| длинною |%d| байт", message, n)

	buffer := make([]byte, 1024)
	n, err = secondPort.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	if n != len(message) {
		log.Fatal("Read did not read enough bytes")
	}
	receivedMessage := string(buffer[:n])
	fmt.Printf("\nПрочитано сообщение |%s| длинною |%d| байт", receivedMessage, n)

}
