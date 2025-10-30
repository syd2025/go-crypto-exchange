package tools

import (
	"errors"
	"reflect"

	"github.com/zeromicro/go-zero/core/logx"
)

func Default(data any) error {
	// 反射处理
	typeOf := reflect.TypeOf(data)
	valueOf := reflect.ValueOf(data)
	if typeOf.Kind() != reflect.Pointer {
		logx.Error("data must be a pointer")
		return errors.New("类型错误")
	}

	ele := typeOf.Elem()
	valueEle := valueOf.Elem()
	for i := 0; i < ele.NumField(); i++ {
		field := ele.Field(i)
		value := valueEle.Field(i)
		// field.Tag.Get("default")
		kind := field.Type.Kind()
		if kind == reflect.Int {
			value.Set(defaultInt())
		}
		if kind == reflect.Int32 {
			value.Set(defaultInt32())
		}
		if kind == reflect.Int64 {
			value.Set(defaultInt64())
		}
		if kind == reflect.String {
			value.Set(defaultString())
		}
		if kind == reflect.Float32 {
			value.Set(defaultFloat32())
		}
		if kind == reflect.Float64 {
			value.Set(defaultFloat64())
		}
	}
	return nil
}

func defaultFloat64() reflect.Value {
	var i float64 = 0
	return reflect.ValueOf(i)
}

func defaultFloat32() reflect.Value {
	var i float32 = 0
	return reflect.ValueOf(i)
}

func defaultString() reflect.Value {
	var i string = ""
	return reflect.ValueOf(i)
}

func defaultInt64() reflect.Value {
	var i int64 = -1
	return reflect.ValueOf(i)
}

func defaultInt32() reflect.Value {
	var i int32 = -1
	return reflect.ValueOf(i)
}

func defaultInt() reflect.Value {
	var i int = -1
	return reflect.ValueOf(i)
}
