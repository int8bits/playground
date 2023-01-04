
from flask import Flask
from flask_restful import Api

from api import Healtheck


app = Flask(__name__)
api = Api(app)


api.add_resource(Healtheck, '/api/healthcheck')

if __name__ == '__main__':
    app.run(debug=True)
