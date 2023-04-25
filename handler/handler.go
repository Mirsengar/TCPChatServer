package handler

import (
          `bufio`
          `fmt`
          `go/types`
          `net`
          
          `TCPChatServer/message`
)

type Handler struct {
          msgService      *message.Service
          messagesBuffer  types.Message
          incomingClients chan types.Client
          leavingClients  chan types.Client
}

func NewHandler(service *message.Service, msg types.Message, incoming chan types.Client, leaving chan types.Client) *Handler {
          return &Handler{
                    msgService:      service,
                    messagesBuffer:  msg,
                    incomingClients: incoming,
                    leavingClients:  leaving,
          }
}
func (h *Handler) HandleConnection(conn net.Conn) {
          defer conn.Close()
          clientMessages := make(chan string)
          go h.msgService.MessageWriter(conn, clientMessages)
          clientName := conn.RemoteAddr().String() // 192.168.1.11:8080
          clientMessages <- fmt.Sprintf("Welcome to the server, %s\n", clientName)
          h.messagesBuffer <- clientName + "has joined!!"
          h.incomingClients <- clientMessages
          inputMessage := bufio.NewScanner(conn)
          for inputMessage.Scan() {
                    h.messagesBuffer <- clientName + ": " + inputMessage.Text()
          }
          h.leavingClients <- clientMessages
          h.messagesBuffer <- clientName + " has left the chat!"
}
