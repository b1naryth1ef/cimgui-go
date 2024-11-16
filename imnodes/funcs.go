// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imnodes

import (
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/internal"
)

// #include "structs_accessor.h"
// #include "wrapper.h"
// #include "../imgui/extra_types.h"
import "C"

func NewEmulateThreeButtonMouse() *EmulateThreeButtonMouse {
	return NewEmulateThreeButtonMouseFromC(C.EmulateThreeButtonMouse_EmulateThreeButtonMouse())
}

func (self *EmulateThreeButtonMouse) Destroy() {
	selfArg, selfFin := self.Handle()
	C.EmulateThreeButtonMouse_destroy(internal.ReinterpretCast[*C.EmulateThreeButtonMouse](selfArg))

	selfFin()
}

func NewIO() *IO {
	return NewIOFromC(C.ImNodesIO_ImNodesIO())
}

func (self *IO) Destroy() {
	selfArg, selfFin := self.Handle()
	C.ImNodesIO_destroy(internal.ReinterpretCast[*C.ImNodesIO](selfArg))

	selfFin()
}

func NewStyle() *Style {
	return NewStyleFromC(C.ImNodesStyle_ImNodesStyle())
}

func (self *Style) Destroy() {
	selfArg, selfFin := self.Handle()
	C.ImNodesStyle_destroy(internal.ReinterpretCast[*C.ImNodesStyle](selfArg))

	selfFin()
}

func NewLinkDetachWithModifierClick() *LinkDetachWithModifierClick {
	return NewLinkDetachWithModifierClickFromC(C.LinkDetachWithModifierClick_LinkDetachWithModifierClick())
}

func (self *LinkDetachWithModifierClick) Destroy() {
	selfArg, selfFin := self.Handle()
	C.LinkDetachWithModifierClick_destroy(internal.ReinterpretCast[*C.LinkDetachWithModifierClick](selfArg))

	selfFin()
}

func NewMultipleSelectModifier() *MultipleSelectModifier {
	return NewMultipleSelectModifierFromC(C.MultipleSelectModifier_MultipleSelectModifier())
}

func (self *MultipleSelectModifier) Destroy() {
	selfArg, selfFin := self.Handle()
	C.MultipleSelectModifier_destroy(internal.ReinterpretCast[*C.MultipleSelectModifier](selfArg))

	selfFin()
}

// BeginInputAttributeV parameter default value hint:
// shape: ImNodesPinShape_CircleFilled
func BeginInputAttributeV(id int32, shape PinShape) {
	C.imnodes_BeginInputAttribute(C.int(id), C.ImNodesPinShape(shape))
}

func BeginNode(id int32) {
	C.imnodes_BeginNode(C.int(id))
}

func BeginNodeEditor() {
	C.imnodes_BeginNodeEditor()
}

func BeginNodeTitleBar() {
	C.imnodes_BeginNodeTitleBar()
}

// BeginOutputAttributeV parameter default value hint:
// shape: ImNodesPinShape_CircleFilled
func BeginOutputAttributeV(id int32, shape PinShape) {
	C.imnodes_BeginOutputAttribute(C.int(id), C.ImNodesPinShape(shape))
}

func BeginStaticAttribute(id int32) {
	C.imnodes_BeginStaticAttribute(C.int(id))
}

func ClearLinkSelectionInt(link_id int32) {
	C.imnodes_ClearLinkSelection_Int(C.int(link_id))
}

func ClearLinkSelection() {
	C.imnodes_ClearLinkSelection_Nil()
}

func ClearNodeSelectionInt(node_id int32) {
	C.imnodes_ClearNodeSelection_Int(C.int(node_id))
}

func ClearNodeSelection() {
	C.imnodes_ClearNodeSelection_Nil()
}

func CreateContext() *Context {
	return NewContextFromC(C.imnodes_CreateContext())
}

// DestroyContextV parameter default value hint:
// ctx: NULL
func DestroyContextV(ctx *Context) {
	ctxArg, ctxFin := ctx.Handle()
	C.imnodes_DestroyContext(internal.ReinterpretCast[*C.ImNodesContext](ctxArg))

	ctxFin()
}

func EditorContextCreate() *EditorContext {
	return NewEditorContextFromC(C.imnodes_EditorContextCreate())
}

func EditorContextFree(noname1 *EditorContext) {
	noname1Arg, noname1Fin := noname1.Handle()
	C.imnodes_EditorContextFree(internal.ReinterpretCast[*C.ImNodesEditorContext](noname1Arg))

	noname1Fin()
}

func EditorContextGetPanning() imgui.Vec2 {
	pOut := new(imgui.Vec2)
	pOutArg, pOutFin := internal.Wrap(pOut)

	C.imnodes_EditorContextGetPanning(internal.ReinterpretCast[*C.ImVec2](pOutArg))

	pOutFin()

	return *pOut
}

func EditorContextMoveToNode(node_id int32) {
	C.imnodes_EditorContextMoveToNode(C.int(node_id))
}

func EditorContextResetPanning(pos imgui.Vec2) {
	C.imnodes_EditorContextResetPanning(internal.ReinterpretCast[C.ImVec2](pos.ToC()))
}

func EditorContextSet(noname1 *EditorContext) {
	noname1Arg, noname1Fin := noname1.Handle()
	C.imnodes_EditorContextSet(internal.ReinterpretCast[*C.ImNodesEditorContext](noname1Arg))

	noname1Fin()
}

func EndInputAttribute() {
	C.imnodes_EndInputAttribute()
}

func EndNode() {
	C.imnodes_EndNode()
}

func EndNodeEditor() {
	C.imnodes_EndNodeEditor()
}

func EndNodeTitleBar() {
	C.imnodes_EndNodeTitleBar()
}

func EndOutputAttribute() {
	C.imnodes_EndOutputAttribute()
}

func EndStaticAttribute() {
	C.imnodes_EndStaticAttribute()
}

func GetCurrentContext() *Context {
	return NewContextFromC(C.imnodes_GetCurrentContext())
}

func GetIO() *IO {
	return NewIOFromC(C.imnodes_GetIO())
}

func GetNodeDimensions(id int32) imgui.Vec2 {
	pOut := new(imgui.Vec2)
	pOutArg, pOutFin := internal.Wrap(pOut)

	C.imnodes_GetNodeDimensions(internal.ReinterpretCast[*C.ImVec2](pOutArg), C.int(id))

	pOutFin()

	return *pOut
}

func GetNodeEditorSpacePos(node_id int32) imgui.Vec2 {
	pOut := new(imgui.Vec2)
	pOutArg, pOutFin := internal.Wrap(pOut)

	C.imnodes_GetNodeEditorSpacePos(internal.ReinterpretCast[*C.ImVec2](pOutArg), C.int(node_id))

	pOutFin()

	return *pOut
}

func GetNodeGridSpacePos(node_id int32) imgui.Vec2 {
	pOut := new(imgui.Vec2)
	pOutArg, pOutFin := internal.Wrap(pOut)

	C.imnodes_GetNodeGridSpacePos(internal.ReinterpretCast[*C.ImVec2](pOutArg), C.int(node_id))

	pOutFin()

	return *pOut
}

func GetNodeScreenSpacePos(node_id int32) imgui.Vec2 {
	pOut := new(imgui.Vec2)
	pOutArg, pOutFin := internal.Wrap(pOut)

	C.imnodes_GetNodeScreenSpacePos(internal.ReinterpretCast[*C.ImVec2](pOutArg), C.int(node_id))

	pOutFin()

	return *pOut
}

func GetSelectedLinks(link_ids *int32) {
	link_idsArg, link_idsFin := internal.WrapNumberPtr[C.int, int32](link_ids)
	C.imnodes_GetSelectedLinks(link_idsArg)

	link_idsFin()
}

func GetSelectedNodes(node_ids *int32) {
	node_idsArg, node_idsFin := internal.WrapNumberPtr[C.int, int32](node_ids)
	C.imnodes_GetSelectedNodes(node_idsArg)

	node_idsFin()
}

func GetStyle() *Style {
	return NewStyleFromC(C.imnodes_GetStyle())
}

// IsAnyAttributeActiveV parameter default value hint:
// attribute_id: NULL
func IsAnyAttributeActiveV(attribute_id *int32) bool {
	attribute_idArg, attribute_idFin := internal.WrapNumberPtr[C.int, int32](attribute_id)

	defer func() {
		attribute_idFin()
	}()
	return C.imnodes_IsAnyAttributeActive(attribute_idArg) == C.bool(true)
}

func IsAttributeActive() bool {
	return C.imnodes_IsAttributeActive() == C.bool(true)
}

func IsEditorHovered() bool {
	return C.imnodes_IsEditorHovered() == C.bool(true)
}

// IsLinkCreatedBoolPtrV parameter default value hint:
// created_from_snap: NULL
func IsLinkCreatedBoolPtrV(started_at_attribute_id *int32, ended_at_attribute_id *int32, created_from_snap *bool) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](ended_at_attribute_id)
	created_from_snapArg, created_from_snapFin := internal.WrapNumberPtr[C.bool, bool](created_from_snap)

	defer func() {
		started_at_attribute_idFin()
		ended_at_attribute_idFin()
		created_from_snapFin()
	}()
	return C.imnodes_IsLinkCreated_BoolPtr(started_at_attribute_idArg, ended_at_attribute_idArg, created_from_snapArg) == C.bool(true)
}

// IsLinkCreatedIntPtrV parameter default value hint:
// created_from_snap: NULL
func IsLinkCreatedIntPtrV(started_at_node_id *int32, started_at_attribute_id *int32, ended_at_node_id *int32, ended_at_attribute_id *int32, created_from_snap *bool) bool {
	started_at_node_idArg, started_at_node_idFin := internal.WrapNumberPtr[C.int, int32](started_at_node_id)
	started_at_attribute_idArg, started_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_node_idArg, ended_at_node_idFin := internal.WrapNumberPtr[C.int, int32](ended_at_node_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](ended_at_attribute_id)
	created_from_snapArg, created_from_snapFin := internal.WrapNumberPtr[C.bool, bool](created_from_snap)

	defer func() {
		started_at_node_idFin()
		started_at_attribute_idFin()
		ended_at_node_idFin()
		ended_at_attribute_idFin()
		created_from_snapFin()
	}()
	return C.imnodes_IsLinkCreated_IntPtr(started_at_node_idArg, started_at_attribute_idArg, ended_at_node_idArg, ended_at_attribute_idArg, created_from_snapArg) == C.bool(true)
}

func IsLinkDestroyed(link_id *int32) bool {
	link_idArg, link_idFin := internal.WrapNumberPtr[C.int, int32](link_id)

	defer func() {
		link_idFin()
	}()
	return C.imnodes_IsLinkDestroyed(link_idArg) == C.bool(true)
}

// IsLinkDroppedV parameter default value hint:
// started_at_attribute_id: NULL
// including_detached_links: true
func IsLinkDroppedV(started_at_attribute_id *int32, including_detached_links bool) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](started_at_attribute_id)

	defer func() {
		started_at_attribute_idFin()
	}()
	return C.imnodes_IsLinkDropped(started_at_attribute_idArg, C.bool(including_detached_links)) == C.bool(true)
}

func IsLinkHovered(link_id *int32) bool {
	link_idArg, link_idFin := internal.WrapNumberPtr[C.int, int32](link_id)

	defer func() {
		link_idFin()
	}()
	return C.imnodes_IsLinkHovered(link_idArg) == C.bool(true)
}

func IsLinkSelected(link_id int32) bool {
	return C.imnodes_IsLinkSelected(C.int(link_id)) == C.bool(true)
}

func IsLinkStarted(started_at_attribute_id *int32) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](started_at_attribute_id)

	defer func() {
		started_at_attribute_idFin()
	}()
	return C.imnodes_IsLinkStarted(started_at_attribute_idArg) == C.bool(true)
}

func IsNodeHovered(node_id *int32) bool {
	node_idArg, node_idFin := internal.WrapNumberPtr[C.int, int32](node_id)

	defer func() {
		node_idFin()
	}()
	return C.imnodes_IsNodeHovered(node_idArg) == C.bool(true)
}

func IsNodeSelected(node_id int32) bool {
	return C.imnodes_IsNodeSelected(C.int(node_id)) == C.bool(true)
}

func IsPinHovered(attribute_id *int32) bool {
	attribute_idArg, attribute_idFin := internal.WrapNumberPtr[C.int, int32](attribute_id)

	defer func() {
		attribute_idFin()
	}()
	return C.imnodes_IsPinHovered(attribute_idArg) == C.bool(true)
}

func Link(id int32, start_attribute_id int32, end_attribute_id int32) {
	C.imnodes_Link(C.int(id), C.int(start_attribute_id), C.int(end_attribute_id))
}

func LoadCurrentEditorStateFromIniFile(file_name string) {
	file_nameArg, file_nameFin := internal.WrapString[C.char](file_name)
	C.imnodes_LoadCurrentEditorStateFromIniFile(file_nameArg)

	file_nameFin()
}

func LoadCurrentEditorStateFromIniString(data string, data_size uint64) {
	dataArg, dataFin := internal.WrapString[C.char](data)
	C.imnodes_LoadCurrentEditorStateFromIniString(dataArg, C.xulong(data_size))

	dataFin()
}

func LoadEditorStateFromIniFile(editor *EditorContext, file_name string) {
	editorArg, editorFin := editor.Handle()
	file_nameArg, file_nameFin := internal.WrapString[C.char](file_name)
	C.imnodes_LoadEditorStateFromIniFile(internal.ReinterpretCast[*C.ImNodesEditorContext](editorArg), file_nameArg)

	editorFin()
	file_nameFin()
}

func LoadEditorStateFromIniString(editor *EditorContext, data string, data_size uint64) {
	editorArg, editorFin := editor.Handle()
	dataArg, dataFin := internal.WrapString[C.char](data)
	C.imnodes_LoadEditorStateFromIniString(internal.ReinterpretCast[*C.ImNodesEditorContext](editorArg), dataArg, C.xulong(data_size))

	editorFin()
	dataFin()
}

// MiniMapV parameter default value hint:
// minimap_size_fraction: 0.2f
// location: ImNodesMiniMapLocation_TopLeft
// node_hovering_callback: NULL
// node_hovering_callback_data: NULL
func MiniMapV(minimap_size_fraction float32, location MiniMapLocation, node_hovering_callback MiniMapNodeHoveringCallback, node_hovering_callback_data MiniMapNodeHoveringCallbackUserData) {
	node_hovering_callbackArg, node_hovering_callbackFin := node_hovering_callback.C()
	node_hovering_callback_dataArg, node_hovering_callback_dataFin := node_hovering_callback_data.C()
	C.imnodes_MiniMap(C.float(minimap_size_fraction), C.ImNodesMiniMapLocation(location), internal.ReinterpretCast[C.ImNodesMiniMapNodeHoveringCallback](node_hovering_callbackArg), internal.ReinterpretCast[C.ImNodesMiniMapNodeHoveringCallbackUserData](node_hovering_callback_dataArg))

	node_hovering_callbackFin()
	node_hovering_callback_dataFin()
}

func NumSelectedLinks() int32 {
	return int32(C.imnodes_NumSelectedLinks())
}

func NumSelectedNodes() int32 {
	return int32(C.imnodes_NumSelectedNodes())
}

func PopAttributeFlag() {
	C.imnodes_PopAttributeFlag()
}

func PopColorStyle() {
	C.imnodes_PopColorStyle()
}

// PopStyleVarV parameter default value hint:
// count: 1
func PopStyleVarV(count int32) {
	C.imnodes_PopStyleVar(C.int(count))
}

func PushAttributeFlag(flag AttributeFlags) {
	C.imnodes_PushAttributeFlag(C.ImNodesAttributeFlags(flag))
}

func PushColorStyle(item Col, color uint32) {
	C.imnodes_PushColorStyle(C.ImNodesCol(item), C.uint(color))
}

func PushStyleVarFloat(style_item StyleVar, value float32) {
	C.imnodes_PushStyleVar_Float(C.ImNodesStyleVar(style_item), C.float(value))
}

func PushStyleVarVec2(style_item StyleVar, value imgui.Vec2) {
	C.imnodes_PushStyleVar_Vec2(C.ImNodesStyleVar(style_item), internal.ReinterpretCast[C.ImVec2](value.ToC()))
}

func SaveCurrentEditorStateToIniFile(file_name string) {
	file_nameArg, file_nameFin := internal.WrapString[C.char](file_name)
	C.imnodes_SaveCurrentEditorStateToIniFile(file_nameArg)

	file_nameFin()
}

// SaveCurrentEditorStateToIniStringV parameter default value hint:
// data_size: NULL
func SaveCurrentEditorStateToIniStringV(data_size *uint64) string {
	return C.GoString(C.imnodes_SaveCurrentEditorStateToIniString((*C.xulong)(data_size)))
}

func SaveEditorStateToIniFile(editor *EditorContext, file_name string) {
	editorArg, editorFin := editor.Handle()
	file_nameArg, file_nameFin := internal.WrapString[C.char](file_name)
	C.imnodes_SaveEditorStateToIniFile(internal.ReinterpretCast[*C.ImNodesEditorContext](editorArg), file_nameArg)

	editorFin()
	file_nameFin()
}

// SaveEditorStateToIniStringV parameter default value hint:
// data_size: NULL
func SaveEditorStateToIniStringV(editor *EditorContext, data_size *uint64) string {
	editorArg, editorFin := editor.Handle()

	defer func() {
		editorFin()
	}()
	return C.GoString(C.imnodes_SaveEditorStateToIniString(internal.ReinterpretCast[*C.ImNodesEditorContext](editorArg), (*C.xulong)(data_size)))
}

func SelectLink(link_id int32) {
	C.imnodes_SelectLink(C.int(link_id))
}

func SelectNode(node_id int32) {
	C.imnodes_SelectNode(C.int(node_id))
}

func SetCurrentContext(ctx *Context) {
	ctxArg, ctxFin := ctx.Handle()
	C.imnodes_SetCurrentContext(internal.ReinterpretCast[*C.ImNodesContext](ctxArg))

	ctxFin()
}

func SetImGuiContext(ctx *imgui.Context) {
	ctxArg, ctxFin := ctx.Handle()
	C.imnodes_SetImGuiContext(internal.ReinterpretCast[*C.ImGuiContext](ctxArg))

	ctxFin()
}

func SetNodeDraggable(node_id int32, draggable bool) {
	C.imnodes_SetNodeDraggable(C.int(node_id), C.bool(draggable))
}

func SetNodeEditorSpacePos(node_id int32, editor_space_pos imgui.Vec2) {
	C.imnodes_SetNodeEditorSpacePos(C.int(node_id), internal.ReinterpretCast[C.ImVec2](editor_space_pos.ToC()))
}

func SetNodeGridSpacePos(node_id int32, grid_pos imgui.Vec2) {
	C.imnodes_SetNodeGridSpacePos(C.int(node_id), internal.ReinterpretCast[C.ImVec2](grid_pos.ToC()))
}

func SetNodeScreenSpacePos(node_id int32, screen_space_pos imgui.Vec2) {
	C.imnodes_SetNodeScreenSpacePos(C.int(node_id), internal.ReinterpretCast[C.ImVec2](screen_space_pos.ToC()))
}

func SnapNodeToGrid(node_id int32) {
	C.imnodes_SnapNodeToGrid(C.int(node_id))
}

// StyleColorsClassicV parameter default value hint:
// dest: NULL
func StyleColorsClassicV(dest *Style) {
	destArg, destFin := dest.Handle()
	C.imnodes_StyleColorsClassic(internal.ReinterpretCast[*C.ImNodesStyle](destArg))

	destFin()
}

// StyleColorsDarkV parameter default value hint:
// dest: NULL
func StyleColorsDarkV(dest *Style) {
	destArg, destFin := dest.Handle()
	C.imnodes_StyleColorsDark(internal.ReinterpretCast[*C.ImNodesStyle](destArg))

	destFin()
}

// StyleColorsLightV parameter default value hint:
// dest: NULL
func StyleColorsLightV(dest *Style) {
	destArg, destFin := dest.Handle()
	C.imnodes_StyleColorsLight(internal.ReinterpretCast[*C.ImNodesStyle](destArg))

	destFin()
}

func BeginInputAttribute(id int32) {
	C.wrap_imnodes_BeginInputAttribute(C.int(id))
}

func BeginOutputAttribute(id int32) {
	C.wrap_imnodes_BeginOutputAttribute(C.int(id))
}

func DestroyContext() {
	C.wrap_imnodes_DestroyContext()
}

func IsAnyAttributeActive() bool {
	return C.wrap_imnodes_IsAnyAttributeActive() == C.bool(true)
}

func IsLinkCreatedBoolPtr(started_at_attribute_id *int32, ended_at_attribute_id *int32) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](ended_at_attribute_id)

	defer func() {
		started_at_attribute_idFin()
		ended_at_attribute_idFin()
	}()
	return C.wrap_imnodes_IsLinkCreated_BoolPtr(started_at_attribute_idArg, ended_at_attribute_idArg) == C.bool(true)
}

func IsLinkCreatedIntPtr(started_at_node_id *int32, started_at_attribute_id *int32, ended_at_node_id *int32, ended_at_attribute_id *int32) bool {
	started_at_node_idArg, started_at_node_idFin := internal.WrapNumberPtr[C.int, int32](started_at_node_id)
	started_at_attribute_idArg, started_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_node_idArg, ended_at_node_idFin := internal.WrapNumberPtr[C.int, int32](ended_at_node_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := internal.WrapNumberPtr[C.int, int32](ended_at_attribute_id)

	defer func() {
		started_at_node_idFin()
		started_at_attribute_idFin()
		ended_at_node_idFin()
		ended_at_attribute_idFin()
	}()
	return C.wrap_imnodes_IsLinkCreated_IntPtr(started_at_node_idArg, started_at_attribute_idArg, ended_at_node_idArg, ended_at_attribute_idArg) == C.bool(true)
}

func IsLinkDropped() bool {
	return C.wrap_imnodes_IsLinkDropped() == C.bool(true)
}

func MiniMap() {
	C.wrap_imnodes_MiniMap()
}

func PopStyleVar() {
	C.wrap_imnodes_PopStyleVar()
}

func SaveCurrentEditorStateToIniString() string {
	return C.GoString(C.wrap_imnodes_SaveCurrentEditorStateToIniString())
}

func SaveEditorStateToIniString(editor *EditorContext) string {
	editorArg, editorFin := editor.Handle()

	defer func() {
		editorFin()
	}()
	return C.GoString(C.wrap_imnodes_SaveEditorStateToIniString(internal.ReinterpretCast[*C.ImNodesEditorContext](editorArg)))
}

func StyleColorsClassic() {
	C.wrap_imnodes_StyleColorsClassic()
}

func StyleColorsDark() {
	C.wrap_imnodes_StyleColorsDark()
}

func StyleColorsLight() {
	C.wrap_imnodes_StyleColorsLight()
}

func (self EmulateThreeButtonMouse) SetModifier(v *bool) {
	vArg, _ := internal.WrapNumberPtr[C.bool, bool](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_EmulateThreeButtonMouse_SetModifier(selfArg, vArg)
}

func (self *EmulateThreeButtonMouse) Modifier() *bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return (*bool)(C.wrap_EmulateThreeButtonMouse_GetModifier(internal.ReinterpretCast[*C.EmulateThreeButtonMouse](selfArg)))
}

func (self IO) SetEmulateThreeButtonMouse(v EmulateThreeButtonMouse) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetEmulateThreeButtonMouse(selfArg, internal.ReinterpretCast[C.EmulateThreeButtonMouse](vArg))
}

func (self *IO) EmulateThreeButtonMouse() EmulateThreeButtonMouse {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_ImNodesIO_GetEmulateThreeButtonMouse(internal.ReinterpretCast[*C.ImNodesIO](selfArg))
	return *NewEmulateThreeButtonMouseFromC(func() *C.EmulateThreeButtonMouse { result := result; return &result }())
}

func (self IO) SetLinkDetachWithModifierClick(v LinkDetachWithModifierClick) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetLinkDetachWithModifierClick(selfArg, internal.ReinterpretCast[C.LinkDetachWithModifierClick](vArg))
}

func (self *IO) LinkDetachWithModifierClick() LinkDetachWithModifierClick {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_ImNodesIO_GetLinkDetachWithModifierClick(internal.ReinterpretCast[*C.ImNodesIO](selfArg))
	return *NewLinkDetachWithModifierClickFromC(func() *C.LinkDetachWithModifierClick { result := result; return &result }())
}

func (self IO) SetMultipleSelectModifier(v MultipleSelectModifier) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetMultipleSelectModifier(selfArg, internal.ReinterpretCast[C.MultipleSelectModifier](vArg))
}

func (self *IO) MultipleSelectModifier() MultipleSelectModifier {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_ImNodesIO_GetMultipleSelectModifier(internal.ReinterpretCast[*C.ImNodesIO](selfArg))
	return *NewMultipleSelectModifierFromC(func() *C.MultipleSelectModifier { result := result; return &result }())
}

func (self IO) SetAltMouseButton(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetAltMouseButton(selfArg, C.int(v))
}

func (self *IO) AltMouseButton() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_ImNodesIO_GetAltMouseButton(internal.ReinterpretCast[*C.ImNodesIO](selfArg)))
}

func (self IO) SetAutoPanningSpeed(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetAutoPanningSpeed(selfArg, C.float(v))
}

func (self *IO) AutoPanningSpeed() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesIO_GetAutoPanningSpeed(internal.ReinterpretCast[*C.ImNodesIO](selfArg)))
}

func (self Style) SetGridSpacing(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetGridSpacing(selfArg, C.float(v))
}

func (self *Style) GridSpacing() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetGridSpacing(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetNodeCornerRounding(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetNodeCornerRounding(selfArg, C.float(v))
}

func (self *Style) NodeCornerRounding() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetNodeCornerRounding(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetNodePadding(v imgui.Vec2) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetNodePadding(selfArg, internal.ReinterpretCast[C.ImVec2](v.ToC()))
}

func (self *Style) NodePadding() imgui.Vec2 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec2 {
		out := C.wrap_ImNodesStyle_GetNodePadding(internal.ReinterpretCast[*C.ImNodesStyle](selfArg))
		return *(&imgui.Vec2{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self Style) SetNodeBorderThickness(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetNodeBorderThickness(selfArg, C.float(v))
}

func (self *Style) NodeBorderThickness() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetNodeBorderThickness(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetLinkThickness(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetLinkThickness(selfArg, C.float(v))
}

func (self *Style) LinkThickness() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetLinkThickness(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetLinkLineSegmentsPerLength(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetLinkLineSegmentsPerLength(selfArg, C.float(v))
}

func (self *Style) LinkLineSegmentsPerLength() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetLinkLineSegmentsPerLength(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetLinkHoverDistance(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetLinkHoverDistance(selfArg, C.float(v))
}

func (self *Style) LinkHoverDistance() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetLinkHoverDistance(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetPinCircleRadius(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinCircleRadius(selfArg, C.float(v))
}

func (self *Style) PinCircleRadius() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinCircleRadius(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetPinQuadSideLength(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinQuadSideLength(selfArg, C.float(v))
}

func (self *Style) PinQuadSideLength() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinQuadSideLength(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetPinTriangleSideLength(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinTriangleSideLength(selfArg, C.float(v))
}

func (self *Style) PinTriangleSideLength() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinTriangleSideLength(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetPinLineThickness(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinLineThickness(selfArg, C.float(v))
}

func (self *Style) PinLineThickness() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinLineThickness(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetPinHoverRadius(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinHoverRadius(selfArg, C.float(v))
}

func (self *Style) PinHoverRadius() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinHoverRadius(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetPinOffset(v float32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinOffset(selfArg, C.float(v))
}

func (self *Style) PinOffset() float32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinOffset(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetMiniMapPadding(v imgui.Vec2) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetMiniMapPadding(selfArg, internal.ReinterpretCast[C.ImVec2](v.ToC()))
}

func (self *Style) MiniMapPadding() imgui.Vec2 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec2 {
		out := C.wrap_ImNodesStyle_GetMiniMapPadding(internal.ReinterpretCast[*C.ImNodesStyle](selfArg))
		return *(&imgui.Vec2{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self Style) SetMiniMapOffset(v imgui.Vec2) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetMiniMapOffset(selfArg, internal.ReinterpretCast[C.ImVec2](v.ToC()))
}

func (self *Style) MiniMapOffset() imgui.Vec2 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec2 {
		out := C.wrap_ImNodesStyle_GetMiniMapOffset(internal.ReinterpretCast[*C.ImNodesStyle](selfArg))
		return *(&imgui.Vec2{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self Style) SetFlags(v StyleFlags) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetFlags(selfArg, C.ImNodesStyleFlags(v))
}

func (self *Style) Flags() StyleFlags {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return StyleFlags(C.wrap_ImNodesStyle_GetFlags(internal.ReinterpretCast[*C.ImNodesStyle](selfArg)))
}

func (self Style) SetColors(v *[29]uint32) {
	vArg := make([]C.uint, len(v))
	for i, vV := range v {
		vArg[i] = C.uint(vV)
	}

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetColors(selfArg, (*C.uint)(&vArg[0]))

	for i, vV := range vArg {
		(*v)[i] = uint32(vV)
	}
}

func (self *Style) Colors() [29]uint32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() [29]uint32 {
		result := [29]uint32{}
		resultMirr := C.wrap_ImNodesStyle_GetColors(internal.ReinterpretCast[*C.ImNodesStyle](selfArg))
		for i := range result {
			result[i] = uint32(C.cimnodes_unsigned_int_GetAtIdx(resultMirr, C.int(i)))
		}

		return result
	}()
}

func (self LinkDetachWithModifierClick) SetModifier(v *bool) {
	vArg, _ := internal.WrapNumberPtr[C.bool, bool](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_LinkDetachWithModifierClick_SetModifier(selfArg, vArg)
}

func (self *LinkDetachWithModifierClick) Modifier() *bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return (*bool)(C.wrap_LinkDetachWithModifierClick_GetModifier(internal.ReinterpretCast[*C.LinkDetachWithModifierClick](selfArg)))
}

func (self MultipleSelectModifier) SetModifier(v *bool) {
	vArg, _ := internal.WrapNumberPtr[C.bool, bool](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MultipleSelectModifier_SetModifier(selfArg, vArg)
}

func (self *MultipleSelectModifier) Modifier() *bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return (*bool)(C.wrap_MultipleSelectModifier_GetModifier(internal.ReinterpretCast[*C.MultipleSelectModifier](selfArg)))
}
