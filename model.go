package main

type Model struct {
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) HasPixel(pos PositionF64, settings *AppSettings) bool {
	return false
}
