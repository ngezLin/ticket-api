package service

import (
	"errors"
	"strings"
	"ticket-api/entity"
	"ticket-api/repository"
	"ticket-api/utils"
)

type UserService interface {
	Register(user entity.User) (*entity.User, error)
	Login(email, password string) (*entity.User, string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Register(user entity.User) (*entity.User, error) {
	user.Email = strings.ToLower(user.Email)

	// validate existing user
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPwd

	// default role
	if user.Role == "" {
		user.Role = "user"
	}

	err = s.repo.Create(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) Login(email, password string) (*entity.User, string, error) {
	user, err := s.repo.FindByEmail(strings.ToLower(email))
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	//check password
	if err := utils.CheckPasswordHash(password, user.Password); err != nil{
		return nil, "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, "", errors.New("error in generating JWT token")
	}

	return user, token, nil
}