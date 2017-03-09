package main

import (
  "github.com/spf13/cobra"
)

const defaultPort = 9487

func main() {
  cmd := &cobra.Command{
    Use: "grpc",
    Short: "grpc skill improve",
  }

  cmd.AddCommand(newServerCmd(), newClientCmd())
  cmd.Execute()
}
