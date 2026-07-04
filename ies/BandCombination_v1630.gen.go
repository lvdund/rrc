// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1630 (line 16708).

var bandCombinationV1630Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1630", Optional: true},
		{Name: "ca-ParametersNRDC-v1630", Optional: true},
		{Name: "mrdc-Parameters-v1630", Optional: true},
		{Name: "supportedTxBandCombListPerBC-Sidelink-r16", Optional: true},
		{Name: "supportedRxBandCombListPerBC-Sidelink-r16", Optional: true},
		{Name: "scalingFactorTxSidelink-r16", Optional: true},
		{Name: "scalingFactorRxSidelink-r16", Optional: true},
	},
}

var bandCombinationV1630SupportedTxBandCombListPerBCSidelinkR16Constraints = per.SizeRange(1, common.MaxBandComb)

var bandCombinationV1630SupportedRxBandCombListPerBCSidelinkR16Constraints = per.SizeRange(1, common.MaxBandComb)

var bandCombinationV1630ScalingFactorTxSidelinkR16Constraints = per.SizeRange(1, common.MaxBandComb)

var bandCombinationV1630ScalingFactorRxSidelinkR16Constraints = per.SizeRange(1, common.MaxBandComb)

type BandCombination_v1630 struct {
	Ca_ParametersNR_v1630                     *CA_ParametersNR_v1630
	Ca_ParametersNRDC_v1630                   *CA_ParametersNRDC_v1630
	Mrdc_Parameters_v1630                     *MRDC_Parameters_v1630
	SupportedTxBandCombListPerBC_Sidelink_r16 *per.BitString
	SupportedRxBandCombListPerBC_Sidelink_r16 *per.BitString
	ScalingFactorTxSidelink_r16               []ScalingFactorSidelink_r16
	ScalingFactorRxSidelink_r16               []ScalingFactorSidelink_r16
}

func (ie *BandCombination_v1630) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1630Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_v1630 != nil, ie.Ca_ParametersNRDC_v1630 != nil, ie.Mrdc_Parameters_v1630 != nil, ie.SupportedTxBandCombListPerBC_Sidelink_r16 != nil, ie.SupportedRxBandCombListPerBC_Sidelink_r16 != nil, ie.ScalingFactorTxSidelink_r16 != nil, ie.ScalingFactorRxSidelink_r16 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1630: ref
	{
		if ie.Ca_ParametersNR_v1630 != nil {
			if err := ie.Ca_ParametersNR_v1630.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1630: ref
	{
		if ie.Ca_ParametersNRDC_v1630 != nil {
			if err := ie.Ca_ParametersNRDC_v1630.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. mrdc-Parameters-v1630: ref
	{
		if ie.Mrdc_Parameters_v1630 != nil {
			if err := ie.Mrdc_Parameters_v1630.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. supportedTxBandCombListPerBC-Sidelink-r16: bit-string
	{
		if ie.SupportedTxBandCombListPerBC_Sidelink_r16 != nil {
			if err := e.EncodeBitString(*ie.SupportedTxBandCombListPerBC_Sidelink_r16, bandCombinationV1630SupportedTxBandCombListPerBCSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. supportedRxBandCombListPerBC-Sidelink-r16: bit-string
	{
		if ie.SupportedRxBandCombListPerBC_Sidelink_r16 != nil {
			if err := e.EncodeBitString(*ie.SupportedRxBandCombListPerBC_Sidelink_r16, bandCombinationV1630SupportedRxBandCombListPerBCSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. scalingFactorTxSidelink-r16: sequence-of
	{
		if ie.ScalingFactorTxSidelink_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1630ScalingFactorTxSidelinkR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ScalingFactorTxSidelink_r16))); err != nil {
				return err
			}
			for i := range ie.ScalingFactorTxSidelink_r16 {
				if err := ie.ScalingFactorTxSidelink_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. scalingFactorRxSidelink-r16: sequence-of
	{
		if ie.ScalingFactorRxSidelink_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1630ScalingFactorRxSidelinkR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ScalingFactorRxSidelink_r16))); err != nil {
				return err
			}
			for i := range ie.ScalingFactorRxSidelink_r16 {
				if err := ie.ScalingFactorRxSidelink_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1630) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1630Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1630: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_v1630 = new(CA_ParametersNR_v1630)
			if err := ie.Ca_ParametersNR_v1630.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1630: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNRDC_v1630 = new(CA_ParametersNRDC_v1630)
			if err := ie.Ca_ParametersNRDC_v1630.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. mrdc-Parameters-v1630: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Mrdc_Parameters_v1630 = new(MRDC_Parameters_v1630)
			if err := ie.Mrdc_Parameters_v1630.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. supportedTxBandCombListPerBC-Sidelink-r16: bit-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeBitString(bandCombinationV1630SupportedTxBandCombListPerBCSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportedTxBandCombListPerBC_Sidelink_r16 = &v
		}
	}

	// 6. supportedRxBandCombListPerBC-Sidelink-r16: bit-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBitString(bandCombinationV1630SupportedRxBandCombListPerBCSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportedRxBandCombListPerBC_Sidelink_r16 = &v
		}
	}

	// 7. scalingFactorTxSidelink-r16: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1630ScalingFactorTxSidelinkR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ScalingFactorTxSidelink_r16 = make([]ScalingFactorSidelink_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ScalingFactorTxSidelink_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. scalingFactorRxSidelink-r16: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1630ScalingFactorRxSidelinkR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ScalingFactorRxSidelink_r16 = make([]ScalingFactorSidelink_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ScalingFactorRxSidelink_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
