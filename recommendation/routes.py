from __init__ import api
from flask_restful import Resource, reqparse
from flask import request, jsonify, abort

from models import Request

import requests
import json

import random

parser = reqparse.RequestParser()
parser.add_argument("Login")
parser.add_argument("Password")
parser.add_argument("ObjectID")
parser.add_argument("ObjectIDs")
parser.add_argument("UserID")

class Recommendation(Resource):
    def get(self):

        req = Request()
 
        req.UserID = request.args.get("UserID")       
        
        # Получение лайков пользователя
        r = requests.get('http://server:8081/api/v1/likes', params={"UserID": req.UserID})
        objects = json.loads(r.content.decode())

        if objects is None:
            abort(500)

        # Пользователи, которые лайкнули контент (случайный, из тех который лайкнули мы), такой же как мы
        content = objects[random.randint(0, len(objects) - 1)]
        users = content["UsersLiked"]

        try:
            users.remove(req.UserID)
        except:
            abort(500)

        if len(users) == 0:
            abort(500)

        # Получение пользователя, которые лайкнул такой же контент,
        # Что и изначальный пользователь
        req.UserID = users[random.randint(0, len(users) - 1)]
        r = requests.get('http://server:8081/api/v1/likes', params={"UserID": req.UserID})
        objects = json.loads(r.content.decode())

        if objects is None:
            abort(500)

        req.ObjectIDs = [objects[i]["ID"] for i in range(len(objects))]

        return jsonify(req.ObjectIDs)

api.add_resource(Recommendation, "/api/v1/recommendation")
