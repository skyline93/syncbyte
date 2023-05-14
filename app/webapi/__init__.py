import logging

from flask import Flask

from .backup import api as backup_api

app = Flask(__name__)
app.register_blueprint(backup_api)

logger = logging.getLogger(__name__)


@app.errorhandler(Exception)
def internal_error(err):
    logger.exception(err)
    return {"error": str(err), "result": None}, 500
