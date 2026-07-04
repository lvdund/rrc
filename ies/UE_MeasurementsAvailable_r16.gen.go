// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MeasurementsAvailable-r16 (line 26667).

var uEMeasurementsAvailableR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "logMeasAvailable-r16", Optional: true},
		{Name: "logMeasAvailableBT-r16", Optional: true},
		{Name: "logMeasAvailableWLAN-r16", Optional: true},
		{Name: "connEstFailInfoAvailable-r16", Optional: true},
		{Name: "rlf-InfoAvailable-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	UE_MeasurementsAvailable_r16_LogMeasAvailable_r16_True = 0
)

var uEMeasurementsAvailableR16LogMeasAvailableR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_MeasurementsAvailable_r16_LogMeasAvailable_r16_True},
}

const (
	UE_MeasurementsAvailable_r16_LogMeasAvailableBT_r16_True = 0
)

var uEMeasurementsAvailableR16LogMeasAvailableBTR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_MeasurementsAvailable_r16_LogMeasAvailableBT_r16_True},
}

const (
	UE_MeasurementsAvailable_r16_LogMeasAvailableWLAN_r16_True = 0
)

var uEMeasurementsAvailableR16LogMeasAvailableWLANR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_MeasurementsAvailable_r16_LogMeasAvailableWLAN_r16_True},
}

const (
	UE_MeasurementsAvailable_r16_ConnEstFailInfoAvailable_r16_True = 0
)

var uEMeasurementsAvailableR16ConnEstFailInfoAvailableR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_MeasurementsAvailable_r16_ConnEstFailInfoAvailable_r16_True},
}

const (
	UE_MeasurementsAvailable_r16_Rlf_InfoAvailable_r16_True = 0
)

var uEMeasurementsAvailableR16RlfInfoAvailableR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_MeasurementsAvailable_r16_Rlf_InfoAvailable_r16_True},
}

const (
	UE_MeasurementsAvailable_r16_Ext_SuccessHO_InfoAvailable_r17_True = 0
)

var uEMeasurementsAvailableR16ExtSuccessHOInfoAvailableR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_MeasurementsAvailable_r16_Ext_SuccessHO_InfoAvailable_r17_True},
}

const (
	UE_MeasurementsAvailable_r16_Ext_SuccessPSCell_InfoAvailable_r18_True = 0
)

var uEMeasurementsAvailableR16ExtSuccessPSCellInfoAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_MeasurementsAvailable_r16_Ext_SuccessPSCell_InfoAvailable_r18_True},
}

type UE_MeasurementsAvailable_r16 struct {
	LogMeasAvailable_r16            *int64
	LogMeasAvailableBT_r16          *int64
	LogMeasAvailableWLAN_r16        *int64
	ConnEstFailInfoAvailable_r16    *int64
	Rlf_InfoAvailable_r16           *int64
	SuccessHO_InfoAvailable_r17     *int64
	SigLogMeasConfigAvailable_r17   *bool
	SuccessPSCell_InfoAvailable_r18 *int64
}

func (ie *UE_MeasurementsAvailable_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMeasurementsAvailableR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.SuccessHO_InfoAvailable_r17 != nil || ie.SigLogMeasConfigAvailable_r17 != nil
	hasExtGroup1 := ie.SuccessPSCell_InfoAvailable_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LogMeasAvailable_r16 != nil, ie.LogMeasAvailableBT_r16 != nil, ie.LogMeasAvailableWLAN_r16 != nil, ie.ConnEstFailInfoAvailable_r16 != nil, ie.Rlf_InfoAvailable_r16 != nil}); err != nil {
		return err
	}

	// 3. logMeasAvailable-r16: enumerated
	{
		if ie.LogMeasAvailable_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LogMeasAvailable_r16, uEMeasurementsAvailableR16LogMeasAvailableR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. logMeasAvailableBT-r16: enumerated
	{
		if ie.LogMeasAvailableBT_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LogMeasAvailableBT_r16, uEMeasurementsAvailableR16LogMeasAvailableBTR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. logMeasAvailableWLAN-r16: enumerated
	{
		if ie.LogMeasAvailableWLAN_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LogMeasAvailableWLAN_r16, uEMeasurementsAvailableR16LogMeasAvailableWLANR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. connEstFailInfoAvailable-r16: enumerated
	{
		if ie.ConnEstFailInfoAvailable_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ConnEstFailInfoAvailable_r16, uEMeasurementsAvailableR16ConnEstFailInfoAvailableR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. rlf-InfoAvailable-r16: enumerated
	{
		if ie.Rlf_InfoAvailable_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Rlf_InfoAvailable_r16, uEMeasurementsAvailableR16RlfInfoAvailableR16Constraints); err != nil {
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
					{Name: "successHO-InfoAvailable-r17", Optional: true},
					{Name: "sigLogMeasConfigAvailable-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SuccessHO_InfoAvailable_r17 != nil, ie.SigLogMeasConfigAvailable_r17 != nil}); err != nil {
				return err
			}

			if ie.SuccessHO_InfoAvailable_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SuccessHO_InfoAvailable_r17, uEMeasurementsAvailableR16ExtSuccessHOInfoAvailableR17Constraints); err != nil {
					return err
				}
			}

			if ie.SigLogMeasConfigAvailable_r17 != nil {
				if err := ex.EncodeBoolean(*ie.SigLogMeasConfigAvailable_r17); err != nil {
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
					{Name: "successPSCell-InfoAvailable-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SuccessPSCell_InfoAvailable_r18 != nil}); err != nil {
				return err
			}

			if ie.SuccessPSCell_InfoAvailable_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SuccessPSCell_InfoAvailable_r18, uEMeasurementsAvailableR16ExtSuccessPSCellInfoAvailableR18Constraints); err != nil {
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

func (ie *UE_MeasurementsAvailable_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMeasurementsAvailableR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. logMeasAvailable-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uEMeasurementsAvailableR16LogMeasAvailableR16Constraints)
			if err != nil {
				return err
			}
			ie.LogMeasAvailable_r16 = &idx
		}
	}

	// 4. logMeasAvailableBT-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uEMeasurementsAvailableR16LogMeasAvailableBTR16Constraints)
			if err != nil {
				return err
			}
			ie.LogMeasAvailableBT_r16 = &idx
		}
	}

	// 5. logMeasAvailableWLAN-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uEMeasurementsAvailableR16LogMeasAvailableWLANR16Constraints)
			if err != nil {
				return err
			}
			ie.LogMeasAvailableWLAN_r16 = &idx
		}
	}

	// 6. connEstFailInfoAvailable-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(uEMeasurementsAvailableR16ConnEstFailInfoAvailableR16Constraints)
			if err != nil {
				return err
			}
			ie.ConnEstFailInfoAvailable_r16 = &idx
		}
	}

	// 7. rlf-InfoAvailable-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(uEMeasurementsAvailableR16RlfInfoAvailableR16Constraints)
			if err != nil {
				return err
			}
			ie.Rlf_InfoAvailable_r16 = &idx
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
				{Name: "successHO-InfoAvailable-r17", Optional: true},
				{Name: "sigLogMeasConfigAvailable-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(uEMeasurementsAvailableR16ExtSuccessHOInfoAvailableR17Constraints)
			if err != nil {
				return err
			}
			ie.SuccessHO_InfoAvailable_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.SigLogMeasConfigAvailable_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "successPSCell-InfoAvailable-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(uEMeasurementsAvailableR16ExtSuccessPSCellInfoAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.SuccessPSCell_InfoAvailable_r18 = &v
		}
	}

	return nil
}
