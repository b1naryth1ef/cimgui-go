package main

import "strings"

// Preset is a set of rules iperative to XXX_templates/ json files.
// They are used to manually set some properties of the generator.
type Preset struct {
	// SkipFuncs functions (from definitions.json) to be skipped
	// e.g. they are temporarily hard-coded
	SkipFuncs map[CIdentifier]bool
	// SkipStructs allows to specify struct names that will be skipped.
	SkipStructs map[CIdentifier]bool
	// SkipMethods struct names from structs_and_enums.json.
	// structures that's METHODS should be skipped
	SkipMethods map[CIdentifier]bool
	// SkipTypedefs typedefs from typedefs_dict.json to be skipped
	// e.g. for hardcoded typedefs or typedefs which are obvious (e.g. ImU16 becomes uint16 without extra type information)
	SkipTypedefs map[CIdentifier]bool
	// TypedefsPoolSize sets a default size for callbacks pool.
	// Rembmber to set this as it defaults to 0 and you'll get no callbacks!
	TypedefsPoolSize int
	// TypedefsCustomPoolSizes allows to override TypedefsPoolSize for certain types.
	TypedefsCustomPoolSizes map[CIdentifier]int
	// Replace is a map for C -> Go names conversion.
	// It allows you to force-rename anything (including functions and enums)
	Replace map[CIdentifier]GoIdentifier
	// TrimPrefix allows to remove unwanted prefixes from everything during C->Go renaming.
	// NOTE: order sensitive!
	// NOTE: Case sensitive
	TrimPrefix []string
	// OriginReplace allows to force-replace function name with some other name.
	// Introduced to replace TextEditor_GetText -> TextEditor_GetText_alloc
	// but could be re-used to force use of another function than json tells us to use.
	//
	// It differs from Replace - Replace renames an identifier in general (changes its name but refers to the same function).
	// This allows to absolutely abandon the source C function and use some OTHER C function.
	OriginReplace map[CIdentifier]CIdentifier
	// DefaultArgReplace is used in C-side default args generation (gencpp.go).
	// cimgui-go uses this to change FLT_MIN with igGet_FLT_MIN()
	// NOTE: Iterated randomly!
	DefaultArgReplace map[CIdentifier]CIdentifier
	// DefaultArgArbitraryValue is similar to the above DefaultArgReplace, but
	// associates default arg name with any arbitrary value.
	// cimgui-go uses this to set text_end to 0
	DefaultArgArbitraryValue map[CIdentifier]CIdentifier
	// ExtraCGOPreamble allows to specify additional C code elements in Go files.
	// For example could be used to include extra files.
	// For ease of use, its in form of []string. These lines will be merged and prefixed (if appliable) with '//'
	ExtraCGOPreamble []string
	// InternalFiles allows to specify files that are considered Internal.
	// If an identifier is found in such a file, it will be generated but its
	// name will be prefixed with InternalPrefix
	InternalFiles []string
	// InternalPrefix is a prefix for identifiers from InternalFiles.
	InternalPrefix string
	// PackagePath is an import path of the package.
	// This is base path. flags.packageName will be added.
	// Example:
	//   "github.com/AllenDang/cimgui-go"
	//   If enerated with -pkg imgui, import path
	//   is supposed to be "github.com/AllenDang/cimgui-go/imgui"
	PackagePath string
	// SimpleTypes are used for simple (go-convertable) custom types.
	// Example:
	//   ImS16 is defined as short in C code, so Go can easily convert it via int16()
	// Expected format is:
	//   "ImS16": ["int16", "C.ImS16", "pkgname"]
	// where:
	// - ImS16 is a C type name
	// - int16 is a Go-friendly type name
	// - C.ImS16 is a cgo compatible type name
	// - pkgname is a source package for the type (in this case int16 is a builtin so it should be empty)
	// See also: simpleW
	SimpleTypes map[CIdentifier][3]GoIdentifier
	// SimplePtrTypes are just like SimpleTypes but for pointer types.
	// Example:
	//   "ImS16": ["int16", "C.ImS16", "pkgname"]
	SimplePtrTypes map[CIdentifier][3]GoIdentifier
}

func (p *Preset) MergeCGoPreamble() string {
	result := ""
	for _, line := range p.ExtraCGOPreamble {
		result += "// " + line + "\n"
	}

	result = strings.TrimSuffix(result, "\n")

	return result
}
