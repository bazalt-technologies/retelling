from __init__ import api
from flask_restful import reqparse

from recommendations import Recommendation
from analytics import Analytics

parser = reqparse.RequestParser()
parser.add_argument("Login")
parser.add_argument("Password")
parser.add_argument("ObjectID")
parser.add_argument("ObjectIDs")
parser.add_argument("UserID")


api.add_resource(Recommendation, "/api/v1/recommendation")
api.add_resource(Analytics, "/api/v1/analytics")
