// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-RemoteUE-ToAddMod-r17 (line 27443).

var sLRemoteUEToAddModR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-L2IdentityRemote-r17"},
		{Name: "sl-SRAP-ConfigRelay-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sLRemoteUEToAddModR17ExtSlSRAPConfigRelayToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofRemoteUE_r17)

var sLRemoteUEToAddModR17ExtSlSRAPConfigRelayToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofRemoteUE_r17)

type SL_RemoteUE_ToAddMod_r17 struct {
	Sl_L2IdentityRemote_r17               SL_DestinationIdentity_r16
	Sl_SRAP_ConfigRelay_r17               *SL_SRAP_Config_r17
	Sl_SRAP_ConfigRelay_ToAddModList_r19  []SL_SRAP_Config_ToAddMod_r19
	Sl_SRAP_ConfigRelay_ToReleaseList_r19 []SL_SRAP_ConfigId_r19
}

func (ie *SL_RemoteUE_ToAddMod_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRemoteUEToAddModR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_SRAP_ConfigRelay_ToAddModList_r19 != nil || ie.Sl_SRAP_ConfigRelay_ToReleaseList_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_SRAP_ConfigRelay_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-L2IdentityRemote-r17: ref
	{
		if err := ie.Sl_L2IdentityRemote_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-SRAP-ConfigRelay-r17: ref
	{
		if ie.Sl_SRAP_ConfigRelay_r17 != nil {
			if err := ie.Sl_SRAP_ConfigRelay_r17.Encode(e); err != nil {
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
					{Name: "sl-SRAP-ConfigRelay-ToAddModList-r19", Optional: true},
					{Name: "sl-SRAP-ConfigRelay-ToReleaseList-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_SRAP_ConfigRelay_ToAddModList_r19 != nil, ie.Sl_SRAP_ConfigRelay_ToReleaseList_r19 != nil}); err != nil {
				return err
			}

			if ie.Sl_SRAP_ConfigRelay_ToAddModList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLRemoteUEToAddModR17ExtSlSRAPConfigRelayToAddModListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_SRAP_ConfigRelay_ToAddModList_r19))); err != nil {
					return err
				}
				for i := range ie.Sl_SRAP_ConfigRelay_ToAddModList_r19 {
					if err := ie.Sl_SRAP_ConfigRelay_ToAddModList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_SRAP_ConfigRelay_ToReleaseList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLRemoteUEToAddModR17ExtSlSRAPConfigRelayToReleaseListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_SRAP_ConfigRelay_ToReleaseList_r19))); err != nil {
					return err
				}
				for i := range ie.Sl_SRAP_ConfigRelay_ToReleaseList_r19 {
					if err := ie.Sl_SRAP_ConfigRelay_ToReleaseList_r19[i].Encode(ex); err != nil {
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

func (ie *SL_RemoteUE_ToAddMod_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRemoteUEToAddModR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-L2IdentityRemote-r17: ref
	{
		if err := ie.Sl_L2IdentityRemote_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-SRAP-ConfigRelay-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_SRAP_ConfigRelay_r17 = new(SL_SRAP_Config_r17)
			if err := ie.Sl_SRAP_ConfigRelay_r17.Decode(d); err != nil {
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
				{Name: "sl-SRAP-ConfigRelay-ToAddModList-r19", Optional: true},
				{Name: "sl-SRAP-ConfigRelay-ToReleaseList-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sLRemoteUEToAddModR17ExtSlSRAPConfigRelayToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_SRAP_ConfigRelay_ToAddModList_r19 = make([]SL_SRAP_Config_ToAddMod_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_SRAP_ConfigRelay_ToAddModList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(sLRemoteUEToAddModR17ExtSlSRAPConfigRelayToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_SRAP_ConfigRelay_ToReleaseList_r19 = make([]SL_SRAP_ConfigId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_SRAP_ConfigRelay_ToReleaseList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
