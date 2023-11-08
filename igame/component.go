package igame

type IComponent interface {
	OnInit(ctx interface{})
	OnLoad(ctx interface{})
	OnSave(ctx interface{})
	ComName() string
	OnActive() error
	OnDeActive() error
	CustomizeSave() bool
}

var componentRoot ComponentRoot

type ComponentRoot struct {
	ComponentRef interface{}
}

func InitComponentRoot(root interface{}) {
	componentRoot = ComponentRoot{}
	componentRoot.ComponentRef = root
}

func NewComponents() ComponentRoot {
	newCom := componentRoot
	return newCom
}
