from flask import Flask
from flask_restful import Api

HOST = "recommendations"
SERVER = "server"

app = Flask(__name__)
api = Api(app)

import routes
