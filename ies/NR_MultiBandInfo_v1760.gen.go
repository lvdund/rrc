// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-MultiBandInfo-v1760 (line 10179).

var nRMultiBandInfoV1760Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nr-NS-PmaxList-v1760", Optional: true},
	},
}

type NR_MultiBandInfo_v1760 struct {
	Nr_NS_PmaxList_v1760 *NR_NS_PmaxList_v1760
}

func (ie *NR_MultiBandInfo_v1760) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRMultiBandInfoV1760Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nr_NS_PmaxList_v1760 != nil}); err != nil {
		return err
	}

	// 2. nr-NS-PmaxList-v1760: ref
	{
		if ie.Nr_NS_PmaxList_v1760 != nil {
			if err := ie.Nr_NS_PmaxList_v1760.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NR_MultiBandInfo_v1760) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRMultiBandInfoV1760Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nr-NS-PmaxList-v1760: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Nr_NS_PmaxList_v1760 = new(NR_NS_PmaxList_v1760)
			if err := ie.Nr_NS_PmaxList_v1760.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
