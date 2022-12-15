from __init__ import SERVER

from models import Request
from flask_restful import Resource

from flask import request, jsonify, abort

import requests
import json

import random
import math

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
            return []
        r = requests.get('http://{}:8081/api/v1/genres'.format(SERVER))
        NumbersOFGenres= json.loads(r.content.decode())
        # Считаем сколько количество жанров лайкнутых пользователем, для построения соотношений
        genres=[0 for i in range(len(NumbersOFGenres))] #Считаем сколько количество жанров лайкнутых пользователем, для построения соотношений
        user_likes=[]
        # запись в список ID объектов ,которые лайкнул пользователь
        for j in range(len(objects)):
            user_likes.append(objects[j]["ID"])
            if objects[j]["GenreID1"]!=0:
                genres[objects[j]["GenreID1"]-1]+=1
            if objects[j]["GenreID2"] != 0:
                genres[objects[j]["GenreID2"] - 1] += 1
            if objects[j]["GenreID3"] != 0:
                genres[objects[j]["GenreID3"] - 1] += 1

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
            #Нужное количество схожих пользователей
            n = len(users) // 2
            if n>50:
                n=50
            # Коэфицент для сравнивания
            SimilarCoef = 0.9
            # словари для записи лайков других пользователей и их коэффицент Танимото
            DictionaryLikes = {}
            DictionaryCoefs = {}
            # Словарь, в котором хранится ID контента , и его жанры
            DictionaryGenres={}
            #Словарь, в котором хранится ID юзера , и кол-во его жанров для дальнейших рекомендаций
            AmmountGenres={}
            DictContent={}
            for i in range(len(users)):
                r = requests.get('http://{}:8081/api/v1/likes'.format(SERVER), params={"UserID": users[i]})
                objects2 = json.loads(r.content.decode())
                if objects2 is None:
                    continue
                genres2 = []
                k=0
                user_likes2=[]
                #Записываем лайки другого пользователя и кол-во его жанров
                for j in range(len(objects2)):
                    user_likes2.append(objects2[j]["ID"])
                    if objects2[j]["ID"] not in DictionaryGenres:
                        if objects2[j]["GenreID1"] != 0:
                            genres2.append(objects2[j]["GenreID1"])
                            k+=1
                        if objects2[j]["GenreID2"] != 0:
                            genres2.append(objects2[j]["GenreID2"])
                            k+=1
                        if objects2[j]["GenreID3"] != 0:
                            genres2.append(objects2[j]["GenreID3"])
                            k+=1
                        DictionaryGenres[objects2[j]["ID"]] = genres2
                    if objects2[j]["ID"] not in DictContent:
                        DictContent[objects2[j]["ID"]]=objects2[j]
                AmmountGenres[users[i]] = k
                # лайки другого пользователя
                #Общие лайки
                SameLikes = list(set(user_likes2) - set(user_likes))
                coef = len(SameLikes) / (len(user_likes) + len(user_likes2) - len(SameLikes))
                DictionaryLikes[users[i]] = SameLikes
                DictionaryCoefs[users[i]] = coef
            #сравнение схожести пользователей
            if len(DictionaryLikes)==0:
                abort(500)
            c=[]
            NumbersOfGenres=0
            k=0
            while SimilarCoef >= 0.3 and k < n:
                for i in range(len(users)):
                    if users[i] in DictionaryCoefs.keys():
                        if DictionaryCoefs.get(users[i]) >= SimilarCoef:
                            c.append(DictionaryLikes.get(users[i]))
                            k += 1
                            NumbersOfGenres+=AmmountGenres.get(users[i])
                            del DictionaryCoefs[users[i]]
                            del DictionaryLikes[users[i]]
                            del AmmountGenres[users[i]]
                SimilarCoef -= 0.2
            #Если не выбралось нужное количество пользователей
            if k < n:
                for j in range(n - k):
                    key = random.choice(list(DictionaryLikes.keys()))
                    c.append(DictionaryLikes.get(key))
                    NumbersOfGenres += AmmountGenres.get(key)
                    del DictionaryCoefs[key]
                    del DictionaryLikes[key]
                    del AmmountGenres[key]
            num=sum(genres)
            genres3=[]
            #Определяем какое кол-во жанров нам нужно рекомендовать пользователю
            for i in range(len(genres)):
                genres3.append(math.trunc(NumbersOfGenres*genres[i]//num))
            for i in range(len(c)):
                for j in range(len(c[i])):
                    #Если контент еще не рекомендован
                    if (c[i])[j] in DictionaryGenres:
                        array=DictionaryGenres.get(c[i][j])
                        zxc=False
                        for r in range (len(array)):
                            # Если есть нужные нам жанры добавляем в рекомендации
                            if genres3[array[r]-1]!=0:
                                genres3[array[r]-1]-=1
                                zxc=True
                        if zxc==True:
                            req.ObjectIDs.append(c[i][j])

            return jsonify([DictContent.get(req.ObjectIDs[i]) for i in range(len(req.ObjectIDs))])
        else:
            r = requests.get('http://{}:8081/api/v1/likes'.format(SERVER), params={"UserID": users[0]})
            objects = json.loads(r.content.decode())

        return objects

