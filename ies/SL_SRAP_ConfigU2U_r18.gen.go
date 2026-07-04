// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-SRAP-ConfigU2U-r18 (line 28294).

var sLSRAPConfigU2UR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MappingToAddMod-U2U-List-r18", Optional: true},
		{Name: "sl-MappingToRelease-U2U-List-r18", Optional: true},
	},
}

var sLSRAPConfigU2UR18SlMappingToAddModU2UListR18Constraints = per.SizeRange(1, common.MaxSL_LCID_r16)

var sLSRAPConfigU2UR18SlMappingToReleaseU2UListR18Constraints = per.SizeRange(1, common.MaxSL_LCID_r16)

type SL_SRAP_ConfigU2U_r18 struct {
	Sl_MappingToAddMod_U2U_List_r18  []SL_MappingConfig_U2U_r18
	Sl_MappingToRelease_U2U_List_r18 []SLRB_Uu_ConfigIndex_r16
}

func (ie *SL_SRAP_ConfigU2U_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSRAPConfigU2UR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_MappingToAddMod_U2U_List_r18 != nil, ie.Sl_MappingToRelease_U2U_List_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-MappingToAddMod-U2U-List-r18: sequence-of
	{
		if ie.Sl_MappingToAddMod_U2U_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLSRAPConfigU2UR18SlMappingToAddModU2UListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_MappingToAddMod_U2U_List_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_MappingToAddMod_U2U_List_r18 {
				if err := ie.Sl_MappingToAddMod_U2U_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-MappingToRelease-U2U-List-r18: sequence-of
	{
		if ie.Sl_MappingToRelease_U2U_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLSRAPConfigU2UR18SlMappingToReleaseU2UListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_MappingToRelease_U2U_List_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_MappingToRelease_U2U_List_r18 {
				if err := ie.Sl_MappingToRelease_U2U_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_SRAP_ConfigU2U_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSRAPConfigU2UR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-MappingToAddMod-U2U-List-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLSRAPConfigU2UR18SlMappingToAddModU2UListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_MappingToAddMod_U2U_List_r18 = make([]SL_MappingConfig_U2U_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_MappingToAddMod_U2U_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-MappingToRelease-U2U-List-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLSRAPConfigU2UR18SlMappingToReleaseU2UListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_MappingToRelease_U2U_List_r18 = make([]SLRB_Uu_ConfigIndex_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_MappingToRelease_U2U_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
