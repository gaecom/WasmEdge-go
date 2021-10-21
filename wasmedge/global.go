package wasmedge

// #include <wasmedge.h>
import "C"

type Global struct {
	_inner *C.WasmEdge_GlobalInstanceContext
}

func NewGlobal(val interface{}, vtype ValMut) *Global {
	cval := toWasmEdgeValue(val)
	var globTypeCxt *C.WasmEdge_GlobalTypeContext = C.WasmEdge_GlobalTypeCreate(cval.Type, C.enum_WasmEdge_Mutability(vtype));
	self := &Global{
		_inner: C.WasmEdge_GlobalInstanceCreate(globTypeCxt, cval),
	}
	if self._inner == nil {
		return nil
	}
	return self
}

func (self *Global) GetValType() ValType {
	var globTypeCxt *C.WasmEdge_GlobalTypeContext = C.WasmEdge_GlobalInstanceGetGlobalType(self._inner);
	return ValType(C.WasmEdge_GlobalTypeGetValType(globTypeCxt))
}

func (self *Global) GetMutability() ValMut {
	var globTypeCxt *C.WasmEdge_GlobalTypeContext = C.WasmEdge_GlobalInstanceGetGlobalType(self._inner);
	return ValMut(C.WasmEdge_GlobalTypeGetMutability(globTypeCxt))
}

func (self *Global) GetValue() interface{} {
	cval := C.WasmEdge_GlobalInstanceGetValue(self._inner)
	return fromWasmEdgeValue(cval, cval.Type)
}

func (self *Global) SetValue(val interface{}) {
	C.WasmEdge_GlobalInstanceSetValue(self._inner, toWasmEdgeValue(val))
}

func (self *Global) Delete() {
	C.WasmEdge_GlobalInstanceDelete(self._inner)
	self._inner = nil
}
