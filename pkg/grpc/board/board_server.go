package board

import (
	"context"
	"errors"
	board "github.com/Task_tracker_back/pkg/board_v1"
	"github.com/Task_tracker_back/pkg/models"
	"github.com/Task_tracker_back/pkg/utils"
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
	name := req.GetName()

	if utils.CheckEmpty(name) {
		return &board.Columns{}, errors.New("New column name is too short!")
	}

	var column models.Column

	database.First(&column, req.GetIdColumn())

	if utils.CheckEmpty(column.ID) {
		return &board.Columns{}, errors.New("Column ID is invalid!")
	}

	database.Model(&column).Update("name", name)

	return &board.Columns{}, nil
}

func (s *server) PutReports(ctx context.Context, req *board.Reports) (*board.Reports, error) {
	text := req.GetReportText()

	if utils.CheckEmpty(text) {
		return &board.Reports{}, errors.New("New report text is too short!")
	}

	var report models.Report

	database.First(&report, req.GetIdReport())

	if utils.CheckEmpty(report.ID) {
		return &board.Reports{}, errors.New("Report ID is invalid!")
	}

	database.Model(&report).Update("ReportText", text)

	return &board.Reports{}, nil
}

func (s *server) PutBoard(ctx context.Context, req *board.Boards) (*board.Boards, error) {
	name := req.GetName()
	private := req.GetPrivate()

	if utils.CheckEmpty(name) {
		return &board.Boards{}, errors.New("New board name is too short!")
	}

	var boardM models.Board

	database.First(&boardM, req.GetIdBoard())

	if utils.CheckEmpty(boardM.ID) {
		return &board.Boards{}, errors.New("Report ID is invalid!")
	}

	if utils.CheckEmpty(private) {
	}

	database.Model(&boardM).Updates(models.Board{
		Name:    name,
		Private: private,
	})

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
	idRole := req.GetIdRole()
	idBoard := req.GetIdBoard()
	idUser := req.GetIdUser()

	if utils.CheckEmpty(idRole) {
		return &board.UsersBoard{}, errors.New("Role ID field is empty!")
	}

	if utils.CheckEmpty(idBoard) {
		return &board.UsersBoard{}, errors.New("Board ID field is empty!")
	}

	if utils.CheckEmpty(idUser) {
		return &board.UsersBoard{}, errors.New("User ID field is empty!")
	}

	newUserBoard := models.NewUserBoard(idUser, idBoard, idRole)
	database.Create(&newUserBoard)

	return &board.UsersBoard{
		IdRole:  idRole,
		IdBoard: idBoard,
		IdUser:  idUser,
	}, nil
}

func (s *server) AddColumn(ctx context.Context, req *board.Columns) (*board.Columns, error) {
	name := req.GetName()
	idBoard := req.GetIdBoard()

	if utils.CheckEmpty(idBoard) {
		return &board.Columns{}, errors.New("Board ID field is empty!")
	}

	if utils.CheckEmpty(name) {
		return &board.Columns{}, errors.New("Name field is empty!")
	}

	newColumn := models.NewColumn(name, idBoard)
	database.Create(&newColumn)

	return &board.Columns{Name: name, IdBoard: idBoard}, nil
}

func (s *server) AddReports(ctx context.Context, req *board.Reports) (*board.Reports, error) {
	text := req.GetReportText()

	if utils.CheckEmpty(text) {
		return &board.Reports{}, errors.New("Report text field is empty!")
	}

	newReport := models.NewReport(text)
	database.Create(&newReport)

	return &board.Reports{
		IdReport:   uint32(newReport.ID),
		ReportText: text,
	}, nil
}

func (s *server) CreateBoard(ctx context.Context, req *board.Boards) (*board.Boards, error) {
	name := req.GetName()
	private := req.GetPrivate()

	if utils.CheckEmpty(name) {
		return &board.Boards{}, errors.New("Name field is empty!")
	}

	newBoard := models.NewBoard(name, private)
	database.Create(&newBoard)

	newColumn := models.NewColumn("Первая", uint32(newBoard.ID))
	database.Create(&newColumn)

	return &board.Boards{IdBoard: uint32(newBoard.ID), Name: name, Private: private}, nil
}

func (s *server) AddTask(ctx context.Context, req *board.Tasks) (*board.Tasks, error) {
	idColumn := req.GetIdColumn()
	name := req.GetName()
	description := req.GetDescription()

	if utils.CheckEmpty(idColumn) {
		return &board.Tasks{}, errors.New("Column id is missing")
	}

	if utils.CheckEmpty(name) {
		return &board.Tasks{}, errors.New("Name field is missing")
	}

	if utils.CheckEmpty(description) {
		return &board.Tasks{}, errors.New("Description field is missing")
	}

	newTask := models.NewTask(name, description, 1, idColumn)
	database.Create(&newTask)

	return &board.Tasks{
		IdTask:      uint32(newTask.ID),
		Name:        newTask.Name,
		Priority:    uint32(newTask.Priority),
		Description: newTask.Description,
		IdStatus:    1,
		IdColumn:    idColumn,
	}, nil
}

func (s *server) AddSubtask(ctx context.Context, req *board.Subtasks) (*board.Subtasks, error) {
	idTask := req.GetIdTask()
	name := req.GetName()
	description := req.GetDescriptions()

	if utils.CheckEmpty(idTask) {
		return &board.Subtasks{}, errors.New("Task id is missing")
	}

	if utils.CheckEmpty(name) {
		return &board.Subtasks{}, errors.New("Name field is missing")
	}

	if utils.CheckEmpty(description) {
		return &board.Subtasks{}, errors.New("Description field is missing")
	}

	newTask := models.NewTask(name, description, 1, idTask)
	database.Create(&newTask)

	return &board.Subtasks{
		IdTask:       uint32(newTask.ID),
		Name:         newTask.Name,
		Priority:     uint32(newTask.Priority),
		Descriptions: newTask.Description,
		IdStatus:     1,
	}, nil
}

func (s *server) AddComments(ctx context.Context, req *board.Comments) (*board.Comments, error) {
	idUserBoard := req.GetIdUserBoard()
	idTask := req.GetIdTask()
	idSubtask := req.GetIdSubtask()
	idReport := req.GetIdReport()
	comment := req.GetComment()

	if utils.CheckEmpty(comment) {
		return &board.Comments{}, errors.New("Comment text is missing")
	}

	newComment := models.NewComment(idUserBoard, idTask, idSubtask, idReport, comment)
	database.Create(&newComment)

	return &board.Comments{
		IdComment:   uint32(newComment.ID),
		IdUserBoard: newComment.UserBoardID,
		IdSubtask:   newComment.SubtaskID,
		IdTask:      newComment.TaskID,
		IdReport:    newComment.ReportID,
		Comment:     comment,
	}, nil
}

//////////////////////////////////////////////////

// DELETE
func (s *server) DeleteColumn(ctx context.Context, req *board.GetRequestColumns) (*emptypb.Empty, error) {
	idColumn := req.GetIdBoard()

	if utils.CheckEmpty(idColumn) {
		return nil, errors.New("Column ID field is missing")
	}

	database.Delete(&models.Column{}, idColumn)

	return nil, nil
}

func (s *server) DeleteReports(ctx context.Context, req *board.GetResponseReport) (*emptypb.Empty, error) {
	report := req.GetReport()

	if utils.CheckEmpty(report.IdReport) {
		return &emptypb.Empty{}, errors.New("Report ID field is missing")
	}

	database.Where("id = ?", report.IdReport).Delete(&models.Report{})

	return &emptypb.Empty{}, nil
}

func (s *server) DeleteBoard(ctx context.Context, req *board.GetResponseBoard) (*emptypb.Empty, error) {
	idBoard := req.GetBoard().IdBoard

	if utils.CheckEmpty(idBoard) {
		return nil, errors.New("Board ID field is missing")
	}

	database.Delete(&models.Board{}, idBoard)

	return nil, nil
}

func (s *server) DeleteTask(ctx context.Context, req *board.GetResponseTask) (*emptypb.Empty, error) {
	idTask := req.GetTask().IdTask

	if utils.CheckEmpty(idTask) {
		return nil, errors.New("Task ID field is missing")
	}

	database.Delete(&models.Task{}, idTask)

	return nil, nil
}

func (s *server) DeleteSubtask(ctx context.Context, req *board.GetResponseSubtasks) (*emptypb.Empty, error) {
	idSubtask := req.GetSubtasks()

	if utils.CheckEmpty(idSubtask) {
		return nil, errors.New("Subtask ID field is missing")
	}

	database.Delete(&models.Subtask{}, idSubtask)

	return nil, nil
}

func (s *server) DeleteComments(ctx context.Context, req *board.GetResponseComments) (*emptypb.Empty, error) {
	idComment := req.GetComments()

	if utils.CheckEmpty(idComment) {
		return nil, errors.New("Comment ID field is missing")
	}

	database.Delete(&models.Comment{}, idComment)

	return nil, nil
}

///////////////////////////////////////////////////
