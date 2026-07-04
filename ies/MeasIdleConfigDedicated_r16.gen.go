// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasIdleConfigDedicated-r16 (line 9250).

var measIdleConfigDedicatedR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measIdleCarrierListNR-r16", Optional: true},
		{Name: "measIdleCarrierListEUTRA-r16", Optional: true},
		{Name: "measIdleDuration-r16"},
		{Name: "validityAreaList-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var measIdleConfigDedicatedR16MeasIdleCarrierListNRR16Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

var measIdleConfigDedicatedR16MeasIdleCarrierListEUTRAR16Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

const (
	MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec10  = 0
	MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec30  = 1
	MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec60  = 2
	MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec120 = 3
	MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec180 = 4
	MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec240 = 5
	MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec300 = 6
	MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Spare  = 7
)

var measIdleConfigDedicatedR16MeasIdleDurationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec10, MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec30, MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec60, MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec120, MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec180, MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec240, MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Sec300, MeasIdleConfigDedicated_r16_MeasIdleDuration_r16_Spare},
}

var measIdleConfigDedicatedR16ExtMeasReselectionCarrierListNRR18Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

type MeasIdleConfigDedicated_r16 struct {
	MeasIdleCarrierListNR_r16           []MeasIdleCarrierNR_r16
	MeasIdleCarrierListEUTRA_r16        []MeasIdleCarrierEUTRA_r16
	MeasIdleDuration_r16                int64
	ValidityAreaList_r16                *ValidityAreaList_r16
	MeasReselectionCarrierListNR_r18    []MeasReselectionCarrierNR_r18
	MeasIdleValidityDuration_r18        *MeasurementValidityDuration_r18
	MeasReselectionValidityDuration_r18 *MeasurementValidityDuration_r18
}

func (ie *MeasIdleConfigDedicated_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measIdleConfigDedicatedR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MeasReselectionCarrierListNR_r18 != nil || ie.MeasIdleValidityDuration_r18 != nil || ie.MeasReselectionValidityDuration_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasIdleCarrierListNR_r16 != nil, ie.MeasIdleCarrierListEUTRA_r16 != nil, ie.ValidityAreaList_r16 != nil}); err != nil {
		return err
	}

	// 3. measIdleCarrierListNR-r16: sequence-of
	{
		if ie.MeasIdleCarrierListNR_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(measIdleConfigDedicatedR16MeasIdleCarrierListNRR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.MeasIdleCarrierListNR_r16))); err != nil {
				return err
			}
			for i := range ie.MeasIdleCarrierListNR_r16 {
				if err := ie.MeasIdleCarrierListNR_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. measIdleCarrierListEUTRA-r16: sequence-of
	{
		if ie.MeasIdleCarrierListEUTRA_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(measIdleConfigDedicatedR16MeasIdleCarrierListEUTRAR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.MeasIdleCarrierListEUTRA_r16))); err != nil {
				return err
			}
			for i := range ie.MeasIdleCarrierListEUTRA_r16 {
				if err := ie.MeasIdleCarrierListEUTRA_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. measIdleDuration-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MeasIdleDuration_r16, measIdleConfigDedicatedR16MeasIdleDurationR16Constraints); err != nil {
			return err
		}
	}

	// 6. validityAreaList-r16: ref
	{
		if ie.ValidityAreaList_r16 != nil {
			if err := ie.ValidityAreaList_r16.Encode(e); err != nil {
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
					{Name: "measReselectionCarrierListNR-r18", Optional: true},
					{Name: "measIdleValidityDuration-r18", Optional: true},
					{Name: "measReselectionValidityDuration-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasReselectionCarrierListNR_r18 != nil, ie.MeasIdleValidityDuration_r18 != nil, ie.MeasReselectionValidityDuration_r18 != nil}); err != nil {
				return err
			}

			if ie.MeasReselectionCarrierListNR_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(measIdleConfigDedicatedR16ExtMeasReselectionCarrierListNRR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.MeasReselectionCarrierListNR_r18))); err != nil {
					return err
				}
				for i := range ie.MeasReselectionCarrierListNR_r18 {
					if err := ie.MeasReselectionCarrierListNR_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MeasIdleValidityDuration_r18 != nil {
				if err := ie.MeasIdleValidityDuration_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MeasReselectionValidityDuration_r18 != nil {
				if err := ie.MeasReselectionValidityDuration_r18.Encode(ex); err != nil {
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

func (ie *MeasIdleConfigDedicated_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measIdleConfigDedicatedR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measIdleCarrierListNR-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(measIdleConfigDedicatedR16MeasIdleCarrierListNRR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MeasIdleCarrierListNR_r16 = make([]MeasIdleCarrierNR_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MeasIdleCarrierListNR_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. measIdleCarrierListEUTRA-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(measIdleConfigDedicatedR16MeasIdleCarrierListEUTRAR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MeasIdleCarrierListEUTRA_r16 = make([]MeasIdleCarrierEUTRA_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MeasIdleCarrierListEUTRA_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. measIdleDuration-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(measIdleConfigDedicatedR16MeasIdleDurationR16Constraints)
		if err != nil {
			return err
		}
		ie.MeasIdleDuration_r16 = v2
	}

	// 6. validityAreaList-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.ValidityAreaList_r16 = new(ValidityAreaList_r16)
			if err := ie.ValidityAreaList_r16.Decode(d); err != nil {
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
				{Name: "measReselectionCarrierListNR-r18", Optional: true},
				{Name: "measIdleValidityDuration-r18", Optional: true},
				{Name: "measReselectionValidityDuration-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(measIdleConfigDedicatedR16ExtMeasReselectionCarrierListNRR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MeasReselectionCarrierListNR_r18 = make([]MeasReselectionCarrierNR_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MeasReselectionCarrierListNR_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.MeasIdleValidityDuration_r18 = new(MeasurementValidityDuration_r18)
			if err := ie.MeasIdleValidityDuration_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.MeasReselectionValidityDuration_r18 = new(MeasurementValidityDuration_r18)
			if err := ie.MeasReselectionValidityDuration_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
