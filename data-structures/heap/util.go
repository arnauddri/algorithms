package heap

type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}
