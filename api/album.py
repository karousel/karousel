from flask import abort, request
from . import CollectionModel, AlbumModel, AuthenticatedResource

class AlbumsResource (AuthenticatedResource):

    def get (self):

        collection = request.args.get('collection')

        if collection is None:

            albums = [{
                        'id':album.id,
                        'name':album.name,
                        'collection': album.collection.name
                      } for album in AlbumModel.select()]

            return albums

        else:

            if CollectionModel.select().where(CollectionModel.name == collection).count() != 1:

                abort(404)

            collection = CollectionModel.get(CollectionModel.name == collection)

            albums = [{
                        'id':album.id,
                        'name':album.name,
                        'collection': album.collection.name
                      } for album in collection.albums]

            return albums

    def post (self):

        collection = request.form.get('collection')
        name = request.form.get('name')

        if not collection or not name:

            print collection
            print name

            abort(400)

        if CollectionModel.select().where(CollectionModel.name == collection).count() != 1:

            abort(404)

        collection = CollectionModel.get(CollectionModel.name == collection)

        for album in collection.albums:

            if album.name == name:

                abort(409)

        AlbumModel.create(
            name = name,
            collection = collection
        )

        albums = [{
                    'id':album.id,
                    'name':album.name,
                    'collection': album.collection.name
                  } for album in AlbumModel.select()]

        return albums
