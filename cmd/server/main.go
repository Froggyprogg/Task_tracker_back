package main

import (
	"context"
	"errors"
	"fmt"
	desc "github.com/Task_tracker_back/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

// GETUser ПОЛУЧЕНИЕ ПОЛЬЗОВАТЕЛЯ
func (s *server) GetUser(ctx context.Context, req *desc.GetRequestUser) (*desc.GetResponseUser, error) {
	log.Printf("User id: %d", req.GetIdUser())

	return &desc.GetResponseUser{ //TODO: прописать запросы к БД, чтобы возвращали данные из БД.
		IdUser:    req.GetIdUser(),
		Login:     "Test Subject",
		IsManager: false,
	}, nil
}

// Валидация логина и пароля пользователя
// TODO: Добавить запрос к бд, проверяющий пароль и логин
func (s *server) ValidateUser(ctx context.Context, req *desc.GetRequestAuth) (*desc.GetResponseAuth, error) {
	if req.GetLogin() == "" {
		return &desc.GetResponseAuth{IsValidated: false}, errors.New("Login invalid")
	}

	if req.GetPassword() == "" {
		return &desc.GetResponseAuth{IsValidated: false}, errors.New("Password invalid")
	}

	return &desc.GetResponseAuth{IsValidated: false}, nil
}

// Изменение email пользователя
// TODO: Запрос к БД (конкретный id, email).
func (s *server) UpdateMail(ctx context.Context, req *desc.PutRequestMail) (*desc.PutResponseMail, error) {
	idUser := req.GetIdUser()
	fmt.Println("Email: ", req.GetEmail(), ", UserID: ", idUser)
	if req.GetEmail() == "" {
		return &desc.PutResponseMail{}, errors.New("Mail invalid or missing!")
	}

	return &desc.PutResponseMail{Email: req.GetEmail()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %d", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %d", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %d", err)
	}
}
