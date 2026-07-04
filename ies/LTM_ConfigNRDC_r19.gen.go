// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-ConfigNRDC-r19 (line 8780).

var lTMConfigNRDCR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-ConfigurationSCG-r19", Optional: true},
		{Name: "ltm-SK-CounterConfigToAddModList-r19", Optional: true},
		{Name: "ltm-SK-CounterConfigToReleaseList-r19", Optional: true},
	},
}

var lTMConfigNRDCR19LtmSKCounterConfigToAddModListR19Constraints = per.SizeRange(1, common.MaxSecurityCellSet_r18)

var lTMConfigNRDCR19LtmSKCounterConfigToReleaseListR19Constraints = per.SizeRange(1, common.MaxSecurityCellSet_r18)

type LTM_ConfigNRDC_r19 struct {
	Ltm_ConfigurationSCG_r19              *LTM_Config_r18
	Ltm_SK_CounterConfigToAddModList_r19  []SK_CounterConfigLTM_r19
	Ltm_SK_CounterConfigToReleaseList_r19 []LTM_NoSecurityChangeId_r19
}

func (ie *LTM_ConfigNRDC_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMConfigNRDCR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ltm_ConfigurationSCG_r19 != nil, ie.Ltm_SK_CounterConfigToAddModList_r19 != nil, ie.Ltm_SK_CounterConfigToReleaseList_r19 != nil}); err != nil {
		return err
	}

	// 3. ltm-ConfigurationSCG-r19: ref
	{
		if ie.Ltm_ConfigurationSCG_r19 != nil {
			if err := ie.Ltm_ConfigurationSCG_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ltm-SK-CounterConfigToAddModList-r19: sequence-of
	{
		if ie.Ltm_SK_CounterConfigToAddModList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMConfigNRDCR19LtmSKCounterConfigToAddModListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_SK_CounterConfigToAddModList_r19))); err != nil {
				return err
			}
			for i := range ie.Ltm_SK_CounterConfigToAddModList_r19 {
				if err := ie.Ltm_SK_CounterConfigToAddModList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. ltm-SK-CounterConfigToReleaseList-r19: sequence-of
	{
		if ie.Ltm_SK_CounterConfigToReleaseList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMConfigNRDCR19LtmSKCounterConfigToReleaseListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_SK_CounterConfigToReleaseList_r19))); err != nil {
				return err
			}
			for i := range ie.Ltm_SK_CounterConfigToReleaseList_r19 {
				if err := ie.Ltm_SK_CounterConfigToReleaseList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *LTM_ConfigNRDC_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMConfigNRDCR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ltm-ConfigurationSCG-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ltm_ConfigurationSCG_r19 = new(LTM_Config_r18)
			if err := ie.Ltm_ConfigurationSCG_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ltm-SK-CounterConfigToAddModList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(lTMConfigNRDCR19LtmSKCounterConfigToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_SK_CounterConfigToAddModList_r19 = make([]SK_CounterConfigLTM_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_SK_CounterConfigToAddModList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. ltm-SK-CounterConfigToReleaseList-r19: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(lTMConfigNRDCR19LtmSKCounterConfigToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_SK_CounterConfigToReleaseList_r19 = make([]LTM_NoSecurityChangeId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_SK_CounterConfigToReleaseList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
