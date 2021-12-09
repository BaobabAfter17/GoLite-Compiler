package typeSys

type Type interface {
	Literal() string
	Equals(Type) bool
}

type StructType struct {
	Name   string
	Fields map[string]Type
}

func (strct *StructType) Insert(field string, typ Type) bool {
	strct.Fields[field] = typ
	return true
}

func (strct *StructType) Contains(name string) bool {
	_, ok := strct.Fields[name]
	return ok
}

func (st *StructType) Literal() string { return st.Name }

// Naming Equivalence between two StructType
func (st1 *StructType) Equals(t2 Type) bool { 
	switch t2.(type) {
	case *StructType:
		st2 := t2.(*StructType)
		return st1.Name == st2.Name 
	case *voidType:
		return true
	default:
		return false
	}
}

// ===========================================================================
// BaseTypes, including int, bool, void and unknown, only export their singletons
type intType struct {
	Type
}

func (it *intType) Literal() string { return "int" }
func (it *intType) Equals(t2 Type) bool { 
	switch t2.(type) {
	case *intType:
		return true
	default:
		return false
	}
}

var IntType = &intType{}

type boolType struct {
	Type
}

func (bt *boolType) Literal() string { return "bool" }
func (bt *boolType) Equals(t2 Type) bool { 
	switch t2.(type) {
	case *boolType:
		return true
	default:
		return false
	}
}

var BoolType = &boolType{}

type voidType struct {
	Type
}

func (vt *voidType) Literal() string { return "void" }
func (vt *voidType) Equals(t2 Type) bool { 
	switch t2.(type) {
	case *voidType:
		return true
	default:
		return false
	}
}

var VoidType = &voidType{}

type unknownType struct {
	Type
}

func (ut *unknownType) Literal() string { return "unknown" }
func (ut *unknownType) Equals(t2 Type) bool { 
	return false
}

var UnknownType = &unknownType{}
