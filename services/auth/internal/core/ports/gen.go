package ports

//go:generate mockgen -destination=auth_usecase_mock.go -package=ports github.com/popeskul/awesome-messanger/services/auth/internal/core/ports AuthUseCase
//go:generate mockgen -destination=token_manager_mock.go -package=ports github.com/popeskul/awesome-messanger/services/auth/internal/core/ports TokenManager
//go:generate mockgen -destination=token_usecase_mock.go -package=ports github.com/popeskul/awesome-messanger/services/auth/internal/core/ports TokenUseCase
//go:generate mockgen -destination=zap_logger_mock.go -package=ports github.com/popeskul/awesome-messanger/services/auth/internal/core/ports Logger
