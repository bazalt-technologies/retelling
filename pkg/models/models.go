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
	Login     string
	Password  string
	ObjectID  int
	ObjectIDs []int
}

type Type struct {
	ID	int
	Type 	string
}

type Genre struct {
	ID	int
	Genre	string
}
