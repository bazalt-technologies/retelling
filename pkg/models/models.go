package models

type User struct {
	ID       int
	Name     string
	Login    string
	Password string
	Age      int
}

type Review struct {
	ID     int
	UserID int
	Type   string
	Genre  string
	Title  string
	Rating int
	Date   int
	Review string
	Likes  int
}

type Request struct {
	UserID    int
	UserIDs   []int
	ReviewID  int
	ReviewIDs []int
}
