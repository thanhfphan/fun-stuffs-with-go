package be

// Environment ...
type Environment struct {
	store map[string]Object
}

// NewEnvironment ...
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Get ...
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set ...
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
