package DataContainers

type NaimedTable[Table any] struct {
	Type  string
	Name  string
	Table []Table
}
