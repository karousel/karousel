from peewee import *
from flask.ext.restful import Resource
from . import database

class UserModel (Model):

    admin = BooleanField()
    name = CharField()
    username = CharField()
    password = CharField()

    class Meta:

        database = database

class UsersResource (Resource):

    def get (self):

        users = UserModel.select()

        users = [{'id':user.id, 'admin':user.admin, 'name':user.name, 'username':user.username} for user in users]

        return users
