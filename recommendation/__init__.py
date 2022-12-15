from flask import Flask
from flask_restful import Api
from redis import StrictRedis

HOST = "recommendations" # recommendations
SERVER = "server" # server
CACHE = "cache" # cache

app = Flask(__name__)
api = Api(app)

cache = StrictRedis(host = CACHE, port = 6379)

import routes
