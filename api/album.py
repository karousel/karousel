from flask import abort, request
from . import AlbumModel, AuthenticatedResource

class AlbumsResource (AuthenticatedResource):

    def get (self):

        collection = request.args.get('collection')

        if collection is None:

            albums = AlbumModel.select()

            albums = [{
                        'id':album.id,
                        'name':album.name,
                        'collection': album.collection.name
                      } for album in AlbumModel.select()]

            return albums

        else:

            albums = AlbumModel.select()

            albums = [{
                        'id':album.id,
                        'name':album.name,
                        'collection': album.collection.name
                      } for album in AlbumModel.select().where(album.collection.name == collection)]

            return albums
