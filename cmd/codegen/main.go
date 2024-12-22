package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kpango/glg"
)

const (
	generatorInfo = "// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.\n// DO NOT EDIT.\n\n"
	cppFileHeader = generatorInfo
)

// this cextracts enums and structs names from json file.
func getEnumAndStructNames(enumJsonBytes []byte, context *Context) (enumNames []EnumIdentifier, typedefsNames []CIdentifier, err error) {
	enums, err := getEnumDefs(enumJsonBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get enum definitions: %w", err)
	}

	structs, err := getStructDefs(enumJsonBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get struct definitions: %w", err)
	}

	for _, e := range enums {
		enumNames = append(enumNames, e.Name)
	}

	for _, s := range structs {
		if shouldSkipStruct := context.preset.SkipStructs[s.Name]; !shouldSkipStruct {
			typedefsNames = append(typedefsNames, s.Name)
		}
	}

	return
}

func validateFiles(f *flags) {
	stat, err := os.Stat(f.DefJsonPath)
	if err != nil || stat.IsDir() {
		glg.Fatal("Invalid definitions json file path")
	}

	stat, err = os.Stat(f.EnumsJsonpath)
	if err != nil || stat.IsDir() {
		glg.Fatal("Invalid enum json file path")
	}

	stat, err = os.Stat(f.TypedefsJsonpath)
	if err != nil || stat.IsDir() {
		glg.Fatalf("Invalid typedefs json file path: %s", f.TypedefsJsonpath)
	}
}

// this stores file bytes of our json files
type jsonData struct {
	structAndEnums,
	typedefs,
	defs,

	refStructAndEnums,
	refTypedefs,

	preset []byte
}

func loadData(f *flags) (*jsonData, error) {
	var err error

	result := &jsonData{}

	result.defs, err = os.ReadFile(f.DefJsonPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read definitions json file: %w", err)
	}

	result.typedefs, err = os.ReadFile(f.TypedefsJsonpath)
	if err != nil {
		return nil, fmt.Errorf("cannot read typedefs json file: %w", err)
	}

	result.structAndEnums, err = os.ReadFile(f.EnumsJsonpath)
	if err != nil {
		glg.Fatalf("cannot read struct and enums json file: %v", err)
	}

	if len(f.RefEnumsJsonPath) > 0 {
		result.refStructAndEnums, err = os.ReadFile(f.RefEnumsJsonPath)
		if err != nil {
			return nil, fmt.Errorf("cannot read reference struct and enums json file: %w", err)
		}
	}

	if len(f.RefTypedefsJsonPath) > 0 {
		result.refTypedefs, err = os.ReadFile(f.RefTypedefsJsonPath)
		if err != nil {
			return nil, fmt.Errorf("cannot read reference typedefs json file: %w", err)
		}
	}

	if len(f.PresetJsonPath) > 0 {
		result.preset, err = os.ReadFile(f.PresetJsonPath)
		if err != nil {
			glg.Fatalf("cannot read preset json file: %v", err)
		}

	}

	return result, nil
}

// this will store json data processed by appropiate pre-rocessors
type Context struct {
	// plain idata loaded from json
	funcs    []FuncDef
	structs  []StructDef
	enums    []EnumDef
	typedefs *Typedefs

	// ghese fields are filled by parser while it generates code.
	enumNames     map[CIdentifier]bool
	typedefsNames map[CIdentifier]bool

	// contains helper C functions to get/set struct fields
	// of array types
	arrayIndexGetters map[CIdentifier]CIdentifier

	// contains identifiers from other package (usually imgui).
	refStructNames map[CIdentifier]bool
	refEnumNames   map[CIdentifier]bool
	refTypedefs    map[CIdentifier]bool

	// TODO: might want to remove this
	flags *flags

	preset *Preset
}

func parseJson(jsonData *jsonData, f *flags) (*Context, error) {
	var err error

	result := &Context{
		arrayIndexGetters: make(map[CIdentifier]CIdentifier),
		preset:            &Preset{},
		flags:             f,
	}

	if len(jsonData.preset) > 0 {
		if err := json.Unmarshal(jsonData.preset, result.preset); err != nil {
			glg.Warnf("Unable to load preset: %v", err)
		}

		if result.flags.Verbose {
			glg.Debugf("Preset loaded: %v", result.preset)
		}
	}

	// get definitions from json file
	result.funcs, err = getFunDefs(jsonData.defs)
	if err != nil {
		return nil, fmt.Errorf("cannot get function definitions: %w", err)
	}

	result.enums, err = getEnumDefs(jsonData.structAndEnums)
	if err != nil {
		return nil, fmt.Errorf("cannot get enum definitions: %w", err)
	}

	result.typedefs, err = getTypedefs(jsonData.typedefs)
	if err != nil {
		return nil, fmt.Errorf("cannot get typedefs: %w", err)
	}

	result.structs, err = getStructDefs(jsonData.structAndEnums)
	if err != nil {
		return nil, fmt.Errorf("cannot get struct definitions: %w", err)
	}

	result.refTypedefs = make(map[CIdentifier]bool)
	if len(jsonData.refTypedefs) > 0 {
		typedefs, err := getTypedefs(jsonData.refTypedefs)
		if err != nil {
			return nil, fmt.Errorf("cannot get reference typedefs: %w", err)
		}

		result.refTypedefs = RemoveMapValues(typedefs.data)
	}

	_, structs, err := getEnumAndStructNames(jsonData.structAndEnums, result)
	result.typedefsNames = SliceToMap(structs)
	if err != nil {
		return nil, fmt.Errorf("cannot get reference struct and enums names: %w", err)
	}

	result.refEnumNames = make(map[CIdentifier]bool)
	result.refStructNames = make(map[CIdentifier]bool)
	if len(jsonData.refStructAndEnums) > 0 {
		refEnums, refStructs, err := getEnumAndStructNames(jsonData.refStructAndEnums, result)
		if err != nil {
			return nil, fmt.Errorf("cannot get reference struct and enums names: %w", err)
		}

		result.refEnumNames = make(map[CIdentifier]bool)
		for _, refEnum := range refEnums {
			result.refEnumNames[refEnum.renameEnum()] = true
		}

		result.refStructNames = SliceToMap(refStructs)
	}

	return result, nil
}

func main() {
	flags := parse()
	validateFiles(flags)

	jsonData, err := loadData(flags)
	if err != nil {
		glg.Fatalf("cannot load data: %v", err)
	}

	context, err := parseJson(jsonData, flags)
	if err != nil {
		glg.Fatalf("cannot parse json: %v", err)
	}

	// 1. Generate code
	// 1.1. Generate Go Enums
	glg.Info("Generating Go Enums")
	enumNames, err := generateGoEnums(context.enums, context)
	if err != nil {
		glg.Fatalf("Generating enum names: %v", err)
	}

	context.enumNames = SliceToMap(enumNames)

	// 1.2. Generate Go typedefs
	glg.Info("Generating Go Typedefs")
	typedefsNames, callbacksToGenerate, err := GenerateTypedefs(context.typedefs, context.structs, context)
	if err != nil {
		glg.Fatalf("Cannot generate typedefs: %v", err)
	}

	context.typedefsNames = SliceToMap(typedefsNames)

	validCallbacks, err := GenerateCallbacks(callbacksToGenerate, context)
	if err != nil {
		glg.Fatalf("Cannot generate callbacks: %v", err)
	}

	context.typedefsNames = MergeMaps(context.typedefsNames, SliceToMap(validCallbacks))

	// 1.3. Generate C wrapper
	glg.Info("Generating C Wrapper")
	validFuncs, err := generateCppWrapper(flags.Include, context.funcs, context)
	if err != nil {
		glg.Fatalf("Cannot generate CPP Wrapper: %v", err)
	}

	// 1.3.1. Generate Struct accessors in C
	glg.Info("Generating C Struct Accessors")
	structAccessorFuncs, err := generateCppStructsAccessor(validFuncs, context.structs, context)
	if err != nil {
		glg.Fatalf("Cannot generate CPP Struct Accessors: %v", err)
	}

	// This variable stores funcs that needs to be written to GO now.
	validFuncs = append(validFuncs, structAccessorFuncs...)

	// 1.4. Generate Go functions
	glg.Info("Generating Go Functions")
	if err := GenerateGoFuncs(validFuncs, context); err != nil {
		glg.Fatalf("Cannot generate Go functions: %v", err)
	}
}

func getGoPackageHeader(ctx *Context) string {
	return fmt.Sprintf(
		`%[1]s
package %[2]s

import (
	"%[4]s/utils"
	"%[4]s/%[3]s"
)
`,
		generatorInfo,
		ctx.flags.PackageName,
		ctx.flags.RefPackageName,
		ctx.preset.PackagePath,
	)
}

func prefixGoPackage(t, sourcePackage GoIdentifier, ctx *Context) GoIdentifier {
	if sourcePackage == GoIdentifier(ctx.flags.PackageName) || sourcePackage == "" {
		return t
	}

	isPtr := HasPrefix(t, "*")
	t = TrimPrefix(t, "*")
	t = sourcePackage + "." + t

	if isPtr {
		t = "*" + t
	}

	return t
}
