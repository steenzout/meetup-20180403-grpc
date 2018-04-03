package start

// Pokemon representation of a Pokemon.
type Pokemon struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Weight int64  `json:"weight"`
}
