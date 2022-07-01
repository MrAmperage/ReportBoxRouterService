package DataContainers

type ResponseContainer[Data any] struct {
	Type      string
	Id        string
	Container Data
}
