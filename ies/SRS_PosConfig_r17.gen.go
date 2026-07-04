// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosConfig-r17 (line 1445).

var sRSPosConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PosResourceSetToReleaseList-r17", Optional: true},
		{Name: "srs-PosResourceSetToAddModList-r17", Optional: true},
		{Name: "srs-PosResourceToReleaseList-r17", Optional: true},
		{Name: "srs-PosResourceToAddModList-r17", Optional: true},
	},
}

var sRSPosConfigR17SrsPosResourceSetToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofSRS_PosResourceSets_r16)

var sRSPosConfigR17SrsPosResourceSetToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofSRS_PosResourceSets_r16)

var sRSPosConfigR17SrsPosResourceToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofSRS_PosResources_r16)

var sRSPosConfigR17SrsPosResourceToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofSRS_PosResources_r16)

type SRS_PosConfig_r17 struct {
	Srs_PosResourceSetToReleaseList_r17 []SRS_PosResourceSetId_r16
	Srs_PosResourceSetToAddModList_r17  []SRS_PosResourceSet_r16
	Srs_PosResourceToReleaseList_r17    []SRS_PosResourceId_r16
	Srs_PosResourceToAddModList_r17     []SRS_PosResource_r16
}

func (ie *SRS_PosConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_PosResourceSetToReleaseList_r17 != nil, ie.Srs_PosResourceSetToAddModList_r17 != nil, ie.Srs_PosResourceToReleaseList_r17 != nil, ie.Srs_PosResourceToAddModList_r17 != nil}); err != nil {
		return err
	}

	// 2. srs-PosResourceSetToReleaseList-r17: sequence-of
	{
		if ie.Srs_PosResourceSetToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sRSPosConfigR17SrsPosResourceSetToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_PosResourceSetToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.Srs_PosResourceSetToReleaseList_r17 {
				if err := ie.Srs_PosResourceSetToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. srs-PosResourceSetToAddModList-r17: sequence-of
	{
		if ie.Srs_PosResourceSetToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sRSPosConfigR17SrsPosResourceSetToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_PosResourceSetToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.Srs_PosResourceSetToAddModList_r17 {
				if err := ie.Srs_PosResourceSetToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. srs-PosResourceToReleaseList-r17: sequence-of
	{
		if ie.Srs_PosResourceToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sRSPosConfigR17SrsPosResourceToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_PosResourceToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.Srs_PosResourceToReleaseList_r17 {
				if err := ie.Srs_PosResourceToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. srs-PosResourceToAddModList-r17: sequence-of
	{
		if ie.Srs_PosResourceToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sRSPosConfigR17SrsPosResourceToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_PosResourceToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.Srs_PosResourceToAddModList_r17 {
				if err := ie.Srs_PosResourceToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SRS_PosConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-PosResourceSetToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sRSPosConfigR17SrsPosResourceSetToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosResourceSetToReleaseList_r17 = make([]SRS_PosResourceSetId_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosResourceSetToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. srs-PosResourceSetToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sRSPosConfigR17SrsPosResourceSetToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosResourceSetToAddModList_r17 = make([]SRS_PosResourceSet_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosResourceSetToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. srs-PosResourceToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sRSPosConfigR17SrsPosResourceToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosResourceToReleaseList_r17 = make([]SRS_PosResourceId_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosResourceToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. srs-PosResourceToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sRSPosConfigR17SrsPosResourceToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosResourceToAddModList_r17 = make([]SRS_PosResource_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosResourceToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
