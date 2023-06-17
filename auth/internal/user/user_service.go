package user

import (
	"auth/util"
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		Id:       strconv.Itoa(int(r.Id)),
		Username: r.Username,
		Email:    r.Email,
		Role:     r.Role,
	}

	return res, nil
}

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}
	expiresAt := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID:       u.Id,
		Username: u.Username,
		Role:     u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.Id)),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})

	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return &LoginUserRes{}, err
	}
	session := Session{
		UserId:       u.Id,
		ExpiresAt:    expiresAt,
		SessionToken: ss,
	}
	err = s.Repository.CreateSession(c, &session)
	if err != nil {
		return nil, err
	}
	return &LoginUserRes{accessToken: ss, Username: u.Username, Id: strconv.Itoa(int(u.Id))}, nil
}

func (s *service) GetClaimsFromToken(tokenString string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("JWT_SECRET"), nil
	})
	if err != nil {
		log.Print(token.Claims.(jwt.Claims))
		return nil, err
	}
	log.Print(token.Valid)
	return &claims, nil
}
