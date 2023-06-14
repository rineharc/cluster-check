from typing import List
import utils

from myTypes import Config, NamespacedEntry
from kubernetes.client import CoreV1Api
from kubernetes.client.rest import ApiException


def validate_service_accounts(config_check: Config, clientset, sce_check: bool, azure_check: bool, oms_check: bool):
    print("---------------------------------------------------------")
    print("Checking service accounts")
    print("---------------------------------------------------------")
    passed = True
    if oms_check:
        if not check_service_accounts(config_check.OMS.service_accounts, clientset):
            passed = False
    if sce_check:
        if not check_service_accounts(config_check.SCE.service_accounts, clientset):
            passed = False
    if azure_check:
        if not check_service_accounts(config_check.Azure.service_accounts, clientset):
            passed = False
    print()
    return passed


def check_service_accounts(service_accounts: List[NamespacedEntry], clientset):
    passed = True
    for sa in service_accounts:
        status = ""
        color = ""
        try:
            clientset.read_namespaced_service_account(sa.name, sa.namespace)
            status = "Found"
            color = utils.Color.GREEN.value
        except ApiException as e:
            status = "Not Found"
            color = utils.Color.RED.value
            passed = False
            utils.check_error(e)
        print(f"{sa.namespace}/{sa.name} - {color}{status}\033[0m")
    return passed