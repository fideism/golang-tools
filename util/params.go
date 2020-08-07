package util

// Params map[string]interface{}
type Params map[string]interface{}

// Set 设置值
func (p Params) Set(k string, v interface{}) Params {
	p[k] = v

	return p
}

// Get 获取值
func (p Params) Get(k string) (v interface{}) {
	v, _ = p[k]

	return
}

// GetString 强制获取k对应的v string类型
func (p Params) GetString(k string) string {
	v, _ := p[k]

	return InterfaceToString(v)
}

// Exists 判断是否存在
func (p Params) Exists(k string) bool {
	_, ok := p[k]

	return ok
}

// InterfaceToString 不定类型强制装换为string
func InterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	key = value.(string)
	return key
}
