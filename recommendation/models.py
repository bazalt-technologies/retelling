class Request:
    Login: str = ""
    Password: str = ""
    ObjectID: int = 0
    ObjectIDs: list = []
    UserID: int = 0

    def __repr__(self):
        return self.UserID


class Content:
    ID: int = 0
    TypeID: int = 0
    GenreID1: int = 0
    GenreID2: int = 0
    GenreID3: int = 0
    Title: str = ""
    Description: str = ""
    UsersLiked: list = []

    def __repr__(self):
        return self.ID
