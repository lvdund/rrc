// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-U2U-RemoteUE-Config-r18 (line 27453).

var sLU2URemoteUEConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-L2IdentityRemoteUE-r18"},
		{Name: "sl-SourceRemoteUE-ToAddModList-r18", Optional: true},
		{Name: "sl-SourceRemoteUE-ToReleaseList-r18", Optional: true},
	},
}

var sLU2URemoteUEConfigR18SlSourceRemoteUEToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

var sLU2URemoteUEConfigR18SlSourceRemoteUEToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SL_U2U_RemoteUE_Config_r18 struct {
	Sl_L2IdentityRemoteUE_r18           SL_DestinationIdentity_r16
	Sl_SourceRemoteUE_ToAddModList_r18  []SL_SourceRemoteUE_Config_r18
	Sl_SourceRemoteUE_ToReleaseList_r18 []SL_SourceIdentity_r17
}

func (ie *SL_U2U_RemoteUE_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLU2URemoteUEConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_SourceRemoteUE_ToAddModList_r18 != nil, ie.Sl_SourceRemoteUE_ToReleaseList_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-L2IdentityRemoteUE-r18: ref
	{
		if err := ie.Sl_L2IdentityRemoteUE_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-SourceRemoteUE-ToAddModList-r18: sequence-of
	{
		if ie.Sl_SourceRemoteUE_ToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLU2URemoteUEConfigR18SlSourceRemoteUEToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_SourceRemoteUE_ToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_SourceRemoteUE_ToAddModList_r18 {
				if err := ie.Sl_SourceRemoteUE_ToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-SourceRemoteUE-ToReleaseList-r18: sequence-of
	{
		if ie.Sl_SourceRemoteUE_ToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLU2URemoteUEConfigR18SlSourceRemoteUEToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_SourceRemoteUE_ToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_SourceRemoteUE_ToReleaseList_r18 {
				if err := ie.Sl_SourceRemoteUE_ToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_U2U_RemoteUE_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLU2URemoteUEConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-L2IdentityRemoteUE-r18: ref
	{
		if err := ie.Sl_L2IdentityRemoteUE_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-SourceRemoteUE-ToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLU2URemoteUEConfigR18SlSourceRemoteUEToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_SourceRemoteUE_ToAddModList_r18 = make([]SL_SourceRemoteUE_Config_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_SourceRemoteUE_ToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-SourceRemoteUE-ToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sLU2URemoteUEConfigR18SlSourceRemoteUEToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_SourceRemoteUE_ToReleaseList_r18 = make([]SL_SourceIdentity_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_SourceRemoteUE_ToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
