package main

import (
	"go/ast"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestProcessDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "structbuilder-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test Go file
	testFile := filepath.Join(tempDir, "test.go")
	testContent := `
package test

type TestStruct struct {
	Field1 string ` + "`structbuilder:\"required\"`" + `
	Field2 int    ` + "`structbuilder:\"optional\"`" + `
}
`
	err = os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test processDirectory function
	packages, err := processDirectory(tempDir)
	if err != nil {
		t.Fatalf("processDirectory failed: %v", err)
	}

	// Verify results
	if len(packages) != 1 {
		t.Errorf("Expected package count: 1, got: %d", len(packages))
	}

	pkg, ok := packages[tempDir]
	if !ok {
		t.Fatalf("Expected package not found")
	}

	if pkg.Name != "test" {
		t.Errorf("Expected package name: test, got: %s", pkg.Name)
	}

	if !pkg.HasStructs {
		t.Errorf("HasStructs should be true")
	}

	if !strings.Contains(pkg.GeneratedCode.String(), "TestStructBuilder") {
		t.Errorf("Generated code does not contain TestStructBuilder")
	}
}

// func TestGenerateBuilder(t *testing.T) {
// 	// Create a test struct
// 	structType := &ast.StructType{
// 		Fields: &ast.FieldList{
// 			List: []*ast.Field{
// 				{
// 					Names: []*ast.Ident{{Name: "Field1"}},
// 					Type:  &ast.Ident{Name: "string"},
// 					Tag:   &ast.BasicLit{Value: "`structbuilder:\"required\"`"},
// 				},
// 				{
// 					Names: []*ast.Ident{{Name: "Field2"}},
// 					Type:  &ast.Ident{Name: "int"},
// 					Tag:   &ast.BasicLit{Value: "`structbuilder:\"optional\"`"},
// 				},
// 			},
// 		},
// 	}

// 	// Test generateBuilder function
// 	result := generateBuilder("TestStruct", structType)

// 	// Verify results
// 	expectedStrings := []string{
// 		"type TestStructBuilder struct",
// 		"type TestStructField1Builder interface",
// 		"type TestStructOptionalBuilder interface",
// 		"func (b *testStructBuilder) SetField1(Field1 string) TestStructOptionalBuilder",
// 		"func (b *testStructBuilder) SetOptField2(Field2 int) TestStructOptionalBuilder",
// 		"func (b *testStructBuilder) Build() *TestStruct",
// 		"func NewTestStructBuilder() TestStructField1Builder",
// 		"func NewTestStruct(Field1 string, Field2 int) *TestStruct",
// 	}

// 	for _, expected := range expectedStrings {
// 		if !strings.Contains(result, expected) {
// 			t.Errorf("Expected string not found: %s", expected)
// 		}
// 	}
// }

func TestHasStructBuilderTag(t *testing.T) {
	tests := []struct {
		name     string
		tagValue string
		expected bool
	}{
		{"With structbuilder tag", "`structbuilder:\"required\"`", true},
		{"Without structbuilder tag", "`json:\"field\"`", false},
		{"No tag", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structType := &ast.StructType{
				Fields: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{{Name: "Field"}},
							Type:  &ast.Ident{Name: "string"},
						},
					},
				},
			}

			if tt.tagValue != "" {
				structType.Fields.List[0].Tag = &ast.BasicLit{Value: tt.tagValue}
			}

			result := hasStructBuilderTag(structType)
			if result != tt.expected {
				t.Errorf("Expected result: %v, got: %v", tt.expected, result)
			}
		})
	}
}

// func TestRun(t *testing.T) {
// 	// Create a temporary directory for testing
// 	tempDir, err := os.MkdirTemp("", "structbuilder-run-test")
// 	if err != nil {
// 		t.Fatalf("Failed to create temp directory: %v", err)
// 	}
// 	defer os.RemoveAll(tempDir)

// 	// Create a test Go file
// 	testFile := filepath.Join(tempDir, "test.go")
// 	testContent := `
// package test

// type TestStruct struct {
// 	Field1 string ` + "`structbuilder:\"required\"`" + `
// 	Field2 int    ` + "`structbuilder:\"optional\"`" + `
// }
// `
// 	err = os.WriteFile(testFile, []byte(testContent), 0644)
// 	if err != nil {
// 		t.Fatalf("Failed to create test file: %v", err)
// 	}

// 	// Set the rootDir to the temp directory
// 	originalRootDir := rootDir
// 	rootDir = tempDir
// 	defer func() { rootDir = originalRootDir }()

// 	// Run the main function
// 	err = run()
// 	if err != nil {
// 		t.Fatalf("run() failed: %v", err)
// 	}

// 	// Check if the generated file exists
// 	generatedFile := filepath.Join(tempDir, generatedFile)
// 	if _, err := os.Stat(generatedFile); os.IsNotExist(err) {
// 		t.Errorf("Generated file does not exist: %s", generatedFile)
// 	}

// 	// Read the content of the generated file
// 	content, err := os.ReadFile(generatedFile)
// 	if err != nil {
// 		t.Fatalf("Failed to read generated file: %v", err)
// 	}

// 	// Check for expected content in the generated file
// 	expectedStrings := []string{
// 		"type TestStructBuilder struct",
// 		"func NewTestStructBuilder() TestStructField1Builder",
// 		"func (b *testStructBuilder) Build() *TestStruct",
// 	}

// 	for _, expected := range expectedStrings {
// 		if !strings.Contains(string(content), expected) {
// 			t.Errorf("Expected string not found in generated file: %s", expected)
// 		}
// 	}
// }

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"world", "World"},
		{"", ""},
		{"A", "A"},
		{"already", "Already"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := capitalize(tt.input)
			if result != tt.expected {
				t.Errorf("capitalize(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLcFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "hello"},
		{"World", "world"},
		{"", ""},
		{"a", "a"},
		{"ALLCAPS", "aLLCAPS"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := lcFirst(tt.input)
			if result != tt.expected {
				t.Errorf("lcFirst(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTypeToString(t *testing.T) {
	tests := []struct {
		name     string
		expr     ast.Expr
		expected string
	}{
		{
			name:     "Identifier",
			expr:     &ast.Ident{Name: "string"},
			expected: "string",
		},
		{
			name:     "Pointer",
			expr:     &ast.StarExpr{X: &ast.Ident{Name: "int"}},
			expected: "*int",
		},
		{
			name:     "Array",
			expr:     &ast.ArrayType{Elt: &ast.Ident{Name: "byte"}},
			expected: "[]byte",
		},
		{
			name: "Map",
			expr: &ast.MapType{
				Key:   &ast.Ident{Name: "string"},
				Value: &ast.Ident{Name: "int"},
			},
			expected: "map[string]int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := typeToString(tt.expr)
			if result != tt.expected {
				t.Errorf("typeToString() = %q, want %q", result, tt.expected)
			}
		})
	}
}
