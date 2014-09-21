from flask import abort, request
from . import CollectionModel, AuthenticatedResource

class CollectionsResource (AuthenticatedResource):

    def get (self):

        collections = [{
                        'id':collection.id,
                        'name':collection.name,
                       } for collection in CollectionModel.select()]

        return collections

    def post (self):

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
