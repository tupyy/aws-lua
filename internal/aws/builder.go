package aws

import (
	"context"
	"errors"

	"github.com/tupyy/aws-lua/internal/lua"
)

type clientBuilder[T, S any] struct {
	config              ClientConfiguration
	client              ClientType
	opType              OpType
	tranformInputFunc   func(o lua.Object) T
	transformOutputFunc func(s S) lua.Object
}

func NewBuilder[T, S any](c ClientConfiguration) *clientBuilder[T, S] {
	return &clientBuilder[T, S]{
		config: c,
	}
}

func (b *clientBuilder[T, S]) Type(t ClientType) *clientBuilder[T, S] {
	b.client = t
	return b
}

func (b *clientBuilder[T, S]) Op(opType OpType) *clientBuilder[T, S] {
	b.opType = opType
	return b
}

func (b *clientBuilder[T, S]) TransformInputFunc(f func(o lua.Object) T) *clientBuilder[T, S] {
	b.tranformInputFunc = f
	return b
}

func (b *clientBuilder[T, S]) TransformOutputFunc(f func(s S) lua.Object) *clientBuilder[T, S] {
	b.transformOutputFunc = f
	return b
}

func (b *clientBuilder[T, S]) Build(ctx context.Context) func(ctx context.Context, o lua.Object) (lua.Object, error) {
	return func(ctx context.Context, o lua.Object) (lua.Object, error) {
		opFunc := b.getOpFunc()
		output, err := opFunc(ctx, b.tranformInputFunc(o))
		if err != nil {
			return lua.Object{}, err
		}
		return b.transformOutputFunc(output.(S)), nil
	}
}

func (b *clientBuilder[T, S]) getOpFunc() opFunc {
	switch b.client {
	case IamClient:
		return iamGetOpFunc(b.opType, b.config)
	case Ec2Client:
		return ec2GetOpFunc(b.opType, b.config)
	}

	return func(ctx context.Context, input interface{}) (interface{}, error) {
		return nil, errors.New("unknown client or op type")
	}
}
