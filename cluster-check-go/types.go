package main

type NamespacedEntry struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ConfigEntry struct {
	Namespaces      []string          `json:"namespaces"`
	ServiceAccounts []NamespacedEntry `json:"serviceAccounts"`
	Pods            []NamespacedEntry `json:"pods"`
}

type Config struct {
	OMS              ConfigEntry `json:"oms"`
	Azure            ConfigEntry `json:"azure"`
	SCE              ConfigEntry `json:"sce"`
	SecurityContexts []string    `json:"securityContexts"`
}

func (*Config) GetConfig() Config {
	return Config{
		OMS: ConfigEntry{
			Namespaces: []string{"camunda", "confluent", "flux-system", "meijer", "meijer-rx", "rng-kafka", "rng-rx-kafka"},
			ServiceAccounts: []NamespacedEntry{
				{
					Name:      "meijer-sa",
					Namespace: "meijer",
				},
				{
					Name:      "flux-applier",
					Namespace: "meijer",
				},
				{
					Name:      "confluent-for-kubernetes",
					Namespace: "confluent",
				},
				{
					Name:      "flux-applier",
					Namespace: "confluent",
				},
				{
					Name:      "confluent-for-kubernetes",
					Namespace: "rng-kafka",
				},
				{
					Name:      "confluent-for-kubernetes",
					Namespace: "rng-rx-kafka",
				},
			},
			Pods: []NamespacedEntry{},
		},
		Azure: ConfigEntry{
			Namespaces:      []string{"arcdataservices", "azure-arc"},
			ServiceAccounts: []NamespacedEntry{},
			Pods: []NamespacedEntry{
				{
					Name:      "store",
					Namespace: "arcdataservices",
				},
			},
		},
		SCE: ConfigEntry{
			Namespaces:      []string{"camunda", "confluent", "flux-system", "meijer", "meijer-rx", "rng-kafka", "rng-rx-kafka"},
			ServiceAccounts: []NamespacedEntry{},
			Pods: []NamespacedEntry{
				{
					Name:      "confluent-operator",
					Namespace: "confluent",
				},
				{
					Name:      "kafka",
					Namespace: "rng-kafka",
				},
				{
					Name:      "kafka",
					Namespace: "rng-rx-kafka",
				},
				{
					Name:      "lightningcart",
					Namespace: "meijer",
				},
			},
		},
		SecurityContexts: []string{"meijer-scc", "confluent-operator"},
	}
}
