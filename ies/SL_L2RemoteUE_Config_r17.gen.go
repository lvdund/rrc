// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-L2RemoteUE-Config-r17 (line 27475).

var sLL2RemoteUEConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-SRAP-ConfigRemote-r17", Optional: true},
		{Name: "sl-UEIdentityRemote-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sLL2RemoteUEConfigR17ExtSlU2URelayUEToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

var sLL2RemoteUEConfigR17ExtSlU2URelayUEToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SL_L2RemoteUE_Config_r17 struct {
	Sl_SRAP_ConfigRemote_r17         *SL_SRAP_Config_r17
	Sl_UEIdentityRemote_r17          *RNTI_Value
	Sl_U2U_RelayUE_ToAddModList_r18  []SL_U2U_RelayUE_Config_r18
	Sl_U2U_RelayUE_ToReleaseList_r18 []SL_DestinationIdentity_r16
}

func (ie *SL_L2RemoteUE_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLL2RemoteUEConfigR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_U2U_RelayUE_ToAddModList_r18 != nil || ie.Sl_U2U_RelayUE_ToReleaseList_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_SRAP_ConfigRemote_r17 != nil, ie.Sl_UEIdentityRemote_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-SRAP-ConfigRemote-r17: ref
	{
		if ie.Sl_SRAP_ConfigRemote_r17 != nil {
			if err := ie.Sl_SRAP_ConfigRemote_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-UEIdentityRemote-r17: ref
	{
		if ie.Sl_UEIdentityRemote_r17 != nil {
			if err := ie.Sl_UEIdentityRemote_r17.Encode(e); err != nil {
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
					{Name: "sl-U2U-RelayUE-ToAddModList-r18", Optional: true},
					{Name: "sl-U2U-RelayUE-ToReleaseList-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_U2U_RelayUE_ToAddModList_r18 != nil, ie.Sl_U2U_RelayUE_ToReleaseList_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_U2U_RelayUE_ToAddModList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLL2RemoteUEConfigR17ExtSlU2URelayUEToAddModListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_U2U_RelayUE_ToAddModList_r18))); err != nil {
					return err
				}
				for i := range ie.Sl_U2U_RelayUE_ToAddModList_r18 {
					if err := ie.Sl_U2U_RelayUE_ToAddModList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_U2U_RelayUE_ToReleaseList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLL2RemoteUEConfigR17ExtSlU2URelayUEToReleaseListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_U2U_RelayUE_ToReleaseList_r18))); err != nil {
					return err
				}
				for i := range ie.Sl_U2U_RelayUE_ToReleaseList_r18 {
					if err := ie.Sl_U2U_RelayUE_ToReleaseList_r18[i].Encode(ex); err != nil {
						return err
					}
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

func (ie *SL_L2RemoteUE_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLL2RemoteUEConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-SRAP-ConfigRemote-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_SRAP_ConfigRemote_r17 = new(SL_SRAP_Config_r17)
			if err := ie.Sl_SRAP_ConfigRemote_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-UEIdentityRemote-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_UEIdentityRemote_r17 = new(RNTI_Value)
			if err := ie.Sl_UEIdentityRemote_r17.Decode(d); err != nil {
				return err
			}
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
				{Name: "sl-U2U-RelayUE-ToAddModList-r18", Optional: true},
				{Name: "sl-U2U-RelayUE-ToReleaseList-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sLL2RemoteUEConfigR17ExtSlU2URelayUEToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_U2U_RelayUE_ToAddModList_r18 = make([]SL_U2U_RelayUE_Config_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_U2U_RelayUE_ToAddModList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(sLL2RemoteUEConfigR17ExtSlU2URelayUEToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_U2U_RelayUE_ToReleaseList_r18 = make([]SL_DestinationIdentity_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_U2U_RelayUE_ToReleaseList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
