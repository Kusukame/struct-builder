package main

import (
	"go/ast"
	"go/token"
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
func TestGenerateBuilder(t *testing.T) {
	structType := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "Field1"}},
					Type:  &ast.Ident{Name: "string"},
					Tag:   &ast.BasicLit{Value: "`structbuilder:\"required\"`"},
				},
				{
					Names: []*ast.Ident{{Name: "Field2"}},
					Type:  &ast.Ident{Name: "int"},
					Tag:   &ast.BasicLit{Value: "`structbuilder:\"optional\"`"},
				},
			},
		},
	}

	result := generateBuilder("TestStruct", structType)

	// 期待される文字列を生成されたコードに合わせて修正
	expectedStrings := []string{
		"type testStructBuilder struct",
		"func (b *testStructBuilder) SetField1(Field1 string) TestStructOptionalBuilder {",
		"func (b *testStructBuilder) SetOptField2(Field2 int) TestStructOptionalBuilder {",
		"func (b *testStructBuilder) Build() *TestStruct {",
		"func NewTestStructBuilder() TestStructField1Builder {",
		"func NewTestStruct(Field1 string, Field2 int) *TestStruct {",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(result, expected) {
			t.Errorf("Expected string not found: %s", expected)
		}
	}
}
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

func TestFormatChanType(t *testing.T) {
	tests := []struct {
		name     string
		chanType *ast.ChanType
		want     string
	}{
		{
			name: "bidirectional channel",
			chanType: &ast.ChanType{
				Dir:   ast.SEND | ast.RECV,
				Value: &ast.Ident{Name: "int"},
			},
			want: "chan int",
		},
		{
			name: "send-only channel",
			chanType: &ast.ChanType{
				Dir:   ast.SEND,
				Value: &ast.Ident{Name: "string"},
			},
			want: "chan<- string",
		},
		{
			name: "receive-only channel",
			chanType: &ast.ChanType{
				Dir:   ast.RECV,
				Value: &ast.Ident{Name: "bool"},
			},
			want: "<-chan bool",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatChanType(tt.chanType)
			if got != tt.want {
				t.Errorf("formatChanType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		item  string
		want  bool
	}{
		{
			name:  "item exists in slice",
			slice: []string{"apple", "banana", "cherry"},
			item:  "banana",
			want:  true,
		},
		{
			name:  "item does not exist in slice",
			slice: []string{"apple", "banana", "cherry"},
			item:  "grape",
			want:  false,
		},
		{
			name:  "empty slice",
			slice: []string{},
			item:  "apple",
			want:  false,
		},
		{
			name:  "nil slice",
			slice: nil,
			item:  "apple",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contains(tt.slice, tt.item)
			if got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteGeneratedFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-structbuilder")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test PackageInfo struct
	pkg := &PackageInfo{
		Name: "testpackage",
		Imports: []*ast.ImportSpec{
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"fmt\"",
				},
			},
		},
		GeneratedCode: strings.Builder{},
	}
	pkg.GeneratedCode.WriteString("func TestFunc() {}\n")

	// Execute the function
	err = writeGeneratedFile(tempDir, pkg)
	if err != nil {
		t.Fatalf("writeGeneratedFile failed: %v", err)
	}

	// Read the generated file
	generatedFilePath := filepath.Join(tempDir, generatedFile)
	content, err := os.ReadFile(generatedFilePath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Verify the content of the generated code
	expectedContent := `// Code generated by struct-builder; DO NOT EDIT.

package testpackage

func TestFunc() {}
`
	if string(content) != expectedContent {
		t.Errorf("Generated content does not match expected content")
	}

	// Check the file permissions
	fileInfo, err := os.Stat(generatedFilePath)
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}
	if fileInfo.Mode().Perm() != os.FileMode(0644) {
		t.Errorf("File permissions are not 0644")
	}
}

func TestGenerateGetterMethods(t *testing.T) {
	structName := "TestStruct"
	requiredFields := []string{"requiredField"}
	optionalFields := []string{"optionalField"}

	// Create a test struct
	structType := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "requiredField"}},
					Type:  &ast.Ident{Name: "string"},
					Tag:   &ast.BasicLit{Value: "`getter:\"true\"`"},
				},
				{
					Names: []*ast.Ident{{Name: "optionalField"}},
					Type:  &ast.Ident{Name: "int"},
					Tag:   &ast.BasicLit{Value: "`getter:\"true\"`"},
				},
				{
					Names: []*ast.Ident{{Name: "PublicField"}},
					Type:  &ast.Ident{Name: "bool"},
					Tag:   &ast.BasicLit{Value: "`getter:\"true\"`"},
				},
				{
					Names: []*ast.Ident{{Name: "ignoredField"}},
					Type:  &ast.Ident{Name: "float64"},
				},
			},
		},
	}

	result := generateGetterMethods(structName, requiredFields, optionalFields, structType)

	expected := `func (s *TestStruct) RequiredField() string {
	return s.requiredField
}

func (s *TestStruct) OptionalField() int {
	return s.optionalField
}

`

	if result != expected {
		t.Errorf("Generated getter methods do not match expected output.\nExpected:\n%s\nGot:\n%s", expected, result)
	}

	// Check that getters for PublicField and ignoredField are not generated
	if strings.Contains(result, "PublicField") {
		t.Error("Getter method was generated for a public field")
	}
	if strings.Contains(result, "ignoredField") {
		t.Error("Getter method was generated for a field without getter tag")
	}
}

func TestGenerateInterfaces(t *testing.T) {
	structType := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "Name"}},
					Type:  &ast.Ident{Name: "string"},
					Tag:   &ast.BasicLit{Value: "`structbuilder:\"required\"`"},
				},
				{
					Names: []*ast.Ident{{Name: "Age"}},
					Type:  &ast.Ident{Name: "int"},
					Tag:   &ast.BasicLit{Value: "`structbuilder:\"optional\"`"},
				},
			},
		},
	}

	result := generateInterfaces("Person", []string{"Name"}, []string{"Age"}, structType)

	expectedInterfaces := []string{
		"type PersonNameBuilder interface {",
		"type PersonOptionalBuilder interface {",
		"SetName(Name string) PersonOptionalBuilder", // TestStructOptionalBuilder -> PersonOptionalBuilder
		"SetOptAge(Age int) PersonOptionalBuilder",   // TestStructOptionalBuilder -> PersonOptionalBuilder
		"Build() *Person",
	}

	for _, expected := range expectedInterfaces {
		if !strings.Contains(result, expected) {
			t.Errorf("Expected interface to contain: %s", expected)
		}
	}
}

func TestGenerateSetterMethods(t *testing.T) {
	structType := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "Name"}},
					Type:  &ast.Ident{Name: "string"},
					Tag:   &ast.BasicLit{Value: "`structbuilder:\"required\"`"},
				},
				{
					Names: []*ast.Ident{{Name: "Age"}},
					Type:  &ast.Ident{Name: "int"},
					Tag:   &ast.BasicLit{Value: "`structbuilder:\"optional\"`"},
				},
			},
		},
	}

	requiredFields := []string{"Name"}
	optionalFields := []string{"Age"}
	result := generateSetterMethods("Person", requiredFields, optionalFields, structType)

	expectedMethods := []string{
		"func (b *personBuilder) SetName(Name string) PersonOptionalBuilder {",
		"func (b *personBuilder) SetOptAge(Age int) PersonOptionalBuilder {",
		"b.Name = Name",
		"b.Age = Age",
		"return b",
	}

	for _, expected := range expectedMethods {
		if !strings.Contains(result, expected) {
			t.Errorf("Expected setter methods to contain %q, but it didn't", expected)
		}
	}
}

func TestGenerateBuildMethod(t *testing.T) {
	requiredFields := []string{"Name"}
	optionalFields := []string{"Age"}
	result := generateBuildMethod("Person", requiredFields, optionalFields)

	expectedContent := []string{
		"func (b *personBuilder) Build() *Person {",
		"return &Person{",
		"Name: b.Name",
		"Age: b.Age",
		"}",
	}

	for _, expected := range expectedContent {
		if !strings.Contains(result, expected) {
			t.Errorf("Expected build method to contain %q, but it didn't", expected)
		}
	}
}

func TestProcessDirectoryErrors(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() string
		expectError bool
	}{
		{
			name: "非存在ディレクトリ",
			setup: func() string {
				return "/non/existent/directory"
			},
			expectError: true,
		},
		{
			name: "無効なGOファイル",
			setup: func() string {
				dir, err := os.MkdirTemp("", "invalid-go-test")
				if err != nil {
					t.Fatal(err)
				}
				err = os.WriteFile(filepath.Join(dir, "invalid.go"), []byte("invalid go code"), 0644)
				if err != nil {
					t.Fatal(err)
				}
				return dir
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.setup()
			if strings.HasPrefix(dir, os.TempDir()) {
				defer os.RemoveAll(dir)
			}

			_, err := processDirectory(dir)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestComplexTypeStructs(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "complex-struct-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// テスト用の複雑な構造体を含むGoファイルを作成
	testContent := `
package test

type Inner struct {
	Value string
}

type ComplexStruct struct {
	// 基本型
	StringField string ` + "`structbuilder:\"required\"`" + `
	IntField    int    ` + "`structbuilder:\"optional\"`" + `
	
	// 複合型
	SliceField  []string ` + "`structbuilder:\"required\"`" + `
	MapField    map[string]int ` + "`structbuilder:\"optional\"`" + `
	PointerField *string ` + "`structbuilder:\"required\"`" + `
	
	// カスタム型
	InnerStruct Inner ` + "`structbuilder:\"required\"`" + `
	
	// チャネル
	ChanField chan string ` + "`structbuilder:\"optional\"`" + `
}
`

	testFile := filepath.Join(tempDir, "complex.go")
	err = os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// ディレクトリを処理
	packages, err := processDirectory(tempDir)
	if err != nil {
		t.Fatalf("processDirectory failed: %v", err)
	}

	// 生成されたコードを検証
	pkg := packages[tempDir]
	generatedCode := pkg.GeneratedCode.String()

	expectedTypes := []string{
		"[]string",
		"map[string]int",
		"*string",
		"Inner",
		"chan string",
	}

	for _, expectedType := range expectedTypes {
		if !strings.Contains(generatedCode, expectedType) {
			t.Errorf("Generated code does not contain type %q", expectedType)
		}
	}

	// ビルダーメソッドの存在を確認
	expectedMethods := []string{
		"SetStringField",
		"SetSliceField",
		"SetPointerField",
		"SetInnerStruct",
		"SetOptIntField",
		"SetOptMapField",
		"SetOptChanField",
	}

	for _, method := range expectedMethods {
		if !strings.Contains(generatedCode, method) {
			t.Errorf("Generated code does not contain method %q", method)
		}
	}
}
