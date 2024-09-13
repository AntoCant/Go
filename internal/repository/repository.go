package repository

import (
	"Api-users/internal/model"
	"Api-users/pkg/db"
)

func GetAllUsers() ([]model.User, error) {
	rows, err := db.Db.Query("SELECT id_rl_user, username, user_password, id_rl_user_role FROM rl_api_users.rl_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.UserPassword, &user.UserRoleID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(user model.User) error {
	_, err := db.Db.Exec("INSERT INTO rl_api_users.rl_user (username, user_password, id_rl_user_role) VALUES ($1, $2, $3)",
		user.Username, user.UserPassword, user.UserRoleID)
	return err
}
