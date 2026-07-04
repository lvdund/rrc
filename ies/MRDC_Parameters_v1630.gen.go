// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v1630 (line 22597).

var mRDCParametersV1630Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxUplinkDutyCycle-interBandENDC-FDD-TDD-PC2-r16", Optional: true},
		{Name: "interBandMRDC-WithOverlapDL-Bands-r16", Optional: true},
	},
}

const (
	MRDC_Parameters_v1630_InterBandMRDC_WithOverlapDL_Bands_r16_Supported = 0
)

var mRDCParametersV1630InterBandMRDCWithOverlapDLBandsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1630_InterBandMRDC_WithOverlapDL_Bands_r16_Supported},
}

var mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxUplinkDutyCycle-FDD-TDD-EN-DC1-r16", Optional: true},
		{Name: "maxUplinkDutyCycle-FDD-TDD-EN-DC2-r16", Optional: true},
	},
}

const (
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N30  = 0
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N40  = 1
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N50  = 2
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N60  = 3
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N70  = 4
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N80  = 5
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N90  = 6
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N100 = 7
)

var mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16MaxUplinkDutyCycleFDDTDDENDC1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N30, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N40, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N50, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N60, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N70, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N80, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N90, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_N100},
}

const (
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N30  = 0
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N40  = 1
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N50  = 2
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N60  = 3
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N70  = 4
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N80  = 5
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N90  = 6
	MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N100 = 7
)

var mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16MaxUplinkDutyCycleFDDTDDENDC2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N30, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N40, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N50, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N60, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N70, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N80, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N90, MRDC_Parameters_v1630_MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16_MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16_N100},
}

type MRDC_Parameters_v1630 struct {
	MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16 *struct {
		MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 *int64
		MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 *int64
	}
	InterBandMRDC_WithOverlapDL_Bands_r16 *int64
}

func (ie *MRDC_Parameters_v1630) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV1630Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16 != nil, ie.InterBandMRDC_WithOverlapDL_Bands_r16 != nil}); err != nil {
		return err
	}

	// 2. maxUplinkDutyCycle-interBandENDC-FDD-TDD-PC2-r16: sequence
	{
		if ie.MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16 != nil {
			c := ie.MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16
			mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16Seq := e.NewSequenceEncoder(mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16Constraints)
			if err := mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16Seq.EncodePreamble([]bool{c.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 != nil, c.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 != nil}); err != nil {
				return err
			}
			if c.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 != nil {
				if err := e.EncodeEnumerated((*c.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16), mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16MaxUplinkDutyCycleFDDTDDENDC1R16Constraints); err != nil {
					return err
				}
			}
			if c.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 != nil {
				if err := e.EncodeEnumerated((*c.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16), mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16MaxUplinkDutyCycleFDDTDDENDC2R16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 3. interBandMRDC-WithOverlapDL-Bands-r16: enumerated
	{
		if ie.InterBandMRDC_WithOverlapDL_Bands_r16 != nil {
			if err := e.EncodeEnumerated(*ie.InterBandMRDC_WithOverlapDL_Bands_r16, mRDCParametersV1630InterBandMRDCWithOverlapDLBandsR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_v1630) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV1630Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxUplinkDutyCycle-interBandENDC-FDD-TDD-PC2-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16 = &struct {
				MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 *int64
				MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 *int64
			}{}
			c := ie.MaxUplinkDutyCycle_InterBandENDC_FDD_TDD_PC2_r16
			mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16Seq := d.NewSequenceDecoder(mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16Constraints)
			if err := mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16Seq.DecodePreamble(); err != nil {
				return err
			}
			if mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16Seq.IsComponentPresent(0) {
				c.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 = new(int64)
				v, err := d.DecodeEnumerated(mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16MaxUplinkDutyCycleFDDTDDENDC1R16Constraints)
				if err != nil {
					return err
				}
				(*c.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16) = v
			}
			if mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16Seq.IsComponentPresent(1) {
				c.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 = new(int64)
				v, err := d.DecodeEnumerated(mRDCParametersV1630MaxUplinkDutyCycleInterBandENDCFDDTDDPC2R16MaxUplinkDutyCycleFDDTDDENDC2R16Constraints)
				if err != nil {
					return err
				}
				(*c.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16) = v
			}
		}
	}

	// 3. interBandMRDC-WithOverlapDL-Bands-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1630InterBandMRDCWithOverlapDLBandsR16Constraints)
			if err != nil {
				return err
			}
			ie.InterBandMRDC_WithOverlapDL_Bands_r16 = &idx
		}
	}

	return nil
}
