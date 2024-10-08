// Code generated by struct-builder; DO NOT EDIT.

package model

import (
	"io"
	"net/http"
	"time"
)

type exampleStructBuilder struct {
	intValue                int
	ptrIntValue             *int
	stringValue             string
	ptrStringValue          *string
	timeValue               time.Time
	ptrTimeValue            *time.Time
	sliceValue              []string
	ptrSliceValue           *[]string
	mapValue                map[string]string
	ptrMapValue             *map[string]string
	sampleInterface         SampleInterface
	otherStruct             http.Request
	otherStructPtr          *http.Request
	optionalIntValue        *int
	optionalStringValue     *string
	optionalTimeValue       *time.Time
	optionalSliceValue      *[]string
	optionalMapValue        *map[string]string
	optionalOtherStruct     http.Request
	optionalOtherStructPtr  *http.Request
	optionalSampleInterface SampleInterface
}

type ExampleStructIntValueBuilder interface {
	SetIntValue(intValue int) ExampleStructPtrIntValueBuilder
}

type ExampleStructPtrIntValueBuilder interface {
	SetPtrIntValue(ptrIntValue *int) ExampleStructStringValueBuilder
}

type ExampleStructStringValueBuilder interface {
	SetStringValue(stringValue string) ExampleStructPtrStringValueBuilder
}

type ExampleStructPtrStringValueBuilder interface {
	SetPtrStringValue(ptrStringValue *string) ExampleStructTimeValueBuilder
}

type ExampleStructTimeValueBuilder interface {
	SetTimeValue(timeValue time.Time) ExampleStructPtrTimeValueBuilder
}

type ExampleStructPtrTimeValueBuilder interface {
	SetPtrTimeValue(ptrTimeValue *time.Time) ExampleStructSliceValueBuilder
}

type ExampleStructSliceValueBuilder interface {
	SetSliceValue(sliceValue []string) ExampleStructPtrSliceValueBuilder
}

type ExampleStructPtrSliceValueBuilder interface {
	SetPtrSliceValue(ptrSliceValue *[]string) ExampleStructMapValueBuilder
}

type ExampleStructMapValueBuilder interface {
	SetMapValue(mapValue map[string]string) ExampleStructPtrMapValueBuilder
}

type ExampleStructPtrMapValueBuilder interface {
	SetPtrMapValue(ptrMapValue *map[string]string) ExampleStructSampleInterfaceBuilder
}

type ExampleStructSampleInterfaceBuilder interface {
	SetSampleInterface(sampleInterface SampleInterface) ExampleStructOtherStructBuilder
}

type ExampleStructOtherStructBuilder interface {
	SetOtherStruct(otherStruct http.Request) ExampleStructOtherStructPtrBuilder
}

type ExampleStructOtherStructPtrBuilder interface {
	SetOtherStructPtr(otherStructPtr *http.Request) ExampleStructOptionalBuilder
}

type ExampleStructOptionalBuilder interface {
	SetOptOptionalIntValue(optionalIntValue *int) ExampleStructOptionalBuilder
	SetOptOptionalStringValue(optionalStringValue *string) ExampleStructOptionalBuilder
	SetOptOptionalTimeValue(optionalTimeValue *time.Time) ExampleStructOptionalBuilder
	SetOptOptionalSliceValue(optionalSliceValue *[]string) ExampleStructOptionalBuilder
	SetOptOptionalMapValue(optionalMapValue *map[string]string) ExampleStructOptionalBuilder
	SetOptOptionalOtherStruct(optionalOtherStruct http.Request) ExampleStructOptionalBuilder
	SetOptOptionalOtherStructPtr(optionalOtherStructPtr *http.Request) ExampleStructOptionalBuilder
	SetOptOptionalSampleInterface(optionalSampleInterface SampleInterface) ExampleStructOptionalBuilder
	Build() *ExampleStruct
}

func (b *exampleStructBuilder) SetIntValue(intValue int) ExampleStructPtrIntValueBuilder {
	b.intValue = intValue
	return b
}

func (b *exampleStructBuilder) SetPtrIntValue(ptrIntValue *int) ExampleStructStringValueBuilder {
	b.ptrIntValue = ptrIntValue
	return b
}

func (b *exampleStructBuilder) SetStringValue(stringValue string) ExampleStructPtrStringValueBuilder {
	b.stringValue = stringValue
	return b
}

func (b *exampleStructBuilder) SetPtrStringValue(ptrStringValue *string) ExampleStructTimeValueBuilder {
	b.ptrStringValue = ptrStringValue
	return b
}

func (b *exampleStructBuilder) SetTimeValue(timeValue time.Time) ExampleStructPtrTimeValueBuilder {
	b.timeValue = timeValue
	return b
}

func (b *exampleStructBuilder) SetPtrTimeValue(ptrTimeValue *time.Time) ExampleStructSliceValueBuilder {
	b.ptrTimeValue = ptrTimeValue
	return b
}

func (b *exampleStructBuilder) SetSliceValue(sliceValue []string) ExampleStructPtrSliceValueBuilder {
	b.sliceValue = sliceValue
	return b
}

func (b *exampleStructBuilder) SetPtrSliceValue(ptrSliceValue *[]string) ExampleStructMapValueBuilder {
	b.ptrSliceValue = ptrSliceValue
	return b
}

func (b *exampleStructBuilder) SetMapValue(mapValue map[string]string) ExampleStructPtrMapValueBuilder {
	b.mapValue = mapValue
	return b
}

func (b *exampleStructBuilder) SetPtrMapValue(ptrMapValue *map[string]string) ExampleStructSampleInterfaceBuilder {
	b.ptrMapValue = ptrMapValue
	return b
}

func (b *exampleStructBuilder) SetSampleInterface(sampleInterface SampleInterface) ExampleStructOtherStructBuilder {
	b.sampleInterface = sampleInterface
	return b
}

func (b *exampleStructBuilder) SetOtherStruct(otherStruct http.Request) ExampleStructOtherStructPtrBuilder {
	b.otherStruct = otherStruct
	return b
}

func (b *exampleStructBuilder) SetOtherStructPtr(otherStructPtr *http.Request) ExampleStructOptionalBuilder {
	b.otherStructPtr = otherStructPtr
	return b
}

func (b *exampleStructBuilder) SetOptOptionalIntValue(optionalIntValue *int) ExampleStructOptionalBuilder {
	b.optionalIntValue = optionalIntValue
	return b
}

func (b *exampleStructBuilder) SetOptOptionalStringValue(optionalStringValue *string) ExampleStructOptionalBuilder {
	b.optionalStringValue = optionalStringValue
	return b
}

func (b *exampleStructBuilder) SetOptOptionalTimeValue(optionalTimeValue *time.Time) ExampleStructOptionalBuilder {
	b.optionalTimeValue = optionalTimeValue
	return b
}

func (b *exampleStructBuilder) SetOptOptionalSliceValue(optionalSliceValue *[]string) ExampleStructOptionalBuilder {
	b.optionalSliceValue = optionalSliceValue
	return b
}

func (b *exampleStructBuilder) SetOptOptionalMapValue(optionalMapValue *map[string]string) ExampleStructOptionalBuilder {
	b.optionalMapValue = optionalMapValue
	return b
}

func (b *exampleStructBuilder) SetOptOptionalOtherStruct(optionalOtherStruct http.Request) ExampleStructOptionalBuilder {
	b.optionalOtherStruct = optionalOtherStruct
	return b
}

func (b *exampleStructBuilder) SetOptOptionalOtherStructPtr(optionalOtherStructPtr *http.Request) ExampleStructOptionalBuilder {
	b.optionalOtherStructPtr = optionalOtherStructPtr
	return b
}

func (b *exampleStructBuilder) SetOptOptionalSampleInterface(optionalSampleInterface SampleInterface) ExampleStructOptionalBuilder {
	b.optionalSampleInterface = optionalSampleInterface
	return b
}

func (s *ExampleStruct) IntValue() int {
	return s.intValue
}

func (s *ExampleStruct) PtrIntValue() *int {
	return s.ptrIntValue
}

func (s *ExampleStruct) StringValue() string {
	return s.stringValue
}

func (s *ExampleStruct) PtrStringValue() *string {
	return s.ptrStringValue
}

func (s *ExampleStruct) TimeValue() time.Time {
	return s.timeValue
}

func (s *ExampleStruct) PtrTimeValue() *time.Time {
	return s.ptrTimeValue
}

func (s *ExampleStruct) SliceValue() []string {
	return s.sliceValue
}

func (s *ExampleStruct) PtrSliceValue() *[]string {
	return s.ptrSliceValue
}

func (b *exampleStructBuilder) Build() *ExampleStruct {
	return &ExampleStruct{
		intValue:                b.intValue,
		ptrIntValue:             b.ptrIntValue,
		stringValue:             b.stringValue,
		ptrStringValue:          b.ptrStringValue,
		timeValue:               b.timeValue,
		ptrTimeValue:            b.ptrTimeValue,
		sliceValue:              b.sliceValue,
		ptrSliceValue:           b.ptrSliceValue,
		mapValue:                b.mapValue,
		ptrMapValue:             b.ptrMapValue,
		sampleInterface:         b.sampleInterface,
		otherStruct:             b.otherStruct,
		otherStructPtr:          b.otherStructPtr,
		optionalIntValue:        b.optionalIntValue,
		optionalStringValue:     b.optionalStringValue,
		optionalTimeValue:       b.optionalTimeValue,
		optionalSliceValue:      b.optionalSliceValue,
		optionalMapValue:        b.optionalMapValue,
		optionalOtherStruct:     b.optionalOtherStruct,
		optionalOtherStructPtr:  b.optionalOtherStructPtr,
		optionalSampleInterface: b.optionalSampleInterface,
	}
}

func NewExampleStructBuilder() ExampleStructIntValueBuilder {
	return &exampleStructBuilder{}
}

func NewExampleStruct(intValue int, ptrIntValue *int, stringValue string, ptrStringValue *string, timeValue time.Time, ptrTimeValue *time.Time, sliceValue []string, ptrSliceValue *[]string, mapValue map[string]string, ptrMapValue *map[string]string, sampleInterface SampleInterface, otherStruct http.Request, otherStructPtr *http.Request, optionalIntValue *int, optionalStringValue *string, optionalTimeValue *time.Time, optionalSliceValue *[]string, optionalMapValue *map[string]string, optionalOtherStruct http.Request, optionalOtherStructPtr *http.Request, optionalSampleInterface SampleInterface) *ExampleStruct {
	return &ExampleStruct{
		intValue:                intValue,
		ptrIntValue:             ptrIntValue,
		stringValue:             stringValue,
		ptrStringValue:          ptrStringValue,
		timeValue:               timeValue,
		ptrTimeValue:            ptrTimeValue,
		sliceValue:              sliceValue,
		ptrSliceValue:           ptrSliceValue,
		mapValue:                mapValue,
		ptrMapValue:             ptrMapValue,
		sampleInterface:         sampleInterface,
		otherStruct:             otherStruct,
		otherStructPtr:          otherStructPtr,
		optionalIntValue:        optionalIntValue,
		optionalStringValue:     optionalStringValue,
		optionalTimeValue:       optionalTimeValue,
		optionalSliceValue:      optionalSliceValue,
		optionalMapValue:        optionalMapValue,
		optionalOtherStruct:     optionalOtherStruct,
		optionalOtherStructPtr:  optionalOtherStructPtr,
		optionalSampleInterface: optionalSampleInterface,
	}
}

type exampleStruct2Builder struct {
	intValue               int
	ptrIntValue            *int
	stringValue            string
	ptrStringValue         *string
	timeValue              time.Time
	ptrTimeValue           *time.Time
	sliceValue             []string
	ptrSliceValue          *[]string
	mapValue               map[string]string
	ptrMapValue            *map[string]string
	otherStruct            io.ByteReader
	otherStructPtr         *io.ByteReader
	optionalIntValue       *int
	optionalStringValue    *string
	optionalTimeValue      *time.Time
	optionalSliceValue     *[]string
	optionalMapValue       *map[string]string
	optionalOtherStruct    io.ByteReader
	optionalOtherStructPtr *io.ByteReader
	chanValue              chan int
	optionalChanValue      chan int
	chanValue2             <-chan int
	optionalChanValue2     <-chan int
	chanValue3             chan<- int
	optionalChanValue3     chan<- int
}

type ExampleStruct2IntValueBuilder interface {
	SetIntValue(intValue int) ExampleStruct2PtrIntValueBuilder
}

type ExampleStruct2PtrIntValueBuilder interface {
	SetPtrIntValue(ptrIntValue *int) ExampleStruct2StringValueBuilder
}

type ExampleStruct2StringValueBuilder interface {
	SetStringValue(stringValue string) ExampleStruct2PtrStringValueBuilder
}

type ExampleStruct2PtrStringValueBuilder interface {
	SetPtrStringValue(ptrStringValue *string) ExampleStruct2TimeValueBuilder
}

type ExampleStruct2TimeValueBuilder interface {
	SetTimeValue(timeValue time.Time) ExampleStruct2PtrTimeValueBuilder
}

type ExampleStruct2PtrTimeValueBuilder interface {
	SetPtrTimeValue(ptrTimeValue *time.Time) ExampleStruct2SliceValueBuilder
}

type ExampleStruct2SliceValueBuilder interface {
	SetSliceValue(sliceValue []string) ExampleStruct2PtrSliceValueBuilder
}

type ExampleStruct2PtrSliceValueBuilder interface {
	SetPtrSliceValue(ptrSliceValue *[]string) ExampleStruct2MapValueBuilder
}

type ExampleStruct2MapValueBuilder interface {
	SetMapValue(mapValue map[string]string) ExampleStruct2PtrMapValueBuilder
}

type ExampleStruct2PtrMapValueBuilder interface {
	SetPtrMapValue(ptrMapValue *map[string]string) ExampleStruct2OtherStructBuilder
}

type ExampleStruct2OtherStructBuilder interface {
	SetOtherStruct(otherStruct io.ByteReader) ExampleStruct2OtherStructPtrBuilder
}

type ExampleStruct2OtherStructPtrBuilder interface {
	SetOtherStructPtr(otherStructPtr *io.ByteReader) ExampleStruct2ChanValueBuilder
}

type ExampleStruct2ChanValueBuilder interface {
	SetChanValue(chanValue chan int) ExampleStruct2ChanValue2Builder
}

type ExampleStruct2ChanValue2Builder interface {
	SetChanValue2(chanValue2 <-chan int) ExampleStruct2ChanValue3Builder
}

type ExampleStruct2ChanValue3Builder interface {
	SetChanValue3(chanValue3 chan<- int) ExampleStruct2OptionalBuilder
}

type ExampleStruct2OptionalBuilder interface {
	SetOptOptionalIntValue(optionalIntValue *int) ExampleStruct2OptionalBuilder
	SetOptOptionalStringValue(optionalStringValue *string) ExampleStruct2OptionalBuilder
	SetOptOptionalTimeValue(optionalTimeValue *time.Time) ExampleStruct2OptionalBuilder
	SetOptOptionalSliceValue(optionalSliceValue *[]string) ExampleStruct2OptionalBuilder
	SetOptOptionalMapValue(optionalMapValue *map[string]string) ExampleStruct2OptionalBuilder
	SetOptOptionalOtherStruct(optionalOtherStruct io.ByteReader) ExampleStruct2OptionalBuilder
	SetOptOptionalOtherStructPtr(optionalOtherStructPtr *io.ByteReader) ExampleStruct2OptionalBuilder
	SetOptOptionalChanValue(optionalChanValue chan int) ExampleStruct2OptionalBuilder
	SetOptOptionalChanValue2(optionalChanValue2 <-chan int) ExampleStruct2OptionalBuilder
	SetOptOptionalChanValue3(optionalChanValue3 chan<- int) ExampleStruct2OptionalBuilder
	Build() *ExampleStruct2
}

func (b *exampleStruct2Builder) SetIntValue(intValue int) ExampleStruct2PtrIntValueBuilder {
	b.intValue = intValue
	return b
}

func (b *exampleStruct2Builder) SetPtrIntValue(ptrIntValue *int) ExampleStruct2StringValueBuilder {
	b.ptrIntValue = ptrIntValue
	return b
}

func (b *exampleStruct2Builder) SetStringValue(stringValue string) ExampleStruct2PtrStringValueBuilder {
	b.stringValue = stringValue
	return b
}

func (b *exampleStruct2Builder) SetPtrStringValue(ptrStringValue *string) ExampleStruct2TimeValueBuilder {
	b.ptrStringValue = ptrStringValue
	return b
}

func (b *exampleStruct2Builder) SetTimeValue(timeValue time.Time) ExampleStruct2PtrTimeValueBuilder {
	b.timeValue = timeValue
	return b
}

func (b *exampleStruct2Builder) SetPtrTimeValue(ptrTimeValue *time.Time) ExampleStruct2SliceValueBuilder {
	b.ptrTimeValue = ptrTimeValue
	return b
}

func (b *exampleStruct2Builder) SetSliceValue(sliceValue []string) ExampleStruct2PtrSliceValueBuilder {
	b.sliceValue = sliceValue
	return b
}

func (b *exampleStruct2Builder) SetPtrSliceValue(ptrSliceValue *[]string) ExampleStruct2MapValueBuilder {
	b.ptrSliceValue = ptrSliceValue
	return b
}

func (b *exampleStruct2Builder) SetMapValue(mapValue map[string]string) ExampleStruct2PtrMapValueBuilder {
	b.mapValue = mapValue
	return b
}

func (b *exampleStruct2Builder) SetPtrMapValue(ptrMapValue *map[string]string) ExampleStruct2OtherStructBuilder {
	b.ptrMapValue = ptrMapValue
	return b
}

func (b *exampleStruct2Builder) SetOtherStruct(otherStruct io.ByteReader) ExampleStruct2OtherStructPtrBuilder {
	b.otherStruct = otherStruct
	return b
}

func (b *exampleStruct2Builder) SetOtherStructPtr(otherStructPtr *io.ByteReader) ExampleStruct2ChanValueBuilder {
	b.otherStructPtr = otherStructPtr
	return b
}

func (b *exampleStruct2Builder) SetChanValue(chanValue chan int) ExampleStruct2ChanValue2Builder {
	b.chanValue = chanValue
	return b
}

func (b *exampleStruct2Builder) SetChanValue2(chanValue2 <-chan int) ExampleStruct2ChanValue3Builder {
	b.chanValue2 = chanValue2
	return b
}

func (b *exampleStruct2Builder) SetChanValue3(chanValue3 chan<- int) ExampleStruct2OptionalBuilder {
	b.chanValue3 = chanValue3
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalIntValue(optionalIntValue *int) ExampleStruct2OptionalBuilder {
	b.optionalIntValue = optionalIntValue
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalStringValue(optionalStringValue *string) ExampleStruct2OptionalBuilder {
	b.optionalStringValue = optionalStringValue
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalTimeValue(optionalTimeValue *time.Time) ExampleStruct2OptionalBuilder {
	b.optionalTimeValue = optionalTimeValue
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalSliceValue(optionalSliceValue *[]string) ExampleStruct2OptionalBuilder {
	b.optionalSliceValue = optionalSliceValue
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalMapValue(optionalMapValue *map[string]string) ExampleStruct2OptionalBuilder {
	b.optionalMapValue = optionalMapValue
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalOtherStruct(optionalOtherStruct io.ByteReader) ExampleStruct2OptionalBuilder {
	b.optionalOtherStruct = optionalOtherStruct
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalOtherStructPtr(optionalOtherStructPtr *io.ByteReader) ExampleStruct2OptionalBuilder {
	b.optionalOtherStructPtr = optionalOtherStructPtr
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalChanValue(optionalChanValue chan int) ExampleStruct2OptionalBuilder {
	b.optionalChanValue = optionalChanValue
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalChanValue2(optionalChanValue2 <-chan int) ExampleStruct2OptionalBuilder {
	b.optionalChanValue2 = optionalChanValue2
	return b
}

func (b *exampleStruct2Builder) SetOptOptionalChanValue3(optionalChanValue3 chan<- int) ExampleStruct2OptionalBuilder {
	b.optionalChanValue3 = optionalChanValue3
	return b
}

func (b *exampleStruct2Builder) Build() *ExampleStruct2 {
	return &ExampleStruct2{
		intValue:               b.intValue,
		ptrIntValue:            b.ptrIntValue,
		stringValue:            b.stringValue,
		ptrStringValue:         b.ptrStringValue,
		timeValue:              b.timeValue,
		ptrTimeValue:           b.ptrTimeValue,
		sliceValue:             b.sliceValue,
		ptrSliceValue:          b.ptrSliceValue,
		mapValue:               b.mapValue,
		ptrMapValue:            b.ptrMapValue,
		otherStruct:            b.otherStruct,
		otherStructPtr:         b.otherStructPtr,
		chanValue:              b.chanValue,
		chanValue2:             b.chanValue2,
		chanValue3:             b.chanValue3,
		optionalIntValue:       b.optionalIntValue,
		optionalStringValue:    b.optionalStringValue,
		optionalTimeValue:      b.optionalTimeValue,
		optionalSliceValue:     b.optionalSliceValue,
		optionalMapValue:       b.optionalMapValue,
		optionalOtherStruct:    b.optionalOtherStruct,
		optionalOtherStructPtr: b.optionalOtherStructPtr,
		optionalChanValue:      b.optionalChanValue,
		optionalChanValue2:     b.optionalChanValue2,
		optionalChanValue3:     b.optionalChanValue3,
	}
}

func NewExampleStruct2Builder() ExampleStruct2IntValueBuilder {
	return &exampleStruct2Builder{}
}

func NewExampleStruct2(intValue int, ptrIntValue *int, stringValue string, ptrStringValue *string, timeValue time.Time, ptrTimeValue *time.Time, sliceValue []string, ptrSliceValue *[]string, mapValue map[string]string, ptrMapValue *map[string]string, otherStruct io.ByteReader, otherStructPtr *io.ByteReader, chanValue chan int, chanValue2 <-chan int, chanValue3 chan<- int, optionalIntValue *int, optionalStringValue *string, optionalTimeValue *time.Time, optionalSliceValue *[]string, optionalMapValue *map[string]string, optionalOtherStruct io.ByteReader, optionalOtherStructPtr *io.ByteReader, optionalChanValue chan int, optionalChanValue2 <-chan int, optionalChanValue3 chan<- int) *ExampleStruct2 {
	return &ExampleStruct2{
		intValue:               intValue,
		ptrIntValue:            ptrIntValue,
		stringValue:            stringValue,
		ptrStringValue:         ptrStringValue,
		timeValue:              timeValue,
		ptrTimeValue:           ptrTimeValue,
		sliceValue:             sliceValue,
		ptrSliceValue:          ptrSliceValue,
		mapValue:               mapValue,
		ptrMapValue:            ptrMapValue,
		otherStruct:            otherStruct,
		otherStructPtr:         otherStructPtr,
		chanValue:              chanValue,
		chanValue2:             chanValue2,
		chanValue3:             chanValue3,
		optionalIntValue:       optionalIntValue,
		optionalStringValue:    optionalStringValue,
		optionalTimeValue:      optionalTimeValue,
		optionalSliceValue:     optionalSliceValue,
		optionalMapValue:       optionalMapValue,
		optionalOtherStruct:    optionalOtherStruct,
		optionalOtherStructPtr: optionalOtherStructPtr,
		optionalChanValue:      optionalChanValue,
		optionalChanValue2:     optionalChanValue2,
		optionalChanValue3:     optionalChanValue3,
	}
}

type exampleStruct3Builder struct {
	optionalIntValue       *int
	optionalStringValue    *string
	optionalTimeValue      *time.Time
	optionalSliceValue     *[]string
	optionalMapValue       *map[string]string
	optionalOtherStruct    io.ByteReader
	optionalOtherStructPtr *io.ByteReader
}

type ExampleStruct3OptionalBuilder interface {
	SetOptOptionalIntValue(optionalIntValue *int) ExampleStruct3OptionalBuilder
	SetOptOptionalStringValue(optionalStringValue *string) ExampleStruct3OptionalBuilder
	SetOptOptionalTimeValue(optionalTimeValue *time.Time) ExampleStruct3OptionalBuilder
	SetOptOptionalSliceValue(optionalSliceValue *[]string) ExampleStruct3OptionalBuilder
	SetOptOptionalMapValue(optionalMapValue *map[string]string) ExampleStruct3OptionalBuilder
	SetOptOptionalOtherStruct(optionalOtherStruct io.ByteReader) ExampleStruct3OptionalBuilder
	SetOptOptionalOtherStructPtr(optionalOtherStructPtr *io.ByteReader) ExampleStruct3OptionalBuilder
	Build() *ExampleStruct3
}

func (b *exampleStruct3Builder) SetOptOptionalIntValue(optionalIntValue *int) ExampleStruct3OptionalBuilder {
	b.optionalIntValue = optionalIntValue
	return b
}

func (b *exampleStruct3Builder) SetOptOptionalStringValue(optionalStringValue *string) ExampleStruct3OptionalBuilder {
	b.optionalStringValue = optionalStringValue
	return b
}

func (b *exampleStruct3Builder) SetOptOptionalTimeValue(optionalTimeValue *time.Time) ExampleStruct3OptionalBuilder {
	b.optionalTimeValue = optionalTimeValue
	return b
}

func (b *exampleStruct3Builder) SetOptOptionalSliceValue(optionalSliceValue *[]string) ExampleStruct3OptionalBuilder {
	b.optionalSliceValue = optionalSliceValue
	return b
}

func (b *exampleStruct3Builder) SetOptOptionalMapValue(optionalMapValue *map[string]string) ExampleStruct3OptionalBuilder {
	b.optionalMapValue = optionalMapValue
	return b
}

func (b *exampleStruct3Builder) SetOptOptionalOtherStruct(optionalOtherStruct io.ByteReader) ExampleStruct3OptionalBuilder {
	b.optionalOtherStruct = optionalOtherStruct
	return b
}

func (b *exampleStruct3Builder) SetOptOptionalOtherStructPtr(optionalOtherStructPtr *io.ByteReader) ExampleStruct3OptionalBuilder {
	b.optionalOtherStructPtr = optionalOtherStructPtr
	return b
}

func (b *exampleStruct3Builder) Build() *ExampleStruct3 {
	return &ExampleStruct3{
		optionalIntValue:       b.optionalIntValue,
		optionalStringValue:    b.optionalStringValue,
		optionalTimeValue:      b.optionalTimeValue,
		optionalSliceValue:     b.optionalSliceValue,
		optionalMapValue:       b.optionalMapValue,
		optionalOtherStruct:    b.optionalOtherStruct,
		optionalOtherStructPtr: b.optionalOtherStructPtr,
	}
}

func NewExampleStruct3Builder() ExampleStruct3OptionalBuilder {
	return &exampleStruct3Builder{}
}

func NewExampleStruct3(optionalIntValue *int, optionalStringValue *string, optionalTimeValue *time.Time, optionalSliceValue *[]string, optionalMapValue *map[string]string, optionalOtherStruct io.ByteReader, optionalOtherStructPtr *io.ByteReader) *ExampleStruct3 {
	return &ExampleStruct3{
		optionalIntValue:       optionalIntValue,
		optionalStringValue:    optionalStringValue,
		optionalTimeValue:      optionalTimeValue,
		optionalSliceValue:     optionalSliceValue,
		optionalMapValue:       optionalMapValue,
		optionalOtherStruct:    optionalOtherStruct,
		optionalOtherStructPtr: optionalOtherStructPtr,
	}
}
