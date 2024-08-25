package validator

import (
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"
)

type validatorBase struct {
	validator *protovalidate.Validator
}

func NewGrpcValidator(messages ...protoreflect.ProtoMessage) (ports.Validator, error) {
	mainValidator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(messages...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create validator: %w", err)
	}

	return &validatorBase{
		validator: mainValidator,
	}, nil
}

func (v validatorBase) Validate(msg protoreflect.ProtoMessage) error {
	return v.validator.Validate(msg)
}
