import ConfigParser

from peewee import *

config = ConfigParser.RawConfigParser()
config.read('server.conf')

database = SqliteDatabase('gallery.db')

from collection import CollectionModel
from album import AlbumModel
from user import UserModel
from photo import PhotoModel

database.create_tables([PhotoModel, AlbumModel, UserModel, CollectionModel], True)

from boto.s3.connection import S3Connection
s3 = S3Connection(config.get('S3', 'AccessKey'), config.get('S3', 'SecretKey'))

if s3.lookup(config.get('S3', 'Bucket')) is None:

    s3.create_bucket(config.get('S3', 'Bucket'))

