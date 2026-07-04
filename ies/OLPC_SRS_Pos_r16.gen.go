// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OLPC-SRS-Pos-r16 (line 22697).

var oLPCSRSPosR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "olpc-SRS-PosBasedOnPRS-Serving-r16", Optional: true},
		{Name: "olpc-SRS-PosBasedOnSSB-Neigh-r16", Optional: true},
		{Name: "olpc-SRS-PosBasedOnPRS-Neigh-r16", Optional: true},
		{Name: "maxNumberPathLossEstimatePerServing-r16", Optional: true},
	},
}

const (
	OLPC_SRS_Pos_r16_Olpc_SRS_PosBasedOnPRS_Serving_r16_Supported = 0
)

var oLPCSRSPosR16OlpcSRSPosBasedOnPRSServingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OLPC_SRS_Pos_r16_Olpc_SRS_PosBasedOnPRS_Serving_r16_Supported},
}

const (
	OLPC_SRS_Pos_r16_Olpc_SRS_PosBasedOnSSB_Neigh_r16_Supported = 0
)

var oLPCSRSPosR16OlpcSRSPosBasedOnSSBNeighR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OLPC_SRS_Pos_r16_Olpc_SRS_PosBasedOnSSB_Neigh_r16_Supported},
}

const (
	OLPC_SRS_Pos_r16_Olpc_SRS_PosBasedOnPRS_Neigh_r16_Supported = 0
)

var oLPCSRSPosR16OlpcSRSPosBasedOnPRSNeighR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OLPC_SRS_Pos_r16_Olpc_SRS_PosBasedOnPRS_Neigh_r16_Supported},
}

const (
	OLPC_SRS_Pos_r16_MaxNumberPathLossEstimatePerServing_r16_N1  = 0
	OLPC_SRS_Pos_r16_MaxNumberPathLossEstimatePerServing_r16_N4  = 1
	OLPC_SRS_Pos_r16_MaxNumberPathLossEstimatePerServing_r16_N8  = 2
	OLPC_SRS_Pos_r16_MaxNumberPathLossEstimatePerServing_r16_N16 = 3
)

var oLPCSRSPosR16MaxNumberPathLossEstimatePerServingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OLPC_SRS_Pos_r16_MaxNumberPathLossEstimatePerServing_r16_N1, OLPC_SRS_Pos_r16_MaxNumberPathLossEstimatePerServing_r16_N4, OLPC_SRS_Pos_r16_MaxNumberPathLossEstimatePerServing_r16_N8, OLPC_SRS_Pos_r16_MaxNumberPathLossEstimatePerServing_r16_N16},
}

type OLPC_SRS_Pos_r16 struct {
	Olpc_SRS_PosBasedOnPRS_Serving_r16      *int64
	Olpc_SRS_PosBasedOnSSB_Neigh_r16        *int64
	Olpc_SRS_PosBasedOnPRS_Neigh_r16        *int64
	MaxNumberPathLossEstimatePerServing_r16 *int64
}

func (ie *OLPC_SRS_Pos_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(oLPCSRSPosR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Olpc_SRS_PosBasedOnPRS_Serving_r16 != nil, ie.Olpc_SRS_PosBasedOnSSB_Neigh_r16 != nil, ie.Olpc_SRS_PosBasedOnPRS_Neigh_r16 != nil, ie.MaxNumberPathLossEstimatePerServing_r16 != nil}); err != nil {
		return err
	}

	// 2. olpc-SRS-PosBasedOnPRS-Serving-r16: enumerated
	{
		if ie.Olpc_SRS_PosBasedOnPRS_Serving_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Olpc_SRS_PosBasedOnPRS_Serving_r16, oLPCSRSPosR16OlpcSRSPosBasedOnPRSServingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. olpc-SRS-PosBasedOnSSB-Neigh-r16: enumerated
	{
		if ie.Olpc_SRS_PosBasedOnSSB_Neigh_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Olpc_SRS_PosBasedOnSSB_Neigh_r16, oLPCSRSPosR16OlpcSRSPosBasedOnSSBNeighR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. olpc-SRS-PosBasedOnPRS-Neigh-r16: enumerated
	{
		if ie.Olpc_SRS_PosBasedOnPRS_Neigh_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Olpc_SRS_PosBasedOnPRS_Neigh_r16, oLPCSRSPosR16OlpcSRSPosBasedOnPRSNeighR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maxNumberPathLossEstimatePerServing-r16: enumerated
	{
		if ie.MaxNumberPathLossEstimatePerServing_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumberPathLossEstimatePerServing_r16, oLPCSRSPosR16MaxNumberPathLossEstimatePerServingR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *OLPC_SRS_Pos_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(oLPCSRSPosR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. olpc-SRS-PosBasedOnPRS-Serving-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(oLPCSRSPosR16OlpcSRSPosBasedOnPRSServingR16Constraints)
			if err != nil {
				return err
			}
			ie.Olpc_SRS_PosBasedOnPRS_Serving_r16 = &idx
		}
	}

	// 3. olpc-SRS-PosBasedOnSSB-Neigh-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(oLPCSRSPosR16OlpcSRSPosBasedOnSSBNeighR16Constraints)
			if err != nil {
				return err
			}
			ie.Olpc_SRS_PosBasedOnSSB_Neigh_r16 = &idx
		}
	}

	// 4. olpc-SRS-PosBasedOnPRS-Neigh-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(oLPCSRSPosR16OlpcSRSPosBasedOnPRSNeighR16Constraints)
			if err != nil {
				return err
			}
			ie.Olpc_SRS_PosBasedOnPRS_Neigh_r16 = &idx
		}
	}

	// 5. maxNumberPathLossEstimatePerServing-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(oLPCSRSPosR16MaxNumberPathLossEstimatePerServingR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberPathLossEstimatePerServing_r16 = &idx
		}
	}

	return nil
}
