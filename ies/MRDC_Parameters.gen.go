// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters (line 22545).

var mRDCParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "singleUL-Transmission", Optional: true},
		{Name: "dynamicPowerSharingENDC", Optional: true},
		{Name: "tdm-Pattern", Optional: true},
		{Name: "ul-SharingEUTRA-NR", Optional: true},
		{Name: "ul-SwitchingTimeEUTRA-NR", Optional: true},
		{Name: "simultaneousRxTxInterBandENDC", Optional: true},
		{Name: "asyncIntraBandENDC", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	MRDC_Parameters_SingleUL_Transmission_Supported = 0
)

var mRDCParametersSingleULTransmissionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_SingleUL_Transmission_Supported},
}

const (
	MRDC_Parameters_DynamicPowerSharingENDC_Supported = 0
)

var mRDCParametersDynamicPowerSharingENDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_DynamicPowerSharingENDC_Supported},
}

const (
	MRDC_Parameters_Tdm_Pattern_Supported = 0
)

var mRDCParametersTdmPatternConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_Tdm_Pattern_Supported},
}

const (
	MRDC_Parameters_Ul_SharingEUTRA_NR_Tdm  = 0
	MRDC_Parameters_Ul_SharingEUTRA_NR_Fdm  = 1
	MRDC_Parameters_Ul_SharingEUTRA_NR_Both = 2
)

var mRDCParametersUlSharingEUTRANRConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_Ul_SharingEUTRA_NR_Tdm, MRDC_Parameters_Ul_SharingEUTRA_NR_Fdm, MRDC_Parameters_Ul_SharingEUTRA_NR_Both},
}

const (
	MRDC_Parameters_Ul_SwitchingTimeEUTRA_NR_Type1 = 0
	MRDC_Parameters_Ul_SwitchingTimeEUTRA_NR_Type2 = 1
)

var mRDCParametersUlSwitchingTimeEUTRANRConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_Ul_SwitchingTimeEUTRA_NR_Type1, MRDC_Parameters_Ul_SwitchingTimeEUTRA_NR_Type2},
}

const (
	MRDC_Parameters_SimultaneousRxTxInterBandENDC_Supported = 0
)

var mRDCParametersSimultaneousRxTxInterBandENDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_SimultaneousRxTxInterBandENDC_Supported},
}

const (
	MRDC_Parameters_AsyncIntraBandENDC_Supported = 0
)

var mRDCParametersAsyncIntraBandENDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_AsyncIntraBandENDC_Supported},
}

const (
	MRDC_Parameters_Ext_DualPA_Architecture_Supported = 0
)

var mRDCParametersExtDualPAArchitectureConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_Ext_DualPA_Architecture_Supported},
}

const (
	MRDC_Parameters_Ext_IntraBandENDC_Support_Non_Contiguous = 0
	MRDC_Parameters_Ext_IntraBandENDC_Support_Both           = 1
)

var mRDCParametersExtIntraBandENDCSupportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_Ext_IntraBandENDC_Support_Non_Contiguous, MRDC_Parameters_Ext_IntraBandENDC_Support_Both},
}

const (
	MRDC_Parameters_Ext_Ul_TimingAlignmentEUTRA_NR_Required = 0
)

var mRDCParametersExtUlTimingAlignmentEUTRANRConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_Ext_Ul_TimingAlignmentEUTRA_NR_Required},
}

type MRDC_Parameters struct {
	SingleUL_Transmission         *int64
	DynamicPowerSharingENDC       *int64
	Tdm_Pattern                   *int64
	Ul_SharingEUTRA_NR            *int64
	Ul_SwitchingTimeEUTRA_NR      *int64
	SimultaneousRxTxInterBandENDC *int64
	AsyncIntraBandENDC            *int64
	DualPA_Architecture           *int64
	IntraBandENDC_Support         *int64
	Ul_TimingAlignmentEUTRA_NR    *int64
}

func (ie *MRDC_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.DualPA_Architecture != nil || ie.IntraBandENDC_Support != nil || ie.Ul_TimingAlignmentEUTRA_NR != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SingleUL_Transmission != nil, ie.DynamicPowerSharingENDC != nil, ie.Tdm_Pattern != nil, ie.Ul_SharingEUTRA_NR != nil, ie.Ul_SwitchingTimeEUTRA_NR != nil, ie.SimultaneousRxTxInterBandENDC != nil, ie.AsyncIntraBandENDC != nil}); err != nil {
		return err
	}

	// 3. singleUL-Transmission: enumerated
	{
		if ie.SingleUL_Transmission != nil {
			if err := e.EncodeEnumerated(*ie.SingleUL_Transmission, mRDCParametersSingleULTransmissionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. dynamicPowerSharingENDC: enumerated
	{
		if ie.DynamicPowerSharingENDC != nil {
			if err := e.EncodeEnumerated(*ie.DynamicPowerSharingENDC, mRDCParametersDynamicPowerSharingENDCConstraints); err != nil {
				return err
			}
		}
	}

	// 5. tdm-Pattern: enumerated
	{
		if ie.Tdm_Pattern != nil {
			if err := e.EncodeEnumerated(*ie.Tdm_Pattern, mRDCParametersTdmPatternConstraints); err != nil {
				return err
			}
		}
	}

	// 6. ul-SharingEUTRA-NR: enumerated
	{
		if ie.Ul_SharingEUTRA_NR != nil {
			if err := e.EncodeEnumerated(*ie.Ul_SharingEUTRA_NR, mRDCParametersUlSharingEUTRANRConstraints); err != nil {
				return err
			}
		}
	}

	// 7. ul-SwitchingTimeEUTRA-NR: enumerated
	{
		if ie.Ul_SwitchingTimeEUTRA_NR != nil {
			if err := e.EncodeEnumerated(*ie.Ul_SwitchingTimeEUTRA_NR, mRDCParametersUlSwitchingTimeEUTRANRConstraints); err != nil {
				return err
			}
		}
	}

	// 8. simultaneousRxTxInterBandENDC: enumerated
	{
		if ie.SimultaneousRxTxInterBandENDC != nil {
			if err := e.EncodeEnumerated(*ie.SimultaneousRxTxInterBandENDC, mRDCParametersSimultaneousRxTxInterBandENDCConstraints); err != nil {
				return err
			}
		}
	}

	// 9. asyncIntraBandENDC: enumerated
	{
		if ie.AsyncIntraBandENDC != nil {
			if err := e.EncodeEnumerated(*ie.AsyncIntraBandENDC, mRDCParametersAsyncIntraBandENDCConstraints); err != nil {
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
					{Name: "dualPA-Architecture", Optional: true},
					{Name: "intraBandENDC-Support", Optional: true},
					{Name: "ul-TimingAlignmentEUTRA-NR", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DualPA_Architecture != nil, ie.IntraBandENDC_Support != nil, ie.Ul_TimingAlignmentEUTRA_NR != nil}); err != nil {
				return err
			}

			if ie.DualPA_Architecture != nil {
				if err := ex.EncodeEnumerated(*ie.DualPA_Architecture, mRDCParametersExtDualPAArchitectureConstraints); err != nil {
					return err
				}
			}

			if ie.IntraBandENDC_Support != nil {
				if err := ex.EncodeEnumerated(*ie.IntraBandENDC_Support, mRDCParametersExtIntraBandENDCSupportConstraints); err != nil {
					return err
				}
			}

			if ie.Ul_TimingAlignmentEUTRA_NR != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_TimingAlignmentEUTRA_NR, mRDCParametersExtUlTimingAlignmentEUTRANRConstraints); err != nil {
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

func (ie *MRDC_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. singleUL-Transmission: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCParametersSingleULTransmissionConstraints)
			if err != nil {
				return err
			}
			ie.SingleUL_Transmission = &idx
		}
	}

	// 4. dynamicPowerSharingENDC: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mRDCParametersDynamicPowerSharingENDCConstraints)
			if err != nil {
				return err
			}
			ie.DynamicPowerSharingENDC = &idx
		}
	}

	// 5. tdm-Pattern: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mRDCParametersTdmPatternConstraints)
			if err != nil {
				return err
			}
			ie.Tdm_Pattern = &idx
		}
	}

	// 6. ul-SharingEUTRA-NR: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(mRDCParametersUlSharingEUTRANRConstraints)
			if err != nil {
				return err
			}
			ie.Ul_SharingEUTRA_NR = &idx
		}
	}

	// 7. ul-SwitchingTimeEUTRA-NR: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(mRDCParametersUlSwitchingTimeEUTRANRConstraints)
			if err != nil {
				return err
			}
			ie.Ul_SwitchingTimeEUTRA_NR = &idx
		}
	}

	// 8. simultaneousRxTxInterBandENDC: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(mRDCParametersSimultaneousRxTxInterBandENDCConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousRxTxInterBandENDC = &idx
		}
	}

	// 9. asyncIntraBandENDC: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(mRDCParametersAsyncIntraBandENDCConstraints)
			if err != nil {
				return err
			}
			ie.AsyncIntraBandENDC = &idx
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
				{Name: "dualPA-Architecture", Optional: true},
				{Name: "intraBandENDC-Support", Optional: true},
				{Name: "ul-TimingAlignmentEUTRA-NR", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mRDCParametersExtDualPAArchitectureConstraints)
			if err != nil {
				return err
			}
			ie.DualPA_Architecture = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mRDCParametersExtIntraBandENDCSupportConstraints)
			if err != nil {
				return err
			}
			ie.IntraBandENDC_Support = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(mRDCParametersExtUlTimingAlignmentEUTRANRConstraints)
			if err != nil {
				return err
			}
			ie.Ul_TimingAlignmentEUTRA_NR = &v
		}
	}

	return nil
}
