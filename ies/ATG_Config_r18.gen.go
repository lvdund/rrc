// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ATG-Config-r18 (line 5039).

var aTGConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "atg-gNB-Location-r18", Optional: true},
		{Name: "height-gNB-r18", Optional: true},
		{Name: "cellSpecificKoffset-r18", Optional: true},
		{Name: "ta-ReportATG-r18", Optional: true},
	},
}

var aTGConfigR18HeightGNBR18Constraints = per.Constrained(-16384, 16383)

var aTGConfigR18CellSpecificKoffsetR18Constraints = per.Constrained(1, 3)

const (
	ATG_Config_r18_Ta_ReportATG_r18_Enabled = 0
)

var aTGConfigR18TaReportATGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ATG_Config_r18_Ta_ReportATG_r18_Enabled},
}

type ATG_Config_r18 struct {
	Atg_GNB_Location_r18    *ReferenceLocation_r17
	Height_GNB_r18          *int64
	CellSpecificKoffset_r18 *int64
	Ta_ReportATG_r18        *int64
}

func (ie *ATG_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(aTGConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Atg_GNB_Location_r18 != nil, ie.Height_GNB_r18 != nil, ie.CellSpecificKoffset_r18 != nil, ie.Ta_ReportATG_r18 != nil}); err != nil {
		return err
	}

	// 2. atg-gNB-Location-r18: ref
	{
		if ie.Atg_GNB_Location_r18 != nil {
			if err := ie.Atg_GNB_Location_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. height-gNB-r18: integer
	{
		if ie.Height_GNB_r18 != nil {
			if err := e.EncodeInteger(*ie.Height_GNB_r18, aTGConfigR18HeightGNBR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. cellSpecificKoffset-r18: integer
	{
		if ie.CellSpecificKoffset_r18 != nil {
			if err := e.EncodeInteger(*ie.CellSpecificKoffset_r18, aTGConfigR18CellSpecificKoffsetR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ta-ReportATG-r18: enumerated
	{
		if ie.Ta_ReportATG_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ta_ReportATG_r18, aTGConfigR18TaReportATGR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ATG_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(aTGConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. atg-gNB-Location-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Atg_GNB_Location_r18 = new(ReferenceLocation_r17)
			if err := ie.Atg_GNB_Location_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. height-gNB-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(aTGConfigR18HeightGNBR18Constraints)
			if err != nil {
				return err
			}
			ie.Height_GNB_r18 = &v
		}
	}

	// 4. cellSpecificKoffset-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(aTGConfigR18CellSpecificKoffsetR18Constraints)
			if err != nil {
				return err
			}
			ie.CellSpecificKoffset_r18 = &v
		}
	}

	// 5. ta-ReportATG-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(aTGConfigR18TaReportATGR18Constraints)
			if err != nil {
				return err
			}
			ie.Ta_ReportATG_r18 = &idx
		}
	}

	return nil
}
