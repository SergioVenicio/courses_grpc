/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/sergiovenicio/courses_grpc/client"
	pb "github.com/sergiovenicio/courses_grpc/gen/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var method string
var name string
var description string
var id string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		switch method {
		case "list":
			list(ctx, conn)
		case "create":
			create(ctx, conn)
		case "find":
			find(ctx, conn)
		}
	},
}

func create(ctx context.Context, conn *grpc.ClientConn) {
	client.CreateCourse(ctx, conn, &pb.CourseCreateRequest{
		Name:        name,
		Description: description,
	})
}

func list(ctx context.Context, conn *grpc.ClientConn) {
	client.ListCourses(ctx, conn)
}

func find(ctx context.Context, conn *grpc.ClientConn) {
	client.FindCourse(ctx, conn, &pb.FindCourseRequest{
		Id: id,
	})
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.PersistentFlags().StringVar(&method, "method", "list", "")
	clientCmd.PersistentFlags().StringVar(&name, "name", "", "")
	clientCmd.PersistentFlags().StringVar(&description, "description", "", "")
	clientCmd.PersistentFlags().StringVar(&id, "id", "", "")
}
