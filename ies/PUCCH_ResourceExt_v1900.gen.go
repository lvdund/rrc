// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-ResourceExt-v1900 (line 12116).

var pUCCHResourceExtV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "startingPRB-SBFD-r19", Optional: true},
		{Name: "secondHopPRB-SBFD-r19", Optional: true},
	},
}

type PUCCH_ResourceExt_v1900 struct {
	StartingPRB_SBFD_r19  *PRB_Id
	SecondHopPRB_SBFD_r19 *PRB_Id
}

func (ie *PUCCH_ResourceExt_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHResourceExtV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.StartingPRB_SBFD_r19 != nil, ie.SecondHopPRB_SBFD_r19 != nil}); err != nil {
		return err
	}

	// 2. startingPRB-SBFD-r19: ref
	{
		if ie.StartingPRB_SBFD_r19 != nil {
			if err := ie.StartingPRB_SBFD_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. secondHopPRB-SBFD-r19: ref
	{
		if ie.SecondHopPRB_SBFD_r19 != nil {
			if err := ie.SecondHopPRB_SBFD_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUCCH_ResourceExt_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHResourceExtV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. startingPRB-SBFD-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.StartingPRB_SBFD_r19 = new(PRB_Id)
			if err := ie.StartingPRB_SBFD_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. secondHopPRB-SBFD-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SecondHopPRB_SBFD_r19 = new(PRB_Id)
			if err := ie.SecondHopPRB_SBFD_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
