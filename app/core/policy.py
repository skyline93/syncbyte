from app.database.session import get_session
from app.database import models


class SyncbyteEntity(object):
    def __init__(self, id, **kwargs):
        self.id = id
        self.session = kwargs.get("session", get_session())


class BackupPolicy(SyncbyteEntity):
    def __init__(self, id, refresh_from_db=False, **kwargs):
        super().__init__(id, **kwargs)

        if refresh_from_db:
            self._refresh()

    def _refresh(self):
        item = self.session.query(models.BackupPolicy).filter_by(id=self.id).first()
        self._load_from_model(item)

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

    def update(self, **kwargs):
        self.session.query(models.BackupPolicy).filter_by(id=self.id).update(kwargs)

    def get_resource(self):
        return Resource(self.resource_id, refresh_from_db=True, session=self.session)


class Resource(SyncbyteEntity):
    def __init__(self, id, refresh_from_db=False, **kwargs):
        super().__init__(id, **kwargs)

        if refresh_from_db:
            self._refresh()

    def _refresh(self):
        item = self.session.query(models.Resource).filter_by(id=self.id).first()
        self._load_from_model(item)

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

    def update(self, **kwargs):
        self.session.query(models.Resource).filter_by(id=self.id).update(kwargs)
