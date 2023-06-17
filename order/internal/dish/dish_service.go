package dish

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
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

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func (s *service) GetDish(c context.Context, id int64) (*Dish, error) {
	d, err := s.Repository.GetDish(c, id)
	if err != nil {
		return &Dish{}, err
	}
	return d, nil
}

func (s *service) CreateDish(c context.Context, req *CreateDishReq) (*Dish, error) {
	dish, err := s.Repository.CreateDish(c, req)
	if err != nil {
		return &Dish{}, err
	}
	return dish, nil
}
func (s *service) UpdateDish(c context.Context, req *UpdateDishReq) (*Dish, error) {
	dish, err := s.Repository.UpdateDish(c, req)
	if err != nil {
		return &Dish{}, err
	}
	return dish, nil
}
func (s *service) DeleteDish(c context.Context, id int64) (*Dish, error) {
	d, err := s.Repository.GetDish(c, id)
	if err != nil {
		return &Dish{}, err
	}
	return d, nil
}

func (s *service) GetAll(c context.Context) ([]Dish, error) {
	dishes, err := s.Repository.GetAll(c)
	if err != nil {
		return dishes, err
	}
	return dishes, nil
}

func (s *service) GetClaimsFromToken(tokenString string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Print(err)
		log.Print(token.Claims)
		return nil, err
	}
	log.Print(token.Valid)
	return &claims, nil
}

func (s *service) CheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(401, err)
			c.Abort()
			return
		}
		claims, err := s.GetClaimsFromToken(token)
		if err != nil {
			c.JSON(500, err)
			c.Abort()
			return
		}
		if claims.Role != role {
			c.JSON(401, err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func (s *service) UseId() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(401, err)
			c.Abort()
			return
		}

		claims, err := s.GetClaimsFromToken(token)
		if err != nil {
			c.JSON(500, err)
			c.Abort()
			return
		}
		c.Set("user_id", claims.ID)
		c.Next()
	}
}
