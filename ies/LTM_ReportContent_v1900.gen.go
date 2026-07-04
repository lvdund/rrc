// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-ReportContent-v1900 (line 8889).

var lTMReportContentV1900Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportQuantity-r19", Optional: false},
	},
}

const (
	LTM_ReportContent_v1900_ReportQuantity_r19_Cri_RSRP       = 0
	LTM_ReportContent_v1900_ReportQuantity_r19_Ssb_Index_RSRP = 1
	LTM_ReportContent_v1900_ReportQuantity_r19_Dummy          = 2
	LTM_ReportContent_v1900_ReportQuantity_r19_Value1         = 3
)

var lTMReportContentV1900ReportQuantityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_ReportContent_v1900_ReportQuantity_r19_Cri_RSRP, LTM_ReportContent_v1900_ReportQuantity_r19_Ssb_Index_RSRP, LTM_ReportContent_v1900_ReportQuantity_r19_Dummy, LTM_ReportContent_v1900_ReportQuantity_r19_Value1},
}

type LTM_ReportContent_v1900 struct {
	ReportQuantity_r19 *int64
}

func (ie *LTM_ReportContent_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMReportContentV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReportQuantity_r19 != nil && *ie.ReportQuantity_r19 != LTM_ReportContent_v1900_ReportQuantity_r19_Ssb_Index_RSRP}); err != nil {
		return err
	}

	// 3. reportQuantity-r19: enumerated
	{
		if ie.ReportQuantity_r19 != nil && *ie.ReportQuantity_r19 != LTM_ReportContent_v1900_ReportQuantity_r19_Ssb_Index_RSRP {
			if err := e.EncodeEnumerated(*ie.ReportQuantity_r19, lTMReportContentV1900ReportQuantityR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_ReportContent_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMReportContentV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. reportQuantity-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(lTMReportContentV1900ReportQuantityR19Constraints)
			if err != nil {
				return err
			}
			ie.ReportQuantity_r19 = &idx
		} else {
			v := int64(LTM_ReportContent_v1900_ReportQuantity_r19_Ssb_Index_RSRP)
			ie.ReportQuantity_r19 = &v
		}
	}

	return nil
}
