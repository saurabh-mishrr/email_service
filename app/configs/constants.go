package configs

type EmailRequestPayLoad struct {
	Name string `json:"name"`
}

type MailerConfig struct {
	Host string
	Port int
	Username string
	Password string
	FromAddr string
	FromAlias string
	UseCommand bool
}
