from flask import request

from app.database.session import get_session

from . import api
from .controller import BackupPolicyController


@api.route("/policy", methods=["POST"])
def create_policy():
    args = request.json
    session = get_session()
    resource = args["resource"]

    ctr = BackupPolicyController(session)
    policy = ctr.create_policy(
        args["retention"], resource["resource_type"], resource["args"]
    )

    session.commit()

    return {"error": "", "result": {"policy_id": policy.id}}
