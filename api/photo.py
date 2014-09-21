from peewee import *
from . import database, UserModel, AlbumModel

class PhotoModel (Model):

    name = CharField()
    date = DateField()
    size = BigIntegerField()
    user = ForeignKeyField(UserModel, related_name='photos')
    album = ForeignKeyField(AlbumModel, related_name='photos')

    class Meta:

        database = database
