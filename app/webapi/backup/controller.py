from app.core.policy import BackupPolicy, Resource


class BackupPolicyController(object):
    def __init__(self, session):
        self.session = session

    def create_policy(self, retention, resource_type, resource_args):
        res = Resource.add(resource_type, resource_args, session=self.session)
        pl = BackupPolicy.add(retention, res.id, session=self.session)

        return pl
