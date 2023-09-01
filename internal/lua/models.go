package lua

import (
	"fmt"
	"reflect"
)

const (
	AwsUser                 string = "aws_user"
	AwsAccessKey            string = "aws_access_key"
	AwsVpc                  string = "aws_vpc"
	AwsSubnet               string = "aws_subnet"
	AwsRoute                string = "aws_route"
	AwsInternetGateway      string = "aws_igw"
	AwsNatGateway           string = "aws_nat"
	AwsAZs                  string = "aws_availability_zones"
	UnknownType             string = "unknown_type"
	AwsProviderType         string = "aws"
	UnsupportedProviderType string = "unsupported_type"
)

type Object map[string]interface{}

func (o Object) Type() string {
	stype, err := getKey[string](o, "type")
	if err != nil {
		return UnknownType
	}
	return stype
}

func (o Object) Provider() string {
	provider, err := getKey[string](o, "provider")
	if err != nil {
		return UnsupportedProviderType
	}
	return provider
}

func (o Object) GetString(key string) string {
	v, err := getKey[string](o, key)
	if err != nil {
		return ""
	}
	return v
}

func (o Object) GetBool(key string) bool {
	v, err := getKey[bool](o, key)
	if err != nil {
		return false
	}
	return v
}

func (o Object) GetObject(key string) map[string]interface{} {
	v, err := getKey[Object](o, key)
	if err != nil {
		return make(map[string]interface{})
	}
	return v
}

func (o Object) GetList(key string) []interface{} {
	v, err := getKey[[]interface{}](o, key)
	if err != nil {
		return nil
	}
	return v
}

func getKey[S any](o Object, key string) (S, error) {
	var emptyVal S
	t, ok := o[key]
	if !ok {
		return emptyVal, fmt.Errorf("no value found for key %q", key)
	}
	stype, ok := t.(S)
	if !ok {
		return emptyVal, fmt.Errorf("value of key %q is not of type %q", key, reflect.TypeOf(emptyVal))
	}
	return stype, nil
}
