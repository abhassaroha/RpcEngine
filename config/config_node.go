package config

type ConfigNode struct {
    key interface{}
    value interface{}
    children []*ConfigNode
}

func (ConfigNode *c) GetKey() interface{} {
    key := c.key
    return key
}

func (ConfigNode *c) GetValue() interface{} {
    value := c.value
    return value
}

func (ConfigNode *c) AddChild(child *ConfigNode) {
    c.children = append(c.children, child)
} 
