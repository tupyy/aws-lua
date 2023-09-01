package lua

import (
	"context"
	"fmt"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

type AwsProvider interface {
	Create(ctx context.Context, resource string, o Object) (Object, error)
	Delete(ctx context.Context, resource string, o Object) (Object, error)
	List(ctx context.Context, resource string, o Object) (Object, error)
}

type LuaInterpreter struct {
	awsProvider AwsProvider
}

func NewAwsModule(awsProvider AwsProvider) *LuaInterpreter {
	return &LuaInterpreter{awsProvider}
}

func (l *LuaInterpreter) Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"create": l.create,
		"delete": l.delete,
		"list":   l.list,
	})

	L.Push(mod)
	return 1
}

func (l *LuaInterpreter) create(L *lua.LState) int {
	respTable := L.NewTable()

	resource, err := getData[string](L, 1)
	if err != nil {
		L.Push(respTable)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	obj, err := getData[Object](L, 2)
	if err != nil {
		L.Push(respTable)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	o, err := l.execute("create", resource, obj)
	if err != nil {
		L.Push(respTable)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	respTable = toLTable(o)
	L.Push(respTable)
	return 1
}

func (l *LuaInterpreter) delete(L *lua.LState) int {
	respTable := L.NewTable()

	resource, err := getData[string](L, 1)
	if err != nil {
		L.Push(respTable)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	obj, err := getData[Object](L, 2)
	if err != nil {
		L.Push(respTable)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	_, err = l.execute("delete", resource, obj)
	if err != nil {
		L.Push(respTable)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(respTable)
	return 1
}

func (l *LuaInterpreter) list(L *lua.LState) int {
	respTable := L.NewTable()

	resource, err := getData[string](L, 1)
	if err != nil {
		L.Push(respTable)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	obj, err := getData[Object](L, 2)
	if err != nil {
		L.Push(respTable)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	o, err := l.execute("list", resource, obj)
	if err != nil {
		L.Push(respTable)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	respTable = toLTable(o)
	L.Push(respTable)
	return 1
}

func getData[T any](L *lua.LState, idx int) (T, error) {
	var t T
	value := L.Get(idx)
	if value == lua.LNil {
		return t, nil
	}
	goVal := toGoValue(value)

	obj, ok := goVal.(T)
	if !ok {
		return t, fmt.Errorf("expected %s. got: %+v", reflect.TypeOf(t), value)
	}

	return obj, nil
}

func (l *LuaInterpreter) execute(name string, resource string, o Object) (Object, error) {
	switch name {
	case "create":
		return l.awsProvider.Create(context.TODO(), resource, o)
	case "list":
		return l.awsProvider.List(context.TODO(), resource, o)
	case "delete":
		return l.awsProvider.Delete(context.TODO(), resource, o)
	default:
		return nil, fmt.Errorf("unknows method")
	}
}
