import ConfigParser

from flask import Flask
from flask.ext.restful import Api, Resource
from peewee import *

from cors import crossdomain

config = ConfigParser.RawConfigParser()
config.read('server.conf')

database = SqliteDatabase('gallery.db', threadlocals=True)

from models import CollectionModel, AlbumModel, UserModel, PhotoModel, TokenModel

database.create_tables([PhotoModel, AlbumModel, UserModel, CollectionModel, TokenModel], True)

if UserModel.select().count() == 0:

    # username = admin; password = password
    UserModel.create(
        admin = True,
        name = 'Administrator',
        username = 'admin',
        password = '$2a$12$pMtKl1b7h1sFKbMdBvPqbuza1tJN2ZNNAFMEs1RQmwqYTbBwrrKpy'
    )

app = Flask(__name__)
api = Api(app)

from authentication import Authenticate, AuthenticatedResource
from user import UserInstance, UsersResource, RegistrationResource
from collection import CollectionsResource, CollectionInstance
from album import AlbumsResource, AlbumInstance
from photo import PhotosResource, PhotoInstance

api.decorators=[crossdomain(origin='*')]

api.add_resource(PhotoInstance, '/photos/<string:id>/')
api.add_resource(PhotosResource, '/photos/')
api.add_resource(UserInstance, '/users/<string:id>/')
api.add_resource(UsersResource, '/users/')
api.add_resource(RegistrationResource, '/users/')
api.add_resource(AlbumInstance, '/albums/<string:id>/')
api.add_resource(AlbumsResource, '/albums/')
api.add_resource(CollectionInstance, '/collections/<string:id>/')
api.add_resource(CollectionsResource, '/collections/')
api.add_resource(Authenticate, '/authenticate/')
