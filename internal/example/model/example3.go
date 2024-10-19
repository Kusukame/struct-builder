package model

import (
	"io"
	"time"
)

type ExampleStruct3 struct {
	optionalIntValue       *int               `structbuilder:"optional"`
	optionalStringValue    *string            `structbuilder:"optional"`
	optionalTimeValue      *time.Time         `structbuilder:"optional"`
	optionalSliceValue     *[]string          `structbuilder:"optional"`
	optionalMapValue       *map[string]string `structbuilder:"optional"`
	optionalOtherStruct    io.ByteReader      `structbuilder:"optional"`
	optionalOtherStructPtr *io.ByteReader     `structbuilder:"optional"`
}
