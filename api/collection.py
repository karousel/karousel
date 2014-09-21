from peewee import *
from . import database

class CollectionModel (Model):

    name = CharField()

    class Meta:

        database = database
