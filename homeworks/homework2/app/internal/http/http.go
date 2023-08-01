package http

import (
	"errors"
	"go-mongo-crud-rest-api/internal/model"
	repository2 "go-mongo-crud-rest-api/internal/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	repository        repository2.Repository
	elasticRepository repository2.Repository
}

func NewServer(repository repository2.Repository, elasticRepository repository2.Repository) *Server {
	return &Server{repository: repository, elasticRepository: elasticRepository}
}

func (s Server) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (s Server) GetUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}
	user, err := s.repository.GetUser(ctx, email)
	if err != nil {
		log.Printf("mongo GetUser: error: %s", err.Error())
		if errors.Is(err, repository2.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = s.elasticRepository.GetUser(ctx, email)
	if err != nil {
		log.Printf("elasticRepository GetUser: error: %s", err.Error())
		if errors.Is(err, repository2.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := s.elasticRepository.CreateUser(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) UpdateUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user.Email = email
	user, err := s.repository.UpdateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository2.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := s.elasticRepository.UpdateUser(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) DeleteUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}
	if err := s.repository.DeleteUser(ctx, email); err != nil {
		if errors.Is(err, repository2.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := s.elasticRepository.DeleteUser(ctx, email); err != nil {
		if errors.Is(err, repository2.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
