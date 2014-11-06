package gowebsockets

type ArcParameters struct {
	X, Y, R          float64
	Counterclockwise bool
}

type Vector struct {
	X, Y float64
}

type LineParameters struct {
	Start Vector
	End   Vector
}

type SetSizeParameters struct {
	Width, Height int
}

type Message struct {
	DrawId     string
	Parameters interface{}
}

func Clear() {
	msg := Message{DrawId: "Clear"}
	h.broadcast <- msg
}

func Line(params LineParameters) {
	msg := Message{DrawId: "Line", Parameters: params}
	h.broadcast <- msg
}

func Arc(params ArcParameters) {
	msg := Message{DrawId: "Arc", Parameters: params}
	h.broadcast <- msg
}

func SetSize(width, height int) {
	params := SetSizeParameters{Width: width, Height: height}
	h.broadcast <- Message{DrawId: "SetSize", Parameters: params}
}
