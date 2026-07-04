// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: N3C-IndirectPathConfigRelay-r18 (line 10246).

var n3CIndirectPathConfigRelayR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "n3c-MappingToReleaseList-r18", Optional: true},
		{Name: "n3c-MappingToAddModList-r18", Optional: true},
	},
}

var n3CIndirectPathConfigRelayR18N3cMappingToReleaseListR18Constraints = per.SizeRange(1, common.MaxLC_ID)

var n3CIndirectPathConfigRelayR18N3cMappingToAddModListR18Constraints = per.SizeRange(1, common.MaxLC_ID)

type N3C_IndirectPathConfigRelay_r18 struct {
	N3c_MappingToReleaseList_r18 []SL_RemoteUE_RB_Identity_r17
	N3c_MappingToAddModList_r18  []N3C_MappingConfig_r18
}

func (ie *N3C_IndirectPathConfigRelay_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(n3CIndirectPathConfigRelayR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.N3c_MappingToReleaseList_r18 != nil, ie.N3c_MappingToAddModList_r18 != nil}); err != nil {
		return err
	}

	// 3. n3c-MappingToReleaseList-r18: sequence-of
	{
		if ie.N3c_MappingToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(n3CIndirectPathConfigRelayR18N3cMappingToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.N3c_MappingToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.N3c_MappingToReleaseList_r18 {
				if err := ie.N3c_MappingToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. n3c-MappingToAddModList-r18: sequence-of
	{
		if ie.N3c_MappingToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(n3CIndirectPathConfigRelayR18N3cMappingToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.N3c_MappingToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.N3c_MappingToAddModList_r18 {
				if err := ie.N3c_MappingToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *N3C_IndirectPathConfigRelay_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(n3CIndirectPathConfigRelayR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. n3c-MappingToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(n3CIndirectPathConfigRelayR18N3cMappingToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.N3c_MappingToReleaseList_r18 = make([]SL_RemoteUE_RB_Identity_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.N3c_MappingToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. n3c-MappingToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(n3CIndirectPathConfigRelayR18N3cMappingToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.N3c_MappingToAddModList_r18 = make([]N3C_MappingConfig_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.N3c_MappingToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
