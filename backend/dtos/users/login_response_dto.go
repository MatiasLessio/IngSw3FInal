package dto

type LoginResponseDto struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
