// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NRDC-Parameters-v1570 (line 22656).

var nRDCParametersV1570Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sfn-SyncNRDC", Optional: true},
	},
}

const (
	NRDC_Parameters_v1570_Sfn_SyncNRDC_Supported = 0
)

var nRDCParametersV1570SfnSyncNRDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NRDC_Parameters_v1570_Sfn_SyncNRDC_Supported},
}

type NRDC_Parameters_v1570 struct {
	Sfn_SyncNRDC *int64
}

func (ie *NRDC_Parameters_v1570) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRDCParametersV1570Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sfn_SyncNRDC != nil}); err != nil {
		return err
	}

	// 2. sfn-SyncNRDC: enumerated
	{
		if ie.Sfn_SyncNRDC != nil {
			if err := e.EncodeEnumerated(*ie.Sfn_SyncNRDC, nRDCParametersV1570SfnSyncNRDCConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NRDC_Parameters_v1570) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRDCParametersV1570Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sfn-SyncNRDC: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(nRDCParametersV1570SfnSyncNRDCConstraints)
			if err != nil {
				return err
			}
			ie.Sfn_SyncNRDC = &idx
		}
	}

	return nil
}
