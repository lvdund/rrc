// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v17d0 (line 19649).

var featureSetDownlinkV17d0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-AntennaSwitching2SP-1Periodic-r17", Optional: true},
		{Name: "srs-ExtensionAperiodicSRS-r17", Optional: true},
		{Name: "srs-OneAP-SRS-r17", Optional: true},
	},
}

const (
	FeatureSetDownlink_V17d0_Srs_AntennaSwitching2SP_1Periodic_r17_Supported = 0
)

var featureSetDownlinkV17d0SrsAntennaSwitching2SP1PeriodicR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V17d0_Srs_AntennaSwitching2SP_1Periodic_r17_Supported},
}

const (
	FeatureSetDownlink_V17d0_Srs_ExtensionAperiodicSRS_r17_Supported = 0
)

var featureSetDownlinkV17d0SrsExtensionAperiodicSRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V17d0_Srs_ExtensionAperiodicSRS_r17_Supported},
}

const (
	FeatureSetDownlink_V17d0_Srs_OneAP_SRS_r17_Supported = 0
)

var featureSetDownlinkV17d0SrsOneAPSRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V17d0_Srs_OneAP_SRS_r17_Supported},
}

type FeatureSetDownlink_V17d0 struct {
	Srs_AntennaSwitching2SP_1Periodic_r17 *int64
	Srs_ExtensionAperiodicSRS_r17         *int64
	Srs_OneAP_SRS_r17                     *int64
}

func (ie *FeatureSetDownlink_V17d0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV17d0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_AntennaSwitching2SP_1Periodic_r17 != nil, ie.Srs_ExtensionAperiodicSRS_r17 != nil, ie.Srs_OneAP_SRS_r17 != nil}); err != nil {
		return err
	}

	// 2. srs-AntennaSwitching2SP-1Periodic-r17: enumerated
	{
		if ie.Srs_AntennaSwitching2SP_1Periodic_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_AntennaSwitching2SP_1Periodic_r17, featureSetDownlinkV17d0SrsAntennaSwitching2SP1PeriodicR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. srs-ExtensionAperiodicSRS-r17: enumerated
	{
		if ie.Srs_ExtensionAperiodicSRS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_ExtensionAperiodicSRS_r17, featureSetDownlinkV17d0SrsExtensionAperiodicSRSR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. srs-OneAP-SRS-r17: enumerated
	{
		if ie.Srs_OneAP_SRS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_OneAP_SRS_r17, featureSetDownlinkV17d0SrsOneAPSRSR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_V17d0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV17d0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-AntennaSwitching2SP-1Periodic-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV17d0SrsAntennaSwitching2SP1PeriodicR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_AntennaSwitching2SP_1Periodic_r17 = &idx
		}
	}

	// 3. srs-ExtensionAperiodicSRS-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV17d0SrsExtensionAperiodicSRSR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_ExtensionAperiodicSRS_r17 = &idx
		}
	}

	// 4. srs-OneAP-SRS-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV17d0SrsOneAPSRSR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_OneAP_SRS_r17 = &idx
		}
	}

	return nil
}
