// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: JitterBound-r18 (line 2789).
// JitterBound-r18 ::= ENUMERATED {ms0, ms0dot5, ms1, ms1dot5, ms2, ms2dot5, ms3, ms3dot5, ms4, ms4dot5, ms5, ms5dot5, ms6, ms6dot5, ms7, beyondMs7}

const (
	JitterBound_r18_Ms0       = 0
	JitterBound_r18_Ms0dot5   = 1
	JitterBound_r18_Ms1       = 2
	JitterBound_r18_Ms1dot5   = 3
	JitterBound_r18_Ms2       = 4
	JitterBound_r18_Ms2dot5   = 5
	JitterBound_r18_Ms3       = 6
	JitterBound_r18_Ms3dot5   = 7
	JitterBound_r18_Ms4       = 8
	JitterBound_r18_Ms4dot5   = 9
	JitterBound_r18_Ms5       = 10
	JitterBound_r18_Ms5dot5   = 11
	JitterBound_r18_Ms6       = 12
	JitterBound_r18_Ms6dot5   = 13
	JitterBound_r18_Ms7       = 14
	JitterBound_r18_BeyondMs7 = 15
)

var jitterBoundR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{JitterBound_r18_Ms0, JitterBound_r18_Ms0dot5, JitterBound_r18_Ms1, JitterBound_r18_Ms1dot5, JitterBound_r18_Ms2, JitterBound_r18_Ms2dot5, JitterBound_r18_Ms3, JitterBound_r18_Ms3dot5, JitterBound_r18_Ms4, JitterBound_r18_Ms4dot5, JitterBound_r18_Ms5, JitterBound_r18_Ms5dot5, JitterBound_r18_Ms6, JitterBound_r18_Ms6dot5, JitterBound_r18_Ms7, JitterBound_r18_BeyondMs7},
}

type JitterBound_r18 struct {
	Value int64
}

func (v *JitterBound_r18) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, jitterBoundR18Constraints)
}

func (v *JitterBound_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(jitterBoundR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
