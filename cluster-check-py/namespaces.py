from typing import List
import utils

from myTypes import Config
from kubernetes.client import CoreV1Api
from kubernetes.client.rest import ApiException



def validate_namespaces(config_check: Config, clientset, sce_check: bool, azure_check: bool, oms_check: bool):
    print("---------------------------------------------------------")
    print("Checking namespaces")
    print("---------------------------------------------------------")
    passed = True
    if oms_check:
        if not check_namespaces(config_check.OMS.namespaces, clientset):
            passed = False
    if sce_check:
        if not check_namespaces(config_check.SCE.namespaces, clientset):
            passed = False
    if azure_check:
        if not check_namespaces(config_check.Azure.namespaces, clientset):
            passed = False
    print()
    return passed


def check_namespaces(namespaces: List[str], clientset):
    passed = True
    for ns in namespaces:
        status = ""
        color = ""
        try:
            clientset.read_namespace(ns)
            status = "Found"
            color = utils.Color.GREEN.value
        except ApiException as e:
            status = "Not Found"
            color = utils.Color.RED.value
            passed = False
            utils.check_error(e)
        print(f"{ns} - {color}{status}\033[0m")
    return passed