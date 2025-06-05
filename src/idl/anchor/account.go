package anchor

type Account struct {
	Name          string        `json:"name"`
	Discriminator Discriminator `json:"discriminator,omitempty"`
}
