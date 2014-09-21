from peewee import *
from . import database

class AlbumModel (Model):

    name = CharField()

    class Meta:

        database = database
