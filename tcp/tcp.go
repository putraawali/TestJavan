package tcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"testjavan/helpers/constants"
	"testjavan/model"
	"testjavan/tcp/handler"
	tcp_model "testjavan/tcp/model"
	"time"
)

type Message struct {
	remoteAddr string
	msg        []byte
}

type Server struct {
	listenAddr string
	listener   net.Listener
	quitChanel chan struct{}
	msgChanel  chan Message
}

func NewServer(listenAdd string) *Server {
	return &Server{
		listenAddr: listenAdd,
		quitChanel: make(chan struct{}),
		msgChanel:  make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	s.listener = listener

	go s.acceptRequest()

	<-s.quitChanel
	close(s.msgChanel)
	return nil
}

func (s *Server) acceptRequest() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Error accept connection: " + err.Error())
			// Error make block for another connection. So we have to continue it so another connection can connect to our server
			continue
		}

		fmt.Println("New connection to server: ", conn.RemoteAddr().String())

		// Run as go routine, so the server able to handle million of requests
		go s.readRequest(conn)
	}
}

func (s *Server) readRequest(conn net.Conn) {
	defer conn.Close()

	for {
		var (
			result  model.Return
			err     error
			request tcp_model.Request
		)

		d := json.NewDecoder(conn)
		err = d.Decode(&request)

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			fmt.Println("Error read connection:" + err.Error())
			continue
		}

		switch request.Category {
		case "family":
			result, err = handler.FamilyHandler(request.Method, request.Args)
		}

		byteRequest, _ := json.Marshal(request)

		s.msgChanel <- Message{
			remoteAddr: conn.RemoteAddr().String(),
			msg:        byteRequest,
		}

		var reply []byte
		if err != nil {
			reply, _ = json.Marshal(err.Error())
		} else {
			reply, _ = json.Marshal(result)
		}

		conn.Write([]byte(reply))
	}
}

func TCP() {
	server := NewServer("localhost:" + os.Getenv("TCP_PORT"))

	go func() {
		for msg := range server.msgChanel {
			fmt.Printf("message from connection (%s): %s \n", string(msg.remoteAddr), string(msg.msg))
		}
	}()

	fmt.Printf("[%s] TCP Service running on port: %s\n", time.Now().Format(constants.TimeFormat), os.Getenv("TCP_PORT"))
	log.Fatal(server.Start())
}
