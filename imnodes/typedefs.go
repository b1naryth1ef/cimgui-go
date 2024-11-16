// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imnodes

// #include <stdlib.h>
// #include <memory.h>
// #include "wrapper.h"
// #include "typedefs.h"
// #include "../imgui/extra_types.h"

import "C"
import "github.com/AllenDang/cimgui-go/internal"

type EmulateThreeButtonMouse struct {
	CData *C.EmulateThreeButtonMouse
}

// Handle returns C version of EmulateThreeButtonMouse and its finalizer func.
func (self *EmulateThreeButtonMouse) Handle() (result *C.EmulateThreeButtonMouse, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self EmulateThreeButtonMouse) C() (C.EmulateThreeButtonMouse, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewEmptyEmulateThreeButtonMouse creates EmulateThreeButtonMouse with its 0 value.
func NewEmptyEmulateThreeButtonMouse() *EmulateThreeButtonMouse {
	return &EmulateThreeButtonMouse{CData: new(C.EmulateThreeButtonMouse)}
}

// NewEmulateThreeButtonMouseFromC creates EmulateThreeButtonMouse from its C pointer.
// SRC ~= *C.EmulateThreeButtonMouse
func NewEmulateThreeButtonMouseFromC[SRC any](cvalue SRC) *EmulateThreeButtonMouse {
	return &EmulateThreeButtonMouse{CData: internal.ReinterpretCast[*C.EmulateThreeButtonMouse](cvalue)}
}

type Context struct {
	CData *C.ImNodesContext
}

// Handle returns C version of Context and its finalizer func.
func (self *Context) Handle() (result *C.ImNodesContext, fin func()) {
	return self.CData, func() {}
}

// NewContextFromC creates Context from its C pointer.
// SRC ~= *C.ImNodesContext
func NewContextFromC[SRC any](cvalue SRC) *Context {
	return &Context{CData: internal.ReinterpretCast[*C.ImNodesContext](cvalue)}
}

type EditorContext struct {
	CData *C.ImNodesEditorContext
}

// Handle returns C version of EditorContext and its finalizer func.
func (self *EditorContext) Handle() (result *C.ImNodesEditorContext, fin func()) {
	return self.CData, func() {}
}

// NewEditorContextFromC creates EditorContext from its C pointer.
// SRC ~= *C.ImNodesEditorContext
func NewEditorContextFromC[SRC any](cvalue SRC) *EditorContext {
	return &EditorContext{CData: internal.ReinterpretCast[*C.ImNodesEditorContext](cvalue)}
}

type IO struct {
	CData *C.ImNodesIO
}

// Handle returns C version of IO and its finalizer func.
func (self *IO) Handle() (result *C.ImNodesIO, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self IO) C() (C.ImNodesIO, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewEmptyIO creates IO with its 0 value.
func NewEmptyIO() *IO {
	return &IO{CData: new(C.ImNodesIO)}
}

// NewIOFromC creates IO from its C pointer.
// SRC ~= *C.ImNodesIO
func NewIOFromC[SRC any](cvalue SRC) *IO {
	return &IO{CData: internal.ReinterpretCast[*C.ImNodesIO](cvalue)}
}

type MiniMapNodeHoveringCallbackUserData struct {
	Data uintptr
}

// Handle returns C version of MiniMapNodeHoveringCallbackUserData and its finalizer func.
func (self *MiniMapNodeHoveringCallbackUserData) Handle() (result *C.ImNodesMiniMapNodeHoveringCallbackUserData, fin func()) {
	r, f := self.C()
	return &r, f
}

// C is like Handle but returns plain type instead of pointer.
func (self MiniMapNodeHoveringCallbackUserData) C() (C.ImNodesMiniMapNodeHoveringCallbackUserData, func()) {
	return (C.ImNodesMiniMapNodeHoveringCallbackUserData)(C.ImNodesMiniMapNodeHoveringCallbackUserData_fromUintptr(C.uintptr_t(self.Data))), func() {}
}

// NewMiniMapNodeHoveringCallbackUserDataFromC creates MiniMapNodeHoveringCallbackUserData from its C pointer.
// SRC ~= *C.ImNodesMiniMapNodeHoveringCallbackUserData
func NewMiniMapNodeHoveringCallbackUserDataFromC[SRC any](cvalue SRC) *MiniMapNodeHoveringCallbackUserData {
	return &MiniMapNodeHoveringCallbackUserData{Data: (uintptr)(C.ImNodesMiniMapNodeHoveringCallbackUserData_toUintptr(*internal.ReinterpretCast[*C.ImNodesMiniMapNodeHoveringCallbackUserData](cvalue)))}
}

type Style struct {
	CData *C.ImNodesStyle
}

// Handle returns C version of Style and its finalizer func.
func (self *Style) Handle() (result *C.ImNodesStyle, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self Style) C() (C.ImNodesStyle, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewEmptyStyle creates Style with its 0 value.
func NewEmptyStyle() *Style {
	return &Style{CData: new(C.ImNodesStyle)}
}

// NewStyleFromC creates Style from its C pointer.
// SRC ~= *C.ImNodesStyle
func NewStyleFromC[SRC any](cvalue SRC) *Style {
	return &Style{CData: internal.ReinterpretCast[*C.ImNodesStyle](cvalue)}
}

type LinkDetachWithModifierClick struct {
	CData *C.LinkDetachWithModifierClick
}

// Handle returns C version of LinkDetachWithModifierClick and its finalizer func.
func (self *LinkDetachWithModifierClick) Handle() (result *C.LinkDetachWithModifierClick, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self LinkDetachWithModifierClick) C() (C.LinkDetachWithModifierClick, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewEmptyLinkDetachWithModifierClick creates LinkDetachWithModifierClick with its 0 value.
func NewEmptyLinkDetachWithModifierClick() *LinkDetachWithModifierClick {
	return &LinkDetachWithModifierClick{CData: new(C.LinkDetachWithModifierClick)}
}

// NewLinkDetachWithModifierClickFromC creates LinkDetachWithModifierClick from its C pointer.
// SRC ~= *C.LinkDetachWithModifierClick
func NewLinkDetachWithModifierClickFromC[SRC any](cvalue SRC) *LinkDetachWithModifierClick {
	return &LinkDetachWithModifierClick{CData: internal.ReinterpretCast[*C.LinkDetachWithModifierClick](cvalue)}
}

type MultipleSelectModifier struct {
	CData *C.MultipleSelectModifier
}

// Handle returns C version of MultipleSelectModifier and its finalizer func.
func (self *MultipleSelectModifier) Handle() (result *C.MultipleSelectModifier, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self MultipleSelectModifier) C() (C.MultipleSelectModifier, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewEmptyMultipleSelectModifier creates MultipleSelectModifier with its 0 value.
func NewEmptyMultipleSelectModifier() *MultipleSelectModifier {
	return &MultipleSelectModifier{CData: new(C.MultipleSelectModifier)}
}

// NewMultipleSelectModifierFromC creates MultipleSelectModifier from its C pointer.
// SRC ~= *C.MultipleSelectModifier
func NewMultipleSelectModifierFromC[SRC any](cvalue SRC) *MultipleSelectModifier {
	return &MultipleSelectModifier{CData: internal.ReinterpretCast[*C.MultipleSelectModifier](cvalue)}
}
