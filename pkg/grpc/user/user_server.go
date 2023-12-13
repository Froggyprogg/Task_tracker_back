package user

import (
	"context"
	"errors"
	"github.com/Task_tracker_back/pkg/models"
	desc "github.com/Task_tracker_back/pkg/user_v1"
	"github.com/Task_tracker_back/pkg/utils"
	"github.com/go-passwd/validator"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"strconv"
)

type server struct {
	desc.UnimplementedUserV1Server
}

var (
	database *gorm.DB
)

func Register(gRPCServer *grpc.Server, db *gorm.DB) {
	desc.RegisterUserV1Server(gRPCServer, &server{})

	database = db
}

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

	if utils.CheckEmpty(login) {
		return &desc.GetResponseAuth{}, errors.New("Login invalid")
	}

	if utils.CheckEmpty(password) {
		return &desc.GetResponseAuth{}, errors.New("Password invalid")
	}

	var user models.User
	database.Where(&models.User{Login: login}).First(&user)

	if utils.CheckEmpty(user.Login) {
		return &desc.GetResponseAuth{}, errors.New("Login incorrect")
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

	if utils.CheckEmpty(user.ID) {
		return &desc.PutResponseMail{}, errors.New("User is not created!")
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
