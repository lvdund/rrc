// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasIdleConfigSIB-r16 (line 9237).

var measIdleConfigSIBR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measIdleCarrierListNR-r16", Optional: true},
		{Name: "measIdleCarrierListEUTRA-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var measIdleConfigSIBR16MeasIdleCarrierListNRR16Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

var measIdleConfigSIBR16MeasIdleCarrierListEUTRAR16Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

var measIdleConfigSIBR16ExtMeasIdleCarrierListNRLessThan5MHzR18Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

var measIdleConfigSIBR16ExtMeasReselectionCarrierListNRR18Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

var measIdleConfigSIBR16ExtMeasReselectionCarrierListNRLessThan5MHzR18Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

type MeasIdleConfigSIB_r16 struct {
	MeasIdleCarrierListNR_r16                     []MeasIdleCarrierNR_r16
	MeasIdleCarrierListEUTRA_r16                  []MeasIdleCarrierEUTRA_r16
	MeasIdleCarrierListNR_LessThan5MHz_r18        []MeasIdleCarrierNR_r16
	MeasReselectionCarrierListNR_r18              []MeasReselectionCarrierNR_r18
	MeasReselectionCarrierListNR_LessThan5MHz_r18 []MeasReselectionCarrierNR_r18
	MeasIdleValidityDuration_r18                  *MeasurementValidityDuration_r18
	MeasReselectionValidityDuration_r18           *MeasurementValidityDuration_r18
}

func (ie *MeasIdleConfigSIB_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measIdleConfigSIBR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MeasIdleCarrierListNR_LessThan5MHz_r18 != nil || ie.MeasReselectionCarrierListNR_r18 != nil || ie.MeasReselectionCarrierListNR_LessThan5MHz_r18 != nil || ie.MeasIdleValidityDuration_r18 != nil || ie.MeasReselectionValidityDuration_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasIdleCarrierListNR_r16 != nil, ie.MeasIdleCarrierListEUTRA_r16 != nil}); err != nil {
		return err
	}

	// 3. measIdleCarrierListNR-r16: sequence-of
	{
		if ie.MeasIdleCarrierListNR_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(measIdleConfigSIBR16MeasIdleCarrierListNRR16Constraints)
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
			seqOf := e.NewSequenceOfEncoder(measIdleConfigSIBR16MeasIdleCarrierListEUTRAR16Constraints)
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

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "measIdleCarrierListNR-LessThan5MHz-r18", Optional: true},
					{Name: "measReselectionCarrierListNR-r18", Optional: true},
					{Name: "measReselectionCarrierListNR-LessThan5MHz-r18", Optional: true},
					{Name: "measIdleValidityDuration-r18", Optional: true},
					{Name: "measReselectionValidityDuration-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasIdleCarrierListNR_LessThan5MHz_r18 != nil, ie.MeasReselectionCarrierListNR_r18 != nil, ie.MeasReselectionCarrierListNR_LessThan5MHz_r18 != nil, ie.MeasIdleValidityDuration_r18 != nil, ie.MeasReselectionValidityDuration_r18 != nil}); err != nil {
				return err
			}

			if ie.MeasIdleCarrierListNR_LessThan5MHz_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(measIdleConfigSIBR16ExtMeasIdleCarrierListNRLessThan5MHzR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.MeasIdleCarrierListNR_LessThan5MHz_r18))); err != nil {
					return err
				}
				for i := range ie.MeasIdleCarrierListNR_LessThan5MHz_r18 {
					if err := ie.MeasIdleCarrierListNR_LessThan5MHz_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MeasReselectionCarrierListNR_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(measIdleConfigSIBR16ExtMeasReselectionCarrierListNRR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.MeasReselectionCarrierListNR_r18))); err != nil {
					return err
				}
				for i := range ie.MeasReselectionCarrierListNR_r18 {
					if err := ie.MeasReselectionCarrierListNR_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MeasReselectionCarrierListNR_LessThan5MHz_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(measIdleConfigSIBR16ExtMeasReselectionCarrierListNRLessThan5MHzR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.MeasReselectionCarrierListNR_LessThan5MHz_r18))); err != nil {
					return err
				}
				for i := range ie.MeasReselectionCarrierListNR_LessThan5MHz_r18 {
					if err := ie.MeasReselectionCarrierListNR_LessThan5MHz_r18[i].Encode(ex); err != nil {
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

func (ie *MeasIdleConfigSIB_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measIdleConfigSIBR16Constraints)

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
			seqOf := d.NewSequenceOfDecoder(measIdleConfigSIBR16MeasIdleCarrierListNRR16Constraints)
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
			seqOf := d.NewSequenceOfDecoder(measIdleConfigSIBR16MeasIdleCarrierListEUTRAR16Constraints)
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
				{Name: "measIdleCarrierListNR-LessThan5MHz-r18", Optional: true},
				{Name: "measReselectionCarrierListNR-r18", Optional: true},
				{Name: "measReselectionCarrierListNR-LessThan5MHz-r18", Optional: true},
				{Name: "measIdleValidityDuration-r18", Optional: true},
				{Name: "measReselectionValidityDuration-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(measIdleConfigSIBR16ExtMeasIdleCarrierListNRLessThan5MHzR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MeasIdleCarrierListNR_LessThan5MHz_r18 = make([]MeasIdleCarrierNR_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MeasIdleCarrierListNR_LessThan5MHz_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(measIdleConfigSIBR16ExtMeasReselectionCarrierListNRR18Constraints)
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

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(measIdleConfigSIBR16ExtMeasReselectionCarrierListNRLessThan5MHzR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MeasReselectionCarrierListNR_LessThan5MHz_r18 = make([]MeasReselectionCarrierNR_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MeasReselectionCarrierListNR_LessThan5MHz_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.MeasIdleValidityDuration_r18 = new(MeasurementValidityDuration_r18)
			if err := ie.MeasIdleValidityDuration_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.MeasReselectionValidityDuration_r18 = new(MeasurementValidityDuration_r18)
			if err := ie.MeasReselectionValidityDuration_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
