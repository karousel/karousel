from . import CollectionModel, AuthenticatedResource

class CollectionsResource (AuthenticatedResource):

    def get (self):

        collections = [{
                        'id':collection.id,
                        'name':collection.name,
                       } for collection in CollectionModel.select()]

        return collections
