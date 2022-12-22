from flask import Flask
from flask_restful import Api
from redis import StrictRedis
from flask_cors import CORS

HOST = "recommendations" # recommendations
SERVER = "server" # server
CACHE = "cache" # cache

app = Flask(__name__)
CORS(app)
api = Api(app)

cache = StrictRedis(host = CACHE, port = 6379)

import routes
