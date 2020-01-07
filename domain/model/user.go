package model

import "time"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`     //名前
	Password string `json:"password"` //パスワード
	Email    string `json:"email"`    //メールアドレス
	Favorite string `json:"favorite"` //お気に入り本
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
