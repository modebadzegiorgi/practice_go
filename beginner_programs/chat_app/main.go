package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

type Message struct {
	User string    `json:"user_name"`
	Text string    `json:"message"`
	Time time.Time `json:"timestamp"`
}

type Server struct {
	conns    map[*websocket.Conn]bool
	messages []Message
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New WebSocket connection:", ws.RemoteAddr())
	s.conns[ws] = true

	defer func() {
		delete(s.conns, ws)
		ws.Close()
	}()
	s.broadcastHistory(ws)
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
			} else {
				fmt.Println("Error reading:", err)
			}
			break
		}

		msg := buf[:n]
		smsg := strings.Split(string(msg), ":")
		fmt.Println("Received message:", smsg)

		message := Message{
			User: smsg[0],
			Text: smsg[1],
			Time: time.Now().UTC(),
		}
		s.messages = append(s.messages, message)

		s.broadcast(message, ws)
	}
}

func (s *Server) broadcast(msg Message, ws *websocket.Conn) {
	messageJson, _ := json.Marshal(msg)
	for conn := range s.conns {
		if conn != ws {
			conn.Write(messageJson)
		}
	}
}

func (s *Server) broadcastHistory(ws *websocket.Conn) {
	if len(s.messages) > 0 {
		for _, msg := range s.messages {
			messageJson, _ := json.Marshal(msg)
			ws.Write(messageJson)
		}
	}
}

func (s *Server) storeHistory() {

	for {
		data, _ := json.Marshal(s.messages)
		os.WriteFile("data.json", data, 0644)
		time.Sleep(time.Second * 10)
	}

}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/index.html")
}

func main() {
	s := NewServer()
	go s.storeHistory()

	http.HandleFunc("/", s.handleIndex)

	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/ws", websocket.Handler(s.handleWS))
	fmt.Println("Server started at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
