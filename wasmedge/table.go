package wasmedge

// #include <wasmedge.h>
import "C"

type Table struct {
	_inner *C.WasmEdge_TableInstanceContext
}

func NewTable(rtype RefType, lim *Limit) *Table {
	climit := C.WasmEdge_Limit{HasMax: C.bool(lim.hasmax), Min: C.uint32_t(lim.min), Max: C.uint32_t(lim.max)}
	crtype := C.enum_WasmEdge_RefType(rtype)
	var tabTypeCxt *C.WasmEdge_TableTypeContext = C.WasmEdge_TableTypeCreate(crtype, climit)
	self := &Table{
		_inner: C.WasmEdge_TableInstanceCreate(tabTypeCxt),
	}
	if self._inner == nil {
		return nil
	}
	return self
}

func (self *Table) GetRefType() RefType {
	tabTypeCxt := C.WasmEdge_TableInstanceGetTableType(self._inner);
	return RefType(C.WasmEdge_TableTypeGetRefType(tabTypeCxt))
}

func (self *Table) GetData(off uint) (interface{}, error) {
	cval := C.WasmEdge_Value{}
	res := C.WasmEdge_TableInstanceGetData(self._inner, &cval, C.uint32_t(off))
	if !C.WasmEdge_ResultOK(res) {
		return nil, newError(res)
	}
	return fromWasmEdgeValue(cval, cval.Type), nil
}

func (self *Table) SetData(data interface{}, off uint) error {
	cval := toWasmEdgeValue(data)
	res := C.WasmEdge_TableInstanceSetData(self._inner, cval, C.uint32_t(off))
	if !C.WasmEdge_ResultOK(res) {
		return newError(res)
	}
	return nil
}

func (self *Table) GetSize() uint {
	return uint(C.WasmEdge_TableInstanceGetSize(self._inner))
}

func (self *Table) Grow(size uint) error {
	res := C.WasmEdge_TableInstanceGrow(self._inner, C.uint32_t(size))
	if !C.WasmEdge_ResultOK(res) {
		return newError(res)
	}
	return nil
}

func (self *Table) Delete() {
	C.WasmEdge_TableInstanceDelete(self._inner)
	self._inner = nil
}
