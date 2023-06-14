import openshift as oc
import utils

from myTypes import Config
from kubernetes.client.rest import ApiException
from kubernetes.dynamic import DynamicClient


def validate_sccs(dyn_client,config_check: Config):
    print("---------------------------------------------------------")
    print("Checking security context constraints")
    print("---------------------------------------------------------")
    passed = True
    scc_client = dyn_client.resources.get(api_version='security.openshift.io/v1', kind='SecurityContextConstraints')
    for scc in config_check.SecurityContexts:
        status = ""
        color = ""
        try:
            temp= scc_client.get(name=scc)
            status = "Found"
            color = utils.Color.GREEN.value
        except ApiException as e:
            status = "Not Found"
            color = utils.Color.RED.value
            passed = False
            utils.check_error(e)
        print(f"{scc} - {color}{status}\033[0m")
    print()
    return passed