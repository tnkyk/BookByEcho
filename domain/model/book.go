package model

type Book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`      //本のタイトル
	Author     string `json:"author"`     //本の著者
	UserID     string `json:"user_id"`    //本を持っているユーザー
	Reservable string `json:"reservable"` //予約可能かどうか
}
