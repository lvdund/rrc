// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SPS-PUCCH-AN-List-r16 (line 15271).
// SPS-PUCCH-AN-List-r16 ::=      SEQUENCE (SIZE(1..4)) OF SPS-PUCCH-AN-r16

var sPSPUCCHANListR16SizeConstraints = per.SizeRange(1, 4)

type SPS_PUCCH_AN_List_r16 []SPS_PUCCH_AN_r16

func (ie *SPS_PUCCH_AN_List_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sPSPUCCHANListR16SizeConstraints)
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

func (ie *SPS_PUCCH_AN_List_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sPSPUCCHANListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SPS_PUCCH_AN_List_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
