package ports

//go:generate mockgen -destination=app_mock.go -package=ports github.com/popeskul/awesome-messanger/services/platform/app/ports App
//go:generate mockgen -destination=logger_mock.go -package=ports github.com/popeskul/awesome-messanger/services/platform/app/ports Logger
//go:generate mockgen -destination=server_mock.go -package=ports github.com/popeskul/awesome-messanger/services/platform/app/ports GRPCServer,HTTPServer,SwaggerServer,ServerFactory
