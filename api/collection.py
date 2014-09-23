from flask import abort, request, g
from . import CollectionModel, AuthenticatedResource

class CollectionInstance (AuthenticatedResource):

    def get (self, id):

        if CollectionModel.select().where(CollectionModel.id == id).count() != 1:

            abort(404)

        collection = CollectionModel.get(CollectionModel.id == id)

        albums = [{
                    'id': album.id,
                    'name': album.name
                  } for album in collection.albums]

        return {'id': collection.id, 'name': collection.name, 'albums': albums}

    def delete (self, id):

        if not g.user.admin:

            abort(401)

        if CollectionModel.select().where(CollectionModel.id == id).count() != 1:

            abort(404)

        collection = CollectionModel.get(CollectionModel.id == id)

        for album in collection.albums:

            for photo in album.photos:

                photo.delete_instance()

            album.delete_instance()

        collection.delete_instance()

        return '', 204

class CollectionsResource (AuthenticatedResource):

    def get (self):

        collections = [{
                        'id':collection.id,
                        'name':collection.name,
                       } for collection in CollectionModel.select()]

        return collections

    def post (self):

        if not g.user.admin:

            abort(401)

        name = request.form.get('name')

        if not name:

            abort(400)

        if CollectionModel.select().where(CollectionModel.name == name).count() == 1:

            abort(409)

        CollectionModel.create(
            name = name
        )

        collections = [{
                        'id':collection.id,
                        'name':collection.name,
                       } for collection in CollectionModel.select()]

        return collections
