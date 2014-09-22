import datetime
from peewee import *
from . import database

class CollectionModel (Model):

    name = CharField()

    class Meta:

        database = database

class AlbumModel (Model):

    name = CharField()
    collection = ForeignKeyField(CollectionModel, related_name='albums')

    class Meta:

        database = database

class UserModel (Model):

    admin = BooleanField()
    name = CharField()
    username = CharField()
    password = CharField()

    class Meta:

        database = database

class PhotoModel (Model):

    name = CharField()
    uploaded = DateTimeField(default=datetime.datetime.now)
    size = IntegerField()
    user = ForeignKeyField(UserModel, related_name='photos')
    album = ForeignKeyField(AlbumModel, related_name='photos')

    class Meta:

        database = database


class Token (Model):

    token = CharField()
    user = IntegerField()

    class Meta:

        database = database
