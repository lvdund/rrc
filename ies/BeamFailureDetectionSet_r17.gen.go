// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BeamFailureDetectionSet-r17 (line 13224).

var beamFailureDetectionSetR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bfdResourcesToAddModList-r17", Optional: true},
		{Name: "bfdResourcesToReleaseList-r17", Optional: true},
		{Name: "beamFailureInstanceMaxCount-r17", Optional: true},
		{Name: "beamFailureDetectionTimer-r17", Optional: true},
	},
}

var beamFailureDetectionSetR17BfdResourcesToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofBFDResourcePerSet_r17)

var beamFailureDetectionSetR17BfdResourcesToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofBFDResourcePerSet_r17)

const (
	BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N1  = 0
	BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N2  = 1
	BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N3  = 2
	BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N4  = 3
	BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N5  = 4
	BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N6  = 5
	BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N8  = 6
	BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N10 = 7
)

var beamFailureDetectionSetR17BeamFailureInstanceMaxCountR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N1, BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N2, BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N3, BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N4, BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N5, BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N6, BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N8, BeamFailureDetectionSet_r17_BeamFailureInstanceMaxCount_r17_N10},
}

const (
	BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd1  = 0
	BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd2  = 1
	BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd3  = 2
	BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd4  = 3
	BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd5  = 4
	BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd6  = 5
	BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd8  = 6
	BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd10 = 7
)

var beamFailureDetectionSetR17BeamFailureDetectionTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd1, BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd2, BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd3, BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd4, BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd5, BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd6, BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd8, BeamFailureDetectionSet_r17_BeamFailureDetectionTimer_r17_Pbfd10},
}

type BeamFailureDetectionSet_r17 struct {
	BfdResourcesToAddModList_r17    []BeamLinkMonitoringRS_r17
	BfdResourcesToReleaseList_r17   []BeamLinkMonitoringRS_Id_r17
	BeamFailureInstanceMaxCount_r17 *int64
	BeamFailureDetectionTimer_r17   *int64
}

func (ie *BeamFailureDetectionSet_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(beamFailureDetectionSetR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BfdResourcesToAddModList_r17 != nil, ie.BfdResourcesToReleaseList_r17 != nil, ie.BeamFailureInstanceMaxCount_r17 != nil, ie.BeamFailureDetectionTimer_r17 != nil}); err != nil {
		return err
	}

	// 3. bfdResourcesToAddModList-r17: sequence-of
	{
		if ie.BfdResourcesToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(beamFailureDetectionSetR17BfdResourcesToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.BfdResourcesToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.BfdResourcesToAddModList_r17 {
				if err := ie.BfdResourcesToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. bfdResourcesToReleaseList-r17: sequence-of
	{
		if ie.BfdResourcesToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(beamFailureDetectionSetR17BfdResourcesToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.BfdResourcesToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.BfdResourcesToReleaseList_r17 {
				if err := ie.BfdResourcesToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. beamFailureInstanceMaxCount-r17: enumerated
	{
		if ie.BeamFailureInstanceMaxCount_r17 != nil {
			if err := e.EncodeEnumerated(*ie.BeamFailureInstanceMaxCount_r17, beamFailureDetectionSetR17BeamFailureInstanceMaxCountR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. beamFailureDetectionTimer-r17: enumerated
	{
		if ie.BeamFailureDetectionTimer_r17 != nil {
			if err := e.EncodeEnumerated(*ie.BeamFailureDetectionTimer_r17, beamFailureDetectionSetR17BeamFailureDetectionTimerR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BeamFailureDetectionSet_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(beamFailureDetectionSetR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. bfdResourcesToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(beamFailureDetectionSetR17BfdResourcesToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.BfdResourcesToAddModList_r17 = make([]BeamLinkMonitoringRS_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.BfdResourcesToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. bfdResourcesToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(beamFailureDetectionSetR17BfdResourcesToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.BfdResourcesToReleaseList_r17 = make([]BeamLinkMonitoringRS_Id_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.BfdResourcesToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. beamFailureInstanceMaxCount-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(beamFailureDetectionSetR17BeamFailureInstanceMaxCountR17Constraints)
			if err != nil {
				return err
			}
			ie.BeamFailureInstanceMaxCount_r17 = &idx
		}
	}

	// 6. beamFailureDetectionTimer-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(beamFailureDetectionSetR17BeamFailureDetectionTimerR17Constraints)
			if err != nil {
				return err
			}
			ie.BeamFailureDetectionTimer_r17 = &idx
		}
	}

	return nil
}
