from flask import Flask
from flask import jsonify
from flask import request

import requests
import json
import random

app = Flask(__name__)


@app.route('/', methods=['GET'])
def returnAll():
    userId1 = request.get_json()["UserID"]
    
    # Получение лайков пользователя
    r = requests.get('http://localhost:8081/api/v1/likes', params = {"UserID" : userId1})
    objects = json.loads(r.content.decode())

    if objects == None:
        return (jsonify(None))

    # Пользователи, которые лайкнули контент (случайный, из тех который лайкнули мы), такой же как мы
    objectId = objects[random.randint(0, len(objects) - 1)]
    users = objectId["UsersLiked"]
    users.remove(userId1)

    if len(users) == 0:
        return (jsonify(None))
    
    # Получение пользователя, которые лайкнул такой же контент,
    # Что и изначальный пользователь
    userId2 = users[random.randint(0, len(users) - 1)]
    r = requests.get('http://localhost:8081/api/v1/likes', params = {"UserID" : userId2})
    objects = json.loads(r.content.decode())

    if objects == None:
        return (jsonify(None))

    objectIds = [objects[i]["ID"] for i in range(len(objects))]

    return jsonify(objectIds)
