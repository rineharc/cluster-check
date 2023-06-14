from typing import List
import utils

from myTypes import Config, NamespacedEntry
from kubernetes.client import CoreV1Api
from kubernetes.client.rest import ApiException


def validate_pods(config_check: Config, clientset, sce_check: bool, azure_check: bool, oms_check: bool):
    print("---------------------------------------------------------")
    print("Checking for pods")
    print("---------------------------------------------------------")
    passed = True
    if oms_check:
        if not check_pods(config_check.OMS.pods, clientset):
            passed = False
    if sce_check:
        if not check_pods(config_check.SCE.pods, clientset):
            passed = False
    if azure_check:
        if not check_pods(config_check.Azure.pods, clientset):
            passed = False
    print()
    return passed


def check_pods(pods: List[NamespacedEntry], clientset):
    passed = True
    for pod_test in pods:
        status = ""
        color = ""
        try:
            pod_list = clientset.list_namespaced_pod(pod_test.namespace)
            found = False
            for pod in pod_list.items:
                if pod_test.name in pod.metadata.name:
                    found = True
                    break
            if found:
                status = "Found"
                color = utils.Color.GREEN.value
            else:
                status = "Not Found"
                color = utils.Color.RED.value
                passed = False
            print(f"{pod_test.namespace}/*{pod_test.name}* - {color}{status}\033[0m")
        except ApiException as e:
            status = "Not Found"
            color = utils.Color.RED
            passed = False
            utils.check_error(e)
            print(f"{pod_test.namespace}/*{pod_test.name}* - {color}{status}\033[0m")
    return passed