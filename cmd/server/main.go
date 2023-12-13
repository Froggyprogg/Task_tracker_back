package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Task_tracker_back/pkg/config"
	"github.com/Task_tracker_back/pkg/db"
	"github.com/Task_tracker_back/pkg/models"
	"github.com/Task_tracker_back/pkg/utils"
	"github.com/Task_tracker_back/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	__.UnimplementedUserV1Server
}

var (
	database *gorm.DB
)

// GETUser ПОЛУЧЕНИЕ ПОЛЬЗОВАТЕЛЯ
func (s *server) GetUser(ctx context.Context, req *__.GetRequestUser) (*__.GetResponseUser, error) {
	idUser := req.GetIdUser()
	log.Printf("User id: %d", idUser)

	var user models.User
	database.First(&user, idUser)

	if user.ID == 0 {
		return &__.GetResponseUser{}, errors.New("Invalid User ID!")
	}

	return &__.GetResponseUser{
		IdUser:    user.ID,
		Login:     user.Login,
		IsManager: user.IsManager,
	}, nil
}

// Валидация логина и пароля пользователя
func (s *server) ValidateUser(ctx context.Context, req *__.GetRequestAuth) (*__.GetResponseAuth, error) {
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

	if user.Password != password {
		return &__.GetResponseAuth{IsValidated: false}, errors.New("Password incorrect")
	}

	return &__.GetResponseAuth{IsValidated: true}, nil
}

// Изменение email пользователя
func (s *server) UpdateMail(ctx context.Context, req *__.PutRequestMail) (*__.PutResponseMail, error) {
	idUser := req.GetIdUser()
	mail := req.GetEmail()
	fmt.Println("Email: ", req.GetEmail(), ", UserID: ", idUser)

	if mail == "" {
		return &__.PutResponseMail{}, errors.New("Mail invalid or missing!")
	}

	var user models.User

	database.First(&user, idUser)

	if user.ID == 0 {
		return &__.PutResponseMail{}, errors.New("User is not created!")
	}

	if utils.ValidateEmail(mail) {
		database.Model(&user).Update("email", mail)
		return &__.PutResponseMail{Email: mail}, nil
	}

	return &__.PutResponseMail{}, errors.New("Mail invalid or missing!")
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
	__.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %d", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %d", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	log.Print("Gracefully stopped")
}
