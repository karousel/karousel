from flask import abort, request, g
from . import CollectionModel, AlbumModel, AuthenticatedResource

class AlbumInstance (AuthenticatedResource):

    def get (self, id):

        if AlbumModel.select().where(AlbumModel.id == id).count() != 1:

            abort(404)

        album = AlbumModel.get(AlbumModel.id == id)

        photos = [{
                    'id': photo.id,
                    'name': photo.name,
                    'uploaded': photo.uploaded.strftime("%Y-%m-%d %H:%M:%S") ,
                    'size': photo.size
                  } for photo in album.photos]

        return {'id': album.id, 'name': album.name, 'collection': {'name': album.collection.name, 'id': album.collection.id}, 'photos': photos}

class AlbumsResource (AuthenticatedResource):

    def get (self):

        collection = request.args.get('collection')

        if collection is None:

            albums = [{
                        'id':album.id,
                        'name':album.name,
                        'collection': {
                            'name': album.collection.name,
                            'id': album.collection.id
                        }
                      } for album in AlbumModel.select()]

            return albums

        else:

            if not collection.isdigit():

                abort(400)

            if CollectionModel.select().where(CollectionModel.id == collection).count() != 1:

                abort(404)

            collection = CollectionModel.get(CollectionModel.id == collection)

            albums = [{
                        'id':album.id,
                        'name':album.name,
                        'collection': {
                            'name': album.collection.name,
                            'id': album.collection.id
                        }
                      } for album in collection.albums]

            return albums

    def post (self):

        if not g.user.admin:

            abort(401)

        collection = request.form.get('collection')
        name = request.form.get('name')

        if not collection or not name or not collection.isdigit():

            abort(400)

        if CollectionModel.select().where(CollectionModel.id == collection).count() != 1:

            abort(404)

        collection = CollectionModel.get(CollectionModel.id == collection)

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
                    'collection': {
                        'name': album.collection.name,
                        'id': album.collection.id
                    }
                  } for album in AlbumModel.select()]

        return albums
