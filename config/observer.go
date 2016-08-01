package config

type Callback func(string)

type Observer interface {
	ObserveNode(path string, callback Callback)
}
