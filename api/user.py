from flask import request, abort, g
from . import config, UserModel, Resource, AuthenticatedResource
import bcrypt

class UserInstance (AuthenticatedResource):

    def get (self, id):

        if UserModel.select().where(UserModel.id == id).count() != 1:

            abort(404)

        user = UserModel.get(UserModel.id == id)

        return {'id': user.id, 'name': user.name, 'username': user.username}

    def delete (self, id):

        if not g.user.admin:

            abort(401)

        if UserModel.select().where(UserModel.id == id).count() != 1:

            abort(404)

        user = UserModel.get(UserModel.id == id)

        for photo in user.photos:

            photo.delete_instance()

        user.delete_instance()

        return '', 204


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

class RegistrationResource (Resource):

    def post (self):

        if not config.getboolean('Users', 'Registration'):

            abort(401)

        name = request.form.get('name').encode('utf-8')
        username = request.form.get('username').encode('utf-8')
        password = request.form.get('password').encode('utf-8')

        if not name or not username or not password:

            abort(400)

        if UserModel.select().where(UserModel.username == username).count() == 1:

            abort(409)

        user = UserModel.create(
            admin = False,
            name = name,
            username = username,
            password = bcrypt.hashpw(password, bcrypt.gensalt())
        )

        return {'id': user.id, 'name': user.name, 'username': user.username}
