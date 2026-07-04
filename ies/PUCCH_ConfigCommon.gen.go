// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-ConfigCommon (line 12188).

var pUCCHConfigCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-ResourceCommon", Optional: true},
		{Name: "pucch-GroupHopping"},
		{Name: "hoppingId", Optional: true},
		{Name: "p0-nominal", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var pUCCHConfigCommonPucchResourceCommonConstraints = per.Constrained(0, 15)

const (
	PUCCH_ConfigCommon_Pucch_GroupHopping_Neither = 0
	PUCCH_ConfigCommon_Pucch_GroupHopping_Enable  = 1
	PUCCH_ConfigCommon_Pucch_GroupHopping_Disable = 2
)

var pUCCHConfigCommonPucchGroupHoppingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_ConfigCommon_Pucch_GroupHopping_Neither, PUCCH_ConfigCommon_Pucch_GroupHopping_Enable, PUCCH_ConfigCommon_Pucch_GroupHopping_Disable},
}

var pUCCHConfigCommonHoppingIdConstraints = per.Constrained(0, 1023)

var pUCCHConfigCommonP0NominalConstraints = per.Constrained(-202, 24)

var pUCCHConfigCommonNrofPRBsConstraints = per.Constrained(1, 16)

const (
	PUCCH_ConfigCommon_Ext_Intra_SlotFH_r17_FromLowerEdge = 0
	PUCCH_ConfigCommon_Ext_Intra_SlotFH_r17_FromUpperEdge = 1
)

var pUCCHConfigCommonExtIntraSlotFHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_ConfigCommon_Ext_Intra_SlotFH_r17_FromLowerEdge, PUCCH_ConfigCommon_Ext_Intra_SlotFH_r17_FromUpperEdge},
}

var pUCCHConfigCommonPucchResourceCommonRedCapR17Constraints = per.Constrained(0, 15)

const (
	PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N2  = 0
	PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N3  = 1
	PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N4  = 2
	PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N6  = 3
	PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N8  = 4
	PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N9  = 5
	PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N10 = 6
	PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N12 = 7
)

var pUCCHConfigCommonExtAdditionalPRBOffsetR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N2, PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N3, PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N4, PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N6, PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N8, PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N9, PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N10, PUCCH_ConfigCommon_Ext_AdditionalPRBOffset_r17_N12},
}

var pUCCHConfigCommonP0NominalSBFDR19Constraints = per.Constrained(-202, 24)

type PUCCH_ConfigCommon struct {
	Pucch_ResourceCommon           *int64
	Pucch_GroupHopping             int64
	HoppingId                      *int64
	P0_Nominal                     *int64
	NrofPRBs                       *int64
	Intra_SlotFH_r17               *int64
	Pucch_ResourceCommonRedCap_r17 *int64
	AdditionalPRBOffset_r17        *int64
	P0_Nominal_SBFD_r19            *int64
}

func (ie *PUCCH_ConfigCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHConfigCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.NrofPRBs != nil || ie.Intra_SlotFH_r17 != nil || ie.Pucch_ResourceCommonRedCap_r17 != nil || ie.AdditionalPRBOffset_r17 != nil
	hasExtGroup1 := ie.P0_Nominal_SBFD_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pucch_ResourceCommon != nil, ie.HoppingId != nil, ie.P0_Nominal != nil}); err != nil {
		return err
	}

	// 3. pucch-ResourceCommon: integer
	{
		if ie.Pucch_ResourceCommon != nil {
			if err := e.EncodeInteger(*ie.Pucch_ResourceCommon, pUCCHConfigCommonPucchResourceCommonConstraints); err != nil {
				return err
			}
		}
	}

	// 4. pucch-GroupHopping: enumerated
	{
		if err := e.EncodeEnumerated(ie.Pucch_GroupHopping, pUCCHConfigCommonPucchGroupHoppingConstraints); err != nil {
			return err
		}
	}

	// 5. hoppingId: integer
	{
		if ie.HoppingId != nil {
			if err := e.EncodeInteger(*ie.HoppingId, pUCCHConfigCommonHoppingIdConstraints); err != nil {
				return err
			}
		}
	}

	// 6. p0-nominal: integer
	{
		if ie.P0_Nominal != nil {
			if err := e.EncodeInteger(*ie.P0_Nominal, pUCCHConfigCommonP0NominalConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "nrofPRBs", Optional: true},
					{Name: "intra-SlotFH-r17", Optional: true},
					{Name: "pucch-ResourceCommonRedCap-r17", Optional: true},
					{Name: "additionalPRBOffset-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.NrofPRBs != nil, ie.Intra_SlotFH_r17 != nil, ie.Pucch_ResourceCommonRedCap_r17 != nil, ie.AdditionalPRBOffset_r17 != nil}); err != nil {
				return err
			}

			if ie.NrofPRBs != nil {
				if err := ex.EncodeInteger(*ie.NrofPRBs, pUCCHConfigCommonNrofPRBsConstraints); err != nil {
					return err
				}
			}

			if ie.Intra_SlotFH_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Intra_SlotFH_r17, pUCCHConfigCommonExtIntraSlotFHR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pucch_ResourceCommonRedCap_r17 != nil {
				if err := ex.EncodeInteger(*ie.Pucch_ResourceCommonRedCap_r17, pUCCHConfigCommonPucchResourceCommonRedCapR17Constraints); err != nil {
					return err
				}
			}

			if ie.AdditionalPRBOffset_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.AdditionalPRBOffset_r17, pUCCHConfigCommonExtAdditionalPRBOffsetR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "p0-nominal-SBFD-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.P0_Nominal_SBFD_r19 != nil}); err != nil {
				return err
			}

			if ie.P0_Nominal_SBFD_r19 != nil {
				if err := ex.EncodeInteger(*ie.P0_Nominal_SBFD_r19, pUCCHConfigCommonP0NominalSBFDR19Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_ConfigCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pucch-ResourceCommon: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pUCCHConfigCommonPucchResourceCommonConstraints)
			if err != nil {
				return err
			}
			ie.Pucch_ResourceCommon = &v
		}
	}

	// 4. pucch-GroupHopping: enumerated
	{
		v1, err := d.DecodeEnumerated(pUCCHConfigCommonPucchGroupHoppingConstraints)
		if err != nil {
			return err
		}
		ie.Pucch_GroupHopping = v1
	}

	// 5. hoppingId: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(pUCCHConfigCommonHoppingIdConstraints)
			if err != nil {
				return err
			}
			ie.HoppingId = &v
		}
	}

	// 6. p0-nominal: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(pUCCHConfigCommonP0NominalConstraints)
			if err != nil {
				return err
			}
			ie.P0_Nominal = &v
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "nrofPRBs", Optional: true},
				{Name: "intra-SlotFH-r17", Optional: true},
				{Name: "pucch-ResourceCommonRedCap-r17", Optional: true},
				{Name: "additionalPRBOffset-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pUCCHConfigCommonNrofPRBsConstraints)
			if err != nil {
				return err
			}
			ie.NrofPRBs = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pUCCHConfigCommonExtIntraSlotFHR17Constraints)
			if err != nil {
				return err
			}
			ie.Intra_SlotFH_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(pUCCHConfigCommonPucchResourceCommonRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_ResourceCommonRedCap_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(pUCCHConfigCommonExtAdditionalPRBOffsetR17Constraints)
			if err != nil {
				return err
			}
			ie.AdditionalPRBOffset_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "p0-nominal-SBFD-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pUCCHConfigCommonP0NominalSBFDR19Constraints)
			if err != nil {
				return err
			}
			ie.P0_Nominal_SBFD_r19 = &v
		}
	}

	return nil
}
