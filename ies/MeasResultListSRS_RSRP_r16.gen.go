// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultListSRS-RSRP-r16 (line 9858).
// MeasResultListSRS-RSRP-r16 ::=   SEQUENCE (SIZE (1.. maxCLI-Report-r16)) OF MeasResultSRS-RSRP-r16

var measResultListSRSRSRPR16SizeConstraints = per.SizeRange(1, common.MaxCLI_Report_r16)

type MeasResultListSRS_RSRP_r16 []MeasResultSRS_RSRP_r16

func (ie *MeasResultListSRS_RSRP_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultListSRSRSRPR16SizeConstraints)
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

func (ie *MeasResultListSRS_RSRP_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultListSRSRSRPR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultListSRS_RSRP_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
