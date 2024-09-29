package auth

import (
	"log"
	"net/http"
	"strings"

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
	r.POST(path+"/login", h.handleLogin)
	r.GET(path+"/verify", h.handleVerify)
}

func (h *Handler) handleRegister(c *gin.Context) {
	var payload pauth.RegisterPayload
	if err := c.BindJSON(&payload); err != nil {
		common.Response(c, http.StatusBadRequest,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{Code: 400, Message: "invalid payload"}})
		return
	}

	payload.Email = nomarlizeEmail(payload.Email)

	// validate payload
	if err := validateRegisterPayload(payload); err != nil {
		common.Response(c, http.StatusBadRequest,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()}})
		return
	}

	// check if email existed
	user, err := h.repo.GetUserByEmail(payload.Email)
	if err != nil {
		log.Printf("error: %s", err.Error())
		common.Response(c, http.StatusInternalServerError,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{Code: http.StatusInternalServerError, Message: "some thing went wrong"}})
		return
	}

	if user != nil {
		common.Response(c, http.StatusBadRequest,
			pauth.RegisterResponse{
				ErrorResponse: &pauth.ErrorResponse{Code: http.StatusBadRequest, Message: "user existed"},
				Email:         payload.Email,
			})
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
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{Code: http.StatusInternalServerError, Message: "some thing went wrong"}})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (h *Handler) handleLogin(c *gin.Context) {
	var payload pauth.LoginPayload
	if err := c.BindJSON(&payload); err != nil {
		common.Response(c, http.StatusBadRequest,
			pauth.LoginResponse{ErrorResponse: &pauth.ErrorResponse{Code: 400, Message: "invalid payload"}})
		return
	}

	payload.Email = nomarlizeEmail(payload.Email)

	user, err := h.repo.GetUserByEmail(payload.Email)
	if err != nil {
		log.Printf("error: %s", err.Error())
		common.Response(c, http.StatusInternalServerError,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{Code: http.StatusInternalServerError, Message: "some thing went wrong"}})
		return
	}

	if user == nil {
		common.Response(c, http.StatusUnauthorized,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{Message: "invalid email or password"}})
		return
	}

	verified := VerifyPassword(payload.Password, user.Password, user.PasswordSalt)

	if !verified {
		common.Response(c, http.StatusUnauthorized,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{Message: "invalid email or password"}})
		return
	}

	token, err := CreateJWTToken(map[string]any{
		"user_id":    user.ID,
		"user_email": user.Email,
	})

	if err != nil {
		log.Printf("error: %s", err.Error())
		common.Response(c, http.StatusInternalServerError,
			pauth.RegisterResponse{ErrorResponse: &pauth.ErrorResponse{Code: http.StatusInternalServerError, Message: "some thing went wrong"}})
		return
	}

	c.Header(KEY_AUTH_TOKEN, "Bearer "+token)
	c.Status(http.StatusOK)
}

func (h *Handler) handleVerify(c *gin.Context) {
	authHeader := c.Request.Header.Get(KEY_AUTH_TOKEN)
	if authHeader == "" {
		log.Printf("auth token is empty")
		common.Response(c, http.StatusBadRequest,
			pauth.GeneralResponse{ErrorResponse: &pauth.ErrorResponse{Code: http.StatusBadRequest, Message: "Request header missing: " + KEY_AUTH_TOKEN}})
		return

	}

	token := strings.ReplaceAll(authHeader, "Bearer ", "")

	claims, err := ParseJWTToken(token)

	if err != nil {
		log.Printf("error: %s", err.Error())
		common.Response(c, http.StatusBadRequest,
			pauth.GeneralResponse{ErrorResponse: &pauth.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()}})
		return
	}

	c.JSON(http.StatusOK, gin.H(claims))
}
