package igame

import (
	"reflect"
)

type IComponent interface {
	OnLoad()
	OnSave()
	SetComName(name string)
	OnActive()
	OnDeActive()
}

type SComBase struct {
	ComName string `json:"-"`
}

func (S SComBase) OnLoad() {
}

func (S SComBase) OnSave() {
}

func (S SComBase) SetComName(name string) {

}

func (S SComBase) OnActive() {
}

func (S SComBase) OnDeActive() {
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
		fieldType := val.Type().Field(i)
		fieldValue := val.Field(i)
		tag := fieldType.Tag.Get("component")
		// 判断字段类型是否为 combase
		if tag == "" {
			// 将字段名称赋值给 ComName 字段
			continue
		}
		if !fieldValue.CanInterface() {
			continue
		}
		iCom := fieldValue.Interface().(IComponent)
		iCom.SetComName(fieldType.Name)
		componentRoot.Components[fieldType.Name] = iCom
	}

}

func NewComponents() ComponentRoot {
	newCom := componentRoot
	return newCom
}
