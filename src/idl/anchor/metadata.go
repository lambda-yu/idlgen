package anchor

type Metadata struct {
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	Spec         string       `json:"spec"`
	Description  string       `json:"description,omitempty"`
	Repository   string       `json:"repository,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
	Contact      string       `json:"contact,omitempty"`
	Deployments  Deployments  `json:"deployments,omitempty"`
}

type Dependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
type Deployments struct {
	Mainnet  string `json:"mainnet,omitempty"`
	Testnet  string `json:"testnet,omitempty"`
	Devnet   string `json:"devnet,omitempty"`
	Localnet string `json:"localnet,omitempty"`
}
