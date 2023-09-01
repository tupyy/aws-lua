package lua

import (
	"context"

	lua "github.com/yuin/gopher-lua"
)

type TwtProvider interface {
	Create(ctx context.Context, o Object) (string, error)
	Verify(ctx context.Context, o Object, twt string) error
}

type TwtInterpreter struct {
	twtProvider TwtProvider
}

func NewTwtModule(twtProvider TwtProvider) *TwtInterpreter {
	return &TwtInterpreter{twtProvider}
}

func (t *TwtInterpreter) Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"create": t.create,
		"verify": t.verify,
	})

	L.Push(mod)
	return 1
}

func (t *TwtInterpreter) create(L *lua.LState) int {
	value := L.Get(1)
	goVal := toGoValue(value)

	obj, ok := goVal.(Object)
	if !ok {
		lerr := lua.LString("expected a table")
		L.Push(lua.LNil)
		L.Push(lerr)
		return 2
	}

	twt, err := t.twtProvider.Create(context.TODO(), obj)
	if err != nil {
		L.Push(lua.LNil)
		lerr := lua.LString(err.Error())
		L.Push(lerr)
		return 2
	}

	L.Push(lua.LString(twt))
	return 1
}

func (t *TwtInterpreter) verify(L *lua.LState) int {
	luaObj := L.Get(1)
	luaTwt := L.Get(2)
	goVal := toGoValue(luaObj)

	obj, ok := goVal.(Object)
	if !ok {
		lerr := lua.LString("expected a table")
		L.Push(lerr)
		return 1
	}

	twt, ok := toGoValue(luaTwt).(string)
	if !ok {
		lerr := lua.LString("expected a string")
		L.Push(lerr)
		return 1
	}

	err := t.twtProvider.Verify(context.TODO(), obj, twt)
	if err != nil {
		lerr := lua.LString(err.Error())
		L.Push(lerr)
		return 1
	}

	return 0
}
