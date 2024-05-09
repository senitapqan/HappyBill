package service

import (
	"errors"
	"happyBill/dtos"
	"happyBill/models"
	"strings"

	"github.com/rs/zerolog/log"
)

func (s *service) CreateClient(client models.User) (int, error) {
	client.Password = s.hashPassword(client.Password)
	log.Info().Msg("service send request to repository: create client request")
	client_id, err := s.repos.CreateClient(client)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "duplicate") {
			if strings.Contains(msg, "email") {
				return -1, errors.New("there is already exist account with such email")
			}
			if strings.Contains(msg, "username") {
				return -1, errors.New("there is already exist account with such username")
			}
		}
		return 0, errors.New("error with repository")
	}
	return client_id, err
}

func (s service) DeleteClient(lessonId int) error {
	return nil
}

func (s service) GetClientById(id int) (dtos.User, error) {
	log.Info().Msg("service send request to repository: get client by id request")
	return s.repos.GetClientById(id)
}

func (s service) GetClientByUserId(id int) (dtos.User, error) {
	log.Info().Msg("service send request to repository: get client by userId request")
	user, err := s.repos.GetClientByUserId(id)
	if err != nil {
		return user, err
	}
	if user.Phone == nil {
		user.Phone = ""
	}
	return user, err
}

func (s service) UpdateMyProfile(userId int, input dtos.UpdateUser) error {
	log.Info().Msg(input.OldPassword)
	input.OldPassword = s.hashPassword(input.OldPassword)
	if input.Password != "" {
		input.Password = s.hashPassword(input.Password)
	}
	log.Info().Msg("service send request to repository: get user info by id request")
	client, _ := s.repos.GetUserById(userId)

	if client.Password != input.OldPassword {
		log.Info().Msg(client.Password + " " + input.OldPassword)
		return errors.New("incorrect password")
	}

	log.Info().Msg("service send request to repository: update profile request")
	err := s.repos.UpdateMyProfile(userId, input)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "duplicate") && strings.Contains(msg, "username") {
			return errors.New("there is already exist account with such username")
		}
		return err
	}

	return nil
}
