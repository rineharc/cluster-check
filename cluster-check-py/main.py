import os
import pathlib
import urllib3

from kubernetes import client, config
from kubernetes.client.rest import ApiException
from kubernetes.dynamic import DynamicClient

from myTypes import Config
from utils import gen_json, load_config
from namespaces import validate_namespaces
from serviceAccounts import validate_service_accounts
from securityContext import validate_sccs
from pods import validate_pods


def main():
    config_file = ""
    oms_check = True
    sce_check = False
    scc_check = True
    azure_check = False
    gen_json_check = False

    # kubeconfig = pathlib.Path.home() / ".kube" / "config"

    urllib3.disable_warnings()

    passed = True

    if gen_json_check:
        base_config = Config.GetConfig()
        gen_json(base_config)
        exit(0)
    config_check = load_config(config_file)

    config.load_kube_config()
    configuration = client.Configuration()
    configuration.verify_ssl = False
    api_client = client.ApiClient(configuration=configuration)

    clientset = client.CoreV1Api()

    if scc_check:
        dyn_client = DynamicClient(client.ApiClient())
        try:
            scc_check = validate_sccs(dyn_client, config_check)
        except ApiException as e:
            print("Error creating OpenShift clientset:", e)
            exit(1)
        if not scc_check:
            passed = False

    ns_check = validate_namespaces(config_check, clientset, sce_check, azure_check, oms_check)
    if not ns_check:
        passed = False
    sa_check = validate_service_accounts(config_check, clientset, sce_check, azure_check, oms_check)
    if not sa_check:
        passed = False
    pod_check = validate_pods(config_check, clientset, sce_check, azure_check, oms_check)
    if not pod_check:
        passed = False

    separator = "---------------------------------------------------------"
    print(f"\n{separator}\n")
    if passed:
        print(f"\033[32mAll checks passed\033[0m")
    else:
        print(f"\033[31mSome checks failed\033[0m")
    print(separator)

if __name__ == "__main__":
    main()