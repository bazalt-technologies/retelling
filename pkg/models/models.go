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
	ReviewID  int
	ContentID int
	UserID    int
	Review    string
	Date      int
}

type Content struct {
	ContentID int
	TypeID    int
	GenreID1  int
	GenreID2  int
	GenreID3  int
	Title     string
	Likes     int
}

type Request struct {
	UserID    int
	UserIDs   []int
	Login     string
	Password  string
	ReviewID  int
	ReviewIDs []int
}
