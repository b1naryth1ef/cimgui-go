// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imnodes

// #include <stdlib.h>
// #include <memory.h>
// #include "../imgui/extra_types.h"
// #include "cimnodes_wrapper.h"
// #include "cimnodes_typedefs.h"
// extern void callbackMiniMapNodeHoveringCallback0(int, void*);
// extern void callbackMiniMapNodeHoveringCallback1(int, void*);
// extern void callbackMiniMapNodeHoveringCallback2(int, void*);
// extern void callbackMiniMapNodeHoveringCallback3(int, void*);
// extern void callbackMiniMapNodeHoveringCallback4(int, void*);
// extern void callbackMiniMapNodeHoveringCallback5(int, void*);
// extern void callbackMiniMapNodeHoveringCallback6(int, void*);
// extern void callbackMiniMapNodeHoveringCallback7(int, void*);
// extern void callbackMiniMapNodeHoveringCallback8(int, void*);
// extern void callbackMiniMapNodeHoveringCallback9(int, void*);
// extern void callbackMiniMapNodeHoveringCallback10(int, void*);
// extern void callbackMiniMapNodeHoveringCallback11(int, void*);
// extern void callbackMiniMapNodeHoveringCallback12(int, void*);
// extern void callbackMiniMapNodeHoveringCallback13(int, void*);
// extern void callbackMiniMapNodeHoveringCallback14(int, void*);
// extern void callbackMiniMapNodeHoveringCallback15(int, void*);
// extern void callbackMiniMapNodeHoveringCallback16(int, void*);
// extern void callbackMiniMapNodeHoveringCallback17(int, void*);
// extern void callbackMiniMapNodeHoveringCallback18(int, void*);
// extern void callbackMiniMapNodeHoveringCallback19(int, void*);
// extern void callbackMiniMapNodeHoveringCallback20(int, void*);
// extern void callbackMiniMapNodeHoveringCallback21(int, void*);
// extern void callbackMiniMapNodeHoveringCallback22(int, void*);
// extern void callbackMiniMapNodeHoveringCallback23(int, void*);
// extern void callbackMiniMapNodeHoveringCallback24(int, void*);
// extern void callbackMiniMapNodeHoveringCallback25(int, void*);
// extern void callbackMiniMapNodeHoveringCallback26(int, void*);
// extern void callbackMiniMapNodeHoveringCallback27(int, void*);
// extern void callbackMiniMapNodeHoveringCallback28(int, void*);
// extern void callbackMiniMapNodeHoveringCallback29(int, void*);
// extern void callbackMiniMapNodeHoveringCallback30(int, void*);
// extern void callbackMiniMapNodeHoveringCallback31(int, void*);
import "C"
import (
	"unsafe"

	"github.com/AllenDang/cimgui-go/internal"
)

type (
	MiniMapNodeHoveringCallback  func(arg0 int32, arg1 unsafe.Pointer)
	cMiniMapNodeHoveringCallback func(arg0 C.int, arg1 unsafe.Pointer)
)

func NewMiniMapNodeHoveringCallbackFromC(cvalue *C.ImNodesMiniMapNodeHoveringCallback) *MiniMapNodeHoveringCallback {
	result := poolMiniMapNodeHoveringCallback.Find(*cvalue)
	return &result
}

func (c MiniMapNodeHoveringCallback) C() (C.ImNodesMiniMapNodeHoveringCallback, func()) {
	return poolMiniMapNodeHoveringCallback.Allocate(c), func() {}
}

func wrapMiniMapNodeHoveringCallback(cb MiniMapNodeHoveringCallback, arg0 C.int, arg1 unsafe.Pointer) {
	cb(int32(arg0), unsafe.Pointer(arg1))
}

//export callbackMiniMapNodeHoveringCallback0
func callbackMiniMapNodeHoveringCallback0(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(0), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback1
func callbackMiniMapNodeHoveringCallback1(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(1), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback2
func callbackMiniMapNodeHoveringCallback2(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(2), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback3
func callbackMiniMapNodeHoveringCallback3(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(3), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback4
func callbackMiniMapNodeHoveringCallback4(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(4), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback5
func callbackMiniMapNodeHoveringCallback5(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(5), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback6
func callbackMiniMapNodeHoveringCallback6(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(6), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback7
func callbackMiniMapNodeHoveringCallback7(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(7), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback8
func callbackMiniMapNodeHoveringCallback8(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(8), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback9
func callbackMiniMapNodeHoveringCallback9(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(9), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback10
func callbackMiniMapNodeHoveringCallback10(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(10), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback11
func callbackMiniMapNodeHoveringCallback11(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(11), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback12
func callbackMiniMapNodeHoveringCallback12(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(12), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback13
func callbackMiniMapNodeHoveringCallback13(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(13), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback14
func callbackMiniMapNodeHoveringCallback14(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(14), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback15
func callbackMiniMapNodeHoveringCallback15(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(15), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback16
func callbackMiniMapNodeHoveringCallback16(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(16), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback17
func callbackMiniMapNodeHoveringCallback17(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(17), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback18
func callbackMiniMapNodeHoveringCallback18(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(18), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback19
func callbackMiniMapNodeHoveringCallback19(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(19), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback20
func callbackMiniMapNodeHoveringCallback20(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(20), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback21
func callbackMiniMapNodeHoveringCallback21(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(21), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback22
func callbackMiniMapNodeHoveringCallback22(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(22), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback23
func callbackMiniMapNodeHoveringCallback23(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(23), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback24
func callbackMiniMapNodeHoveringCallback24(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(24), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback25
func callbackMiniMapNodeHoveringCallback25(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(25), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback26
func callbackMiniMapNodeHoveringCallback26(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(26), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback27
func callbackMiniMapNodeHoveringCallback27(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(27), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback28
func callbackMiniMapNodeHoveringCallback28(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(28), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback29
func callbackMiniMapNodeHoveringCallback29(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(29), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback30
func callbackMiniMapNodeHoveringCallback30(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(30), arg0, arg1)
}

//export callbackMiniMapNodeHoveringCallback31
func callbackMiniMapNodeHoveringCallback31(arg0 C.int, arg1 unsafe.Pointer) {
	wrapMiniMapNodeHoveringCallback(poolMiniMapNodeHoveringCallback.Get(31), arg0, arg1)
}

var poolMiniMapNodeHoveringCallback *internal.Pool[MiniMapNodeHoveringCallback, C.ImNodesMiniMapNodeHoveringCallback]

func init() {
	poolMiniMapNodeHoveringCallback = internal.NewPool[MiniMapNodeHoveringCallback, C.ImNodesMiniMapNodeHoveringCallback](
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback0),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback1),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback2),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback3),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback4),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback5),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback6),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback7),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback8),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback9),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback10),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback11),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback12),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback13),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback14),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback15),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback16),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback17),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback18),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback19),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback20),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback21),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback22),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback23),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback24),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback25),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback26),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback27),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback28),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback29),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback30),
		C.ImNodesMiniMapNodeHoveringCallback(C.callbackMiniMapNodeHoveringCallback31),
	)
}

func ClearMiniMapNodeHoveringCallbackPool() {
	poolMiniMapNodeHoveringCallback.Clear()
}
