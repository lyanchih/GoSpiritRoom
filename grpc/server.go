package main

import (
  "log"
  "fmt"
  "github.com/spf13/cobra"
  "google.golang.org/grpc"
  "golang.org/x/net/context"
  "net"
)

var (
  listenAddress string
)

func newServerCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use: "server",
    Short: "grpc server",
    RunE: runServer,
  }

  f := cmd.Flags()
  f.StringVar(&listenAddress, "address", fmt.Sprintf(":%d", defaultPort), "Server listen address")
  return cmd
}

type BallBox struct {
  n int
}

func (bb *BallBox) SayPing(ctx context.Context, p *Ping) (po *Pong, err error) {
  log.Println(ctx)
  if p.GetName() == "NOONE" {
    bb.n = 0
    po = &Pong{Message: "No!!! Ball is gone!!"}
  } else {
    po = &Pong{Message: fmt.Sprintf("Hey, %s. You got %d", p.GetName(), bb.n)}
    bb.n += 1
  }
  return po, nil
}

func runServer(cmd *cobra.Command, args []string) error {
  l, err := net.Listen("tcp", listenAddress)
  if err != nil {
    log.Println(err)
    return err
  }

  s := grpc.NewServer()
  
  RegisterFoobarServer(s, &BallBox{n: 100})
   
  return s.Serve(l)
}
