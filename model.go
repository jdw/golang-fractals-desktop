package main

type Model struct {
}

func NewModel() *Model {
	return &Model{}
}

var kaka = 0

func (m *Model) HasPixel(pos PositionI64, settings *AppSettings) bool {

	return false
}
