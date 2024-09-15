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
	Email          string `json:"email,omitempty"`
}

type LoginPayload struct {
	*ErrorResponse `json:"error,omitempty"`
	Region         `json:"regison,omitempty"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

type LoginResponse struct {
	*ErrorResponse `json:"error,omitempty"`
}

type GeneralResponse struct {
	*ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    uint   `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
