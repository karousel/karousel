from functools import wraps
import random
import string
import bcrypt
from flask import request, abort, g
from flask.ext.restful import Resource
from . import TokenModel, UserModel

def verify (function):

    @wraps(function)
    def wrapper(*args, **kwargs):
        token = request.headers.get('Auth-Token')

        if token is None:

              abort(401)

        else:

            if TokenModel.select().where(TokenModel.token == token).count() != 1:

                  abort(401)

            else:

                token = TokenModel.get(TokenModel.token == token)
                user = UserModel.get(UserModel.id == token.user)

                g.user = user

        return function(*args, **kwargs)
    return wrapper

class AuthenticatedResource (Resource):

    method_decorators = [verify]

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

            TokenModel.create(
                token = token,
                user = user.id
            )

            return {'token': token}
