package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
)

const (
	rootDir       = "."
	generatedFile = "structbuilder_gen.go"
)

type PackageInfo struct {
	Name          string
	GeneratedCode strings.Builder
	HasStructs    bool
	Imports       []*ast.ImportSpec
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	packages, err := processDirectory(rootDir)
	if err != nil {
		return fmt.Errorf("error processing directory: %w", err)
	}
	return generateFiles(packages)
}

func processDirectory(dir string) (map[string]*PackageInfo, error) {
	packages := make(map[string]*PackageInfo)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if isGoFile(path, info) {
			if err := processFile(path, packages); err != nil {
				return fmt.Errorf("error processing file %s: %w", path, err)
			}
		}
		return nil
	})
	return packages, err
}

func isGoFile(path string, info os.FileInfo) bool {
	return !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, generatedFile)
}

func processFile(filePath string, packages map[string]*PackageInfo) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}

	dir := filepath.Dir(filePath)
	packageInfo, ok := packages[dir]
	if !ok {
		packageInfo = &PackageInfo{Name: node.Name.Name}
		packages[dir] = packageInfo
	}

	packageInfo.Imports = node.Imports

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			if hasStructBuilderTag(structType) {
				fmt.Printf("Found struct to build: %s in %s\n", typeSpec.Name.Name, filePath)
				builderCode := generateBuilder(typeSpec.Name.Name, structType)
				packageInfo.GeneratedCode.WriteString(builderCode)
				packageInfo.HasStructs = true
			}
		}
	}

	return nil
}

func generateFiles(packages map[string]*PackageInfo) error {
	for dir, pkg := range packages {
		if pkg.HasStructs {
			if err := writeGeneratedFile(dir, pkg); err != nil {
				return fmt.Errorf("error generating file for %s: %w", dir, err)
			}
		}
	}
	return nil
}

func writeGeneratedFile(dir string, pkg *PackageInfo) error {
	var buf bytes.Buffer

	buf.WriteString("// Code generated by struct-builder; DO NOT EDIT.\n\n")

	buf.WriteString(fmt.Sprintf("package %s\n\n", pkg.Name))

	fset := token.NewFileSet()
	file := &ast.File{
		Name:  ast.NewIdent(pkg.Name),
		Decls: []ast.Decl{},
	}

	if pkg.Imports != nil {
		for _, imp := range pkg.Imports {
			if imp.Name != nil && imp.Path != nil {
				astutil.AddNamedImport(fset, file, imp.Name.Name, strings.Trim(imp.Path.Value, "\""))
			}
		}

		if len(file.Imports) > 0 {
			importDecl := &ast.GenDecl{
				Tok:   token.IMPORT,
				Specs: make([]ast.Spec, len(file.Imports)),
			}
			for i, imp := range file.Imports {
				importDecl.Specs[i] = imp
			}
			printer.Fprint(&buf, fset, importDecl)
			buf.WriteString("\n\n")
		}
	}

	buf.WriteString(pkg.GeneratedCode.String())

	outputPath := filepath.Join(dir, generatedFile)

	formattedCode, err := imports.Process(outputPath, buf.Bytes(), nil)
	if err != nil {
		return fmt.Errorf("error formatting code: %w", err)
	}

	if err := os.WriteFile(outputPath, formattedCode, 0644); err != nil {
		return fmt.Errorf("error writing generated code: %w", err)
	}
	fmt.Printf("Builder code generated: %s\n", outputPath)
	return nil
}

func generateBuilder(structName string, structType *ast.StructType) string {
	requiredFields, optionalFields := getFields(structType)
	return strings.Join([]string{
		generateStructDefinition(structName, structType),
		generateInterfaces(structName, requiredFields, optionalFields, structType),
		generateSetterMethods(structName, requiredFields, optionalFields, structType),
		generateGetterMethods(structName, requiredFields, optionalFields, structType),
		generateBuildMethod(structName, requiredFields, optionalFields),
		generateNewBuilderFunction(structName, requiredFields),
		generateNewStructFunction(structName, requiredFields, optionalFields, structType),
	}, "\n")
}

func generateStructDefinition(structName string, structType *ast.StructType) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %sBuilder struct {\n", lcFirst(structName)))
	for _, field := range structType.Fields.List {
		if len(field.Names) > 0 {
			fieldName := field.Names[0].Name
			fieldType := getFieldType(structType, fieldName)
			builder.WriteString(fmt.Sprintf("\t%s %s\n", fieldName, fieldType))
		}
	}
	builder.WriteString("}\n")
	return builder.String()
}

func generateInterfaces(structName string, requiredFields, optionalFields []string, structType *ast.StructType) string {
	var builder strings.Builder
	builder.WriteString(generateRequiredInterfaces(structName, requiredFields, structType))
	builder.WriteString(generateOptionalInterface(structName, optionalFields, structType))
	return builder.String()
}

func generateRequiredInterfaces(structName string, requiredFields []string, structType *ast.StructType) string {
	var builder strings.Builder
	for i, field := range requiredFields {
		interfaceName := fmt.Sprintf("%s%sBuilder", structName, capitalize(field))
		nextInterface := getNextInterface(structName, requiredFields, i)
		builder.WriteString(fmt.Sprintf("type %s interface {\n", interfaceName))
		builder.WriteString(fmt.Sprintf("\tSet%s(%s %s) %s\n", capitalize(field), field, getFieldType(structType, field), nextInterface))
		builder.WriteString("}\n\n")
	}
	return builder.String()
}

func getNextInterface(structName string, requiredFields []string, currentIndex int) string {
	if currentIndex < len(requiredFields)-1 {
		return fmt.Sprintf("%s%sBuilder", structName, capitalize(requiredFields[currentIndex+1]))
	}
	return fmt.Sprintf("%sOptionalBuilder", structName)
}

func generateOptionalInterface(structName string, optionalFields []string, structType *ast.StructType) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %sOptionalBuilder interface {\n", structName))
	for _, field := range optionalFields {
		builder.WriteString(fmt.Sprintf("\tSetOpt%s(%s %s) %sOptionalBuilder\n", capitalize(field), field, getFieldType(structType, field), structName))
	}
	builder.WriteString(fmt.Sprintf("\tBuild() *%s\n", structName))
	builder.WriteString("}\n")
	return builder.String()
}

func generateSetterMethods(structName string, requiredFields, optionalFields []string, structType *ast.StructType) string {
	var builder strings.Builder
	builder.WriteString(generateRequiredSetterMethods(structName, requiredFields, structType))
	builder.WriteString(generateOptionalSetterMethods(structName, optionalFields, structType))
	return builder.String()
}

func generateRequiredSetterMethods(structName string, requiredFields []string, structType *ast.StructType) string {
	var builder strings.Builder
	for i, field := range requiredFields {
		fieldType := getFieldType(structType, field)
		nextInterface := getNextInterface(structName, requiredFields, i)
		builder.WriteString(fmt.Sprintf("func (b *%sBuilder) Set%s(%s %s) %s {\n", lcFirst(structName), capitalize(field), field, fieldType, nextInterface))
		builder.WriteString(fmt.Sprintf("\tb.%s = %s\n", field, field))
		builder.WriteString("\treturn b\n}\n\n")
	}
	return builder.String()
}

func generateOptionalSetterMethods(structName string, optionalFields []string, structType *ast.StructType) string {
	var builder strings.Builder
	for _, field := range optionalFields {
		fieldType := getFieldType(structType, field)
		builder.WriteString(fmt.Sprintf("func (b *%sBuilder) SetOpt%s(%s %s) %sOptionalBuilder {\n", lcFirst(structName), capitalize(field), field, fieldType, structName))
		builder.WriteString(fmt.Sprintf("\tb.%s = %s\n", field, field))
		builder.WriteString("\treturn b\n}\n\n")
	}
	return builder.String()
}

func generateGetterMethods(structName string, requiredFields, optionalFields []string, structType *ast.StructType) string {
	var builder strings.Builder
	allFields := append(requiredFields, optionalFields...)
	for _, field := range structType.Fields.List {
		if len(field.Names) > 0 {
			fieldName := field.Names[0].Name
			if hasGetterTag(field) && contains(allFields, fieldName) {
				fieldType := getFieldType(structType, fieldName)
				if !isPublicField(fieldName) {
					builder.WriteString(fmt.Sprintf("func (s *%s) %s() %s {\n", structName, capitalize(fieldName), fieldType))
					builder.WriteString(fmt.Sprintf("\treturn s.%s\n", fieldName))
					builder.WriteString("}\n\n")
				}
			}
		}
	}
	return builder.String()
}

func generateBuildMethod(structName string, requiredFields, optionalFields []string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("func (b *%sBuilder) Build() *%s {\n", lcFirst(structName), structName))
	builder.WriteString(fmt.Sprintf("\treturn &%s{\n", structName))
	for _, field := range append(requiredFields, optionalFields...) {
		builder.WriteString(fmt.Sprintf("\t\t%s: b.%s,\n", field, field))
	}
	builder.WriteString("\t}\n}\n\n")
	return builder.String()
}

func generateNewBuilderFunction(structName string, requiredFields []string) string {
	var builder strings.Builder
	var firstInterface string
	if len(requiredFields) > 0 {
		firstInterface = fmt.Sprintf("%s%sBuilder", structName, capitalize(requiredFields[0]))
	} else {
		firstInterface = fmt.Sprintf("%sOptionalBuilder", structName)
	}
	builder.WriteString(fmt.Sprintf("func New%sBuilder() %s {\n", structName, firstInterface))
	builder.WriteString(fmt.Sprintf("\treturn &%sBuilder{}\n", lcFirst(structName)))
	builder.WriteString("}\n\n")
	return builder.String()
}

func generateNewStructFunction(structName string, requiredFields, optionalFields []string, structType *ast.StructType) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("func New%s(", structName))

	allFields := append(requiredFields, optionalFields...)
	for i, field := range allFields {
		fieldType := getFieldType(structType, field)
		builder.WriteString(fmt.Sprintf("%s %s", field, fieldType))
		if i < len(allFields)-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteString(fmt.Sprintf(") *%s {\n", structName))
	builder.WriteString(fmt.Sprintf("\treturn &%s{\n", structName))

	for _, field := range allFields {
		builder.WriteString(fmt.Sprintf("\t\t%s: %s,\n", field, field))
	}

	builder.WriteString("\t}\n")
	builder.WriteString("}\n\n")
	return builder.String()
}

func getFields(structType *ast.StructType) (required, optional []string) {
	for _, field := range structType.Fields.List {
		if len(field.Names) > 0 {
			fieldName := field.Names[0].Name
			if hasRequiredTag(field) {
				required = append(required, fieldName)
			} else if hasOptionalTag(field) {
				optional = append(optional, fieldName)
			}
		}
	}
	return
}

func hasRequiredTag(field *ast.Field) bool {
	if field.Tag != nil {
		tagValue := field.Tag.Value
		return strings.Contains(tagValue, "required")
	}
	return false
}

func hasOptionalTag(field *ast.Field) bool {
	if field.Tag != nil {
		tagValue := field.Tag.Value
		return strings.Contains(tagValue, "optional")
	}
	return false
}

func hasGetterTag(field *ast.Field) bool {
	if field.Tag != nil {
		tagValue := field.Tag.Value
		return strings.Contains(tagValue, "getter")
	}
	return false
}

func isPublicField(fieldName string) bool {
	return fieldName != "" && unicode.IsUpper(rune(fieldName[0]))
}

func getFieldType(structType *ast.StructType, fieldName string) string {
	for _, field := range structType.Fields.List {
		if len(field.Names) > 0 && field.Names[0].Name == fieldName {
			return typeToString(field.Type)
		}
	}
	return ""
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func typeToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", typeToString(t.X), t.Sel.Name)
	case *ast.StarExpr:
		return "*" + typeToString(t.X)
	case *ast.ArrayType:
		return "[]" + typeToString(t.Elt)
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", typeToString(t.Key), typeToString(t.Value))
	case *ast.ChanType:
		return formatChanType(t)
	default:
		return fmt.Sprintf("Unsupported type: %T", expr)
	}
}

func formatChanType(t *ast.ChanType) string {
	switch t.Dir {
	case ast.SEND:
		return fmt.Sprintf("chan<- %s", typeToString(t.Value))
	case ast.RECV:
		return fmt.Sprintf("<-chan %s", typeToString(t.Value))
	default:
		return fmt.Sprintf("chan %s", typeToString(t.Value))
	}
}

func hasStructBuilderTag(structType *ast.StructType) bool {
	for _, field := range structType.Fields.List {
		if field.Tag != nil {
			tagValue := field.Tag.Value
			if strings.Contains(tagValue, "structbuilder") {
				return true
			}
		}
	}
	return false
}

func lcFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
