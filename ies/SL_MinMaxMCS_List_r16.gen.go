// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MinMaxMCS-List-r16 (line 28096).
// SL-MinMaxMCS-List-r16 ::=              SEQUENCE (SIZE (1..3)) OF SL-MinMaxMCS-Config-r16

var sLMinMaxMCSListR16SizeConstraints = per.SizeRange(1, 3)

type SL_MinMaxMCS_List_r16 []SL_MinMaxMCS_Config_r16

func (ie *SL_MinMaxMCS_List_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLMinMaxMCSListR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SL_MinMaxMCS_List_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLMinMaxMCSListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_MinMaxMCS_List_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
