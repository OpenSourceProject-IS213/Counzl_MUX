package new_user

import (
	"log"
	"crypto/tls"
	"bufio"
	"net"
	"fmt"
	"homemade/converter"
)

const (
	server_crt = "./database/.certificate/server.crt"
	server_key = "./database/.certificate/server.key"
	LN_PORT = ":8081"
	network = "tcp"
)

func Initialise_User_Listener() {
	log.SetFlags(log.Lshortfile)

	certs, err := tls.LoadX509KeyPair(server_crt, server_key)
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{certs}}
	ln, err := tls.Listen(network, LN_PORT, config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

// Funksjon for å ta seg av trafikk til :8081 og sjekke om klient har riktig nøkkel.
func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		username_0a, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		username := converter.Remove_0a(username_0a)
		id := newUser(username)
		fmt.Println(username)

		n, err := conn.Write([]byte(id + "\n"))
		if err != nil {
			log.Println(n, err)
			return
		}

	}
}