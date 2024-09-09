package model

import (
	"fmt"
	"net/http"
	"time"
)

type SampleInterface interface {
	SampleMethod()
}

type ExampleStruct struct {
	intValue int  `structbuilder:"required;getter"`
	ptrIntValue *int  `structbuilder:"required;getter"`
	stringValue string  `structbuilder:"required;getter"`
	ptrStringValue *string  `structbuilder:"required;getter"`
	timeValue time.Time  `structbuilder:"required;getter"`
	ptrTimeValue *time.Time  `structbuilder:"required;getter"`
	sliceValue []string  `structbuilder:"required;getter"`
	ptrSliceValue *[]string  `structbuilder:"required;getter"`
	mapValue map[string]string  `structbuilder:"required"`
	ptrMapValue *map[string]string  `structbuilder:"required"`
	sampleInterface SampleInterface  `structbuilder:"required"`
	otherStruct http.Request  `structbuilder:"required"`
	otherStructPtr *http.Request  `structbuilder:"required"`
	optionalIntValue *int  `structbuilder:"optional"`
	optionalStringValue *string  `structbuilder:"optional"`
	optionalTimeValue *time.Time  `structbuilder:"optional"`
	optionalSliceValue *[]string  `structbuilder:"optional"`
	optionalMapValue *map[string]string  `structbuilder:"optional"`
	optionalOtherStruct http.Request  `structbuilder:"optional"`
	optionalOtherStructPtr *http.Request  `structbuilder:"optional"`
	optionalSampleInterface SampleInterface  `structbuilder:"optional"`
}


func (b *ExampleStruct) SetIntValue(intValue int)  {
	b.intValue = intValue
	fmt.Println(b.intValue)
}
