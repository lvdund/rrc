// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReferenceTime-r16 (line 13328).

var referenceTimeR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "refDays-r16"},
		{Name: "refSeconds-r16"},
		{Name: "refMilliSeconds-r16"},
		{Name: "refTenNanoSeconds-r16"},
	},
}

var referenceTimeR16RefDaysR16Constraints = per.Constrained(0, 72999)

var referenceTimeR16RefSecondsR16Constraints = per.Constrained(0, 86399)

var referenceTimeR16RefMilliSecondsR16Constraints = per.Constrained(0, 999)

var referenceTimeR16RefTenNanoSecondsR16Constraints = per.Constrained(0, 99999)

type ReferenceTime_r16 struct {
	RefDays_r16           int64
	RefSeconds_r16        int64
	RefMilliSeconds_r16   int64
	RefTenNanoSeconds_r16 int64
}

func (ie *ReferenceTime_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(referenceTimeR16Constraints)
	_ = seq

	// 1. refDays-r16: integer
	{
		if err := e.EncodeInteger(ie.RefDays_r16, referenceTimeR16RefDaysR16Constraints); err != nil {
			return err
		}
	}

	// 2. refSeconds-r16: integer
	{
		if err := e.EncodeInteger(ie.RefSeconds_r16, referenceTimeR16RefSecondsR16Constraints); err != nil {
			return err
		}
	}

	// 3. refMilliSeconds-r16: integer
	{
		if err := e.EncodeInteger(ie.RefMilliSeconds_r16, referenceTimeR16RefMilliSecondsR16Constraints); err != nil {
			return err
		}
	}

	// 4. refTenNanoSeconds-r16: integer
	{
		if err := e.EncodeInteger(ie.RefTenNanoSeconds_r16, referenceTimeR16RefTenNanoSecondsR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ReferenceTime_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(referenceTimeR16Constraints)
	_ = seq

	// 1. refDays-r16: integer
	{
		v0, err := d.DecodeInteger(referenceTimeR16RefDaysR16Constraints)
		if err != nil {
			return err
		}
		ie.RefDays_r16 = v0
	}

	// 2. refSeconds-r16: integer
	{
		v1, err := d.DecodeInteger(referenceTimeR16RefSecondsR16Constraints)
		if err != nil {
			return err
		}
		ie.RefSeconds_r16 = v1
	}

	// 3. refMilliSeconds-r16: integer
	{
		v2, err := d.DecodeInteger(referenceTimeR16RefMilliSecondsR16Constraints)
		if err != nil {
			return err
		}
		ie.RefMilliSeconds_r16 = v2
	}

	// 4. refTenNanoSeconds-r16: integer
	{
		v3, err := d.DecodeInteger(referenceTimeR16RefTenNanoSecondsR16Constraints)
		if err != nil {
			return err
		}
		ie.RefTenNanoSeconds_r16 = v3
	}

	return nil
}
