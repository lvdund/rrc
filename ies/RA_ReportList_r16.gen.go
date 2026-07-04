// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RA-ReportList-r16 (line 3059).
// RA-ReportList-r16 ::= SEQUENCE (SIZE (1..maxRAReport-r16)) OF RA-Report-r16

var rAReportListR16SizeConstraints = per.SizeRange(1, common.MaxRAReport_r16)

type RA_ReportList_r16 []RA_Report_r16

func (ie *RA_ReportList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(rAReportListR16SizeConstraints)
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

func (ie *RA_ReportList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(rAReportListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(RA_ReportList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
