// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NZP-CSI-RS-Resource (line 10819).

var nZPCSIRSResourceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nzp-CSI-RS-ResourceId"},
		{Name: "resourceMapping"},
		{Name: "powerControlOffset"},
		{Name: "powerControlOffsetSS", Optional: true},
		{Name: "scramblingID"},
		{Name: "periodicityAndOffset", Optional: true},
		{Name: "qcl-InfoPeriodicCSI-RS", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var nZPCSIRSResourcePowerControlOffsetConstraints = per.Constrained(-8, 15)

const (
	NZP_CSI_RS_Resource_PowerControlOffsetSS_Db_3 = 0
	NZP_CSI_RS_Resource_PowerControlOffsetSS_Db0  = 1
	NZP_CSI_RS_Resource_PowerControlOffsetSS_Db3  = 2
	NZP_CSI_RS_Resource_PowerControlOffsetSS_Db6  = 3
)

var nZPCSIRSResourcePowerControlOffsetSSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NZP_CSI_RS_Resource_PowerControlOffsetSS_Db_3, NZP_CSI_RS_Resource_PowerControlOffsetSS_Db0, NZP_CSI_RS_Resource_PowerControlOffsetSS_Db3, NZP_CSI_RS_Resource_PowerControlOffsetSS_Db6},
}

const (
	NZP_CSI_RS_Resource_Ext_CyclicPrefix_r18_Extended = 0
)

var nZPCSIRSResourceExtCyclicPrefixR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NZP_CSI_RS_Resource_Ext_CyclicPrefix_r18_Extended},
}

const (
	NZP_CSI_RS_Resource_Ext_AdditionalOneSlotOffset_r19_Enabled = 0
)

var nZPCSIRSResourceExtAdditionalOneSlotOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NZP_CSI_RS_Resource_Ext_AdditionalOneSlotOffset_r19_Enabled},
}

var nZPCSIRSResourceAdditionalSlotOffsetR19Constraints = per.Constrained(0, 7)

type NZP_CSI_RS_Resource struct {
	Nzp_CSI_RS_ResourceId       NZP_CSI_RS_ResourceId
	ResourceMapping             CSI_RS_ResourceMapping
	PowerControlOffset          int64
	PowerControlOffsetSS        *int64
	ScramblingID                ScramblingId
	PeriodicityAndOffset        *CSI_ResourcePeriodicityAndOffset
	Qcl_InfoPeriodicCSI_RS      *TCI_StateId
	SubcarrierSpacing_r18       *SubcarrierSpacing
	AbsoluteFrequencyPointA_r18 *ARFCN_ValueNR
	CyclicPrefix_r18            *int64
	AdditionalOneSlotOffset_r19 *int64
	AdditionalSlotOffset_r19    *int64
}

func (ie *NZP_CSI_RS_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nZPCSIRSResourceConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.SubcarrierSpacing_r18 != nil || ie.AbsoluteFrequencyPointA_r18 != nil || ie.CyclicPrefix_r18 != nil
	hasExtGroup1 := ie.AdditionalOneSlotOffset_r19 != nil || ie.AdditionalSlotOffset_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PowerControlOffsetSS != nil, ie.PeriodicityAndOffset != nil, ie.Qcl_InfoPeriodicCSI_RS != nil}); err != nil {
		return err
	}

	// 3. nzp-CSI-RS-ResourceId: ref
	{
		if err := ie.Nzp_CSI_RS_ResourceId.Encode(e); err != nil {
			return err
		}
	}

	// 4. resourceMapping: ref
	{
		if err := ie.ResourceMapping.Encode(e); err != nil {
			return err
		}
	}

	// 5. powerControlOffset: integer
	{
		if err := e.EncodeInteger(ie.PowerControlOffset, nZPCSIRSResourcePowerControlOffsetConstraints); err != nil {
			return err
		}
	}

	// 6. powerControlOffsetSS: enumerated
	{
		if ie.PowerControlOffsetSS != nil {
			if err := e.EncodeEnumerated(*ie.PowerControlOffsetSS, nZPCSIRSResourcePowerControlOffsetSSConstraints); err != nil {
				return err
			}
		}
	}

	// 7. scramblingID: ref
	{
		if err := ie.ScramblingID.Encode(e); err != nil {
			return err
		}
	}

	// 8. periodicityAndOffset: ref
	{
		if ie.PeriodicityAndOffset != nil {
			if err := ie.PeriodicityAndOffset.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. qcl-InfoPeriodicCSI-RS: ref
	{
		if ie.Qcl_InfoPeriodicCSI_RS != nil {
			if err := ie.Qcl_InfoPeriodicCSI_RS.Encode(e); err != nil {
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
					{Name: "subcarrierSpacing-r18", Optional: true},
					{Name: "absoluteFrequencyPointA-r18", Optional: true},
					{Name: "cyclicPrefix-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SubcarrierSpacing_r18 != nil, ie.AbsoluteFrequencyPointA_r18 != nil, ie.CyclicPrefix_r18 != nil}); err != nil {
				return err
			}

			if ie.SubcarrierSpacing_r18 != nil {
				if err := ie.SubcarrierSpacing_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.AbsoluteFrequencyPointA_r18 != nil {
				if err := ie.AbsoluteFrequencyPointA_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CyclicPrefix_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.CyclicPrefix_r18, nZPCSIRSResourceExtCyclicPrefixR18Constraints); err != nil {
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
					{Name: "additionalOneSlotOffset-r19", Optional: true},
					{Name: "additionalSlotOffset-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AdditionalOneSlotOffset_r19 != nil, ie.AdditionalSlotOffset_r19 != nil}); err != nil {
				return err
			}

			if ie.AdditionalOneSlotOffset_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.AdditionalOneSlotOffset_r19, nZPCSIRSResourceExtAdditionalOneSlotOffsetR19Constraints); err != nil {
					return err
				}
			}

			if ie.AdditionalSlotOffset_r19 != nil {
				if err := ex.EncodeInteger(*ie.AdditionalSlotOffset_r19, nZPCSIRSResourceAdditionalSlotOffsetR19Constraints); err != nil {
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

func (ie *NZP_CSI_RS_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nZPCSIRSResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. nzp-CSI-RS-ResourceId: ref
	{
		if err := ie.Nzp_CSI_RS_ResourceId.Decode(d); err != nil {
			return err
		}
	}

	// 4. resourceMapping: ref
	{
		if err := ie.ResourceMapping.Decode(d); err != nil {
			return err
		}
	}

	// 5. powerControlOffset: integer
	{
		v2, err := d.DecodeInteger(nZPCSIRSResourcePowerControlOffsetConstraints)
		if err != nil {
			return err
		}
		ie.PowerControlOffset = v2
	}

	// 6. powerControlOffsetSS: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(nZPCSIRSResourcePowerControlOffsetSSConstraints)
			if err != nil {
				return err
			}
			ie.PowerControlOffsetSS = &idx
		}
	}

	// 7. scramblingID: ref
	{
		if err := ie.ScramblingID.Decode(d); err != nil {
			return err
		}
	}

	// 8. periodicityAndOffset: ref
	{
		if seq.IsComponentPresent(5) {
			ie.PeriodicityAndOffset = new(CSI_ResourcePeriodicityAndOffset)
			if err := ie.PeriodicityAndOffset.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. qcl-InfoPeriodicCSI-RS: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Qcl_InfoPeriodicCSI_RS = new(TCI_StateId)
			if err := ie.Qcl_InfoPeriodicCSI_RS.Decode(d); err != nil {
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
				{Name: "subcarrierSpacing-r18", Optional: true},
				{Name: "absoluteFrequencyPointA-r18", Optional: true},
				{Name: "cyclicPrefix-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SubcarrierSpacing_r18 = new(SubcarrierSpacing)
			if err := ie.SubcarrierSpacing_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.AbsoluteFrequencyPointA_r18 = new(ARFCN_ValueNR)
			if err := ie.AbsoluteFrequencyPointA_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(nZPCSIRSResourceExtCyclicPrefixR18Constraints)
			if err != nil {
				return err
			}
			ie.CyclicPrefix_r18 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "additionalOneSlotOffset-r19", Optional: true},
				{Name: "additionalSlotOffset-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(nZPCSIRSResourceExtAdditionalOneSlotOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.AdditionalOneSlotOffset_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(nZPCSIRSResourceAdditionalSlotOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.AdditionalSlotOffset_r19 = &v
		}
	}

	return nil
}
