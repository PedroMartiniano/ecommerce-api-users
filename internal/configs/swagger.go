package configs

import "github.com/swaggo/swag"

func SwaggerConfigure(infos *swag.Spec) {
	infos.Title = "E-commerce API Users"
	infos.Description = "API Rest for users management"
	infos.Host = "localhost:8090"
	infos.BasePath = "/"
	infos.Version = "1.0"
	infos.Schemes = []string{"http", "https"}
}
