// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Phy-ParametersFR2 (line 23233).

var phyParametersFR2Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dummy", Optional: true},
		{Name: "pdsch-RE-MappingFR2-PerSymbol", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	Phy_ParametersFR2_Dummy_Supported = 0
)

var phyParametersFR2DummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR2_Dummy_Supported},
}

const (
	Phy_ParametersFR2_Pdsch_RE_MappingFR2_PerSymbol_N6  = 0
	Phy_ParametersFR2_Pdsch_RE_MappingFR2_PerSymbol_N20 = 1
)

var phyParametersFR2PdschREMappingFR2PerSymbolConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR2_Pdsch_RE_MappingFR2_PerSymbol_N6, Phy_ParametersFR2_Pdsch_RE_MappingFR2_PerSymbol_N20},
}

const (
	Phy_ParametersFR2_Ext_PCell_FR2_Supported = 0
)

var phyParametersFR2ExtPCellFR2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR2_Ext_PCell_FR2_Supported},
}

const (
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N16  = 0
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N32  = 1
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N48  = 2
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N64  = 3
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N80  = 4
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N96  = 5
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N112 = 6
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N128 = 7
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N144 = 8
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N160 = 9
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N176 = 10
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N192 = 11
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N208 = 12
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N224 = 13
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N240 = 14
	Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N256 = 15
)

var phyParametersFR2ExtPdschREMappingFR2PerSlotConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N16, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N32, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N48, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N64, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N80, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N96, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N112, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N128, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N144, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N160, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N176, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N192, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N208, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N224, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N240, Phy_ParametersFR2_Ext_Pdsch_RE_MappingFR2_PerSlot_N256},
}

const (
	Phy_ParametersFR2_Ext_DefaultSpatialRelationPathlossRS_r16_Supported = 0
)

var phyParametersFR2ExtDefaultSpatialRelationPathlossRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR2_Ext_DefaultSpatialRelationPathlossRS_r16_Supported},
}

const (
	Phy_ParametersFR2_Ext_SpatialRelationUpdateAP_SRS_r16_Supported = 0
)

var phyParametersFR2ExtSpatialRelationUpdateAPSRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR2_Ext_SpatialRelationUpdateAP_SRS_r16_Supported},
}

const (
	Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N0  = 0
	Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N1  = 1
	Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N2  = 2
	Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N4  = 3
	Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N8  = 4
	Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N16 = 5
)

var phyParametersFR2ExtMaxNumberSRSPosSpatialRelationsAllServingCellsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N0, Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N1, Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N2, Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N4, Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N8, Phy_ParametersFR2_Ext_MaxNumberSRS_PosSpatialRelationsAllServingCells_r16_N16},
}

const (
	Phy_ParametersFR2_Ext_MultiRxPreferenceIndication_r18_Supported = 0
)

var phyParametersFR2ExtMultiRxPreferenceIndicationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR2_Ext_MultiRxPreferenceIndication_r18_Supported},
}

type Phy_ParametersFR2 struct {
	Dummy                                               *int64
	Pdsch_RE_MappingFR2_PerSymbol                       *int64
	PCell_FR2                                           *int64
	Pdsch_RE_MappingFR2_PerSlot                         *int64
	DefaultSpatialRelationPathlossRS_r16                *int64
	SpatialRelationUpdateAP_SRS_r16                     *int64
	MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 *int64
	MultiRxPreferenceIndication_r18                     *int64
}

func (ie *Phy_ParametersFR2) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersFR2Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.PCell_FR2 != nil || ie.Pdsch_RE_MappingFR2_PerSlot != nil
	hasExtGroup1 := ie.DefaultSpatialRelationPathlossRS_r16 != nil || ie.SpatialRelationUpdateAP_SRS_r16 != nil || ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil
	hasExtGroup2 := ie.MultiRxPreferenceIndication_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dummy != nil, ie.Pdsch_RE_MappingFR2_PerSymbol != nil}); err != nil {
		return err
	}

	// 3. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, phyParametersFR2DummyConstraints); err != nil {
				return err
			}
		}
	}

	// 4. pdsch-RE-MappingFR2-PerSymbol: enumerated
	{
		if ie.Pdsch_RE_MappingFR2_PerSymbol != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_RE_MappingFR2_PerSymbol, phyParametersFR2PdschREMappingFR2PerSymbolConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "pCell-FR2", Optional: true},
					{Name: "pdsch-RE-MappingFR2-PerSlot", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PCell_FR2 != nil, ie.Pdsch_RE_MappingFR2_PerSlot != nil}); err != nil {
				return err
			}

			if ie.PCell_FR2 != nil {
				if err := ex.EncodeEnumerated(*ie.PCell_FR2, phyParametersFR2ExtPCellFR2Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_RE_MappingFR2_PerSlot != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_RE_MappingFR2_PerSlot, phyParametersFR2ExtPdschREMappingFR2PerSlotConstraints); err != nil {
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
					{Name: "defaultSpatialRelationPathlossRS-r16", Optional: true},
					{Name: "spatialRelationUpdateAP-SRS-r16", Optional: true},
					{Name: "maxNumberSRS-PosSpatialRelationsAllServingCells-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DefaultSpatialRelationPathlossRS_r16 != nil, ie.SpatialRelationUpdateAP_SRS_r16 != nil, ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil}); err != nil {
				return err
			}

			if ie.DefaultSpatialRelationPathlossRS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.DefaultSpatialRelationPathlossRS_r16, phyParametersFR2ExtDefaultSpatialRelationPathlossRSR16Constraints); err != nil {
					return err
				}
			}

			if ie.SpatialRelationUpdateAP_SRS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SpatialRelationUpdateAP_SRS_r16, phyParametersFR2ExtSpatialRelationUpdateAPSRSR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16, phyParametersFR2ExtMaxNumberSRSPosSpatialRelationsAllServingCellsR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "multiRxPreferenceIndication-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MultiRxPreferenceIndication_r18 != nil}); err != nil {
				return err
			}

			if ie.MultiRxPreferenceIndication_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiRxPreferenceIndication_r18, phyParametersFR2ExtMultiRxPreferenceIndicationR18Constraints); err != nil {
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

func (ie *Phy_ParametersFR2) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersFR2Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dummy: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(phyParametersFR2DummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	// 4. pdsch-RE-MappingFR2-PerSymbol: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(phyParametersFR2PdschREMappingFR2PerSymbolConstraints)
			if err != nil {
				return err
			}
			ie.Pdsch_RE_MappingFR2_PerSymbol = &idx
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
				{Name: "pCell-FR2", Optional: true},
				{Name: "pdsch-RE-MappingFR2-PerSlot", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFR2ExtPCellFR2Constraints)
			if err != nil {
				return err
			}
			ie.PCell_FR2 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersFR2ExtPdschREMappingFR2PerSlotConstraints)
			if err != nil {
				return err
			}
			ie.Pdsch_RE_MappingFR2_PerSlot = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "defaultSpatialRelationPathlossRS-r16", Optional: true},
				{Name: "spatialRelationUpdateAP-SRS-r16", Optional: true},
				{Name: "maxNumberSRS-PosSpatialRelationsAllServingCells-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFR2ExtDefaultSpatialRelationPathlossRSR16Constraints)
			if err != nil {
				return err
			}
			ie.DefaultSpatialRelationPathlossRS_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersFR2ExtSpatialRelationUpdateAPSRSR16Constraints)
			if err != nil {
				return err
			}
			ie.SpatialRelationUpdateAP_SRS_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersFR2ExtMaxNumberSRSPosSpatialRelationsAllServingCellsR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "multiRxPreferenceIndication-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFR2ExtMultiRxPreferenceIndicationR18Constraints)
			if err != nil {
				return err
			}
			ie.MultiRxPreferenceIndication_r18 = &v
		}
	}

	return nil
}
