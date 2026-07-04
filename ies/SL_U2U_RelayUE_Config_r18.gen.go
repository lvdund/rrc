// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-U2U-RelayUE-Config-r18 (line 27485).

var sLU2URelayUEConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-L2IdentityRelay-r18"},
		{Name: "sl-TargetRemoteUE-ToAddModList-r18", Optional: true},
		{Name: "sl-TargetRemoteUE-ToReleaseList-r18", Optional: true},
	},
}

var sLU2URelayUEConfigR18SlTargetRemoteUEToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

var sLU2URelayUEConfigR18SlTargetRemoteUEToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SL_U2U_RelayUE_Config_r18 struct {
	Sl_L2IdentityRelay_r18              SL_DestinationIdentity_r16
	Sl_TargetRemoteUE_ToAddModList_r18  []SL_TargetRemoteUE_Config_r18
	Sl_TargetRemoteUE_ToReleaseList_r18 []SL_DestinationIdentity_r16
}

func (ie *SL_U2U_RelayUE_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLU2URelayUEConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_TargetRemoteUE_ToAddModList_r18 != nil, ie.Sl_TargetRemoteUE_ToReleaseList_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-L2IdentityRelay-r18: ref
	{
		if err := ie.Sl_L2IdentityRelay_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-TargetRemoteUE-ToAddModList-r18: sequence-of
	{
		if ie.Sl_TargetRemoteUE_ToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLU2URelayUEConfigR18SlTargetRemoteUEToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_TargetRemoteUE_ToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_TargetRemoteUE_ToAddModList_r18 {
				if err := ie.Sl_TargetRemoteUE_ToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-TargetRemoteUE-ToReleaseList-r18: sequence-of
	{
		if ie.Sl_TargetRemoteUE_ToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLU2URelayUEConfigR18SlTargetRemoteUEToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_TargetRemoteUE_ToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_TargetRemoteUE_ToReleaseList_r18 {
				if err := ie.Sl_TargetRemoteUE_ToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_U2U_RelayUE_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLU2URelayUEConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-L2IdentityRelay-r18: ref
	{
		if err := ie.Sl_L2IdentityRelay_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-TargetRemoteUE-ToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLU2URelayUEConfigR18SlTargetRemoteUEToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_TargetRemoteUE_ToAddModList_r18 = make([]SL_TargetRemoteUE_Config_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_TargetRemoteUE_ToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-TargetRemoteUE-ToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sLU2URelayUEConfigR18SlTargetRemoteUEToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_TargetRemoteUE_ToReleaseList_r18 = make([]SL_DestinationIdentity_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_TargetRemoteUE_ToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
