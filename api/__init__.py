import ConfigParser

from boto.s3.connection import S3Connection
from flask import Flask
from flask.ext.restful import Api, Resource
from peewee import *

from cors import crossdomain

config = ConfigParser.RawConfigParser()
config.read('server.conf')

database = SqliteDatabase('gallery.db', threadlocals=True)

from models import CollectionModel, AlbumModel, UserModel, PhotoModel, Token

database.create_tables([PhotoModel, AlbumModel, UserModel, CollectionModel, Token], True)

if UserModel.select().count() == 0:

    # username = admin; password = password
    UserModel.create(
        admin = True,
        name = 'Administrator',
	    username = 'admin',
	    password = '$2a$12$pMtKl1b7h1sFKbMdBvPqbuza1tJN2ZNNAFMEs1RQmwqYTbBwrrKpy'
    )

s3 = S3Connection(config.get('S3', 'AccessKey'), config.get('S3', 'SecretKey'))

if s3.lookup(config.get('S3', 'Bucket')) is None:

    s3.create_bucket(config.get('S3', 'Bucket'))

store = s3.get_bucket(config.get('S3', 'Bucket'))

app = Flask(__name__)
api = Api(app)

from authentication import Authenticate, AuthenticatedResource
from user import UserInstance, UsersResource, RegistrationResource
from collection import CollectionsResource
from album import AlbumsResource, AlbumInstance
from photo import PhotosResource

api.decorators=[crossdomain(origin='*')]

api.add_resource(PhotosResource, '/photos/')
api.add_resource(UserInstance, '/users/<string:id>/')
api.add_resource(UsersResource, '/users/')
api.add_resource(RegistrationResource, '/users/')
api.add_resource(AlbumInstance, '/albums/<string:id>/')
api.add_resource(AlbumsResource, '/albums/')
api.add_resource(CollectionsResource, '/collections/')
api.add_resource(Authenticate, '/authenticate/')
