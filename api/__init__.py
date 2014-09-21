import ConfigParser

from peewee import *

config = ConfigParser.RawConfigParser()
config.read('server.conf')

database = SqliteDatabase('gallery.db', threadlocals=True)

from collection import CollectionModel
from album import AlbumModel
from user import UserModel, UsersResource
from photo import PhotoModel

database.create_tables([PhotoModel, AlbumModel, UserModel, CollectionModel], True)

if UserModel.select().count() == 0:

	  UserModel.create(
	      admin = True,
	      name = 'Admin',
	      username = 'Admin',
	      password = '$2a$12$pMtKl1b7h1sFKbMdBvPqbuza1tJN2ZNNAFMEs1RQmwqYTbBwrrKpy'
    )

from boto.s3.connection import S3Connection
s3 = S3Connection(config.get('S3', 'AccessKey'), config.get('S3', 'SecretKey'))

if s3.lookup(config.get('S3', 'Bucket')) is None:

    s3.create_bucket(config.get('S3', 'Bucket'))

from flask import Flask
from flask.ext.restful import Api

app = Flask(__name__)
api = Api(app)

api.add_resource(UsersResource, '/users/')
