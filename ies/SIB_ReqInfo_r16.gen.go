// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB-ReqInfo-r16 (line 281).
// SIB-ReqInfo-r16 ::=                   ENUMERATED { sib12, sib13, sib14, sib20-v1700, sib21-v1700, sib23-v1810, sib26-v1900, sib27-v1900 }

const (
	SIB_ReqInfo_r16_Sib12       = 0
	SIB_ReqInfo_r16_Sib13       = 1
	SIB_ReqInfo_r16_Sib14       = 2
	SIB_ReqInfo_r16_Sib20_v1700 = 3
	SIB_ReqInfo_r16_Sib21_v1700 = 4
	SIB_ReqInfo_r16_Sib23_v1810 = 5
	SIB_ReqInfo_r16_Sib26_v1900 = 6
	SIB_ReqInfo_r16_Sib27_v1900 = 7
)

var sIBReqInfoR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB_ReqInfo_r16_Sib12, SIB_ReqInfo_r16_Sib13, SIB_ReqInfo_r16_Sib14, SIB_ReqInfo_r16_Sib20_v1700, SIB_ReqInfo_r16_Sib21_v1700, SIB_ReqInfo_r16_Sib23_v1810, SIB_ReqInfo_r16_Sib26_v1900, SIB_ReqInfo_r16_Sib27_v1900},
}

type SIB_ReqInfo_r16 struct {
	Value int64
}

func (v *SIB_ReqInfo_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sIBReqInfoR16Constraints)
}

func (v *SIB_ReqInfo_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sIBReqInfoR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
