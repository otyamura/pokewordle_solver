package types

type Response struct {
	Error string   `json:"error"`
	Names []string `json:"names"`
}
