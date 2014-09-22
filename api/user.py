from flask import request, abort
from . import UserModel, AuthenticatedResource
import bcrypt

class UserInstance (AuthenticatedResource):

    def get (self, id):

        if UserModel.select().where(UserModel.id == id).count() != 1:

            abort(404)

        user = UserModel.get(UserModel.id == id)

        return {'id': user.id, 'name': user.name, 'username': user.username}

class UsersResource (AuthenticatedResource):

    def get (self):

        users = UserModel.select()

        users = [{
                    'id':user.id,
                    'admin':user.admin,
                    'name':user.name,
                    'username':user.username
                 } for user in users]

        return users

    def post (self):

        name = request.form.get('name').encode('utf-8')
        username = request.form.get('username').encode('utf-8')
        password = request.form.get('password').encode('utf-8')

        if not name or not username or not password:

            abort(400)

        if UserModel.select().where(UserModel.username == username).count() == 1:

            abort(409)

        UserModel.create(
            admin = False,
            name = name,
            username = username,
            password = bcrypt.hashpw(password, bcrypt.gensalt())
        )

        users =  [{
                    'id': user.id,
                    'admin': user.admin,
                    'name': user.name,
                    'username': user.username
                  } for user in UserModel.select()]

        return users
