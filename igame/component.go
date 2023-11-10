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
