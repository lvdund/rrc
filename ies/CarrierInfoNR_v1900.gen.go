// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CarrierInfoNR-v1900 (line 1281).

var carrierInfoNRV1900Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r19"},
		{Name: "ssbSubcarrierSpacing-r19"},
		{Name: "smtc-r19", Optional: true},
		{Name: "satAssistanceInfoList-r19", Optional: true},
	},
}

var carrierInfoNRV1900SatAssistanceInfoListR19Constraints = per.SizeRange(1, common.MaxCellNTN_r17)

type CarrierInfoNR_v1900 struct {
	CarrierFreq_r19           ARFCN_ValueNR
	SsbSubcarrierSpacing_r19  SubcarrierSpacing
	Smtc_r19                  *SSB_MTC
	SatAssistanceInfoList_r19 []NTN_Config_r17
}

func (ie *CarrierInfoNR_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(carrierInfoNRV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Smtc_r19 != nil, ie.SatAssistanceInfoList_r19 != nil}); err != nil {
		return err
	}

	// 3. carrierFreq-r19: ref
	{
		if err := ie.CarrierFreq_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. ssbSubcarrierSpacing-r19: ref
	{
		if err := ie.SsbSubcarrierSpacing_r19.Encode(e); err != nil {
			return err
		}
	}

	// 5. smtc-r19: ref
	{
		if ie.Smtc_r19 != nil {
			if err := ie.Smtc_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. satAssistanceInfoList-r19: sequence-of
	{
		if ie.SatAssistanceInfoList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(carrierInfoNRV1900SatAssistanceInfoListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SatAssistanceInfoList_r19))); err != nil {
				return err
			}
			for i := range ie.SatAssistanceInfoList_r19 {
				if err := ie.SatAssistanceInfoList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CarrierInfoNR_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(carrierInfoNRV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. carrierFreq-r19: ref
	{
		if err := ie.CarrierFreq_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. ssbSubcarrierSpacing-r19: ref
	{
		if err := ie.SsbSubcarrierSpacing_r19.Decode(d); err != nil {
			return err
		}
	}

	// 5. smtc-r19: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Smtc_r19 = new(SSB_MTC)
			if err := ie.Smtc_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. satAssistanceInfoList-r19: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(carrierInfoNRV1900SatAssistanceInfoListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SatAssistanceInfoList_r19 = make([]NTN_Config_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SatAssistanceInfoList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
