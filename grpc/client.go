package main

import (
  "log"
  "fmt"
  "strings"
  "github.com/spf13/cobra"
  "context"
  "google.golang.org/grpc"
)

var (
  targetAddress string
)

func newClientCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use: "client",
    Short: "grpc client",
    RunE: runClient,
  }

  f := cmd.Flags()
  f.StringVar(&targetAddress, "addres", fmt.Sprintf("%s:%d", "localhost", defaultPort), "Client target address")
  return cmd
}

func runClient(cmd *cobra.Command, args []string) error {
  cc, err := grpc.Dial(targetAddress, grpc.WithInsecure())
  if err != nil {
    log.Println(err)
    return err
  }
  defer cc.Close()

  var n string
  if len(args) == 0 || (len(args) == 1 && len(args[0]) == 0) {
    n = "NOONE"
  } else {
    n = strings.Join(args, ",")
  }

  c := NewFoobarClient(cc)
  p, err := c.SayPing(context.Background(), &Ping{Name: n})
  if err != nil {
    log.Println(err)
    return err
  }

  fmt.Println(p.GetMessage())
  return nil
}
