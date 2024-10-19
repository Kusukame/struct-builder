package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kusukame/struct-builder/internal/example/model"
)

func SomeUseCase() {
	example := model.NewExampleStructBuilder().
		SetIntValue(1).
		SetPtrIntValue(nil).
		SetStringValue("Hello, World!").
		SetPtrStringValue(nil).
		SetTimeValue(time.Now()).
		SetPtrTimeValue(nil).
		SetSliceValue([]string{"Hello", "World"}).
		SetPtrSliceValue(nil).
		SetMapValue(map[string]string{"Hello": "World"}).
		SetPtrMapValue(nil).
		SetSampleInterface(nil).
		SetOtherStruct(http.Request{}).
		SetOtherStructPtr(nil).
		Build()
	fmt.Printf("%+v\n", example)
}

func main() {
	SomeUseCase()
}
