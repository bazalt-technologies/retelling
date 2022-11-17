from __init__ import api, SERVER
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

# Затычка аналитики (перенести в отдельный модуль)
class Analytics(Resource):
    def get(self):
            
        req = Request()
        args = request.args
        
        try:
            if args.get("ObjectID") != None:
                req.ObjectID = int(args.get("ObjectID"))
            elif args.get("UserID") != None:
                req.UserID = int(args.get("UserID"))
            else:
                abort(404)
        except TypeError:
            abort(415)
        except ValueError:
            abort(415)
        except:
            abort(400)

        # Данные по пользователям (затычка)
        if req.UserID != 0:
            r = requests.get('http://{}:8081/api/v1/users'.format(SERVER), params={"ObjectID": req.UserID})
            users = json.loads(r.content.decode())
            if users is None:
                abort(500)

            return jsonify(users)

        # Данные по объектам (затычка)
        '''
        elif req.ObjectID != 0:
            r = requests.get('http://{}:8081/api/v1/content'.format(SERVER), params={"ObjectID": req.ObjectID})
            objects = json.loads(r.content.decode())
            if objects is None:
                abort(500)

            return jsonify(objects)
        '''
        
api.add_resource(Recommendation, "/api/v1/recommendation")
api.add_resource(Analytics, "/api/v1/analytics")
