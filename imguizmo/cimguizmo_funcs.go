// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imguizmo

import (
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/internal"
)

// #include "../imgui/extra_types.h"
// #include "cimguizmo_structs_accessor.h"
// #include "cimguizmo_wrapper.h"
import "C"

func AllowAxisFlip(value bool) {
	C.ImGuizmo_AllowAxisFlip(C.bool(value))
}

func BeginFrame() {
	C.ImGuizmo_BeginFrame()
}

func DecomposeMatrixToComponents(matrix *float32, translation *float32, rotation *float32, scale *float32) {
	matrixArg, matrixFin := internal.WrapNumberPtr[C.float, float32](matrix)
	translationArg, translationFin := internal.WrapNumberPtr[C.float, float32](translation)
	rotationArg, rotationFin := internal.WrapNumberPtr[C.float, float32](rotation)
	scaleArg, scaleFin := internal.WrapNumberPtr[C.float, float32](scale)
	C.ImGuizmo_DecomposeMatrixToComponents(matrixArg, translationArg, rotationArg, scaleArg)

	matrixFin()
	translationFin()
	rotationFin()
	scaleFin()
}

func DrawCubes(view *float32, projection *float32, matrices *float32, matrixCount int32) {
	viewArg, viewFin := internal.WrapNumberPtr[C.float, float32](view)
	projectionArg, projectionFin := internal.WrapNumberPtr[C.float, float32](projection)
	matricesArg, matricesFin := internal.WrapNumberPtr[C.float, float32](matrices)
	C.ImGuizmo_DrawCubes(viewArg, projectionArg, matricesArg, C.int(matrixCount))

	viewFin()
	projectionFin()
	matricesFin()
}

func DrawGrid(view *float32, projection *float32, matrix *float32, gridSize float32) {
	viewArg, viewFin := internal.WrapNumberPtr[C.float, float32](view)
	projectionArg, projectionFin := internal.WrapNumberPtr[C.float, float32](projection)
	matrixArg, matrixFin := internal.WrapNumberPtr[C.float, float32](matrix)
	C.ImGuizmo_DrawGrid(viewArg, projectionArg, matrixArg, C.float(gridSize))

	viewFin()
	projectionFin()
	matrixFin()
}

func Enable(enable bool) {
	C.ImGuizmo_Enable(C.bool(enable))
}

func GetIDPtr(ptr_id uintptr) imgui.ID {
	return *imgui.NewIDFromC(func() *C.ImGuiID { result := C.wrap_ImGuizmo_GetID_Ptr(C.uintptr_t(ptr_id)); return &result }())
}

// calculate unique ID (hash of whole ID stack + given parameter). e.g. if you want to query into ImGuiStorage yourself
func GetIDStr(str_id string) imgui.ID {
	str_idArg, str_idFin := internal.WrapString[C.char](str_id)

	defer func() {
		str_idFin()
	}()
	return *imgui.NewIDFromC(func() *C.ImGuiID { result := C.ImGuizmo_GetID_Str(str_idArg); return &result }())
}

func GetIDStrStr(str_id_begin string, str_id_end string) imgui.ID {
	str_id_beginArg, str_id_beginFin := internal.WrapString[C.char](str_id_begin)
	str_id_endArg, str_id_endFin := internal.WrapString[C.char](str_id_end)

	defer func() {
		str_id_beginFin()
		str_id_endFin()
	}()
	return *imgui.NewIDFromC(func() *C.ImGuiID { result := C.ImGuizmo_GetID_StrStr(str_id_beginArg, str_id_endArg); return &result }())
}

func GetStyle() *Style {
	return NewStyleFromC(C.ImGuizmo_GetStyle())
}

func IsOverFloatPtr(position *float32, pixelRadius float32) bool {
	positionArg, positionFin := internal.WrapNumberPtr[C.float, float32](position)

	defer func() {
		positionFin()
	}()
	return C.ImGuizmo_IsOver_FloatPtr(positionArg, C.float(pixelRadius)) == C.bool(true)
}

func IsOver() bool {
	return C.ImGuizmo_IsOver_Nil() == C.bool(true)
}

func IsOverOPERATION(op OPERATION) bool {
	return C.ImGuizmo_IsOver_OPERATION(C.OPERATION(op)) == C.bool(true)
}

func IsUsing() bool {
	return C.ImGuizmo_IsUsing() == C.bool(true)
}

func IsUsingAny() bool {
	return C.ImGuizmo_IsUsingAny() == C.bool(true)
}

func IsUsingViewManipulate() bool {
	return C.ImGuizmo_IsUsingViewManipulate() == C.bool(true)
}

func IsViewManipulateHovered() bool {
	return C.ImGuizmo_IsViewManipulateHovered() == C.bool(true)
}

// ManipulateV parameter default value hint:
// deltaMatrix: NULL
// snap: NULL
// localBounds: NULL
// boundsSnap: NULL
func ManipulateV(view *float32, projection *float32, operation OPERATION, mode MODE, matrix *float32, deltaMatrix *float32, snap *float32, localBounds *float32, boundsSnap *float32) bool {
	viewArg, viewFin := internal.WrapNumberPtr[C.float, float32](view)
	projectionArg, projectionFin := internal.WrapNumberPtr[C.float, float32](projection)
	matrixArg, matrixFin := internal.WrapNumberPtr[C.float, float32](matrix)
	deltaMatrixArg, deltaMatrixFin := internal.WrapNumberPtr[C.float, float32](deltaMatrix)
	snapArg, snapFin := internal.WrapNumberPtr[C.float, float32](snap)
	localBoundsArg, localBoundsFin := internal.WrapNumberPtr[C.float, float32](localBounds)
	boundsSnapArg, boundsSnapFin := internal.WrapNumberPtr[C.float, float32](boundsSnap)

	defer func() {
		viewFin()
		projectionFin()
		matrixFin()
		deltaMatrixFin()
		snapFin()
		localBoundsFin()
		boundsSnapFin()
	}()
	return C.ImGuizmo_Manipulate(viewArg, projectionArg, C.OPERATION(operation), C.MODE(mode), matrixArg, deltaMatrixArg, snapArg, localBoundsArg, boundsSnapArg) == C.bool(true)
}

// pop from the ID stack.
func PopID() {
	C.ImGuizmo_PopID()
}

// push integer into the ID stack (will hash integer).
func PushIDInt(int_id int32) {
	C.ImGuizmo_PushID_Int(C.int(int_id))
}

// push pointer into the ID stack (will hash pointer).
func PushIDPtr(ptr_id uintptr) {
	C.wrap_ImGuizmo_PushID_Ptr(C.uintptr_t(ptr_id))
}

// push string into the ID stack (will hash string).
func PushIDStr(str_id string) {
	str_idArg, str_idFin := internal.WrapString[C.char](str_id)
	C.ImGuizmo_PushID_Str(str_idArg)

	str_idFin()
}

// push string into the ID stack (will hash string).
func PushIDStrStr(str_id_begin string, str_id_end string) {
	str_id_beginArg, str_id_beginFin := internal.WrapString[C.char](str_id_begin)
	str_id_endArg, str_id_endFin := internal.WrapString[C.char](str_id_end)
	C.ImGuizmo_PushID_StrStr(str_id_beginArg, str_id_endArg)

	str_id_beginFin()
	str_id_endFin()
}

func RecomposeMatrixFromComponents(translation *float32, rotation *float32, scale *float32, matrix *float32) {
	translationArg, translationFin := internal.WrapNumberPtr[C.float, float32](translation)
	rotationArg, rotationFin := internal.WrapNumberPtr[C.float, float32](rotation)
	scaleArg, scaleFin := internal.WrapNumberPtr[C.float, float32](scale)
	matrixArg, matrixFin := internal.WrapNumberPtr[C.float, float32](matrix)
	C.ImGuizmo_RecomposeMatrixFromComponents(translationArg, rotationArg, scaleArg, matrixArg)

	translationFin()
	rotationFin()
	scaleFin()
	matrixFin()
}

func SetAlternativeWindow(window *imgui.Window) {
	windowArg, windowFin := window.Handle()
	C.ImGuizmo_SetAlternativeWindow(internal.ReinterpretCast[*C.ImGuiWindow](windowArg))

	windowFin()
}

func SetAxisLimit(value float32) {
	C.ImGuizmo_SetAxisLimit(C.float(value))
}

func SetAxisMask(x bool, y bool, z bool) {
	C.ImGuizmo_SetAxisMask(C.bool(x), C.bool(y), C.bool(z))
}

// SetDrawlistV parameter default value hint:
// drawlist: nullptr
func SetDrawlistV(drawlist *imgui.DrawList) {
	drawlistArg, drawlistFin := drawlist.Handle()
	C.ImGuizmo_SetDrawlist(internal.ReinterpretCast[*C.ImDrawList](drawlistArg))

	drawlistFin()
}

func SetGizmoSizeClipSpace(value float32) {
	C.ImGuizmo_SetGizmoSizeClipSpace(C.float(value))
}

func SetID(id int32) {
	C.ImGuizmo_SetID(C.int(id))
}

func SetImGuiContext(ctx *imgui.Context) {
	ctxArg, ctxFin := ctx.Handle()
	C.ImGuizmo_SetImGuiContext(internal.ReinterpretCast[*C.ImGuiContext](ctxArg))

	ctxFin()
}

func SetOrthographic(isOrthographic bool) {
	C.ImGuizmo_SetOrthographic(C.bool(isOrthographic))
}

func SetPlaneLimit(value float32) {
	C.ImGuizmo_SetPlaneLimit(C.float(value))
}

func SetRect(x float32, y float32, width float32, height float32) {
	C.ImGuizmo_SetRect(C.float(x), C.float(y), C.float(width), C.float(height))
}

func ViewManipulateFloat(view *float32, length float32, position imgui.Vec2, size imgui.Vec2, backgroundColor uint32) {
	viewArg, viewFin := internal.WrapNumberPtr[C.float, float32](view)
	C.ImGuizmo_ViewManipulate_Float(viewArg, C.float(length), internal.ReinterpretCast[C.ImVec2](position.ToC()), internal.ReinterpretCast[C.ImVec2](size.ToC()), C.ImU32(backgroundColor))

	viewFin()
}

func ViewManipulateFloatPtr(view *float32, projection *float32, operation OPERATION, mode MODE, matrix *float32, length float32, position imgui.Vec2, size imgui.Vec2, backgroundColor uint32) {
	viewArg, viewFin := internal.WrapNumberPtr[C.float, float32](view)
	projectionArg, projectionFin := internal.WrapNumberPtr[C.float, float32](projection)
	matrixArg, matrixFin := internal.WrapNumberPtr[C.float, float32](matrix)
	C.ImGuizmo_ViewManipulate_FloatPtr(viewArg, projectionArg, C.OPERATION(operation), C.MODE(mode), matrixArg, C.float(length), internal.ReinterpretCast[C.ImVec2](position.ToC()), internal.ReinterpretCast[C.ImVec2](size.ToC()), C.ImU32(backgroundColor))

	viewFin()
	projectionFin()
	matrixFin()
}

func NewStyle() *Style {
	return NewStyleFromC(C.Style_Style())
}

func (self *Style) Destroy() {
	selfArg, selfFin := self.Handle()
	C.Style_destroy(internal.ReinterpretCast[*C.Style](selfArg))

	selfFin()
}

func Manipulate(view *float32, projection *float32, operation OPERATION, mode MODE, matrix *float32) bool {
	viewArg, viewFin := internal.WrapNumberPtr[C.float, float32](view)
	projectionArg, projectionFin := internal.WrapNumberPtr[C.float, float32](projection)
	matrixArg, matrixFin := internal.WrapNumberPtr[C.float, float32](matrix)

	defer func() {
		viewFin()
		projectionFin()
		matrixFin()
	}()
	return C.wrap_ImGuizmo_Manipulate(viewArg, projectionArg, C.OPERATION(operation), C.MODE(mode), matrixArg) == C.bool(true)
}

func SetDrawlist() {
	C.wrap_ImGuizmo_SetDrawlist()
}

func (self Style) SetTranslationLineThickness(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Style_SetTranslationLineThickness(selfArg, C.float(v))
}

func (self *Style) TranslationLineThickness() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_Style_GetTranslationLineThickness(internal.ReinterpretCast[*C.Style](selfArg)))
}

func (self Style) SetTranslationLineArrowSize(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Style_SetTranslationLineArrowSize(selfArg, C.float(v))
}

func (self *Style) TranslationLineArrowSize() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_Style_GetTranslationLineArrowSize(internal.ReinterpretCast[*C.Style](selfArg)))
}

func (self Style) SetRotationLineThickness(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Style_SetRotationLineThickness(selfArg, C.float(v))
}

func (self *Style) RotationLineThickness() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_Style_GetRotationLineThickness(internal.ReinterpretCast[*C.Style](selfArg)))
}

func (self Style) SetRotationOuterLineThickness(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Style_SetRotationOuterLineThickness(selfArg, C.float(v))
}

func (self *Style) RotationOuterLineThickness() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_Style_GetRotationOuterLineThickness(internal.ReinterpretCast[*C.Style](selfArg)))
}

func (self Style) SetScaleLineThickness(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Style_SetScaleLineThickness(selfArg, C.float(v))
}

func (self *Style) ScaleLineThickness() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_Style_GetScaleLineThickness(internal.ReinterpretCast[*C.Style](selfArg)))
}

func (self Style) SetScaleLineCircleSize(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Style_SetScaleLineCircleSize(selfArg, C.float(v))
}

func (self *Style) ScaleLineCircleSize() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_Style_GetScaleLineCircleSize(internal.ReinterpretCast[*C.Style](selfArg)))
}

func (self Style) SetHatchedAxisLineThickness(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Style_SetHatchedAxisLineThickness(selfArg, C.float(v))
}

func (self *Style) HatchedAxisLineThickness() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_Style_GetHatchedAxisLineThickness(internal.ReinterpretCast[*C.Style](selfArg)))
}

func (self Style) SetCenterCircleSize(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Style_SetCenterCircleSize(selfArg, C.float(v))
}

func (self *Style) CenterCircleSize() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_Style_GetCenterCircleSize(internal.ReinterpretCast[*C.Style](selfArg)))
}

func (self Style) SetColors(v *[15]imgui.Vec4) {
	vArg := make([]C.ImVec4, len(v))
	for i, vV := range v {
		vArg[i] = internal.ReinterpretCast[C.ImVec4](vV.ToC())
	}

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Style_SetColors(selfArg, (*C.ImVec4)(&vArg[0]))

	for i, vV := range vArg {
		(*v)[i] = func() imgui.Vec4 { out := vV; return *(&imgui.Vec4{}).FromC(unsafe.Pointer(&out)) }()
	}
}

func (self *Style) Colors() [15]imgui.Vec4 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() [15]imgui.Vec4 {
		result := [15]imgui.Vec4{}
		resultMirr := C.wrap_Style_GetColors(internal.ReinterpretCast[*C.Style](selfArg))
		for i := range result {
			result[i] = func() imgui.Vec4 {
				out := C.cimguizmo_ImVec4_GetAtIdx(resultMirr, C.int(i))
				return *(&imgui.Vec4{}).FromC(unsafe.Pointer(&out))
			}()
		}

		return result
	}()
}
