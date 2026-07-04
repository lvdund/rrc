// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UL-PDCP-ExcessDelayResultList-r17 (line 9880).
// UL-PDCP-ExcessDelayResultList-r17 ::= SEQUENCE (SIZE (1..maxDRB)) OF UL-PDCP-ExcessDelayResult-r17

var uLPDCPExcessDelayResultListR17SizeConstraints = per.SizeRange(1, common.MaxDRB)

type UL_PDCP_ExcessDelayResultList_r17 []UL_PDCP_ExcessDelayResult_r17

func (ie *UL_PDCP_ExcessDelayResultList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uLPDCPExcessDelayResultListR17SizeConstraints)
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

func (ie *UL_PDCP_ExcessDelayResultList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uLPDCPExcessDelayResultListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UL_PDCP_ExcessDelayResultList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
