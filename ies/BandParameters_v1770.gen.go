// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParameters-v1770 (line 17081).

var bandParametersV1770Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-BandwidthClassDL-NR-r17", Optional: true},
		{Name: "ca-BandwidthClassUL-NR-r17", Optional: true},
	},
}

type BandParameters_v1770 struct {
	Ca_BandwidthClassDL_NR_r17 *CA_BandwidthClassNR_r17
	Ca_BandwidthClassUL_NR_r17 *CA_BandwidthClassNR_r17
}

func (ie *BandParameters_v1770) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersV1770Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_BandwidthClassDL_NR_r17 != nil, ie.Ca_BandwidthClassUL_NR_r17 != nil}); err != nil {
		return err
	}

	// 2. ca-BandwidthClassDL-NR-r17: ref
	{
		if ie.Ca_BandwidthClassDL_NR_r17 != nil {
			if err := ie.Ca_BandwidthClassDL_NR_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-BandwidthClassUL-NR-r17: ref
	{
		if ie.Ca_BandwidthClassUL_NR_r17 != nil {
			if err := ie.Ca_BandwidthClassUL_NR_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandParameters_v1770) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersV1770Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-BandwidthClassDL-NR-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_BandwidthClassDL_NR_r17 = new(CA_BandwidthClassNR_r17)
			if err := ie.Ca_BandwidthClassDL_NR_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-BandwidthClassUL-NR-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_BandwidthClassUL_NR_r17 = new(CA_BandwidthClassNR_r17)
			if err := ie.Ca_BandwidthClassUL_NR_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
