// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UL-PDCP-DelayValueResultList-r16 (line 9872).
// UL-PDCP-DelayValueResultList-r16 ::= SEQUENCE (SIZE (1..maxDRB)) OF UL-PDCP-DelayValueResult-r16

var uLPDCPDelayValueResultListR16SizeConstraints = per.SizeRange(1, common.MaxDRB)

type UL_PDCP_DelayValueResultList_r16 []UL_PDCP_DelayValueResult_r16

func (ie *UL_PDCP_DelayValueResultList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uLPDCPDelayValueResultListR16SizeConstraints)
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

func (ie *UL_PDCP_DelayValueResultList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uLPDCPDelayValueResultListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UL_PDCP_DelayValueResultList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
