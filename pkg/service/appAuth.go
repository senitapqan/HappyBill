package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"happyBill/models"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	Roles  []models.RolesHeaders
}

func (s *service) hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (s *service) GenerateToken(input, password string) ([]models.RolesHeaders, string, error) {
	password = s.hashPassword(password)
	user, err := s.repos.GetUser(input)

	if err != nil {
		return nil, "", err
	}

	if user.Id == 0 {
		return nil, "", fmt.Errorf("there is no such user with username/email: %s", input)
	}

	if user.Password != password {
		return nil, "", errors.New("incorrect password")
	}

	var rolesHeaders []models.RolesHeaders
	roles, err := s.repos.GetRoles(user.Id)

	if err != nil {
		return nil, "", err
	}

	for _, role := range roles {
		id, err := s.repos.GetRoleId(role, user.Id)
		if err != nil {
			return nil, "", err
		}
		rolesHeaders = append(rolesHeaders, models.RolesHeaders{Role: role, Id: id})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		rolesHeaders,
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	return rolesHeaders, tokenString, err
}

func (s *service) ParseToken(accessToken string) (int, []models.RolesHeaders, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, nil, errors.New("token claims are not type of *tokenClaims")
	}

	return claims.UserId, claims.Roles, nil
}
