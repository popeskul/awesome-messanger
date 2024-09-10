package ports

//go:generate mockgen -destination=usecase_mock.go -package=ports github.com/popeskul/awesome-messanger/services/notification/internal/core/ports NotificationUseCase
//go:generate mockgen -destination=zap_logger_mock.go -package=ports github.com/popeskul/awesome-messanger/services/notification/internal/core/ports Logger
