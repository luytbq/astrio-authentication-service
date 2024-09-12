package auth

type Region string

type RegisterPayload struct {
	Region         `json:"region"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

type RegisterResponse struct {
	*ErrorResponse `json:"error,omitempty"`
	Region         `json:"regison,omitempty"`
	ID             string `json:"id,omitempty"`
}

type LoginPayload struct {
	*ErrorResponse `json:"error,omitempty"`
	Region         `json:"regison,omitempty"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

type ErrorResponse struct {
	ErrorCode    uint   `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
