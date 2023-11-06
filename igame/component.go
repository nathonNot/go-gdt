package igame

import (
	"reflect"
)

type IComponent interface {
	OnInit(ctx interface{})
	OnLoad(ctx interface{})
	OnSave(ctx interface{})
	ComName() string
	OnActive() error
	OnDeActive() error
}

var componentRoot ComponentRoot

type ComponentRoot struct {
	Components   map[string]IComponent
	ComponentRef interface{}
}

func InitComponentRoot(root interface{}) {
	componentRoot = ComponentRoot{}
	componentRoot.Components = make(map[string]IComponent)
	componentRoot.ComponentRef = root
	val := reflect.Indirect(reflect.ValueOf(root))
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		// 判断字段类型是否为 combase
		// 将字段名称赋值给 ComName 字段
		if !fieldValue.CanInterface() {
			continue
		}
		iCom := fieldValue.Interface().(IComponent)
		componentRoot.Components[iCom.ComName()] = iCom
	}

}

func NewComponents() ComponentRoot {
	newCom := componentRoot
	return newCom
}
