package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luytbq/astrio-authentication-service/config"
	pauth "github.com/luytbq/astrio-authentication-service/pkg/auth"
	"github.com/luytbq/astrio-authentication-service/pkg/common"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	path := config.App.SERVER_API_PREFIX + "/api/v1/users"
	r.POST(path, h.handleRegister)
	// r.GET(path+"/:id", h.GetUserByID)
	// r.POST(path+"/login", h.handleLogin)
}

func (h *Handler) handleRegister(c *gin.Context) {
	var payload pauth.RegisterPayload
	if err := c.BindJSON(&payload); err != nil {
		common.Response(c, http.StatusBadRequest,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{ErrorCode: 400, ErrorMessage: "invalid payload"}})
		return
	}

	// validate payload
	if err := validateRegisterPayload(payload); err != nil {
		common.Response(c, http.StatusBadRequest,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{ErrorCode: http.StatusBadRequest, ErrorMessage: err.Error()}})
		return
	}

	// check if email existed
	user, err := h.repo.GetUserByEmail(nomarlizeEmail(payload.Email))
	if err != nil {
		log.Printf("error: %s", err.Error())
		common.Response(c, http.StatusInternalServerError,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{ErrorCode: http.StatusInternalServerError, ErrorMessage: "some thing went wrong"}})
		return
	}

	if user != nil {
		common.Response(c, http.StatusBadRequest,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{ErrorCode: http.StatusBadRequest, ErrorMessage: "user existed"}})
		return
	}

	hashPassword, salt := HashPassword(payload.Password)

	newUser := User{Email: payload.Email, Password: hashPassword, PasswordSalt: salt}

	// create user
	err = h.repo.InsertUser(&newUser)
	if err != nil || newUser.ID == 0 {
		log.Printf("error: %s", err.Error())
		log.Printf("error: %d", newUser.ID)
		common.Response(c, http.StatusInternalServerError,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{ErrorCode: http.StatusInternalServerError, ErrorMessage: "some thing went wrong"}})
		return
	}

	c.JSON(http.StatusCreated, newUser)
	// common.Response(c, http.StatusOK, &pauth.RegisterResponse{ID: "12345"})
}
