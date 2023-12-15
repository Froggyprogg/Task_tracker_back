package board

import (
	"context"
	board "github.com/Task_tracker_back/pkg/board_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type server struct {
	board.UnimplementedBoardV1Server
}

var (
	database *gorm.DB
)

func Register(gRPCServer *grpc.Server, db *gorm.DB) {
	board.RegisterBoardV1Server(gRPCServer, &server{})
	database = db
}

// GET
func (s *server) GetBoard(ctx context.Context, req *board.GetRequestBoard) (*board.GetResponseBoard, error) {
	return &board.GetResponseBoard{}, nil
}

func (s *server) GetTask(ctx context.Context, req *board.GetRequestTask) (*board.GetResponseTask, error) {
	return &board.GetResponseTask{}, nil
}

func (s *server) GetTasks(ctx context.Context, req *board.GetRequestTasks) (*board.GetResponseTasks, error) {
	return &board.GetResponseTasks{}, nil
}

func (s *server) GetSubtasks(ctx context.Context, req *board.GetRequestSubtasks) (*board.GetResponseSubtasks, error) {
	return &board.GetResponseSubtasks{}, nil
}

func (s *server) GetUserAtBoard(ctx context.Context, req *board.GetRequestUserAtBoard) (*board.GetResponseUserAtBoard, error) {
	return &board.GetResponseUserAtBoard{}, nil
}

func (s *server) GetColumns(ctx context.Context, req *board.GetRequestColumns) (*board.GetResponseColumns, error) {
	return &board.GetResponseColumns{}, nil
}

func (s *server) GetStatus(ctx context.Context, req *board.GetRequestStatus) (*board.GetResponseStatus, error) {
	return &board.GetResponseStatus{}, nil
}

func (s *server) GetAllStatuses(ctx context.Context, req *board.GetRequestStatuses) (*board.GetResponseStatuses, error) {
	return &board.GetResponseStatuses{}, nil
}

func (s *server) GetReports(ctx context.Context, req *board.GetRequestStatuses) (*board.GetResponseStatuses, error) {
	return &board.GetResponseStatuses{}, nil
}

func (s *server) GetRoles(ctx context.Context, req *board.GetRequestRoles) (*board.GetResponseRoles, error) {
	return &board.GetResponseRoles{}, nil
}

////////////////////////////////////////////////////////

// PUT
func (s *server) PutColumn(ctx context.Context, req *board.Columns) (*board.Columns, error) {
	return &board.Columns{}, nil
}

func (s *server) PutReports(ctx context.Context, req *board.Reports) (*board.Reports, error) {
	return &board.Reports{}, nil
}

func (s *server) PutBoard(ctx context.Context, req *board.Boards) (*board.Boards, error) {
	return &board.Boards{}, nil
}

func (s *server) PutTask(ctx context.Context, req *board.Tasks) (*board.Tasks, error) {
	return &board.Tasks{}, nil
}

func (s *server) PutSubtask(ctx context.Context, req *board.Subtasks) (*board.Subtasks, error) {
	return &board.Subtasks{}, nil
}

func (s *server) PutComments(ctx context.Context, req *board.Comments) (*board.Comments, error) {
	return &board.Comments{}, nil
}

//////////////////////////////////////////////////

// POST
func (s *server) AddUserToBoard(ctx context.Context, req *board.UsersBoard) (*board.UsersBoard, error) {
	return &board.UsersBoard{}, nil
}

func (s *server) AddColumn(ctx context.Context, req *board.Columns) (*board.Columns, error) {
	return &board.Columns{}, nil
}

func (s *server) AddReports(ctx context.Context, req *board.Reports) (*board.Reports, error) {
	return &board.Reports{}, nil
}

func (s *server) CreateBoard(ctx context.Context, req *board.Boards) (*board.Boards, error) {
	return &board.Boards{}, nil
}

func (s *server) AddTask(ctx context.Context, req *board.Tasks) (*board.Tasks, error) {
	return &board.Tasks{}, nil
}

func (s *server) AddSubtask(ctx context.Context, req *board.Subtasks) (*board.Subtasks, error) {
	return &board.Subtasks{}, nil
}

func (s *server) AddComments(ctx context.Context, req *board.Comments) (*board.Comments, error) {
	return &board.Comments{}, nil
}

//////////////////////////////////////////////////

// DELETE
func (s *server) DeleteColumn(ctx context.Context, req *board.GetRequestColumns) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) DeleteReports(ctx context.Context, req *board.GetResponseReport) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) DeleteBoard(ctx context.Context, req *board.GetResponseBoard) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) DeleteTask(ctx context.Context, req *board.GetResponseTask) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) DeleteSubtask(ctx context.Context, req *board.GetResponseSubtasks) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) DeleteComments(ctx context.Context, req *board.GetResponseComments) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

///////////////////////////////////////////////////
