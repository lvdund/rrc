// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggingInterval-r16 (line 26243).

const (
	LoggingInterval_r16_Ms320    = 0
	LoggingInterval_r16_Ms640    = 1
	LoggingInterval_r16_Ms1280   = 2
	LoggingInterval_r16_Ms2560   = 3
	LoggingInterval_r16_Ms5120   = 4
	LoggingInterval_r16_Ms10240  = 5
	LoggingInterval_r16_Ms20480  = 6
	LoggingInterval_r16_Ms30720  = 7
	LoggingInterval_r16_Ms40960  = 8
	LoggingInterval_r16_Ms61440  = 9
	LoggingInterval_r16_Infinity = 10
)

var loggingIntervalR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LoggingInterval_r16_Ms320, LoggingInterval_r16_Ms640, LoggingInterval_r16_Ms1280, LoggingInterval_r16_Ms2560, LoggingInterval_r16_Ms5120, LoggingInterval_r16_Ms10240, LoggingInterval_r16_Ms20480, LoggingInterval_r16_Ms30720, LoggingInterval_r16_Ms40960, LoggingInterval_r16_Ms61440, LoggingInterval_r16_Infinity},
}

type LoggingInterval_r16 struct {
	Value int64
}

func (v *LoggingInterval_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, loggingIntervalR16Constraints)
}

func (v *LoggingInterval_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(loggingIntervalR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
