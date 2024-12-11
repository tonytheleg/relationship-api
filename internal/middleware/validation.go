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
				if err := ValidateRelationshipCount(v); err != nil {
					return nil, errors.BadRequest("VALIDATOR", err.Error()).WithCause(err)
				}
			}
			return handler(ctx, req)
		}
	}
}

// ValidateRelationshipCount ensures that only one relationship is defined in a request
func ValidateRelationshipCount(msg protoreflect.ProtoMessage) error {
	var relationshipCount map[string]interface{}

	data, err := protojson.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshall message: %w", err)
	}
	err = json.Unmarshal(data, &relationshipCount)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}
	if len(relationshipCount) > 2 {
		return fmt.Errorf("only one relationship can be defined per request")
	}
	return nil
}
