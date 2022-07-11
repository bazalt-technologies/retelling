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
	TypeID	int
	Genre1ID  int
	Genre2ID  int
	Genre3ID  int
	Title  string
	Rating int
	Date   int
	Review string
	ImageLink string
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

type Type struct {
	ID	int
	Type 	string
}

type Genre struct {
	ID	int
	Genre	string
}
