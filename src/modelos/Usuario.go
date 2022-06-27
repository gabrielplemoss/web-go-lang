package modelos

type Usuario struct {
	ID      uint64 `json:"id,omitempty"`
	Usuario string `json:"usuario,omitempty"`
	Email   string `json:"email,omitempty"`
	Senha   string `json:"senha,omitempty"`
}
