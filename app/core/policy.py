from app.database import models


class BackupPolicy(object):
    def __init__(self, id):
        self.id = id

    def _load_from_model(self, m):
        self.resource_id = m.resource_id
        self.retention = m.retention

    @classmethod
    def add(cls, retention, resource_id, **kwargs):
        session = kwargs["session"]

        pl = models.BackupPolicy(
            retention=retention, resource_id=resource_id, status="enabled"
        )

        session.add(pl)
        session.flush()

        self = cls(pl.id)
        self._load_from_model(pl)
        return self


class Resource(object):
    def __init__(self, id):
        self.id = id

    def _load_from_model(self, m):
        self.type = m.resource_type
        self.args = m.args

        return self

    @classmethod
    def add(cls, resource_type, resource_args, **kwargs):
        session = kwargs["session"]

        res = models.Resource(
            resource_type=resource_type,
            args=resource_args,
        )

        session.add(res)
        session.flush()

        self = cls(res.id)
        self._load_from_model(res)
        return self
