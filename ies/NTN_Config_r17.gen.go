// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NTN-Config-r17 (line 10780).

var nTNConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "epochTime-r17", Optional: true},
		{Name: "ntn-UlSyncValidityDuration-r17", Optional: true},
		{Name: "cellSpecificKoffset-r17", Optional: true},
		{Name: "kmac-r17", Optional: true},
		{Name: "ta-Info-r17", Optional: true},
		{Name: "ntn-PolarizationDL-r17", Optional: true},
		{Name: "ntn-PolarizationUL-r17", Optional: true},
		{Name: "ephemerisInfo-r17", Optional: true},
		{Name: "ta-Report-r17", Optional: true},
	},
}

const (
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S5   = 0
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S10  = 1
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S15  = 2
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S20  = 3
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S25  = 4
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S30  = 5
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S35  = 6
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S40  = 7
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S45  = 8
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S50  = 9
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S55  = 10
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S60  = 11
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S120 = 12
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S180 = 13
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S240 = 14
	NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S900 = 15
)

var nTNConfigR17NtnUlSyncValidityDurationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S5, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S10, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S15, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S20, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S25, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S30, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S35, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S40, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S45, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S50, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S55, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S60, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S120, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S180, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S240, NTN_Config_r17_Ntn_UlSyncValidityDuration_r17_S900},
}

var nTNConfigR17CellSpecificKoffsetR17Constraints = per.Constrained(1, 1023)

var nTNConfigR17KmacR17Constraints = per.Constrained(1, 512)

const (
	NTN_Config_r17_Ntn_PolarizationDL_r17_Rhcp   = 0
	NTN_Config_r17_Ntn_PolarizationDL_r17_Lhcp   = 1
	NTN_Config_r17_Ntn_PolarizationDL_r17_Linear = 2
)

var nTNConfigR17NtnPolarizationDLR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NTN_Config_r17_Ntn_PolarizationDL_r17_Rhcp, NTN_Config_r17_Ntn_PolarizationDL_r17_Lhcp, NTN_Config_r17_Ntn_PolarizationDL_r17_Linear},
}

const (
	NTN_Config_r17_Ntn_PolarizationUL_r17_Rhcp   = 0
	NTN_Config_r17_Ntn_PolarizationUL_r17_Lhcp   = 1
	NTN_Config_r17_Ntn_PolarizationUL_r17_Linear = 2
)

var nTNConfigR17NtnPolarizationULR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NTN_Config_r17_Ntn_PolarizationUL_r17_Rhcp, NTN_Config_r17_Ntn_PolarizationUL_r17_Lhcp, NTN_Config_r17_Ntn_PolarizationUL_r17_Linear},
}

const (
	NTN_Config_r17_Ta_Report_r17_Enabled = 0
)

var nTNConfigR17TaReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NTN_Config_r17_Ta_Report_r17_Enabled},
}

type NTN_Config_r17 struct {
	EpochTime_r17                  *EpochTime_r17
	Ntn_UlSyncValidityDuration_r17 *int64
	CellSpecificKoffset_r17        *int64
	Kmac_r17                       *int64
	Ta_Info_r17                    *TA_Info_r17
	Ntn_PolarizationDL_r17         *int64
	Ntn_PolarizationUL_r17         *int64
	EphemerisInfo_r17              *EphemerisInfo_r17
	Ta_Report_r17                  *int64
}

func (ie *NTN_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nTNConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EpochTime_r17 != nil, ie.Ntn_UlSyncValidityDuration_r17 != nil, ie.CellSpecificKoffset_r17 != nil, ie.Kmac_r17 != nil, ie.Ta_Info_r17 != nil, ie.Ntn_PolarizationDL_r17 != nil, ie.Ntn_PolarizationUL_r17 != nil, ie.EphemerisInfo_r17 != nil, ie.Ta_Report_r17 != nil}); err != nil {
		return err
	}

	// 3. epochTime-r17: ref
	{
		if ie.EpochTime_r17 != nil {
			if err := ie.EpochTime_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ntn-UlSyncValidityDuration-r17: enumerated
	{
		if ie.Ntn_UlSyncValidityDuration_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_UlSyncValidityDuration_r17, nTNConfigR17NtnUlSyncValidityDurationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. cellSpecificKoffset-r17: integer
	{
		if ie.CellSpecificKoffset_r17 != nil {
			if err := e.EncodeInteger(*ie.CellSpecificKoffset_r17, nTNConfigR17CellSpecificKoffsetR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. kmac-r17: integer
	{
		if ie.Kmac_r17 != nil {
			if err := e.EncodeInteger(*ie.Kmac_r17, nTNConfigR17KmacR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ta-Info-r17: ref
	{
		if ie.Ta_Info_r17 != nil {
			if err := ie.Ta_Info_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. ntn-PolarizationDL-r17: enumerated
	{
		if ie.Ntn_PolarizationDL_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_PolarizationDL_r17, nTNConfigR17NtnPolarizationDLR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. ntn-PolarizationUL-r17: enumerated
	{
		if ie.Ntn_PolarizationUL_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_PolarizationUL_r17, nTNConfigR17NtnPolarizationULR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. ephemerisInfo-r17: ref
	{
		if ie.EphemerisInfo_r17 != nil {
			if err := ie.EphemerisInfo_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. ta-Report-r17: enumerated
	{
		if ie.Ta_Report_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ta_Report_r17, nTNConfigR17TaReportR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NTN_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nTNConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. epochTime-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.EpochTime_r17 = new(EpochTime_r17)
			if err := ie.EpochTime_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ntn-UlSyncValidityDuration-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(nTNConfigR17NtnUlSyncValidityDurationR17Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_UlSyncValidityDuration_r17 = &idx
		}
	}

	// 5. cellSpecificKoffset-r17: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(nTNConfigR17CellSpecificKoffsetR17Constraints)
			if err != nil {
				return err
			}
			ie.CellSpecificKoffset_r17 = &v
		}
	}

	// 6. kmac-r17: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(nTNConfigR17KmacR17Constraints)
			if err != nil {
				return err
			}
			ie.Kmac_r17 = &v
		}
	}

	// 7. ta-Info-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Ta_Info_r17 = new(TA_Info_r17)
			if err := ie.Ta_Info_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. ntn-PolarizationDL-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(nTNConfigR17NtnPolarizationDLR17Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_PolarizationDL_r17 = &idx
		}
	}

	// 9. ntn-PolarizationUL-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(nTNConfigR17NtnPolarizationULR17Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_PolarizationUL_r17 = &idx
		}
	}

	// 10. ephemerisInfo-r17: ref
	{
		if seq.IsComponentPresent(7) {
			ie.EphemerisInfo_r17 = new(EphemerisInfo_r17)
			if err := ie.EphemerisInfo_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. ta-Report-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(nTNConfigR17TaReportR17Constraints)
			if err != nil {
				return err
			}
			ie.Ta_Report_r17 = &idx
		}
	}

	return nil
}
