// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlinkPerCC-v1620 (line 19888).

var featureSetDownlinkPerCCV1620Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "multiDCI-MultiTRP-r16", Optional: true},
		{Name: "supportFDM-SchemeB-r16", Optional: true},
	},
}

const (
	FeatureSetDownlinkPerCC_v1620_SupportFDM_SchemeB_r16_Supported = 0
)

var featureSetDownlinkPerCCV1620SupportFDMSchemeBR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1620_SupportFDM_SchemeB_r16_Supported},
}

type FeatureSetDownlinkPerCC_v1620 struct {
	MultiDCI_MultiTRP_r16  *MultiDCI_MultiTRP_r16
	SupportFDM_SchemeB_r16 *int64
}

func (ie *FeatureSetDownlinkPerCC_v1620) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkPerCCV1620Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MultiDCI_MultiTRP_r16 != nil, ie.SupportFDM_SchemeB_r16 != nil}); err != nil {
		return err
	}

	// 2. multiDCI-MultiTRP-r16: ref
	{
		if ie.MultiDCI_MultiTRP_r16 != nil {
			if err := ie.MultiDCI_MultiTRP_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. supportFDM-SchemeB-r16: enumerated
	{
		if ie.SupportFDM_SchemeB_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SupportFDM_SchemeB_r16, featureSetDownlinkPerCCV1620SupportFDMSchemeBR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1620) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkPerCCV1620Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. multiDCI-MultiTRP-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MultiDCI_MultiTRP_r16 = new(MultiDCI_MultiTRP_r16)
			if err := ie.MultiDCI_MultiTRP_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. supportFDM-SchemeB-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1620SupportFDMSchemeBR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportFDM_SchemeB_r16 = &idx
		}
	}

	return nil
}
