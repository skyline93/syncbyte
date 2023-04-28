from flask import Flask

from app.core.config import settings

from .backup import api as backup_api

app = Flask(__name__)
# app.config.from_object(settings)
app.register_blueprint(backup_api)


@app.errorhandler(Exception)
def page_not_found(err):
    return f"err: {err}", 500
