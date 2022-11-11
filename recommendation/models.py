class User:
    def __init__(self):
        self.ID = 0
        self.Login = ""
        self.Password = ""
        self.Data = UserData()


class UserData:
    def __init__(self):
        self.Name
        self.Age
        self.ReviewCount
        self.Rating
        self.Profession
        self.Likes


class Review:
    def __init__(self):
        self.ID
        self.ContentID
        self.UserID
        self.Review
        self.Date


class User:
    def __init__(self):
        self.ID
        self.TypeID
        self.GenreID1
        self.GenreID2
        self.GenreID3
        self.Title
        self.Description
        self.UserLiked


class Request:
    def __init__(self):
        self.Login = ""
        self.Password = ""
        self.ObjectID = 0
        self.ObjectIDs = []
        self.UserID = 0


class Type:
    def __init__(self):
        self.ID
        self.Type


class Genre:
    def __init__(self):
        self.ID
        self.Genre
