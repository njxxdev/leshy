package leshy_component

type Component interface {
	GetInstance() Component
	GetName() string
}
