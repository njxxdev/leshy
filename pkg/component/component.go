package component

type Component interface {
	GetInstance() Component
	GetName() string
}
