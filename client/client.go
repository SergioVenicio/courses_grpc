package client

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/sergiovenicio/courses_grpc/gen/proto"
	"google.golang.org/grpc"
)

type RequestError struct {
	Error string
}

func CreateCourse(ctx context.Context, conn *grpc.ClientConn, r *pb.CourseCreateRequest) {
	client := pb.NewCourseServiceClient(conn)
	course, err := client.CreateCourse(ctx, r)
	if err != nil {
		jsonError, _ := json.Marshal(RequestError{
			Error: err.Error(),
		})
		fmt.Println(string(jsonError))
		return
	}

	jsonCourse, _ := json.Marshal(course)
	fmt.Println(string(jsonCourse))
}

func ListCourses(ctx context.Context, conn *grpc.ClientConn) {
	client := pb.NewCourseServiceClient(conn)
	courses, err := client.ListCourses(ctx, &pb.Empty{})
	if err != nil {
		jsonError, _ := json.Marshal(RequestError{
			Error: err.Error(),
		})
		fmt.Println(string(jsonError))
		return
	}

	for _, course := range courses.Courses {
		jsonCourse, _ := json.Marshal(course)
		fmt.Println(string(jsonCourse))
	}
}

func FindCourse(ctx context.Context, conn *grpc.ClientConn, r *pb.FindCourseRequest) {
	client := pb.NewCourseServiceClient(conn)
	course, err := client.FindCourse(ctx, r)
	if err != nil {
		jsonError, _ := json.Marshal(RequestError{
			Error: err.Error(),
		})
		fmt.Println(string(jsonError))
		return
	}

	jsonCourse, _ := json.Marshal(course)
	fmt.Println(string(jsonCourse))
}
