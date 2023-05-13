from app.core.database import models
from app.core.policy import BackupPolicy, Resource, BackupSchedule


class BackupPolicyController(object):
    def __init__(self, session):
        self.session = session

    def create_policy(self, retention, cron, resource_type, resource_args):
        res = Resource.add(resource_type, resource_args, session=self.session)
        pl = BackupPolicy.add(retention, res.id, session=self.session)
        BackupSchedule.add(pl.id, cron, session=self.session)

        return pl

    def update_policy(self, policy_id, retention=None, cron=None, resource_args=None):
        pl = BackupPolicy(policy_id, refresh_from_db=True, session=self.session)
        pl.update(retention=retention)

        res = pl.get_resource()
        res.update(args=resource_args)

        sch = pl.get_backup_schedule()
        sch.update(cron=cron)

    def enable_policy(self, policy_id):
        pl = BackupPolicy(policy_id, session=self.session)
        pl.enable()

    def disable_policy(self, policy_id):
        pl = BackupPolicy(policy_id, session=self.session)
        pl.disable()

    def get_policies_all(self, status):
        query = self.session.query(models.BackupPolicy.id)

        if status is not None:
            query = query.filter_by(status=status)

        policy_ids = query.all()

        policies = [
            BackupPolicy(i, refresh_from_db=True, session=self.session).to_json()
            for (i,) in policy_ids
        ]

        return policies
