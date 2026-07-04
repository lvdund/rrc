// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-Thres-RSRP-List-r16 (line 28338).
// SL-Thres-RSRP-List-r16 ::=    SEQUENCE (SIZE (64)) OF SL-Thres-RSRP-r16

var sLThresRSRPListR16SizeConstraints = per.FixedSize(64)

type SL_Thres_RSRP_List_r16 []SL_Thres_RSRP_r16

func (ie *SL_Thres_RSRP_List_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLThresRSRPListR16SizeConstraints)
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

func (ie *SL_Thres_RSRP_List_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLThresRSRPListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_Thres_RSRP_List_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
