// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MIMOParam-v1850 (line 14815).

var mIMOParamV1850Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalTDDConfig-perPCI-ToAddModList-r18", Optional: true},
		{Name: "additionalTDDConfig-perPCI-ToReleaseList-r18", Optional: true},
	},
}

var mIMOParamV1850AdditionalTDDConfigPerPCIToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofAdditionalPCI_r17)

var mIMOParamV1850AdditionalTDDConfigPerPCIToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofAdditionalPCI_r17)

type MIMOParam_v1850 struct {
	AdditionalTDDConfig_PerPCI_ToAddModList_r18  []AdditionalTDDConfig_PerPCI_ToAddMod_r18
	AdditionalTDDConfig_PerPCI_ToReleaseList_r18 []AdditionalPCIIndex_r17
}

func (ie *MIMOParam_v1850) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mIMOParamV1850Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalTDDConfig_PerPCI_ToAddModList_r18 != nil, ie.AdditionalTDDConfig_PerPCI_ToReleaseList_r18 != nil}); err != nil {
		return err
	}

	// 2. additionalTDDConfig-perPCI-ToAddModList-r18: sequence-of
	{
		if ie.AdditionalTDDConfig_PerPCI_ToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mIMOParamV1850AdditionalTDDConfigPerPCIToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AdditionalTDDConfig_PerPCI_ToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.AdditionalTDDConfig_PerPCI_ToAddModList_r18 {
				if err := ie.AdditionalTDDConfig_PerPCI_ToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. additionalTDDConfig-perPCI-ToReleaseList-r18: sequence-of
	{
		if ie.AdditionalTDDConfig_PerPCI_ToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mIMOParamV1850AdditionalTDDConfigPerPCIToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AdditionalTDDConfig_PerPCI_ToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.AdditionalTDDConfig_PerPCI_ToReleaseList_r18 {
				if err := ie.AdditionalTDDConfig_PerPCI_ToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MIMOParam_v1850) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mIMOParamV1850Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. additionalTDDConfig-perPCI-ToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(mIMOParamV1850AdditionalTDDConfigPerPCIToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AdditionalTDDConfig_PerPCI_ToAddModList_r18 = make([]AdditionalTDDConfig_PerPCI_ToAddMod_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AdditionalTDDConfig_PerPCI_ToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. additionalTDDConfig-perPCI-ToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(mIMOParamV1850AdditionalTDDConfigPerPCIToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AdditionalTDDConfig_PerPCI_ToReleaseList_r18 = make([]AdditionalPCIIndex_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AdditionalTDDConfig_PerPCI_ToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
