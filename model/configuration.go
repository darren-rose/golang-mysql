package model

type Configuration struct {
	User     string `required:"true"`
	Password string `required:"true"`
	Host     string `required:"true"`
	Port     int    `required:"true"`
	Database string `required:"true"`
	AppPort  int    `required:"true"`
}
