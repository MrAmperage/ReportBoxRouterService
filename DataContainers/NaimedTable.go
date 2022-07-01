package DataContainers

type NaimedTable[Table any] struct {
	Name  string
	Table []Table
}
