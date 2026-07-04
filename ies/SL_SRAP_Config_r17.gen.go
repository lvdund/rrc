// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-SRAP-Config-r17 (line 28262).

var sLSRAPConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-LocalIdentity-r17", Optional: true},
		{Name: "sl-MappingToAddModList-r17", Optional: true},
		{Name: "sl-MappingToReleaseList-r17", Optional: true},
	},
}

var sLSRAPConfigR17SlLocalIdentityR17Constraints = per.Constrained(0, 255)

var sLSRAPConfigR17SlMappingToAddModListR17Constraints = per.SizeRange(1, common.MaxLC_ID)

var sLSRAPConfigR17SlMappingToReleaseListR17Constraints = per.SizeRange(1, common.MaxLC_ID)

type SL_SRAP_Config_r17 struct {
	Sl_LocalIdentity_r17        *int64
	Sl_MappingToAddModList_r17  []SL_MappingToAddMod_r17
	Sl_MappingToReleaseList_r17 []SL_RemoteUE_RB_Identity_r17
}

func (ie *SL_SRAP_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSRAPConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_LocalIdentity_r17 != nil, ie.Sl_MappingToAddModList_r17 != nil, ie.Sl_MappingToReleaseList_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-LocalIdentity-r17: integer
	{
		if ie.Sl_LocalIdentity_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_LocalIdentity_r17, sLSRAPConfigR17SlLocalIdentityR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-MappingToAddModList-r17: sequence-of
	{
		if ie.Sl_MappingToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLSRAPConfigR17SlMappingToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_MappingToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_MappingToAddModList_r17 {
				if err := ie.Sl_MappingToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-MappingToReleaseList-r17: sequence-of
	{
		if ie.Sl_MappingToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLSRAPConfigR17SlMappingToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_MappingToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_MappingToReleaseList_r17 {
				if err := ie.Sl_MappingToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_SRAP_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSRAPConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-LocalIdentity-r17: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLSRAPConfigR17SlLocalIdentityR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LocalIdentity_r17 = &v
		}
	}

	// 4. sl-MappingToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLSRAPConfigR17SlMappingToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_MappingToAddModList_r17 = make([]SL_MappingToAddMod_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_MappingToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-MappingToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sLSRAPConfigR17SlMappingToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_MappingToReleaseList_r17 = make([]SL_RemoteUE_RB_Identity_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_MappingToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
