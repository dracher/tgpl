package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.Int("port", 8000, "the port listen on")

func localTimeString(tz string) string {
	loc, _ := time.LoadLocation(tz)
	return time.Now().In(loc).Format(time.RubyDate)
}

func handleConn(c net.Conn, tz string) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, fmt.Sprintf("%s\n", localTimeString(tz)))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(5 * time.Second)
	}
}

func currentTimeZone() string {
	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "Asia/Shanghai"
	}
	log.Printf("Current timezone is %s", tz)
	return tz
}

func main() {
	flag.Parse()

	tz := currentTimeZone()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, tz)
	}
}
