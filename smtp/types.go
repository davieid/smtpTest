package smtp

type SMTPConfig struct{
	Host string `json:"host"`
	Port int `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}