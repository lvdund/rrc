// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NCR-Parameters-r18 (line 22636).

var nCRParametersR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inactiveStateNCR-r18", Optional: true},
		{Name: "supportedNumberOfDRBs-NCR-r18", Optional: true},
		{Name: "dummy", Optional: true},
	},
}

const (
	NCR_Parameters_r18_InactiveStateNCR_r18_Supported = 0
)

var nCRParametersR18InactiveStateNCRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NCR_Parameters_r18_InactiveStateNCR_r18_Supported},
}

const (
	NCR_Parameters_r18_SupportedNumberOfDRBs_NCR_r18_N1  = 0
	NCR_Parameters_r18_SupportedNumberOfDRBs_NCR_r18_N16 = 1
)

var nCRParametersR18SupportedNumberOfDRBsNCRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NCR_Parameters_r18_SupportedNumberOfDRBs_NCR_r18_N1, NCR_Parameters_r18_SupportedNumberOfDRBs_NCR_r18_N16},
}

const (
	NCR_Parameters_r18_Dummy_Supported = 0
)

var nCRParametersR18DummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NCR_Parameters_r18_Dummy_Supported},
}

type NCR_Parameters_r18 struct {
	InactiveStateNCR_r18          *int64
	SupportedNumberOfDRBs_NCR_r18 *int64
	Dummy                         *int64
}

func (ie *NCR_Parameters_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nCRParametersR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InactiveStateNCR_r18 != nil, ie.SupportedNumberOfDRBs_NCR_r18 != nil, ie.Dummy != nil}); err != nil {
		return err
	}

	// 2. inactiveStateNCR-r18: enumerated
	{
		if ie.InactiveStateNCR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.InactiveStateNCR_r18, nCRParametersR18InactiveStateNCRR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. supportedNumberOfDRBs-NCR-r18: enumerated
	{
		if ie.SupportedNumberOfDRBs_NCR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SupportedNumberOfDRBs_NCR_r18, nCRParametersR18SupportedNumberOfDRBsNCRR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, nCRParametersR18DummyConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NCR_Parameters_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nCRParametersR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. inactiveStateNCR-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(nCRParametersR18InactiveStateNCRR18Constraints)
			if err != nil {
				return err
			}
			ie.InactiveStateNCR_r18 = &idx
		}
	}

	// 3. supportedNumberOfDRBs-NCR-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(nCRParametersR18SupportedNumberOfDRBsNCRR18Constraints)
			if err != nil {
				return err
			}
			ie.SupportedNumberOfDRBs_NCR_r18 = &idx
		}
	}

	// 4. dummy: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(nCRParametersR18DummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	return nil
}
