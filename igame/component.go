package igame

type IComponent interface {
	OnInit(ctx interface{})
	OnLoad(ctx interface{})
	ComName() string
	OnActive() error
	OnDeActive() error
	CustomizeSave() bool
	CustomizeLoad() bool
	Load(ctx interface{})
	Save(ctx interface{})
	ToJson() []byte
	FromJson(data []byte)
}

var componentRoot ComponentRoot

type ComponentRoot struct {
	ComponentRef interface{}
	ComponentMap map[string]IComponent
}

func InitComponentRoot(root interface{}) {
	componentRoot = ComponentRoot{}
	componentRoot.ComponentRef = root
	componentRoot.ComponentMap = make(map[string]IComponent)
	//val := reflect.Indirect(reflect.ValueOf(root))
	//for i := 0; i < val.NumField(); i++ {
	//	fieldValue := val.Field(i)
	//	if !fieldValue.CanInterface() {
	//		continue
	//	}
	//	iCom := fieldValue.Interface().(IComponent)
	//	componentRoot.ComponentMap[iCom.ComName()] = iCom
	//}
}

func NewComponents() ComponentRoot {
	newCom := componentRoot
	return newCom
}
