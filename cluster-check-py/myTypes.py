from typing import List


class NamespacedEntry:
    def __init__(self, name: str, namespace: str):
        self.name = name
        self.namespace = namespace

class ConfigEntry:
    def __init__(self, namespaces: List[str], service_accounts: List[NamespacedEntry], pods: List[NamespacedEntry]):
        self.namespaces = namespaces
        self.service_accounts = service_accounts
        self.pods = pods
    
    def to_dict(self):
        return {
            "namespaces": self.namespaces,
            "service_accounts": [sa.__dict__ for sa in self.service_accounts],
            "pods": [pod.__dict__ for pod in self.pods]
        }

class Config:
    def __init__(self, oms: ConfigEntry, azure: ConfigEntry, sce: ConfigEntry, security_contexts: List[str]):
        self.OMS = oms
        self.Azure = azure
        self.SCE = sce
        self.SecurityContexts = security_contexts

    def to_dict(self):
        return {
            "OMS": self.OMS.to_dict(),
            "Azure": self.Azure.to_dict(),
            "SCE": self.SCE.to_dict(),
            "SecurityContexts": self.SecurityContexts
        }

    @staticmethod
    def GetConfig():
        return Config(
            oms=ConfigEntry(
                namespaces=["camunda", "confluent", "flux-system", "meijer", "meijer-rx", "rng-kafka", "rng-rx-kafka"],
                service_accounts=[
                    NamespacedEntry(name="meijer-sa", namespace="meijer"),
                    NamespacedEntry(name="flux-applier", namespace="meijer"),
                    NamespacedEntry(name="confluent-for-kubernetes", namespace="confluent"),
                    NamespacedEntry(name="flux-applier", namespace="confluent"),
                    NamespacedEntry(name="confluent-for-kubernetes", namespace="rng-kafka"),
                    NamespacedEntry(name="confluent-for-kubernetes", namespace="rng-rx-kafka")
                ],
                pods=[]
            ),
            azure=ConfigEntry(
                namespaces=["arcdataservices", "azure-arc"],
                service_accounts=[],
                pods=[NamespacedEntry(name="store", namespace="arcdataservices")]
            ),
            sce=ConfigEntry(
                namespaces=["camunda", "confluent", "flux-system", "meijer", "meijer-rx", "rng-kafka", "rng-rx-kafka"],
                service_accounts=[],
                pods=[
                    NamespacedEntry(name="confluent-operator", namespace="confluent"),
                    NamespacedEntry(name="kafka", namespace="rng-kafka"),
                    NamespacedEntry(name="kafka", namespace="rng-rx-kafka"),
                    NamespacedEntry(name="lightningcart", namespace="meijer")
                ]
            ),
            security_contexts=["meijer-scc", "confluent-operator"]
        )