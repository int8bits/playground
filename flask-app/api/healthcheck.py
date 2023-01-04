
from flask_restful import Resource


class Healtheck(Resource):
    def get(self):
        return {'status': 'ok'}