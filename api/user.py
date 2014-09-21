from . import UserModel, AuthenticatedResource

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
