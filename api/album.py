from peewee import *
from . import database, CollectionModel

class AlbumModel (Model):

    name = CharField()
    collection = ForeignKeyField(CollectionModel, related_name='albums')

    class Meta:

        database = database
