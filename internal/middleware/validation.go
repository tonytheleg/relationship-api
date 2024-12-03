package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Validation(validator *protovalidate.Validator) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if v, ok := req.(proto.Message); ok {
				if err := validator.Validate(v); err != nil {
					return nil, errors.BadRequest("VALIDATOR", err.Error()).WithCause(err)
				}
				if err := ValidateResourceCount(v); err != nil {
					return nil, errors.BadRequest("VALIDATOR", err.Error()).WithCause(err)
				}
			}
			return handler(ctx, req)
		}
	}
}

// ValidateResourceCount ensures that only one resource is defined in a request
func ValidateResourceCount(msg protoreflect.ProtoMessage) error {
	var resourceCount map[string]interface{}

	data, err := protojson.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshall message: %w", err)
	}
	err = json.Unmarshal(data, &resourceCount)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}
	if len(resourceCount) > 1 {
		return fmt.Errorf("only one resource can be defined per request")
	}
	return nil
}
