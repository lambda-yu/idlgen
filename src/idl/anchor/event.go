package anchor

type Event struct {
	Name          string   `json:"name"`
	Discriminator *[8]byte `json:"discriminator,omitempty"`
}
