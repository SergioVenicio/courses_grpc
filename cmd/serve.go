/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"

	"github.com/sergiovenicio/courses_grpc/serverGrpc"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := serverGrpc.NewGrpcServer()
		list, _ := net.Listen("tcp", ":50051")
		fmt.Println("running on :50051")
		if err := server.Serve(list); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
