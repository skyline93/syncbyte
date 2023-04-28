from flask import Flask

from .backup import api as backup_api

app = Flask(__name__)
app.register_blueprint(backup_api)


@app.errorhandler(Exception)
def internal_error(err):
    return {"error": err, "result": None}, 500
