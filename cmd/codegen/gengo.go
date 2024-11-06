package main

import (
	"regexp"
	"strings"
)

type (
	// EnumIdentifier is in theory EnumName_
	EnumIdentifier string
	// CIdentifier is a string representing name of struct/func/enum in C
	CIdentifier string
	// GoIdentifier is a string representing name of struct/func/enum in Go
	// ATTENTION: GoIdentifier indicates that the name represents a valid Go name.
	// Should be created only by renameGoIdentifier
	GoIdentifier string
)

// structures that's methods should be skipped
var skippedStructs = map[CIdentifier]bool{
	"ImVec1":     true,
	"ImVec2":     true,
	"ImVec2ih":   true,
	"ImVec4":     true,
	"ImColor":    true,
	"ImRect":     true,
	"ImPlotTime": true,
}

var skippedTypedefs = map[CIdentifier]bool{
	"ImU8":                   true,
	"ImU16":                  true,
	"ImU32":                  true,
	"ImU64":                  true,
	"ImS8":                   true,
	"ImS16":                  true,
	"ImS32":                  true,
	"ImS64":                  true,
	"ImGuiInputTextCallback": true,
}

const TypedefsPoolSize = 32

var customPoolSize = map[CIdentifier]int{
	// nothing here for now
}

var replace = map[CIdentifier]GoIdentifier{
	"igGetDrawData":           "CurrentDrawData",
	"igGetDrawListSharedData": "CurrentDrawListSharedData",
	"igGetFont":               "CurrentFont",
	"igGetIO":                 "CurrentIO",
	"igGetPlatformIO":         "CurrentPlatformIO",
	"igGetStyle":              "CurrentStyle",
	"igGetMouseCursor":        "CurrentMouseCursor",
	"ImAxis":                  "AxisEnum",
	"GetItem_ID":              "ItemByID",
}

func (c CIdentifier) trimImGuiPrefix() CIdentifier {
	// order sensitive!
	prefixes := []string{
		"ImGuizmo",
		"imnodes",
		"ImNodes",
		"ImPlot",
		"ImGui",
		"Im",
		"ig",
	}

	for _, prefix := range prefixes {
		if HasPrefix(c, prefix) &&
			len(c) > len(prefix) && // check if removing it will not result in an empty string
			strings.ToUpper(string(c[len(prefix)])) == string(c[len(prefix)]) { // check if the method will still be exported
			c = TrimPrefix(c, prefix)
		}
	}

	return c
}

func (c CIdentifier) renameGoIdentifier() GoIdentifier {
	if r, ok := replace[c]; ok {
		c = CIdentifier(r)
	}

	c = TrimSuffix(c, "_Nil")

	c = c.trimImGuiPrefix()
	switch {
	case HasPrefix(c, "New"):
		c = "New" + c[3:].trimImGuiPrefix()
	case HasPrefix(c, "new"):
		c = "new" + c[3:].trimImGuiPrefix()
	case HasPrefix(c, "*"):
		c = "*" + c[1:].trimImGuiPrefix()
	}

	c = TrimPrefix(c, "Get")
	if c != "_" {
		c = ReplaceAll(c, "_", "")
	}

	return GoIdentifier(c)
}

func (c EnumIdentifier) renameEnum() CIdentifier {
	return CIdentifier(TrimSuffix(c, "_"))
}

// returns true if s is of form TypeName<*> <(>Name<*><)>(args)
// (fragments in <> are optional)
func IsCallbackTypedef(s string) bool {
	pattern := `\w*\**\(*\w*\**\)*\(.*\);`
	b, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}

	return b
}

func IsStructName(name CIdentifier, ctx *Context) bool {
	_, ok := ctx.structNames[name]
	return ok
}

func IsEnumName(name CIdentifier, enums map[CIdentifier]bool) bool {
	_, ok := enums[name]
	return ok
}

func IsTemplateTypedef(s string) bool {
	return strings.Contains(s, "<")
}
