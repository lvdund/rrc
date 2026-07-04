// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MTCH-SSB-MappingWindowList-r17 (line 28612).
// MTCH-SSB-MappingWindowList-r17 ::= SEQUENCE (SIZE (1..maxNrofMTCH-SSB-MappingWindow-r17)) OF MTCH-SSB-MappingWindowCycleOffset-r17

var mTCHSSBMappingWindowListR17SizeConstraints = per.SizeRange(1, common.MaxNrofMTCH_SSB_MappingWindow_r17)

type MTCH_SSB_MappingWindowList_r17 []MTCH_SSB_MappingWindowCycleOffset_r17

func (ie *MTCH_SSB_MappingWindowList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mTCHSSBMappingWindowListR17SizeConstraints)
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

func (ie *MTCH_SSB_MappingWindowList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mTCHSSBMappingWindowListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MTCH_SSB_MappingWindowList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
