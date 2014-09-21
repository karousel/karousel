from peewee import *
from flask import abort
from flask.ext.restful import Resource
from . import database

class UserModel (Model):

    admin = BooleanField()
    name = CharField()
    username = CharField()
    password = CharField()

    class Meta:

        database = database

class UserInstance (Resource):

    def get (self, id):

        if UserModel.select().where(UserModel.id == id).count() != 1:

            abort(404)

        user = UserModel.get(UserModel.id == id)

        return {'id': user.id, 'name': user.name, 'username': user.username}

class UsersResource (Resource):

    def get (self):

        users = UserModel.select()

        users = [{'id':user.id, 'admin':user.admin, 'name':user.name, 'username':user.username} for user in users]

        return users
