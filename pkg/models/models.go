package models

type User struct {
	ID       int
	Login    string
	Password string
	Data     *UserData
}

type UserData struct {
	Name        string
	Age         int
	ReviewCount int
	Rating      int
	Profession  string
}

type Review struct {
	ID     int
	UserID int
	Type   string
	Genre1 string
	Genre2 string
	Genre3 string
	Title  string
	Rating int
	Date   int
	Review string
	Likes  int
}

type Request struct {
	UserID    int
	UserIDs   []int
	Login     string
	Password  string
	ReviewID  int
	ReviewIDs []int
}
