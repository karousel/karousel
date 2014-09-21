from peewee import *
from . import database

class UserModel (Model):

    admin = BooleanField()
    name = CharField()
    username = CharField()
    password = CharField()

    class Meta:

        database = database

