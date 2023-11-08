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
