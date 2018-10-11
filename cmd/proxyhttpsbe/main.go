//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var (
	flagHost     string
	flagPort     int
	httpsBackend string
)

func init() {
	flag.StringVar(&flagHost, "host", "github.com", "Host of the proxied https backend")
	flag.IntVar(&flagPort, "port", 443, "Port of the proxied https backend")
}

func handleConn(from net.Conn, backend string) {
	config := tls.Config{InsecureSkipVerify: true}
	to, err := tls.Dial("tcp", backend, &config)
	if err != nil {
		log.Printf("%v", err)
	} else {
		done := make(chan struct{})
		go func() {
			defer from.Close()
			defer to.Close()
			io.Copy(from, to)
			done <- struct{}{}
		}()

		go func() {
			defer from.Close()
			defer to.Close()
			io.Copy(to, from)
			done <- struct{}{}
		}()

		<-done
		<-done
	}
}

func proxy(backend string) {
	listener, err := net.Listen("tcp", ":8443")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("%v", err)
		} else {
			go handleConn(conn, backend)
		}
	}
}

func main() {
	flag.Parse()
	httpsBackend = fmt.Sprintf("%s:%d", flagHost, flagPort)

	log.Printf("Proxy server is serving at port 8443 to backend %s", httpsBackend)
	proxy(httpsBackend)
}
