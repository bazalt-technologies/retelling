from flask_restful import abort, Api, Resource, reqparse

from flask import Flask
from flask import jsonify
from flask import request

import requests
import json
import random

import models

app = Flask(__name__)
api = Api(app)

parser = reqparse.RequestParser()
parser.add_argument('UserID', type=int, help='User id')

class Recommendation(Resource):
    def get(self, UserId):
      
        args = parser.parse_args()

        print("---------")
        print(args)
        print("--------")

        '''
        req = models.Request()
        # req.UserID = request.get_json()["UserID"]
        req.UserID = 1
        
        
        # Получение лайков пользователя
        r = requests.get('http://server:8081/api/v1/likes', params={"UserID": req.UserID})
        objects = json.loads(r.content.decode())

        if objects is None:
            abort(500)

        # Пользователи, которые лайкнули контент (случайный, из тех который лайкнули мы), такой же как мы
        content = objects[random.randint(0, len(objects) - 1)]
        users = content["UsersLiked"]
        users.remove(req.UserID)

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
        '''

api.add_resource(Recommendation, '/api/v1/recommendation')

if __name__ == "__main__":
    app.run(debug=True)
