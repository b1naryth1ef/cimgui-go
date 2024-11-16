// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package immarkdown

import (
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/internal"
)

// #include "structs_accessor.h"
// #include "wrapper.h"
// #include "../imgui/extra_types.h"

import "C"

func IsCharInsideWord(c_ rune) bool {
	return C.IsCharInsideWord(C.char(c_)) == C.bool(true)
}

func Markdown(markdown_ string, markdownLength_ uint64, mdConfig_ MarkdownConfig) {
	markdown_Arg, markdown_Fin := internal.WrapString[C.char](markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	C.Markdown(markdown_Arg, C.xulong(markdownLength_), internal.ReinterpretCast[C.MarkdownConfig](mdConfig_Arg))

	markdown_Fin()
	mdConfig_Fin()
}

func RenderLine(markdown_ string, line_ *Line, textRegion_ *TextRegion, mdConfig_ MarkdownConfig) {
	markdown_Arg, markdown_Fin := internal.WrapString[C.char](markdown_)
	line_Arg, line_Fin := line_.Handle()
	textRegion_Arg, textRegion_Fin := textRegion_.Handle()
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	C.RenderLine(markdown_Arg, internal.ReinterpretCast[*C.Line](line_Arg), internal.ReinterpretCast[*C.TextRegion](textRegion_Arg), internal.ReinterpretCast[C.MarkdownConfig](mdConfig_Arg))

	markdown_Fin()
	line_Fin()
	textRegion_Fin()
	mdConfig_Fin()
}

func RenderLinkText(self *TextRegion, text_ string, link_ Link, markdown_ string, mdConfig_ MarkdownConfig, linkHoverStart_ []string) bool {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := internal.WrapString[C.char](text_)
	link_Arg, link_Fin := link_.C()
	markdown_Arg, markdown_Fin := internal.WrapString[C.char](markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	linkHoverStart_Arg, linkHoverStart_Fin := internal.WrapStringList[C.char](linkHoverStart_)

	defer func() {
		selfFin()
		text_Fin()
		link_Fin()
		markdown_Fin()
		mdConfig_Fin()
		linkHoverStart_Fin()
	}()
	return C.wrap_RenderLinkText(internal.ReinterpretCast[*C.TextRegion](selfArg), text_Arg, internal.ReinterpretCast[C.Link](link_Arg), markdown_Arg, internal.ReinterpretCast[C.MarkdownConfig](mdConfig_Arg), linkHoverStart_Arg) == C.bool(true)
}

// RenderLinkTextWrappedV parameter default value hint:
// bIndentToHere_: false
func RenderLinkTextWrappedV(self *TextRegion, text_ string, link_ Link, markdown_ string, mdConfig_ MarkdownConfig, linkHoverStart_ []string, bIndentToHere_ bool) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := internal.WrapString[C.char](text_)
	link_Arg, link_Fin := link_.C()
	markdown_Arg, markdown_Fin := internal.WrapString[C.char](markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	linkHoverStart_Arg, linkHoverStart_Fin := internal.WrapStringList[C.char](linkHoverStart_)
	C.wrap_RenderLinkTextWrappedV(internal.ReinterpretCast[*C.TextRegion](selfArg), text_Arg, internal.ReinterpretCast[C.Link](link_Arg), markdown_Arg, internal.ReinterpretCast[C.MarkdownConfig](mdConfig_Arg), linkHoverStart_Arg, C.bool(bIndentToHere_))

	selfFin()
	text_Fin()
	link_Fin()
	markdown_Fin()
	mdConfig_Fin()
	linkHoverStart_Fin()
}

func RenderListTextWrapped(self *TextRegion, text_ string) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := internal.WrapString[C.char](text_)
	C.wrap_RenderListTextWrapped(internal.ReinterpretCast[*C.TextRegion](selfArg), text_Arg)

	selfFin()
	text_Fin()
}

// RenderTextWrappedV parameter default value hint:
// bIndentToHere_: false
func RenderTextWrappedV(self *TextRegion, text_ string, bIndentToHere_ bool) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := internal.WrapString[C.char](text_)
	C.wrap_RenderTextWrappedV(internal.ReinterpretCast[*C.TextRegion](selfArg), text_Arg, C.bool(bIndentToHere_))

	selfFin()
	text_Fin()
}

func ResetIndent(self *TextRegion) {
	selfArg, selfFin := self.Handle()
	C.ResetIndent(internal.ReinterpretCast[*C.TextRegion](selfArg))

	selfFin()
}

func NewTextRegion() *TextRegion {
	return NewTextRegionFromC(C.TextRegion_TextRegion())
}

func (self *TextRegion) Destroy() {
	selfArg, selfFin := self.Handle()
	C.TextRegion_destroy(internal.ReinterpretCast[*C.TextRegion](selfArg))

	selfFin()
}

func UnderLine(col_ imgui.Color) {
	C.UnderLine(internal.ReinterpretCast[C.ImColor](col_.ToC()))
}

func DefaultMarkdownFormatCallback(markdownFormatInfo_ MarkdownFormatInfo, start_ bool) {
	markdownFormatInfo_Arg, markdownFormatInfo_Fin := markdownFormatInfo_.C()
	C.defaultMarkdownFormatCallback(internal.ReinterpretCast[C.MarkdownFormatInfo](markdownFormatInfo_Arg), C.bool(start_))

	markdownFormatInfo_Fin()
}

func DefaultMarkdownTooltipCallback(data_ MarkdownTooltipCallbackData) {
	data_Arg, data_Fin := data_.C()
	C.defaultMarkdownTooltipCallback(internal.ReinterpretCast[C.MarkdownTooltipCallbackData](data_Arg))

	data_Fin()
}

func Size(self *TextBlock) int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.size(internal.ReinterpretCast[*C.TextBlock](selfArg)))
}

func RenderLinkTextWrapped(self *TextRegion, text_ string, link_ Link, markdown_ string, mdConfig_ MarkdownConfig, linkHoverStart_ []string) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := internal.WrapString[C.char](text_)
	link_Arg, link_Fin := link_.C()
	markdown_Arg, markdown_Fin := internal.WrapString[C.char](markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	linkHoverStart_Arg, linkHoverStart_Fin := internal.WrapStringList[C.char](linkHoverStart_)
	C.wrap_RenderLinkTextWrapped(internal.ReinterpretCast[*C.TextRegion](selfArg), text_Arg, internal.ReinterpretCast[C.Link](link_Arg), markdown_Arg, internal.ReinterpretCast[C.MarkdownConfig](mdConfig_Arg), linkHoverStart_Arg)

	selfFin()
	text_Fin()
	link_Fin()
	markdown_Fin()
	mdConfig_Fin()
	linkHoverStart_Fin()
}

func RenderTextWrapped(self *TextRegion, text_ string) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := internal.WrapString[C.char](text_)
	C.wrap_RenderTextWrapped(internal.ReinterpretCast[*C.TextRegion](selfArg), text_Arg)

	selfFin()
	text_Fin()
}

func (self Emphasis) SetState(v EmphasisState) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Emphasis_SetState(selfArg, C.EmphasisState(v))
}

func (self *Emphasis) State() EmphasisState {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return EmphasisState(C.wrap_Emphasis_GetState(internal.ReinterpretCast[*C.Emphasis](selfArg)))
}

func (self Emphasis) SetText(v TextBlock) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Emphasis_SetText(selfArg, internal.ReinterpretCast[C.TextBlock](vArg))
}

func (self *Emphasis) Text() TextBlock {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_Emphasis_GetText(internal.ReinterpretCast[*C.Emphasis](selfArg))
	return *NewTextBlockFromC(func() *C.TextBlock { result := result; return &result }())
}

func (self Emphasis) SetSym(v rune) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Emphasis_SetSym(selfArg, C.char(v))
}

func (self *Emphasis) Sym() rune {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return rune(C.wrap_Emphasis_GetSym(internal.ReinterpretCast[*C.Emphasis](selfArg)))
}

func (self Line) SetIsHeading(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetIsHeading(selfArg, C.bool(v))
}

func (self *Line) IsHeading() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsHeading(internal.ReinterpretCast[*C.Line](selfArg)) == C.bool(true)
}

func (self Line) SetIsEmphasis(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetIsEmphasis(selfArg, C.bool(v))
}

func (self *Line) IsEmphasis() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsEmphasis(internal.ReinterpretCast[*C.Line](selfArg)) == C.bool(true)
}

func (self Line) SetIsUnorderedListStart(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetIsUnorderedListStart(selfArg, C.bool(v))
}

func (self *Line) IsUnorderedListStart() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsUnorderedListStart(internal.ReinterpretCast[*C.Line](selfArg)) == C.bool(true)
}

func (self Line) SetIsLeadingSpace(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetIsLeadingSpace(selfArg, C.bool(v))
}

func (self *Line) IsLeadingSpace() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsLeadingSpace(internal.ReinterpretCast[*C.Line](selfArg)) == C.bool(true)
}

func (self Line) SetLeadSpaceCount(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetLeadSpaceCount(selfArg, C.int(v))
}

func (self *Line) LeadSpaceCount() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLeadSpaceCount(internal.ReinterpretCast[*C.Line](selfArg)))
}

func (self Line) SetHeadingCount(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetHeadingCount(selfArg, C.int(v))
}

func (self *Line) HeadingCount() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetHeadingCount(internal.ReinterpretCast[*C.Line](selfArg)))
}

func (self Line) SetEmphasisCount(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetEmphasisCount(selfArg, C.int(v))
}

func (self *Line) EmphasisCount() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetEmphasisCount(internal.ReinterpretCast[*C.Line](selfArg)))
}

func (self Line) SetLineStart(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetLineStart(selfArg, C.int(v))
}

func (self *Line) LineStart() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLineStart(internal.ReinterpretCast[*C.Line](selfArg)))
}

func (self Line) SetLineEnd(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetLineEnd(selfArg, C.int(v))
}

func (self *Line) LineEnd() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLineEnd(internal.ReinterpretCast[*C.Line](selfArg)))
}

func (self Line) SetLastRenderPosition(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetLastRenderPosition(selfArg, C.int(v))
}

func (self *Line) LastRenderPosition() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLastRenderPosition(internal.ReinterpretCast[*C.Line](selfArg)))
}

func (self Link) SetState(v LinkState) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetState(selfArg, C.LinkState(v))
}

func (self *Link) State() LinkState {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return LinkState(C.wrap_Link_GetState(internal.ReinterpretCast[*C.Link](selfArg)))
}

func (self Link) SetText(v TextBlock) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetText(selfArg, internal.ReinterpretCast[C.TextBlock](vArg))
}

func (self *Link) Text() TextBlock {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_Link_GetText(internal.ReinterpretCast[*C.Link](selfArg))
	return *NewTextBlockFromC(func() *C.TextBlock { result := result; return &result }())
}

func (self Link) SetUrl(v TextBlock) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetUrl(selfArg, internal.ReinterpretCast[C.TextBlock](vArg))
}

func (self *Link) Url() TextBlock {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_Link_GetUrl(internal.ReinterpretCast[*C.Link](selfArg))
	return *NewTextBlockFromC(func() *C.TextBlock { result := result; return &result }())
}

func (self Link) SetIsImage(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetIsImage(selfArg, C.bool(v))
}

func (self *Link) IsImage() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Link_GetIsImage(internal.ReinterpretCast[*C.Link](selfArg)) == C.bool(true)
}

func (self Link) SetNumbracketsopen(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetNum_brackets_open(selfArg, C.int(v))
}

func (self *Link) Numbracketsopen() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Link_GetNum_brackets_open(internal.ReinterpretCast[*C.Link](selfArg)))
}

func (self MarkdownConfig) SetLinkCallback(v *MarkdownLinkCallback) {
	vArg, _ := v.Handle()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetLinkCallback(selfArg, internal.ReinterpretCast[*C.MarkdownLinkCallback](vArg))
}

func (self *MarkdownConfig) LinkCallback() *MarkdownLinkCallback {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return NewMarkdownLinkCallbackFromC(C.wrap_MarkdownConfig_GetLinkCallback(internal.ReinterpretCast[*C.MarkdownConfig](selfArg)))
}

func (self MarkdownConfig) SetTooltipCallback(v *MarkdownTooltipCallback) {
	vArg, _ := v.Handle()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetTooltipCallback(selfArg, internal.ReinterpretCast[*C.MarkdownTooltipCallback](vArg))
}

func (self *MarkdownConfig) TooltipCallback() *MarkdownTooltipCallback {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return NewMarkdownTooltipCallbackFromC(C.wrap_MarkdownConfig_GetTooltipCallback(internal.ReinterpretCast[*C.MarkdownConfig](selfArg)))
}

func (self MarkdownConfig) SetImageCallback(v *MarkdownImageCallback) {
	vArg, _ := v.Handle()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetImageCallback(selfArg, internal.ReinterpretCast[*C.MarkdownImageCallback](vArg))
}

func (self *MarkdownConfig) ImageCallback() *MarkdownImageCallback {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return NewMarkdownImageCallbackFromC(C.wrap_MarkdownConfig_GetImageCallback(internal.ReinterpretCast[*C.MarkdownConfig](selfArg)))
}

func (self MarkdownConfig) SetLinkIcon(v string) {
	vArg, _ := internal.WrapString[C.char](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetLinkIcon(selfArg, vArg)
}

func (self *MarkdownConfig) LinkIcon() string {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownConfig_GetLinkIcon(internal.ReinterpretCast[*C.MarkdownConfig](selfArg)))
}

func (self MarkdownConfig) SetHeadingFormats(v *[3]MarkdownHeadingFormat) {
	vArg := make([]C.MarkdownHeadingFormat, len(v))
	for i, vV := range v {
		vVArg, _ := vV.C()
		vArg[i] = internal.ReinterpretCast[C.MarkdownHeadingFormat](vVArg)
	}

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetHeadingFormats(selfArg, (*C.MarkdownHeadingFormat)(&vArg[0]))

	for i, vV := range vArg {
		(*v)[i] = *NewMarkdownHeadingFormatFromC(func() *C.MarkdownHeadingFormat { result := vV; return &result }())
	}
}

func (self *MarkdownConfig) HeadingFormats() [3]MarkdownHeadingFormat {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() [3]MarkdownHeadingFormat {
		result := [3]MarkdownHeadingFormat{}
		resultMirr := C.wrap_MarkdownConfig_GetHeadingFormats(internal.ReinterpretCast[*C.MarkdownConfig](selfArg))
		for i := range result {
			result[i] = *NewMarkdownHeadingFormatFromC(func() *C.MarkdownHeadingFormat {
				result := C.cimmarkdown_MarkdownHeadingFormat_GetAtIdx(resultMirr, C.int(i))
				return &result
			}())
		}

		return result
	}()
}

func (self MarkdownConfig) SetUserData(v uintptr) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetUserData(selfArg, C.uintptr_t(v))
}

func (self *MarkdownConfig) UserData() uintptr {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return uintptr(C.wrap_MarkdownConfig_GetUserData(internal.ReinterpretCast[*C.MarkdownConfig](selfArg)))
}

func (self MarkdownConfig) SetFormatCallback(v *MarkdownFormalCallback) {
	vArg, _ := v.Handle()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetFormatCallback(selfArg, internal.ReinterpretCast[*C.MarkdownFormalCallback](vArg))
}

func (self *MarkdownConfig) FormatCallback() *MarkdownFormalCallback {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return NewMarkdownFormalCallbackFromC(C.wrap_MarkdownConfig_GetFormatCallback(internal.ReinterpretCast[*C.MarkdownConfig](selfArg)))
}

func (self MarkdownFormatInfo) SetType(v MarkdownFormatType) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetType(selfArg, C.MarkdownFormatType(v))
}

func (self *MarkdownFormatInfo) Type() MarkdownFormatType {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return MarkdownFormatType(C.wrap_MarkdownFormatInfo_GetType(internal.ReinterpretCast[*C.MarkdownFormatInfo](selfArg)))
}

func (self MarkdownFormatInfo) SetLevel(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetLevel(selfArg, C.int32_t(v))
}

func (self *MarkdownFormatInfo) Level() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_MarkdownFormatInfo_GetLevel(internal.ReinterpretCast[*C.MarkdownFormatInfo](selfArg)))
}

func (self MarkdownFormatInfo) SetItemHovered(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetItemHovered(selfArg, C.bool(v))
}

func (self *MarkdownFormatInfo) ItemHovered() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownFormatInfo_GetItemHovered(internal.ReinterpretCast[*C.MarkdownFormatInfo](selfArg)) == C.bool(true)
}

func (self MarkdownFormatInfo) SetConfig(v *MarkdownConfig) {
	vArg, _ := v.Handle()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetConfig(selfArg, internal.ReinterpretCast[*C.MarkdownConfig](vArg))
}

func (self *MarkdownFormatInfo) Config() *MarkdownConfig {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return NewMarkdownConfigFromC(C.wrap_MarkdownFormatInfo_GetConfig(internal.ReinterpretCast[*C.MarkdownFormatInfo](selfArg)))
}

func (self MarkdownHeadingFormat) SetFont(v *imgui.Font) {
	vArg, _ := v.Handle()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownHeadingFormat_SetFont(selfArg, internal.ReinterpretCast[*C.ImFont](vArg))
}

func (self *MarkdownHeadingFormat) Font() *imgui.Font {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return imgui.NewFontFromC(C.wrap_MarkdownHeadingFormat_GetFont(internal.ReinterpretCast[*C.MarkdownHeadingFormat](selfArg)))
}

func (self MarkdownHeadingFormat) SetSeparator(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownHeadingFormat_SetSeparator(selfArg, C.bool(v))
}

func (self *MarkdownHeadingFormat) Separator() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownHeadingFormat_GetSeparator(internal.ReinterpretCast[*C.MarkdownHeadingFormat](selfArg)) == C.bool(true)
}

func (self MarkdownImageData) SetIsValid(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetIsValid(selfArg, C.bool(v))
}

func (self *MarkdownImageData) IsValid() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownImageData_GetIsValid(internal.ReinterpretCast[*C.MarkdownImageData](selfArg)) == C.bool(true)
}

func (self MarkdownImageData) SetUseLinkCallback(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUseLinkCallback(selfArg, C.bool(v))
}

func (self *MarkdownImageData) UseLinkCallback() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownImageData_GetUseLinkCallback(internal.ReinterpretCast[*C.MarkdownImageData](selfArg)) == C.bool(true)
}

func (self MarkdownImageData) SetUsertextureid(v imgui.TextureID) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUser_texture_id(selfArg, internal.ReinterpretCast[C.ImTextureID](vArg))
}

func (self *MarkdownImageData) Usertextureid() imgui.TextureID {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return *imgui.NewTextureIDFromC(func() *C.ImTextureID {
		result := C.wrap_MarkdownImageData_GetUser_texture_id(internal.ReinterpretCast[*C.MarkdownImageData](selfArg))
		return &result
	}())
}

func (self MarkdownImageData) SetSize(v imgui.Vec2) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetSize(selfArg, internal.ReinterpretCast[C.ImVec2](v.ToC()))
}

func (self *MarkdownImageData) Size() imgui.Vec2 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec2 {
		out := C.wrap_MarkdownImageData_GetSize(internal.ReinterpretCast[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec2{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownImageData) SetUv0(v imgui.Vec2) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUv0(selfArg, internal.ReinterpretCast[C.ImVec2](v.ToC()))
}

func (self *MarkdownImageData) Uv0() imgui.Vec2 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec2 {
		out := C.wrap_MarkdownImageData_GetUv0(internal.ReinterpretCast[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec2{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownImageData) SetUv1(v imgui.Vec2) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUv1(selfArg, internal.ReinterpretCast[C.ImVec2](v.ToC()))
}

func (self *MarkdownImageData) Uv1() imgui.Vec2 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec2 {
		out := C.wrap_MarkdownImageData_GetUv1(internal.ReinterpretCast[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec2{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownImageData) SetTintcol(v imgui.Vec4) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetTint_col(selfArg, internal.ReinterpretCast[C.ImVec4](v.ToC()))
}

func (self *MarkdownImageData) Tintcol() imgui.Vec4 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec4 {
		out := C.wrap_MarkdownImageData_GetTint_col(internal.ReinterpretCast[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec4{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownImageData) SetBordercol(v imgui.Vec4) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetBorder_col(selfArg, internal.ReinterpretCast[C.ImVec4](v.ToC()))
}

func (self *MarkdownImageData) Bordercol() imgui.Vec4 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec4 {
		out := C.wrap_MarkdownImageData_GetBorder_col(internal.ReinterpretCast[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec4{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownLinkCallbackData) SetText(v string) {
	vArg, _ := internal.WrapString[C.char](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetText(selfArg, vArg)
}

func (self *MarkdownLinkCallbackData) Text() string {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownLinkCallbackData_GetText(internal.ReinterpretCast[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetTextLength(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetTextLength(selfArg, C.int(v))
}

func (self *MarkdownLinkCallbackData) TextLength() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_MarkdownLinkCallbackData_GetTextLength(internal.ReinterpretCast[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetLink(v string) {
	vArg, _ := internal.WrapString[C.char](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetLink(selfArg, vArg)
}

func (self *MarkdownLinkCallbackData) Link() string {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownLinkCallbackData_GetLink(internal.ReinterpretCast[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetLinkLength(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetLinkLength(selfArg, C.int(v))
}

func (self *MarkdownLinkCallbackData) LinkLength() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_MarkdownLinkCallbackData_GetLinkLength(internal.ReinterpretCast[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetUserData(v uintptr) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetUserData(selfArg, C.uintptr_t(v))
}

func (self *MarkdownLinkCallbackData) UserData() uintptr {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return uintptr(C.wrap_MarkdownLinkCallbackData_GetUserData(internal.ReinterpretCast[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetIsImage(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetIsImage(selfArg, C.bool(v))
}

func (self *MarkdownLinkCallbackData) IsImage() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownLinkCallbackData_GetIsImage(internal.ReinterpretCast[*C.MarkdownLinkCallbackData](selfArg)) == C.bool(true)
}

func (self MarkdownTooltipCallbackData) SetLinkData(v MarkdownLinkCallbackData) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownTooltipCallbackData_SetLinkData(selfArg, internal.ReinterpretCast[C.MarkdownLinkCallbackData](vArg))
}

func (self *MarkdownTooltipCallbackData) LinkData() MarkdownLinkCallbackData {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_MarkdownTooltipCallbackData_GetLinkData(internal.ReinterpretCast[*C.MarkdownTooltipCallbackData](selfArg))
	return *NewMarkdownLinkCallbackDataFromC(func() *C.MarkdownLinkCallbackData { result := result; return &result }())
}

func (self MarkdownTooltipCallbackData) SetLinkIcon(v string) {
	vArg, _ := internal.WrapString[C.char](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownTooltipCallbackData_SetLinkIcon(selfArg, vArg)
}

func (self *MarkdownTooltipCallbackData) LinkIcon() string {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownTooltipCallbackData_GetLinkIcon(internal.ReinterpretCast[*C.MarkdownTooltipCallbackData](selfArg)))
}

func (self TextBlock) SetStart(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_TextBlock_SetStart(selfArg, C.int(v))
}

func (self *TextBlock) Start() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_TextBlock_GetStart(internal.ReinterpretCast[*C.TextBlock](selfArg)))
}

func (self TextBlock) SetStop(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_TextBlock_SetStop(selfArg, C.int(v))
}

func (self *TextBlock) Stop() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_TextBlock_GetStop(internal.ReinterpretCast[*C.TextBlock](selfArg)))
}
