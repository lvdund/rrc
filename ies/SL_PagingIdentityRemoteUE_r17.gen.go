// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PagingIdentityRemoteUE-r17 (line 27570).

var sLPagingIdentityRemoteUER17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ng-5G-S-TMSI-r17"},
		{Name: "fullI-RNTI-r17", Optional: true},
	},
}

type SL_PagingIdentityRemoteUE_r17 struct {
	Ng_5G_S_TMSI_r17 NG_5G_S_TMSI
	FullI_RNTI_r17   *I_RNTI_Value
}

func (ie *SL_PagingIdentityRemoteUE_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPagingIdentityRemoteUER17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FullI_RNTI_r17 != nil}); err != nil {
		return err
	}

	// 2. ng-5G-S-TMSI-r17: ref
	{
		if err := ie.Ng_5G_S_TMSI_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. fullI-RNTI-r17: ref
	{
		if ie.FullI_RNTI_r17 != nil {
			if err := ie.FullI_RNTI_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PagingIdentityRemoteUE_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPagingIdentityRemoteUER17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ng-5G-S-TMSI-r17: ref
	{
		if err := ie.Ng_5G_S_TMSI_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. fullI-RNTI-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FullI_RNTI_r17 = new(I_RNTI_Value)
			if err := ie.FullI_RNTI_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
