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
        # запись в список ID объектов ,которые лайкнул пользователь
        user_likes = [objects[j]["ID"] for j in range(len(objects))]
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

        if len(users) > 1:
            #коэффициент Танимото
            coef = 0
            k = 0
            #Нужное количество схожих пользователей
            n = len(users) // 2
            # Коэфицент для сравнивания
            SimilarCoef = 0.9
            # словари для записи лайков других пользователей и их коэффицент Танимото
            DictionaryLikes = {}
            DictionaryCoefs = {}
            for i in range(len(users)):
                r = requests.get('http://{}:8081/api/v1/likes'.format(SERVER), params={"UserID": users[i]})
                objects2 = json.loads(r.content.decode())
                if objects2 is None:
                    continue
                #лайки другого пользователя
                user_likes2 = [objects2[j]["ID"] for j in range(len(objects2))]
                #Общие лайки
                SameLikes = (set(user_likes2) - set(user_likes))
                coef = len(SameLikes) / (len(user_likes) + len(user_likes2) - len(SameLikes))
                DictionaryLikes[users[i]] = SameLikes
                DictionaryCoefs[users[i]] = coef
            #сравнение схожести пользователей
            while SimilarCoef != 0.3 and k < n:
                for i in range(len(users)):
                    if DictionaryCoefs.get(users[i]) >= SimilarCoef:
                        req.ObjectIDs += DictionaryLikes.get(users[i])
                        k += 1
                        del DictionaryCoefs[users[i]]
                        del DictionaryLikes[users[i]]
                SimilarCoef -= 0.2
            #Если не выбралось нужное количество пользователей
            if k < n:
                for j in range(n - k):
                    key = random.choice(list(DictionaryLikes.keys()))
                    req.ObjectIDs += DictionaryLikes.get(users[key])
                    del DictionaryLikes[key]
        else:
            r = requests.get('http://{}:8081/api/v1/likes'.format(SERVER), params={"UserID": users[0]})
            objects = json.loads(r.content.decode())
            if objects is None:
                abort(500)
            req.ObjectIDs = [objects[i]["ID"] for i in range(len(objects))]
        return jsonify(list(set(req.ObjectIDs)))

