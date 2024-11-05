package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kpango/glg"
)

const (
	cimguiGoPackagePath = "github.com/AllenDang/cimgui-go"
	generatorInfo       = "// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.\n// DO NOT EDIT.\n\n"
	goPackageHeader     = generatorInfo + "package %[1]s\n\nimport (\n\"" + cimguiGoPackagePath + "/datautils\"\n\t\"" + cimguiGoPackagePath + "/%[2]s\"\n)\n\n"
	cppFileHeader       = generatorInfo
)

// this cextracts enums and structs names from json file.
func getEnumAndStructNames(enumJsonBytes []byte) (enumNames []GoIdentifier, structNames []CIdentifier, err error) {
	enums, err := getEnumDefs(enumJsonBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get enum definitions: %w", err)
	}

	structs, err := getStructDefs(enumJsonBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get struct definitions: %w", err)
	}

	for _, e := range enums {
		enumNames = append(enumNames, e.Name.renameEnum())
	}

	for _, s := range structs {
		if !shouldSkipStruct(s.Name) {
			structNames = append(structNames, s.Name)
		}
	}

	return
}

func validateFiles(f *flags) {
	stat, err := os.Stat(f.defJsonPath)
	if err != nil || stat.IsDir() {
		glg.Fatal("Invalid definitions json file path")
	}

	stat, err = os.Stat(f.enumsJsonpath)
	if err != nil || stat.IsDir() {
		glg.Fatal("Invalid enum json file path")
	}

	stat, err = os.Stat(f.typedefsJsonpath)
	if err != nil || stat.IsDir() {
		glg.Fatalf("Invalid typedefs json file path: %s", f.typedefsJsonpath)
	}
}

// this stores file bytes of our json files
type jsonData struct {
	structAndEnums,
	typedefs,
	defs,

	refStructAndEnums,
	refTypedefs []byte
}

func loadData(f *flags) (*jsonData, error) {
	var err error

	result := &jsonData{}

	result.defs, err = os.ReadFile(f.defJsonPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read definitions json file: %w", err)
	}

	result.typedefs, err = os.ReadFile(f.typedefsJsonpath)
	if err != nil {
		return nil, fmt.Errorf("cannot read typedefs json file: %w", err)
	}

	result.structAndEnums, err = os.ReadFile(f.enumsJsonpath)
	if err != nil {
		log.Panic(err)
	}

	if len(f.refEnumsJsonPath) > 0 {
		result.refStructAndEnums, err = os.ReadFile(f.refEnumsJsonPath)
		if err != nil {
			return nil, fmt.Errorf("cannot read reference struct and enums json file: %w", err)
		}
	}

	if len(f.refTypedefsJsonPath) > 0 {
		result.refTypedefs, err = os.ReadFile(f.refTypedefsJsonPath)
		if err != nil {
			return nil, fmt.Errorf("cannot read reference typedefs json file: %w", err)
		}
	}

	return result, nil
}

// this will store json data processed by appropiate pre-rocessors
type Context struct {
	funcs    []FuncDef
	structs  []StructDef
	enums    []EnumDef
	typedefs *Typedefs

	prefix string

	funcNames map[CIdentifier]bool
	enumNames map[GoIdentifier]bool

	structNames map[CIdentifier]bool

	typedefsNames map[CIdentifier]bool

	arrayIndexGetters map[CIdentifier]CIdentifier

	refStructNames map[CIdentifier]bool
	refEnumNames   map[GoIdentifier]bool
	refTypedefs    map[CIdentifier]bool

	// TODO: might want to remove this
	flags *flags
}

func parseJson(jsonData *jsonData) (*Context, error) {
	var err error

	result := &Context{
		arrayIndexGetters: make(map[CIdentifier]CIdentifier),
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

	_, structs, err := getEnumAndStructNames(jsonData.structAndEnums)
	result.structNames = SliceToMap(structs)
	if err != nil {
		return nil, fmt.Errorf("cannot get reference struct and enums names: %w", err)
	}

	result.refEnumNames = make(map[GoIdentifier]bool)
	result.refStructNames = make(map[CIdentifier]bool)
	if len(jsonData.refStructAndEnums) > 0 {
		refEnums, refStructs, err := getEnumAndStructNames(jsonData.refStructAndEnums)
		if err != nil {
			return nil, fmt.Errorf("cannot get reference struct and enums names: %w", err)
		}

		result.refEnumNames = SliceToMap(refEnums)
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

	context, err := parseJson(jsonData)
	if err != nil {
		glg.Fatalf("cannot parse json: %v", err)
	}

	context.prefix = flags.prefix
	context.flags = flags

	// 1. Generate code
	// 1.1. Generate Go Enums
	enumNames := generateGoEnums(flags.prefix, context.enums, context)
	context.enumNames = MergeMaps(SliceToMap(enumNames), context.refEnumNames)

	// 1.2. Generate Go typedefs
	callbacks, err := GenerateTypedefs(context.typedefs, context.structs, context)
	if err != nil {
		log.Panic(err)
	}

	context.structNames = SliceToMap(callbacks)

	// 1.3. Generate C wrapper
	validFuncs, err := generateCppWrapper(flags.prefix, flags.include, context.funcs)
	if err != nil {
		log.Panic(err)
	}

	// 1.3.1. Generate Struct accessors in C
	structAccessorFuncs, err := generateCppStructsAccessor(flags.prefix, validFuncs, context.structs, context)
	if err != nil {
		log.Panic(err)
	}

	// This variable stores funcs that needs to be written to GO now.
	validFuncs = append(validFuncs, structAccessorFuncs...)

	if err := GenerateGoFuncs(validFuncs, context); err != nil {
		log.Panic(err)
	}
}

func getGoPackageHeader(ctx *Context) string {
	return fmt.Sprintf(goPackageHeader, ctx.flags.packageName, ctx.flags.refPackageName)
}

func prefixGoPackage(t, sourcePackage GoIdentifier, ctx *Context) GoIdentifier {
	if sourcePackage == GoIdentifier(ctx.flags.packageName) || sourcePackage == "" {
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
