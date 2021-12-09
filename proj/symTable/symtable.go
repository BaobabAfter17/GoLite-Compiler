package symTable

import (
	"bytes"
	tp "proj/type"
	"reflect"
)

type Entry interface {
}

type VarEntry struct {
	Typ tp.Type
	Reg int
}

type StructEntry struct {
	Typ *tp.StructType
}

type FuncEntry struct {
	Scope *SymbolTable
	ParaNames []string
	RetTyp tp.Type
}

type SymbolTable struct {
	Attributes map[string]Entry
	Parent     *SymbolTable
}

func (st *SymbolTable) String() string {
	var out bytes.Buffer

	for name, val := range st.Attributes {
		out.WriteString(name + " : ")
		out.WriteString(reflect.TypeOf(val).String())
		out.WriteString("\n")
	}

	return out.String()
}

func (st *SymbolTable) LookUpLocal(name string) Entry {
	entry, _ := st.Attributes[name]
	return entry
}

func (st *SymbolTable) LookUpAll(name string) Entry {
	// var curr *SymbolTable
	curr := st
	for curr != nil {
		if entry, ok := curr.Attributes[name]; ok {
			return entry
		}
		curr = curr.Parent
	}

	return nil
}

func (st *SymbolTable) Insert(name string, record Entry) bool {
	st.Attributes[name] = record
	return true
}

func (st *SymbolTable) InitializeScope(name string) *SymbolTable {
	newTable := &SymbolTable{
		Attributes: make(map[string]Entry),
		Parent:     st,
	}
	entry := &FuncEntry{
		Scope: newTable,
	}
	st.Insert(name, entry)
	return newTable
}

func (st *SymbolTable) FinalizeScope() *SymbolTable {
	return st.Parent
}


