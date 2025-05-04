package models

type Host struct {
	Alias string   `json:"alias"`
	Host  string   `json:"host"`
	User  string   `json:"user"`
	Port  int      `json:"port"`
	Tags  []string `json:"tags"`
	Group string   `json:"group"`
	Password string `json:"password"`
	CertPath string `json:"certPath"`
}