package brodcast

import (
          `go/types`
)

type BroadcastService struct {
          messages       types.Message
          incoming       chan types.Client
          leavingClients chan types.Client
}

func NewBroadcastService(message types.Message, incoming chan types.Client, leaving chan types.Client) *BroadcastService {
          
          return &BroadcastService{
                    messages:       message,
                    incoming:       incoming,
                    leavingClients: leaving,
          }
}
func (s *BroadcastService) Broadcast() {
          ClientConnState := make(map[types.Client]bool)
          for {
                    select {
                    case msg := <-s.messages:
                              // send message to all clients
                              for currentClient := range ClientConnState {
                                        currentClient <- msg
                              }
                    case currentClient := <-s.incomingClients:
                              ClientConnState[currentClient] = true
                    case currentClient := <-s.leavingClients:
                              delete(ClientConnState, currentClient)
                              close(currentClient)
                    }
          }
}
