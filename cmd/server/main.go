package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Task_tracker_back/pkg/db"
	"github.com/Task_tracker_back/pkg/utils"
	"github.com/Task_tracker_back/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	__.UnimplementedUserV1Server
}

// GETUser ПОЛУЧЕНИЕ ПОЛЬЗОВАТЕЛЯ
func (s *server) GetUser(ctx context.Context, req *__.GetRequestUser) (*__.GetResponseUser, error) {
	log.Printf("User id: %d", req.GetIdUser())

	return &__.GetResponseUser{ //TODO: прописать запросы к БД, чтобы возвращали данные из БД.
		IdUser:    req.GetIdUser(),
		Login:     "Test Subject",
		IsManager: false,
	}, nil
}

// Валидация логина и пароля пользователя
// TODO: Добавить запрос к бд, проверяющий пароль и логин
func (s *server) ValidateUser(ctx context.Context, req *__.GetRequestAuth) (*__.GetResponseAuth, error) {
	login := req.GetLogin()
	password := req.GetPassword()

	if login == "" {
		return &__.GetResponseAuth{IsValidated: false}, errors.New("Login invalid")
	}

	if password == "" {
		return &__.GetResponseAuth{IsValidated: false}, errors.New("Password invalid")
	}

	return &__.GetResponseAuth{IsValidated: false}, nil
}

// Изменение email пользователя
// TODO: Запрос к БД (конкретный id, email).
func (s *server) UpdateMail(ctx context.Context, req *__.PutRequestMail) (*__.PutResponseMail, error) {
	idUser := req.GetIdUser()
	mail := req.GetEmail()
	fmt.Println("Email: ", req.GetEmail(), ", UserID: ", idUser)

	if mail == "" {
		return &__.PutResponseMail{}, errors.New("Mail invalid or missing!")
	}

	if utils.ValidateEmail(mail) {
		return &__.PutResponseMail{Email: mail}, nil
	}

	return &__.PutResponseMail{}, errors.New("Mail invalid or missing!")
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %d", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	__.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %d", lis.Addr())
	db.DatabaseConnection()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %d", err)
	}
}
