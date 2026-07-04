// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RadioLinkMonitoringConfig (line 13195).

var radioLinkMonitoringConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "failureDetectionResourcesToAddModList", Optional: true},
		{Name: "failureDetectionResourcesToReleaseList", Optional: true},
		{Name: "beamFailureInstanceMaxCount", Optional: true},
		{Name: "beamFailureDetectionTimer", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var radioLinkMonitoringConfigFailureDetectionResourcesToAddModListConstraints = per.SizeRange(1, common.MaxNrofFailureDetectionResources)

var radioLinkMonitoringConfigFailureDetectionResourcesToReleaseListConstraints = per.SizeRange(1, common.MaxNrofFailureDetectionResources)

const (
	RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N1  = 0
	RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N2  = 1
	RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N3  = 2
	RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N4  = 3
	RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N5  = 4
	RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N6  = 5
	RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N8  = 6
	RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N10 = 7
)

var radioLinkMonitoringConfigBeamFailureInstanceMaxCountConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N1, RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N2, RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N3, RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N4, RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N5, RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N6, RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N8, RadioLinkMonitoringConfig_BeamFailureInstanceMaxCount_N10},
}

const (
	RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd1  = 0
	RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd2  = 1
	RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd3  = 2
	RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd4  = 3
	RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd5  = 4
	RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd6  = 5
	RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd8  = 6
	RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd10 = 7
)

var radioLinkMonitoringConfigBeamFailureDetectionTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd1, RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd2, RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd3, RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd4, RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd5, RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd6, RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd8, RadioLinkMonitoringConfig_BeamFailureDetectionTimer_Pbfd10},
}

type RadioLinkMonitoringConfig struct {
	FailureDetectionResourcesToAddModList  []RadioLinkMonitoringRS
	FailureDetectionResourcesToReleaseList []RadioLinkMonitoringRS_Id
	BeamFailureInstanceMaxCount            *int64
	BeamFailureDetectionTimer              *int64
	BeamFailure_r17                        *BeamFailureDetection_r17
}

func (ie *RadioLinkMonitoringConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(radioLinkMonitoringConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.BeamFailure_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureDetectionResourcesToAddModList != nil, ie.FailureDetectionResourcesToReleaseList != nil, ie.BeamFailureInstanceMaxCount != nil, ie.BeamFailureDetectionTimer != nil}); err != nil {
		return err
	}

	// 3. failureDetectionResourcesToAddModList: sequence-of
	{
		if ie.FailureDetectionResourcesToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(radioLinkMonitoringConfigFailureDetectionResourcesToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.FailureDetectionResourcesToAddModList))); err != nil {
				return err
			}
			for i := range ie.FailureDetectionResourcesToAddModList {
				if err := ie.FailureDetectionResourcesToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. failureDetectionResourcesToReleaseList: sequence-of
	{
		if ie.FailureDetectionResourcesToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(radioLinkMonitoringConfigFailureDetectionResourcesToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.FailureDetectionResourcesToReleaseList))); err != nil {
				return err
			}
			for i := range ie.FailureDetectionResourcesToReleaseList {
				if err := ie.FailureDetectionResourcesToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. beamFailureInstanceMaxCount: enumerated
	{
		if ie.BeamFailureInstanceMaxCount != nil {
			if err := e.EncodeEnumerated(*ie.BeamFailureInstanceMaxCount, radioLinkMonitoringConfigBeamFailureInstanceMaxCountConstraints); err != nil {
				return err
			}
		}
	}

	// 6. beamFailureDetectionTimer: enumerated
	{
		if ie.BeamFailureDetectionTimer != nil {
			if err := e.EncodeEnumerated(*ie.BeamFailureDetectionTimer, radioLinkMonitoringConfigBeamFailureDetectionTimerConstraints); err != nil {
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
					{Name: "beamFailure-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.BeamFailure_r17 != nil}); err != nil {
				return err
			}

			if ie.BeamFailure_r17 != nil {
				if err := ie.BeamFailure_r17.Encode(ex); err != nil {
					return err
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

func (ie *RadioLinkMonitoringConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(radioLinkMonitoringConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. failureDetectionResourcesToAddModList: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(radioLinkMonitoringConfigFailureDetectionResourcesToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FailureDetectionResourcesToAddModList = make([]RadioLinkMonitoringRS, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FailureDetectionResourcesToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. failureDetectionResourcesToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(radioLinkMonitoringConfigFailureDetectionResourcesToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FailureDetectionResourcesToReleaseList = make([]RadioLinkMonitoringRS_Id, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FailureDetectionResourcesToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. beamFailureInstanceMaxCount: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(radioLinkMonitoringConfigBeamFailureInstanceMaxCountConstraints)
			if err != nil {
				return err
			}
			ie.BeamFailureInstanceMaxCount = &idx
		}
	}

	// 6. beamFailureDetectionTimer: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(radioLinkMonitoringConfigBeamFailureDetectionTimerConstraints)
			if err != nil {
				return err
			}
			ie.BeamFailureDetectionTimer = &idx
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
				{Name: "beamFailure-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.BeamFailure_r17 = new(BeamFailureDetection_r17)
			if err := ie.BeamFailure_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
