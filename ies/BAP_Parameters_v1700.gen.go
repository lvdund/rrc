// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BAP-Parameters-v1700 (line 25937).

var bAPParametersV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bapHeaderRewriting-Rerouting-r17", Optional: true},
		{Name: "bapHeaderRewriting-Routing-r17", Optional: true},
	},
}

const (
	BAP_Parameters_v1700_BapHeaderRewriting_Rerouting_r17_Supported = 0
)

var bAPParametersV1700BapHeaderRewritingReroutingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BAP_Parameters_v1700_BapHeaderRewriting_Rerouting_r17_Supported},
}

const (
	BAP_Parameters_v1700_BapHeaderRewriting_Routing_r17_Supported = 0
)

var bAPParametersV1700BapHeaderRewritingRoutingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BAP_Parameters_v1700_BapHeaderRewriting_Routing_r17_Supported},
}

type BAP_Parameters_v1700 struct {
	BapHeaderRewriting_Rerouting_r17 *int64
	BapHeaderRewriting_Routing_r17   *int64
}

func (ie *BAP_Parameters_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bAPParametersV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BapHeaderRewriting_Rerouting_r17 != nil, ie.BapHeaderRewriting_Routing_r17 != nil}); err != nil {
		return err
	}

	// 2. bapHeaderRewriting-Rerouting-r17: enumerated
	{
		if ie.BapHeaderRewriting_Rerouting_r17 != nil {
			if err := e.EncodeEnumerated(*ie.BapHeaderRewriting_Rerouting_r17, bAPParametersV1700BapHeaderRewritingReroutingR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. bapHeaderRewriting-Routing-r17: enumerated
	{
		if ie.BapHeaderRewriting_Routing_r17 != nil {
			if err := e.EncodeEnumerated(*ie.BapHeaderRewriting_Routing_r17, bAPParametersV1700BapHeaderRewritingRoutingR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BAP_Parameters_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bAPParametersV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bapHeaderRewriting-Rerouting-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(bAPParametersV1700BapHeaderRewritingReroutingR17Constraints)
			if err != nil {
				return err
			}
			ie.BapHeaderRewriting_Rerouting_r17 = &idx
		}
	}

	// 3. bapHeaderRewriting-Routing-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(bAPParametersV1700BapHeaderRewritingRoutingR17Constraints)
			if err != nil {
				return err
			}
			ie.BapHeaderRewriting_Routing_r17 = &idx
		}
	}

	return nil
}
