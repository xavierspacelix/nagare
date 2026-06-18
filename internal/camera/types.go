package camera

type Info struct {
	ID   int
	Name string
}

type Frame struct {
	Data   []byte
	Width  int
	Height int
}
