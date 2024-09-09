package model

import (
	"fmt"
	"io"
	"time"
)

type ExampleStruct2 struct {
	intValue int  `structbuilder:"required"`
	ptrIntValue *int  `structbuilder:"required"`
	stringValue string  `structbuilder:"required"`
	ptrStringValue *string  `structbuilder:"required"`
	timeValue time.Time  `structbuilder:"required"`
	ptrTimeValue *time.Time  `structbuilder:"required"`
	sliceValue []string  `structbuilder:"required"`
	ptrSliceValue *[]string  `structbuilder:"required"`
	mapValue map[string]string  `structbuilder:"required"`
	ptrMapValue *map[string]string  `structbuilder:"required"`
	otherStruct io.ByteReader  `structbuilder:"required"`
	otherStructPtr *io.ByteReader  `structbuilder:"required"`
	optionalIntValue *int  `structbuilder:"optional"`
	optionalStringValue *string  `structbuilder:"optional"`
	optionalTimeValue *time.Time  `structbuilder:"optional"`
	optionalSliceValue *[]string  `structbuilder:"optional"`
	optionalMapValue *map[string]string  `structbuilder:"optional"`
	optionalOtherStruct io.ByteReader  `structbuilder:"optional"`
	optionalOtherStructPtr *io.ByteReader  `structbuilder:"optional"`
	chanValue chan int  `structbuilder:"required"`
	optionalChanValue chan int  `structbuilder:"optional"`
	chanValue2 <-chan int  `structbuilder:"required"`
	optionalChanValue2 <-chan int  `structbuilder:"optional"`
	chanValue3 chan<- int  `structbuilder:"required"`
	optionalChanValue3 chan<- int  `structbuilder:"optional"`
}


func (b *ExampleStruct2) SetIntValue(intValue int)  {
	b.intValue = intValue
	fmt.Println(b.intValue)
}
