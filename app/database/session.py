import sqlalchemy
import sqlalchemy.orm

from app.core.config import settings

_ENGINE = None
_MAKER = None


def get_engine():
    global _ENGINE

    if _ENGINE is None:
        engine_args = {
            "echo": False,
            "pool_recycle": 3600,
        }

        _ENGINE = sqlalchemy.create_engine(settings.database_uri, **engine_args)

        _ENGINE.connect()

    return _ENGINE


def get_session(expire_on_commit=False):
    global _MAKER

    if _MAKER is None:
        engine = get_engine()
        _MAKER = sqlalchemy.orm.sessionmaker(
            bind=engine, expire_on_commit=expire_on_commit
        )

    return _MAKER()
