from datetime import datetime

from sqlalchemy import JSON, Boolean, Column, DateTime, Integer, String
from sqlalchemy.ext.declarative import as_declarative


@as_declarative()
class ModelBase:
    id = Column(Integer, primary_key=True)
    created_at = Column(DateTime, default=datetime.now)
    updated_at = Column(DateTime, onupdate=datetime.now)
    deleted = Column(Boolean, default=False)


class Resource(ModelBase):
    __tablename__ = "resource"

    resource_type = Column(String(20))
    args = Column(JSON)


class BackupPolicy(ModelBase):
    __tablename__ = "backup_policy"

    resource_id = Column(Integer)
    retention = Column(Integer)
    status = Column(String(20))


class BackupSchedule(ModelBase):
    __tablename__ = "backup_schedule"

    cron = Column(String(60))
    is_active = Column(Boolean, default=True)
    policy_id = Column(Integer)


class BackupSet(ModelBase):
    __tablename__ = "backup_set"

    is_valid = Column(Boolean, default=False)
    size = Column(Integer)
    backup_time = Column(DateTime)
    resource_id = Column(Integer)
    retention = Column(Integer)


class ScheduledJob(ModelBase):
    __tablename__ = "scheduled_job"

    start_time = Column(DateTime)
    end_time = Column(DateTime)
    resource_id = Column(Integer)
    job_type = Column(String(20))
    status = Column(String(20))
    args = Column(JSON)
