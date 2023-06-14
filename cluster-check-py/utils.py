import json

from enum import Enum
from myTypes import Config
from kubernetes.client import ApiClient, Configuration
from kubernetes.client.rest import ApiException

class Color(Enum):
    RED = "\033[31m"
    GREEN = "\033[32m"
    YELLOW = "\033[33m"
    BLUE = "\033[34m"
    CYAN = "\033[36m"
    WHITE = "\033[37m"
    RESET = "\033[0m"

def load_config(config_file: str):
    if config_file != "":
        with open(config_file, "r") as f:
            data = f.read()
        config_check = json.loads(data)
    else:
        config_check = Config.GetConfig()
    return config_check


def check_error(err: ApiException):
    if isinstance(err, ApiException):
        color = Color.RED
        if err.reason == "Forbidden" or err.reason == "Unauthorized":
            print(f"{color}{err.reason}: Make sure you are logged in\033[0m")
            exit(0)


def gen_json(data: Config):
    json_data = json.dumps(data.to_dict(), indent=4)
    with open("example-config.json", "w") as f:
        f.write(json_data)