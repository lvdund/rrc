// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SliceInfoDedicated-r17 (line 8365).

var sliceInfoDedicatedR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nsag-IdentityInfo-r17"},
		{Name: "nsag-CellReselectionPriority-r17", Optional: true},
		{Name: "nsag-CellReselectionSubPriority-r17", Optional: true},
	},
}

type SliceInfoDedicated_r17 struct {
	Nsag_IdentityInfo_r17               NSAG_IdentityInfo_r17
	Nsag_CellReselectionPriority_r17    *CellReselectionPriority
	Nsag_CellReselectionSubPriority_r17 *CellReselectionSubPriority
}

func (ie *SliceInfoDedicated_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sliceInfoDedicatedR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nsag_CellReselectionPriority_r17 != nil, ie.Nsag_CellReselectionSubPriority_r17 != nil}); err != nil {
		return err
	}

	// 2. nsag-IdentityInfo-r17: ref
	{
		if err := ie.Nsag_IdentityInfo_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. nsag-CellReselectionPriority-r17: ref
	{
		if ie.Nsag_CellReselectionPriority_r17 != nil {
			if err := ie.Nsag_CellReselectionPriority_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. nsag-CellReselectionSubPriority-r17: ref
	{
		if ie.Nsag_CellReselectionSubPriority_r17 != nil {
			if err := ie.Nsag_CellReselectionSubPriority_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SliceInfoDedicated_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sliceInfoDedicatedR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nsag-IdentityInfo-r17: ref
	{
		if err := ie.Nsag_IdentityInfo_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. nsag-CellReselectionPriority-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Nsag_CellReselectionPriority_r17 = new(CellReselectionPriority)
			if err := ie.Nsag_CellReselectionPriority_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nsag-CellReselectionSubPriority-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Nsag_CellReselectionSubPriority_r17 = new(CellReselectionSubPriority)
			if err := ie.Nsag_CellReselectionSubPriority_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
