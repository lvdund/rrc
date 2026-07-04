// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ReportConfigToAddModList (line 13972).
// ReportConfigToAddModList ::=        SEQUENCE (SIZE (1..maxReportConfigId)) OF ReportConfigToAddMod

var reportConfigToAddModListSizeConstraints = per.SizeRange(1, common.MaxReportConfigId)

type ReportConfigToAddModList []ReportConfigToAddMod

func (ie *ReportConfigToAddModList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(reportConfigToAddModListSizeConstraints)
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

func (ie *ReportConfigToAddModList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(reportConfigToAddModListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ReportConfigToAddModList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
