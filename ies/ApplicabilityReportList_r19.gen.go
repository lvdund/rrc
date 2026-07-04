// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ApplicabilityReportList-r19 (line 4998).
// ApplicabilityReportList-r19 ::=      SEQUENCE (SIZE (1..maxNrofServingCells)) OF ApplicabilityReport-r19

var applicabilityReportListR19SizeConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type ApplicabilityReportList_r19 []ApplicabilityReport_r19

func (ie *ApplicabilityReportList_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(applicabilityReportListR19SizeConstraints)
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

func (ie *ApplicabilityReportList_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(applicabilityReportListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ApplicabilityReportList_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
