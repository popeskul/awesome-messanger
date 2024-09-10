package ports

//go:generate mockgen -destination=usecase_mock.go -package=ports github.com/popeskul/awesome-messanger/services/message/internal/core/ports MessageUseCase
//go:generate mockgen -destination=zap_logger_mock.go -package=ports github.com/popeskul/awesome-messanger/services/message/internal/core/ports Logger
