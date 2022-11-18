from __init__ import SERVER, cache

from models import Request
from flask_restful import Resource

from flask import request, jsonify, abort

import requests
import json


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
            got = cache.get(req.UserID)
            if got == None:
                r = requests.get('http://{}:8081/api/v1/users'.format(SERVER), params={"ObjectID": req.UserID})
                users = json.loads(r.content.decode())

                if users is None:
                    abort(500)

                got = cache.set(req.UserID, str(users), ex=300)
                if got == None:
                    pass
                else:
                    pass
            else:
                users = got.decode()

            return jsonify(users)
        else:
            abort(404)