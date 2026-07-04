// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BeamFailureRecoveryRSConfig-r16 (line 5158).

var beamFailureRecoveryRSConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrp-ThresholdBFR-r16", Optional: true},
		{Name: "candidateBeamRS-List-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var beamFailureRecoveryRSConfigR16CandidateBeamRSListR16Constraints = per.SizeRange(1, common.MaxNrofCandidateBeams_r16)

var beamFailureRecoveryRSConfigR16ExtCandidateBeamRSList2R17Constraints = per.SizeRange(1, common.MaxNrofCandidateBeams_r16)

type BeamFailureRecoveryRSConfig_r16 struct {
	Rsrp_ThresholdBFR_r16     *RSRP_Range
	CandidateBeamRS_List_r16  []CandidateBeamRS_r16
	CandidateBeamRS_List2_r17 []CandidateBeamRS_r16
}

func (ie *BeamFailureRecoveryRSConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(beamFailureRecoveryRSConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.CandidateBeamRS_List2_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rsrp_ThresholdBFR_r16 != nil, ie.CandidateBeamRS_List_r16 != nil}); err != nil {
		return err
	}

	// 3. rsrp-ThresholdBFR-r16: ref
	{
		if ie.Rsrp_ThresholdBFR_r16 != nil {
			if err := ie.Rsrp_ThresholdBFR_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. candidateBeamRS-List-r16: sequence-of
	{
		if ie.CandidateBeamRS_List_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(beamFailureRecoveryRSConfigR16CandidateBeamRSListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.CandidateBeamRS_List_r16))); err != nil {
				return err
			}
			for i := range ie.CandidateBeamRS_List_r16 {
				if err := ie.CandidateBeamRS_List_r16[i].Encode(e); err != nil {
					return err
				}
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
					{Name: "candidateBeamRS-List2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CandidateBeamRS_List2_r17 != nil}); err != nil {
				return err
			}

			if ie.CandidateBeamRS_List2_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(beamFailureRecoveryRSConfigR16ExtCandidateBeamRSList2R17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.CandidateBeamRS_List2_r17))); err != nil {
					return err
				}
				for i := range ie.CandidateBeamRS_List2_r17 {
					if err := ie.CandidateBeamRS_List2_r17[i].Encode(ex); err != nil {
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

func (ie *BeamFailureRecoveryRSConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(beamFailureRecoveryRSConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rsrp-ThresholdBFR-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Rsrp_ThresholdBFR_r16 = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdBFR_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. candidateBeamRS-List-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(beamFailureRecoveryRSConfigR16CandidateBeamRSListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CandidateBeamRS_List_r16 = make([]CandidateBeamRS_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CandidateBeamRS_List_r16[i].Decode(d); err != nil {
					return err
				}
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
				{Name: "candidateBeamRS-List2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(beamFailureRecoveryRSConfigR16ExtCandidateBeamRSList2R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CandidateBeamRS_List2_r17 = make([]CandidateBeamRS_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CandidateBeamRS_List2_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
