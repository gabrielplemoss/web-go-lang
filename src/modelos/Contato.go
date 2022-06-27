package modelos

type Contato struct {
	ID       uint64 `json:"id,omitempty"`
	Dono     uint64 `json:"dono,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Apelido  string `json:"apelido,omitempty"`
	Site     string `json:"site,omitempty"`
	Email    string `json:"email,omitempty"`
	Telefone string `json:"telefone,omitempty"`
	Endereco string `json:"endereco,omitempty"`
}
