package main

import (
	"fmt"
	"os"
	"strings"
)

// Generate enums and return enum type names
func generateGoEnums(prefix string, enums []EnumDef, ctx *Context) ([]CIdentifier, error) {
	var sb strings.Builder

	sb.WriteString(getGoPackageHeader(ctx))

	var enumNames []CIdentifier
	for _, e := range enums {
		originalName := e.Name
		eName := e.Name.renameEnum().renameGoIdentifier(ctx)

		enumNames = append(enumNames, e.Name.renameEnum())

		sb.WriteString(fmt.Sprintf("%s\n", e.CommentAbove))
		sb.WriteString(fmt.Sprintf("// original name: %s\n", originalName))
		sb.WriteString(fmt.Sprintf("type %s int32\n", eName))
		sb.WriteString("const (\n")

		for _, v := range e.Values {
			vName := TrimSuffix(v.Name, "_")
			if v.Comment != "" {
				sb.WriteString(fmt.Sprintf("%s\n", v.Comment))
			}
			sb.WriteString(fmt.Sprintf("\t%s %s = %d\n", vName.renameGoIdentifier(ctx), eName, v.Value))
		}

		sb.WriteString(")\n\n")
	}

	enumFile, err := os.Create(fmt.Sprintf("%s_enums.go", prefix))
	if err != nil {
		return nil, err
	}

	defer enumFile.Close()

	_, _ = enumFile.WriteString(sb.String())

	return enumNames, nil
}
