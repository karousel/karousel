import datetime
from peewee import *
from . import database

class CollectionModel (Model):

    name = CharField()
    created = DateTimeField(default=datetime.datetime.now)

    class Meta:

        database = database

class AlbumModel (Model):

    name = CharField()
    created = DateTimeField(default=datetime.datetime.now)
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
    
    # 0 - Waiting, 1 - Processing, 2 - Finalizing, 3 - Uploaded
    status = IntegerField(default=0)

    class Meta:

        database = database


class TokenModel (Model):

    token = CharField()
    user = IntegerField()
    created = DateTimeField(default=datetime.datetime.now)

    class Meta:

        database = database
