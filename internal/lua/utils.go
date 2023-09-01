package lua

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// toGoValue converts the given LValue to a Go object.
func toGoValue(lv lua.LValue) interface{} {
	switch v := lv.(type) {
	case *lua.LNilType:
		return nil
	case lua.LBool:
		return bool(v)
	case lua.LString:
		return string(v)
	case lua.LNumber:
		return float64(v)
	case *lua.LTable:
		maxn := v.MaxN()
		if maxn == 0 { // table
			ret := Object{}
			v.ForEach(func(key, value lua.LValue) {
				keystr := fmt.Sprint(toGoValue(key))
				ret[keystr] = toGoValue(value)
			})
			return ret
		} else { // array
			ret := make([]interface{}, 0, maxn)
			for i := 1; i <= maxn; i++ {
				ret = append(ret, toGoValue(v.RawGetInt(i)))
			}
			return ret
		}
	default:
		return v
	}
}

func toLTable(o Object) *lua.LTable {
	t := &lua.LTable{}
	for k, v := range o {
		switch val := v.(type) {
		case bool:
			t.RawSetH(lua.LString(k), lua.LBool(val))
		case int:
			t.RawSetString(k, lua.LNumber(val))
		case string:
			t.RawSetString(k, lua.LString(val))
		case []string:
			list := &lua.LTable{}
			for _, v := range val {
				list.Append(lua.LString(v))
			}
			t.RawSetString(k, list)
		case []interface{}:
			list := &lua.LTable{}
			for _, v := range val {
				list.Append(toLTable(v.(Object)))
			}
			if list.Len() > 0 {
				t.RawSetString(k, list)
			}
		case map[string]interface{}:
			tbl := toLTable(val)
			t.RawSetString(k, tbl)
		case Object:
			tbl := toLTable(val)
			t.RawSetString(k, tbl)
		}
	}
	return t
}
