package model

type UserRole struct {
	ID          int    `json:"id_rl_user_role"`
	Descripcion string `json:"descripcion"`
}

type User struct {
	ID           int    `json:"id_rl_user"`
	Username     string `json:"username"`
	UserPassword string `json:"user_password"`
	UserRoleID   int    `json:"id_rl_user_role"`
}
