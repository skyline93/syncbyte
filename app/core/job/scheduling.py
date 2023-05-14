from enum import Enum

from app.core.database import models
from app.core.database.session import get_session
from app.core.policy import BackupPolicy


class JobType(str, Enum):
    BACKUP = "backup"


class JobStatus(str, Enum):
    QUEUED = "queued"
    RUNNING = "running"
    FAILED = "failed"
    SUCCESSED = "successed"


class JobConflict(Exception):
    pass


class ScheduledJob(object):
    init_status = JobStatus.QUEUED

    def __init__(self, id, refresh_from_db=False, **kwargs):
        self.id = id
        self.session = kwargs.get("sesson", get_session())

        if refresh_from_db:
            self._refresh()

    def _refresh(self):
        item = self.session.query(models.ScheduledJob).filter_by(id=self.id).first()
        self._load_from_model(item)

    def _load_from_model(self, m):
        self.resource_id = m.resource_id
        self.job_type = m.job_type
        self.args = m.args
        self.status = m.status

    @classmethod
    def add(cls, resource_id, job_type, args, **kwargs):
        session = kwargs["session"]

        job = models.ScheduledJob(
            resource_id=resource_id,
            job_type=job_type,
            args=args,
            status=cls.init_status,
        )

        session.add(job)
        session.flush()

        self = cls(job.id)
        self._load_from_model(job)
        return self


def schedule_backup_job(policy_id, **kwargs):
    session = kwargs["session"]
    job_type = JobType.BACKUP

    policy = BackupPolicy(policy_id, refresh_from_db=True, session=session)
    resource = policy.get_resource()

    if check_backup_job_conflict(resource.id, session=session):
        raise JobConflict(f"backup job conflict in resource {resource.id}")

    job = ScheduledJob.add(policy.resource_id, job_type, resource.args, session=session)

    return job


def check_backup_job_conflict(resource_id, **kwargs):
    session = kwargs["session"]

    result = (
        session.query(models.ScheduledJob)
        .filter(
            models.ScheduledJob.resource_id == resource_id,
            models.ScheduledJob.job_type.in_([JobType.BACKUP]),
            models.ScheduledJob.status.in_([JobStatus.QUEUED, JobStatus.RUNNING]),
        )
        .count()
    )

    if result == 0:
        return False

    return True
