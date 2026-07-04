// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SPS-ConfigMulticastToReleaseList-r17 (line 5907).
// SPS-ConfigMulticastToReleaseList-r17 ::= SEQUENCE (SIZE (1..8)) OF SPS-ConfigIndex-r16

var sPSConfigMulticastToReleaseListR17SizeConstraints = per.SizeRange(1, 8)

type SPS_ConfigMulticastToReleaseList_r17 []SPS_ConfigIndex_r16

func (ie *SPS_ConfigMulticastToReleaseList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sPSConfigMulticastToReleaseListR17SizeConstraints)
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

func (ie *SPS_ConfigMulticastToReleaseList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sPSConfigMulticastToReleaseListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SPS_ConfigMulticastToReleaseList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
