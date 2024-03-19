package leshy_component

type Component interface {
	Instance() Component
	Name() string
}
