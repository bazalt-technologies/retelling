import requests
import random
import json


# Заполнение тестовым контентом
url = 'http://localhost:8081/api/v1/content'
for _ in range(30):
    data = {
        "typeid" : random.randint(1, 4),
        "genreid1" : random.randint(1, 7),
        "genreid2" : random.randint(1, 7),
        "genreid3" : random.randint(1, 7),
        "title" : "TEST_TITLE",
        "description" : "TEST_DESCRIPTION"

        #"data" : {
        #    "name" : "DAAD",
        #    "age" : 213
        #}
    }

    r = requests.post(url, json = data)

    print(r.content.decode())


# Лайки тестовые
url = 'http://localhost:8081/api/v1/likes'
for _ in range(30):
    data = {
        "userid" : random.randint(1, 5),
        "objectid" : random.randint(1, 30) 
    }

    r = requests.post(url, json = data)

    print(r.content.decode())

'''
url = 'http://localhost:8081/api/v1/likes'

data = {
    "UserID" : 7
}

# Получение лайков пользователя
r = requests.get(url, params = data)

data = json.loads(r.content.decode())

# Контент, который пользователь лайкнул
usersLiked = data[random.randint(0, len(data) - 1)]["UsersLiked"]

print("Контент, который пользователь лайкнул")
print(usersLiked)

data = {
    "ObjectID" : usersLiked[random.randint(0, len(usersLiked) - 1)]
}

# Получение пользователей, которые лайкнули такой же контент,
# Что и изначальный пользователь
r = requests.get(url, params = data)

data = json.loads(r.content.decode())

contentId = data[random.randint(0, len(data) - 1)]["Data"]["Likes"]

recommendation = list(set(contentId) - set(usersLiked))

print("Рекомендованный контент")
print(recommendation)
'''