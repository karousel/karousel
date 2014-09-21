from peewee import *
from flask import request, abort
from flask.ext.restful import Resource
import bcrypt
from . import database, UserModel
import random
import string

class Token (Model):

    token = CharField()
    user = IntegerField()

    class Meta:

        database = database

class Authenticate (Resource):

    def post (self):

        username = request.form.get('username').encode('utf-8')
        password = request.form.get('password').encode('utf-8')

        def verify (username, password):

            try:
              
                user = UserModel.get(UserModel.username == username)
                computed = user.password.encode('utf-8')

                if bcrypt.hashpw(password, computed) == computed:

                    return True
            
            except Exception:

                pass
              
            return False

        if not username or not password or not verify(username, password):

            abort(401)

        else:

            token = ''.join(random.choice(string.ascii_uppercase + string.digits) for _ in range(30))
            user = UserModel.get(UserModel.username == username)

            Token.create(
                token = token,
                user = user.id
            )

            return {'token': token}
