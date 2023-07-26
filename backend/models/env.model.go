package models

type EnvConfig struct {
	Port           int    `required:"true"`
	Host           string `required:"false" default:"localhost"`
	KeyPath        string `required:"true"`
	Mongo_Uri      string `required:"true"`
	Redis_Host     string `required:"true"`
	Redis_Port     string `required:"true"`
	Redis_Password string `required:"true"`
}
