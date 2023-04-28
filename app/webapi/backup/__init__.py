from flask import Blueprint

api = Blueprint("backup", __name__, url_prefix="/backup")

from .backup import *
