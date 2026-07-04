// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-DRX-Config-r17 (line 27116).

var sLDRXConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DRX-ConfigGC-BC-r17", Optional: true},
		{Name: "sl-DRX-ConfigUC-ToReleaseList-r17", Optional: true},
		{Name: "sl-DRX-ConfigUC-ToAddModList-r17", Optional: true},
	},
}

var sLDRXConfigR17SlDRXConfigUCToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

var sLDRXConfigR17SlDRXConfigUCToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SL_DRX_Config_r17 struct {
	Sl_DRX_ConfigGC_BC_r17            *SL_DRX_ConfigGC_BC_r17
	Sl_DRX_ConfigUC_ToReleaseList_r17 []SL_DestinationIndex_r16
	Sl_DRX_ConfigUC_ToAddModList_r17  []SL_DRX_ConfigUC_Info_r17
}

func (ie *SL_DRX_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDRXConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DRX_ConfigGC_BC_r17 != nil, ie.Sl_DRX_ConfigUC_ToReleaseList_r17 != nil, ie.Sl_DRX_ConfigUC_ToAddModList_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-DRX-ConfigGC-BC-r17: ref
	{
		if ie.Sl_DRX_ConfigGC_BC_r17 != nil {
			if err := ie.Sl_DRX_ConfigGC_BC_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-DRX-ConfigUC-ToReleaseList-r17: sequence-of
	{
		if ie.Sl_DRX_ConfigUC_ToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLDRXConfigR17SlDRXConfigUCToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_DRX_ConfigUC_ToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_DRX_ConfigUC_ToReleaseList_r17 {
				if err := ie.Sl_DRX_ConfigUC_ToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-DRX-ConfigUC-ToAddModList-r17: sequence-of
	{
		if ie.Sl_DRX_ConfigUC_ToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLDRXConfigR17SlDRXConfigUCToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_DRX_ConfigUC_ToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_DRX_ConfigUC_ToAddModList_r17 {
				if err := ie.Sl_DRX_ConfigUC_ToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_DRX_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDRXConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DRX-ConfigGC-BC-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_DRX_ConfigGC_BC_r17 = new(SL_DRX_ConfigGC_BC_r17)
			if err := ie.Sl_DRX_ConfigGC_BC_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-DRX-ConfigUC-ToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLDRXConfigR17SlDRXConfigUCToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_DRX_ConfigUC_ToReleaseList_r17 = make([]SL_DestinationIndex_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_DRX_ConfigUC_ToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-DRX-ConfigUC-ToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sLDRXConfigR17SlDRXConfigUCToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_DRX_ConfigUC_ToAddModList_r17 = make([]SL_DRX_ConfigUC_Info_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_DRX_ConfigUC_ToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
