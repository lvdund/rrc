// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ReportConfigToRemoveList (line 9129).
// ReportConfigToRemoveList ::=            SEQUENCE (SIZE (1..maxReportConfigId)) OF ReportConfigId

var reportConfigToRemoveListSizeConstraints = per.SizeRange(1, common.MaxReportConfigId)

type ReportConfigToRemoveList []ReportConfigId

func (ie *ReportConfigToRemoveList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(reportConfigToRemoveListSizeConstraints)
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

func (ie *ReportConfigToRemoveList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(reportConfigToRemoveListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ReportConfigToRemoveList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
