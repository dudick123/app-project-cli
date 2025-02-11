package cmd

type Destination struct {
	Namespace string `yaml:"namespace"`
	Server    string `yaml:"server"`
	Name      string `yaml:"name"`
}

// AppProject represents the ArgoCD AppProject YAML structure
type AppProject struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name       string   `yaml:"name"`
		Namespace  string   `yaml:"namespace"`
		Finalizers []string `yaml:"finalizers"`
	} `yaml:"metadata"`
	Spec struct {
		Description              string        `yaml:"description"`
		SourceRepos              []string      `yaml:"sourceRepos"`
		Destinations             []Destination `yaml:"destinations"`
		ClusterResourceWhitelist []struct {
			Group string `yaml:"group"`
			Kind  string `yaml:"kind"`
		} `yaml:"clusterResourceWhitelist"`
		NamespaceResourceBlacklist []struct {
			Group string `yaml:"group"`
			Kind  string `yaml:"kind"`
		} `yaml:"namespaceResourceBlacklist"`
		NamespaceResourceWhitelist []struct {
			Group string `yaml:"group"`
			Kind  string `yaml:"kind"`
		} `yaml:"namespaceResourceWhitelist"`
		OrphanedResources struct {
			Warn bool `yaml:"warn"`
		} `yaml:"orphanedResources"`
		Roles []struct {
			Name        string   `yaml:"name"`
			Description string   `yaml:"description"`
			Policies    []string `yaml:"policies"`
			Groups      []string `yaml:"groups"`
			JwtTokens   []struct {
				Iat int64 `yaml:"iat"`
			} `yaml:"jwtTokens"`
		} `yaml:"roles"`
		SyncWindows []struct {
			Kind         string   `yaml:"kind"`
			Schedule     string   `yaml:"schedule"`
			Duration     string   `yaml:"duration"`
			Applications []string `yaml:"applications"`
			ManualSync   bool     `yaml:"manualSync"`
			Namespaces   []string `yaml:"namespaces"`
			Clusters     []string `yaml:"clusters"`
		} `yaml:"syncWindows"`
		PermitOnlyProjectScopedClusters bool     `yaml:"permitOnlyProjectScopedClusters"`
		SourceNamespaces                []string `yaml:"sourceNamespaces"`
	} `yaml:"spec"`
}
