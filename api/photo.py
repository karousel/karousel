import datetime
from flask import abort, request, g
from werkzeug import secure_filename
from . import PhotoModel, AlbumModel, AuthenticatedResource, store

class PhotosResource (AuthenticatedResource):

    def post (self):

        album = request.form.get('album')
        file = request.files.get('file')

        if not album or not file:

            abort(400)

        if AlbumModel.select().where(AlbumModel.id == album).count() != 1:

            abort(404)

        album = AlbumModel.get(AlbumModel.id == album)

        filename = secure_filename(file.filename)

        if store.get_key(filename) is not None:

            abort(409)

        key = store.new_key(filename)
        key.set_contents_from_file(file)

        photo = PhotoModel.create(
            name = filename,
            size = len(file.read()),
            album = album,
            user = g.user
        )

        return {'id': photo.id, 'name': photo.name, 'date': photo.uploaded.strftime("%Y-%m-%d %H:%M:%S"), 'size': photo.size}
