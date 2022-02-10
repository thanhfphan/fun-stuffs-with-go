package be

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

// ObjectType ...
type ObjectType string
type BuiltinFunction func(args ...Object) Object

const (
	NULL_OBJ    = "NULL"
	ERROR_OBJ   = "ERROR"
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	STRING_OBJ  = "STRING"

	BUILTIN_OBJ = "BUILTIN"

	ARRAY_OBJ = "ARRAY"

	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

// Object ...
type Object interface {
	Type() ObjectType
	ToString() string
}

// HashKey ...
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// Null ...
type Null struct{}

// Type ...
func (n *Null) Type() ObjectType { return NULL_OBJ }

// ToString ...
func (n *Null) ToString() string { return "null" }

// ReturnValue ...
type ReturnValue struct {
	Value Object
}

// Type ...
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// ToString ...
func (rv *ReturnValue) ToString() string { return rv.Value.ToString() }

// Error ...
type Error struct {
	Message string
}

// Type ...
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// ToString ...
func (e *Error) ToString() string { return "ERROR: " + e.Message }

// Integer ...
type Integer struct {
	Value int64
}

// Type ...
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// ToString ...
func (i *Integer) ToString() string { return fmt.Sprintf("%d", i.Value) }

// HashKey ...
func (i *Integer) HashKey() HashKey { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }

// Boolean ...
type Boolean struct {
	Value bool
}

// Type ...
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// ToString ...
func (b *Boolean) ToString() string { return fmt.Sprintf("%t", b.Value) }

// HashKey ...
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// String ...
type String struct {
	Value string
}

// Type ...
func (s *String) Type() ObjectType { return STRING_OBJ }

// ToString ...
func (s *String) ToString() string { return s.Value }

// HashKey ...
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// Builtin ...
type Builtin struct {
	Fn BuiltinFunction
}

// Type ...
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }

// ToString ...
func (b *Builtin) ToString() string { return "builtin function" }

// Array ...
type Array struct {
	Elements []Object
}

// Type ...
func (a *Array) Type() ObjectType { return ARRAY_OBJ }

// ToString ...
func (a *Array) ToString() string {
	var elements []string
	for _, e := range a.Elements {
		elements = append(elements, e.ToString())
	}

	var out bytes.Buffer
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ","))
	out.WriteString("}")

	return out.String()
}
