// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CMRGroupingAndPairing-r17 (line 10879).

var cMRGroupingAndPairingR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofResourcesGroup1-r17"},
		{Name: "pair1OfNZP-CSI-RS-r17", Optional: true},
		{Name: "pair2OfNZP-CSI-RS-r17", Optional: true},
	},
}

var cMRGroupingAndPairingR17NrofResourcesGroup1R17Constraints = per.Constrained(1, 7)

type CMRGroupingAndPairing_r17 struct {
	NrofResourcesGroup1_r17 int64
	Pair1OfNZP_CSI_RS_r17   *NZP_CSI_RS_Pairing_r17
	Pair2OfNZP_CSI_RS_r17   *NZP_CSI_RS_Pairing_r17
}

func (ie *CMRGroupingAndPairing_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cMRGroupingAndPairingR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pair1OfNZP_CSI_RS_r17 != nil, ie.Pair2OfNZP_CSI_RS_r17 != nil}); err != nil {
		return err
	}

	// 2. nrofResourcesGroup1-r17: integer
	{
		if err := e.EncodeInteger(ie.NrofResourcesGroup1_r17, cMRGroupingAndPairingR17NrofResourcesGroup1R17Constraints); err != nil {
			return err
		}
	}

	// 3. pair1OfNZP-CSI-RS-r17: ref
	{
		if ie.Pair1OfNZP_CSI_RS_r17 != nil {
			if err := ie.Pair1OfNZP_CSI_RS_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. pair2OfNZP-CSI-RS-r17: ref
	{
		if ie.Pair2OfNZP_CSI_RS_r17 != nil {
			if err := ie.Pair2OfNZP_CSI_RS_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CMRGroupingAndPairing_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cMRGroupingAndPairingR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nrofResourcesGroup1-r17: integer
	{
		v0, err := d.DecodeInteger(cMRGroupingAndPairingR17NrofResourcesGroup1R17Constraints)
		if err != nil {
			return err
		}
		ie.NrofResourcesGroup1_r17 = v0
	}

	// 3. pair1OfNZP-CSI-RS-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Pair1OfNZP_CSI_RS_r17 = new(NZP_CSI_RS_Pairing_r17)
			if err := ie.Pair1OfNZP_CSI_RS_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. pair2OfNZP-CSI-RS-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Pair2OfNZP_CSI_RS_r17 = new(NZP_CSI_RS_Pairing_r17)
			if err := ie.Pair2OfNZP_CSI_RS_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
