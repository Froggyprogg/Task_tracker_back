package main

import (
	"context"
	"errors"
	"fmt"
	board "github.com/Task_tracker_back/pkg/board_v1"
	"github.com/Task_tracker_back/pkg/config"
	"github.com/Task_tracker_back/pkg/db"
	"github.com/Task_tracker_back/pkg/models"
	desc "github.com/Task_tracker_back/pkg/user_v1"
	"github.com/Task_tracker_back/pkg/utils"
	"github.com/go-passwd/validator"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type server struct {
	desc.UnimplementedUserV1Server
	board.UnimplementedBoardV1Server
}

var (
	database *gorm.DB
)

// GETUser ПОЛУЧЕНИЕ ПОЛЬЗОВАТЕЛЯ
func (s *server) GetUser(ctx context.Context, req *desc.GetRequestUser) (*desc.GetResponseUser, error) {
	idUser := req.GetIdUser()

	var user models.User
	database.First(&user, idUser)

	if user.ID == 0 {
		return &desc.GetResponseUser{}, errors.New("Invalid User ID!")
	}

	return &desc.GetResponseUser{
		IdUser:    user.ID,
		Login:     user.Login,
		IsManager: user.IsManager,
	}, nil
}

// Валидация логина и пароля пользователя
func (s *server) ValidateUser(ctx context.Context, req *desc.GetRequestAuth) (*desc.GetResponseAuth, error) {
	login := req.GetLogin()
	password := req.GetPassword()

	if login == "" {
		return &__.GetResponseAuth{IsValidated: false}, errors.New("Login invalid")
	}

	if password == "" {
		return &__.GetResponseAuth{IsValidated: false}, errors.New("Password invalid")
	}

	var user models.User
	database.Where(&models.User{Login: login}).First(&user)

	if user.Login == "" {
		return &__.GetResponseAuth{IsValidated: false}, errors.New("Login incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return &desc.GetResponseAuth{}, errors.New("Password incorrect")
	}

	return &desc.GetResponseAuth{IsValidated: true}, nil
}

// Изменение email пользователя
func (s *server) UpdateMail(ctx context.Context, req *desc.PutRequestMail) (*desc.PutResponseMail, error) {
	idUser := req.GetIdUser()
	mail := req.GetEmail()

	if utils.ValidateEmail(mail) {
		return &desc.PutResponseMail{}, errors.New("Mail invalid or missing!")
	}

	var user models.User
	database.First(&user, idUser)

	if user.ID == 0 {
		return &__.PutResponseMail{}, errors.New("User is not created!")
	}

	database.Model(&user).Update("email", mail)
	return &desc.PutResponseMail{Email: mail}, nil
}

func (s *server) CreateUser(ctx context.Context, req *desc.PostRequestUser) (*desc.PostResponseUser, error) {
	login := req.GetLogin()
	password := req.GetPassword()
	mail := req.GetEmail()
	isManager := req.GetIsManager()

	if utils.ValidateEmail(mail) == false {
		return &desc.PostResponseUser{}, errors.New("Mail invalid or missing!")
	}

	if utils.CheckEmpty(login) {
		return &desc.PostResponseUser{}, errors.New("Login invalid or missing!")
	}

	if utils.CheckEmpty(password) {
		return &desc.PostResponseUser{}, errors.New("Password invalid or missing!")
	}

	var user models.User
	database.Where(&models.User{Login: login}).Or(&models.User{Email: mail}).First(&user)

	if utils.CheckEmpty(user.ID) {
		return &desc.PostResponseUser{}, errors.New("Login or email is already taken or exists!")
	}

	passwordValidator := validator.New(validator.MinLength(8, errors.New("Password is too short")), validator.MaxLength(32, errors.New("Password is too long")))
	err := passwordValidator.Validate(password)
	if err != nil {
		return &desc.PostResponseUser{}, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		return &desc.PostResponseUser{}, errors.New("Hash error")
	}

	newUser := models.NewUser(login, string(hashed), mail, isManager) //models.User{}
	database.Create(&newUser)

	return &desc.PostResponseUser{IdUser: strconv.Itoa(int(newUser.ID))}, nil
}

func main() {
	cfg := config.MustLoad()

	database = db.NewDatabaseConnection(cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %d", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})
	board.RegisterBoardV1Server(s, &server{})

	log.Printf("server listening at %d", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %d", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	log.Print("Gracefully stopped")
}
