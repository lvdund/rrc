// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-ConfigCommon (line 11469).

var pDSCHConfigCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdsch-TimeDomainAllocationList", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	PDSCH_ConfigCommon_Ext_Pdsch_AggregationFactor_r19_N2 = 0
	PDSCH_ConfigCommon_Ext_Pdsch_AggregationFactor_r19_N4 = 1
)

var pDSCHConfigCommonExtPdschAggregationFactorR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ConfigCommon_Ext_Pdsch_AggregationFactor_r19_N2, PDSCH_ConfigCommon_Ext_Pdsch_AggregationFactor_r19_N4},
}

type PDSCH_ConfigCommon struct {
	Pdsch_TimeDomainAllocationList *PDSCH_TimeDomainResourceAllocationList
	Pdsch_AggregationFactor_r19    *int64
}

func (ie *PDSCH_ConfigCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHConfigCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Pdsch_AggregationFactor_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdsch_TimeDomainAllocationList != nil}); err != nil {
		return err
	}

	// 3. pdsch-TimeDomainAllocationList: ref
	{
		if ie.Pdsch_TimeDomainAllocationList != nil {
			if err := ie.Pdsch_TimeDomainAllocationList.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "pdsch-AggregationFactor-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdsch_AggregationFactor_r19 != nil}); err != nil {
				return err
			}

			if ie.Pdsch_AggregationFactor_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_AggregationFactor_r19, pDSCHConfigCommonExtPdschAggregationFactorR19Constraints); err != nil {
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

func (ie *PDSCH_ConfigCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pdsch-TimeDomainAllocationList: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Pdsch_TimeDomainAllocationList = new(PDSCH_TimeDomainResourceAllocationList)
			if err := ie.Pdsch_TimeDomainAllocationList.Decode(d); err != nil {
				return err
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
				{Name: "pdsch-AggregationFactor-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDSCHConfigCommonExtPdschAggregationFactorR19Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_AggregationFactor_r19 = &v
		}
	}

	return nil
}
