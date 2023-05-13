from flask import request

from app.core.database.session import get_session

from . import api
from .controller import BackupPolicyController


@api.route("/policy", methods=["POST"])
def create_policy():
    args = request.json
    session = get_session()
    resource = args["resource"]

    ctr = BackupPolicyController(session)
    policy = ctr.create_policy(
        args["retention"], args["cron"], resource["resource_type"], resource["args"]
    )

    session.commit()

    return {"error": "", "result": {"policy_id": policy.id}}


@api.route("/policy/<int:policy_id>", methods=["POST"])
def update_policy(policy_id):
    args = request.json
    session = get_session()

    ctr = BackupPolicyController(session)
    ctr.update_policy(policy_id, args["retention"], args["cron"], args["resource_args"])

    session.commit()

    return {"error": "", "result": None}


@api.route("/policy/<int:policy_id>/enable", methods=["POST"])
def enable_policy(policy_id):
    session = get_session()

    ctr = BackupPolicyController(session)
    ctr.enable_policy(policy_id)

    session.commit()

    return {"error": "", "result": None}


@api.route("/policy/<int:policy_id>/disable", methods=["POST"])
def disable_policy(policy_id):
    session = get_session()

    ctr = BackupPolicyController(session)
    ctr.disable_policy(policy_id)

    session.commit()

    return {"error": "", "result": None}


@api.route("/policy", methods=["GET"])
def get_policy_all():
    args = request.args
    session = get_session()

    ctr = BackupPolicyController(session)
    policies = ctr.get_policies_all(status=args["status"])

    session.commit()

    return {"error": "", "result": policies}
