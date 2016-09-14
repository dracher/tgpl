package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"

	"fmt"
)

var portList = []string{"8000", "8010", "8020"}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {

	for _, port := range portList {
		go func(port string) {
			log.Println(port)
			conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
			if err != nil {
				log.Fatal(err)
			}

			defer conn.Close()

			mustCopy(os.Stdout, conn)
		}(port)
	}

	time.Sleep(10)
}
