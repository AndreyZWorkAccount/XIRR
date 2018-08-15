package secantAuto

//Border struct
type Border struct {
	left float64
	right float64
}

func NewBorder(left float64, right float64) Border  {
	return Border{left:left, right:right}
}


func (b *Border) Left() float64{
	return b.left
}

func (b *Border) Right() float64{
	return b.right
}
//end Border struct