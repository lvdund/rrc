// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-RxInterestedGC-BC-DestList-r17 (line 2194).
// SL-RxInterestedGC-BC-DestList-r17 ::=  SEQUENCE (SIZE (1..maxNrofSL-Dest-r16)) OF SL-RxInterestedGC-BC-Dest-r17

var sLRxInterestedGCBCDestListR17SizeConstraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SL_RxInterestedGC_BC_DestList_r17 []SL_RxInterestedGC_BC_Dest_r17

func (ie *SL_RxInterestedGC_BC_DestList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLRxInterestedGCBCDestListR17SizeConstraints)
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

func (ie *SL_RxInterestedGC_BC_DestList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLRxInterestedGCBCDestListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_RxInterestedGC_BC_DestList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
