from __init__ import SERVER

from models import Request
from flask_restful import Resource

from flask import request, jsonify, abort

import requests
import json

import random

class Recommendation(Resource):
    def get(self):

        req = Request()

        try:
            req.UserID = int(request.args.get("UserID"))
        # Отправлено не число    
        except ValueError:
            abort(415)
        except:
            abort(400)

        # Получение лайков пользователя
        r = requests.get('http://{}:8081/api/v1/likes'.format(SERVER), params={"UserID": req.UserID})
        objects = json.loads(r.content.decode())

        # Пользователя с UserID нету или у него нет likes 
        if objects is None:
            abort(404)

        # Пользователи, которые лайкнули контент (случайный, из тех который лайкнули мы), такой же как мы
        users = []
        stop = 0
        while len(users) <= 1:
            content = objects[random.randint(0, len(objects) - 1)]
            users = content["UsersLiked"]
            stop += 1
            if stop > 50:
                abort(508)

        try:
            users.remove(req.UserID)
        except ValueError:
            abort(500)

        # Получение пользователя, которые лайкнул такой же контент,
        # Что и изначальный пользователь
        req.UserID = users[random.randint(0, len(users) - 1)]
        r = requests.get('http://{}:8081/api/v1/likes'.format(SERVER), params={"UserID": req.UserID})
        objects = json.loads(r.content.decode())

        if objects is None:
            abort(500)

        # req.ObjectIDs = [objects[i]["ID"] for i in range(len(objects))]

        return jsonify(objects)