pacakge config

type SimpleConfig struct {
    data map[string]interface{}
    observers map[string][]chan interface{}
    observer Observer
}

func (c *SimpleConfig) GetValue(key string) interface{} {
}

func (c *SimpleConfig) SetValue(key string, value interface{}) {
    c[key] = value
}

func (c *SimpleConfig) AddObserver(key string) chan interface{} {
    ch := make(chan interface{})
    observers[key] = append(observers[key], ch) 
    return ch
}

