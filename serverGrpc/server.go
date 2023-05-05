package serverGrpc

import (
	"context"
	"errors"

	"github.com/google/uuid"
	pb "github.com/sergiovenicio/courses_grpc/gen/proto"
	"google.golang.org/grpc"
)

type Course struct {
	Id          string
	Name        string
	Description string
}

func NewCourse(id string, name string, description string) (*Course, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}

	if name == "" {
		return nil, errors.New("invalid name")
	}

	if description == "" {
		return nil, errors.New("invalid description")
	}

	courseId, _ := uuid.NewUUID()
	return &Course{
		courseId.String(),
		description,
		name,
	}, nil
}

type server struct {
	pb.CourseServiceServer
}

var courses []Course

func (s *server) CreateCourse(ctx context.Context, in *pb.CourseCreateRequest) (*pb.Course, error) {
	c, err := NewCourse("1", in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	courses = append(courses, *c)
	return &pb.Course{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
	}, nil

}

func (s *server) ListCourses(ctx context.Context, in *pb.Empty) (*pb.ListCourseResponse, error) {
	var pbCourses []*pb.Course
	for _, c := range courses {
		pbCourses = append(pbCourses, &pb.Course{
			Id:          c.Id,
			Name:        c.Name,
			Description: c.Description,
		})
	}
	return &pb.ListCourseResponse{
		Courses: pbCourses,
	}, nil
}

func (s *server) FindCourse(ctx context.Context, in *pb.FindCourseRequest) (*pb.Course, error) {
	var course Course
	for _, c := range courses {
		if c.Id == in.Id {
			course = c
		}
	}

	if course.Id == "" {
		return nil, errors.New("course not found")
	}

	return &pb.Course{
		Id:          course.Id,
		Name:        course.Name,
		Description: course.Description,
	}, nil
}

func NewGrpcServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterCourseServiceServer(grpcServer, &server{})
	return grpcServer
}
