package config

type Config struct {
	HttpHost string `json:"http_host"`
	HttpPort int    `json:"http_port"`
	TLSCert  string `json:"tls_cert"`
	TLSKey   string `json:"tls_key"`
}
