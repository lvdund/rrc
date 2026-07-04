// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CrossCarrierSchedulingConfig (line 6817).

var crossCarrierSchedulingConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "schedulingCellInfo"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var crossCarrierSchedulingConfigSchedulingCellInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "own"},
		{Name: "other"},
	},
}

const (
	CrossCarrierSchedulingConfig_SchedulingCellInfo_Own   = 0
	CrossCarrierSchedulingConfig_SchedulingCellInfo_Other = 1
)

const (
	CrossCarrierSchedulingConfig_Ext_EnableDefaultBeamForCCS_r16_Enabled = 0
)

var crossCarrierSchedulingConfigExtEnableDefaultBeamForCCSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CrossCarrierSchedulingConfig_Ext_EnableDefaultBeamForCCS_r16_Enabled},
}

const (
	CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_OneSeventh      = 0
	CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_ThreeFourteenth = 1
	CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_TwoSeventh      = 2
	CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_ThreeSeventh    = 3
	CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_OneHalf         = 4
	CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_FourSeventh     = 5
	CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_FiveSeventh     = 6
	CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_Spare1          = 7
)

var crossCarrierSchedulingConfigExtCcsBlindDetectionSplitR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_OneSeventh, CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_ThreeFourteenth, CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_TwoSeventh, CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_ThreeSeventh, CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_OneHalf, CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_FourSeventh, CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_FiveSeventh, CrossCarrierSchedulingConfig_Ext_Ccs_BlindDetectionSplit_r17_Spare1},
}

type CrossCarrierSchedulingConfig struct {
	SchedulingCellInfo struct {
		Choice int
		Own    *struct{ Cif_Presence bool }
		Other  *struct {
			SchedulingCellId     ServCellIndex
			Cif_InSchedulingCell int64
		}
	}
	CarrierIndicatorSize_r16 *struct {
		CarrierIndicatorSizeDCI_1_2_r16 int64
		CarrierIndicatorSizeDCI_0_2_r16 int64
	}
	EnableDefaultBeamForCCS_r16 *int64
	Ccs_BlindDetectionSplit_r17 *int64
}

func (ie *CrossCarrierSchedulingConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(crossCarrierSchedulingConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.CarrierIndicatorSize_r16 != nil || ie.EnableDefaultBeamForCCS_r16 != nil
	hasExtGroup1 := ie.Ccs_BlindDetectionSplit_r17 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. schedulingCellInfo: choice
	{
		choiceEnc := e.NewChoiceEncoder(crossCarrierSchedulingConfigSchedulingCellInfoConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.SchedulingCellInfo.Choice), false, nil); err != nil {
			return err
		}
		switch ie.SchedulingCellInfo.Choice {
		case CrossCarrierSchedulingConfig_SchedulingCellInfo_Own:
			c := (*ie.SchedulingCellInfo.Own)
			if err := e.EncodeBoolean(c.Cif_Presence); err != nil {
				return err
			}
		case CrossCarrierSchedulingConfig_SchedulingCellInfo_Other:
			c := (*ie.SchedulingCellInfo.Other)
			if err := c.SchedulingCellId.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Cif_InSchedulingCell, per.Constrained(1, 7)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.SchedulingCellInfo.Choice), Constraint: "index out of range"}
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
					{Name: "carrierIndicatorSize-r16", Optional: true},
					{Name: "enableDefaultBeamForCCS-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CarrierIndicatorSize_r16 != nil, ie.EnableDefaultBeamForCCS_r16 != nil}); err != nil {
				return err
			}

			if ie.CarrierIndicatorSize_r16 != nil {
				c := ie.CarrierIndicatorSize_r16
				if err := ex.EncodeInteger(c.CarrierIndicatorSizeDCI_1_2_r16, per.Constrained(0, 3)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.CarrierIndicatorSizeDCI_0_2_r16, per.Constrained(0, 3)); err != nil {
					return err
				}
			}

			if ie.EnableDefaultBeamForCCS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableDefaultBeamForCCS_r16, crossCarrierSchedulingConfigExtEnableDefaultBeamForCCSR16Constraints); err != nil {
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
					{Name: "ccs-BlindDetectionSplit-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ccs_BlindDetectionSplit_r17 != nil}); err != nil {
				return err
			}

			if ie.Ccs_BlindDetectionSplit_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ccs_BlindDetectionSplit_r17, crossCarrierSchedulingConfigExtCcsBlindDetectionSplitR17Constraints); err != nil {
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

func (ie *CrossCarrierSchedulingConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(crossCarrierSchedulingConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. schedulingCellInfo: choice
	{
		choiceDec := d.NewChoiceDecoder(crossCarrierSchedulingConfigSchedulingCellInfoConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.SchedulingCellInfo.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CrossCarrierSchedulingConfig_SchedulingCellInfo_Own:
			ie.SchedulingCellInfo.Own = &struct{ Cif_Presence bool }{}
			c := (*ie.SchedulingCellInfo.Own)
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Cif_Presence = v
			}
		case CrossCarrierSchedulingConfig_SchedulingCellInfo_Other:
			ie.SchedulingCellInfo.Other = &struct {
				SchedulingCellId     ServCellIndex
				Cif_InSchedulingCell int64
			}{}
			c := (*ie.SchedulingCellInfo.Other)
			{
				if err := c.SchedulingCellId.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 7))
				if err != nil {
					return err
				}
				c.Cif_InSchedulingCell = v
			}
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
				{Name: "carrierIndicatorSize-r16", Optional: true},
				{Name: "enableDefaultBeamForCCS-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.CarrierIndicatorSize_r16 = &struct {
				CarrierIndicatorSizeDCI_1_2_r16 int64
				CarrierIndicatorSizeDCI_0_2_r16 int64
			}{}
			c := ie.CarrierIndicatorSize_r16
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				c.CarrierIndicatorSizeDCI_1_2_r16 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				c.CarrierIndicatorSizeDCI_0_2_r16 = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(crossCarrierSchedulingConfigExtEnableDefaultBeamForCCSR16Constraints)
			if err != nil {
				return err
			}
			ie.EnableDefaultBeamForCCS_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ccs-BlindDetectionSplit-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(crossCarrierSchedulingConfigExtCcsBlindDetectionSplitR17Constraints)
			if err != nil {
				return err
			}
			ie.Ccs_BlindDetectionSplit_r17 = &v
		}
	}

	return nil
}
