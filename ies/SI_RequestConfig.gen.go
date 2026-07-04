// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SI-RequestConfig (line 15015).

var sIRequestConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rach-OccasionsSI", Optional: true},
		{Name: "si-RequestPeriod", Optional: true},
		{Name: "si-RequestResources"},
	},
}

const (
	SI_RequestConfig_Si_RequestPeriod_One     = 0
	SI_RequestConfig_Si_RequestPeriod_Two     = 1
	SI_RequestConfig_Si_RequestPeriod_Four    = 2
	SI_RequestConfig_Si_RequestPeriod_Six     = 3
	SI_RequestConfig_Si_RequestPeriod_Eight   = 4
	SI_RequestConfig_Si_RequestPeriod_Ten     = 5
	SI_RequestConfig_Si_RequestPeriod_Twelve  = 6
	SI_RequestConfig_Si_RequestPeriod_Sixteen = 7
)

var sIRequestConfigSiRequestPeriodConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SI_RequestConfig_Si_RequestPeriod_One, SI_RequestConfig_Si_RequestPeriod_Two, SI_RequestConfig_Si_RequestPeriod_Four, SI_RequestConfig_Si_RequestPeriod_Six, SI_RequestConfig_Si_RequestPeriod_Eight, SI_RequestConfig_Si_RequestPeriod_Ten, SI_RequestConfig_Si_RequestPeriod_Twelve, SI_RequestConfig_Si_RequestPeriod_Sixteen},
}

var sIRequestConfigSiRequestResourcesConstraints = per.SizeRange(1, common.MaxSI_Message)

const (
	SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_OneEighth = 0
	SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_OneFourth = 1
	SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_OneHalf   = 2
	SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_One       = 3
	SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_Two       = 4
	SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_Four      = 5
	SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_Eight     = 6
	SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_Sixteen   = 7
)

var sIRequestConfigRachOccasionsSISsbPerRACHOccasionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_OneEighth, SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_OneFourth, SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_OneHalf, SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_One, SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_Two, SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_Four, SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_Eight, SI_RequestConfig_Rach_OccasionsSI_Ssb_PerRACH_Occasion_Sixteen},
}

type SI_RequestConfig struct {
	Rach_OccasionsSI *struct {
		Rach_ConfigSI        RACH_ConfigGeneric
		Ssb_PerRACH_Occasion int64
	}
	Si_RequestPeriod    *int64
	Si_RequestResources []SI_RequestResources
}

func (ie *SI_RequestConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIRequestConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rach_OccasionsSI != nil, ie.Si_RequestPeriod != nil}); err != nil {
		return err
	}

	// 2. rach-OccasionsSI: sequence
	{
		if ie.Rach_OccasionsSI != nil {
			c := ie.Rach_OccasionsSI
			if err := c.Rach_ConfigSI.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Ssb_PerRACH_Occasion, sIRequestConfigRachOccasionsSISsbPerRACHOccasionConstraints); err != nil {
				return err
			}
		}
	}

	// 3. si-RequestPeriod: enumerated
	{
		if ie.Si_RequestPeriod != nil {
			if err := e.EncodeEnumerated(*ie.Si_RequestPeriod, sIRequestConfigSiRequestPeriodConstraints); err != nil {
				return err
			}
		}
	}

	// 4. si-RequestResources: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sIRequestConfigSiRequestResourcesConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Si_RequestResources))); err != nil {
			return err
		}
		for i := range ie.Si_RequestResources {
			if err := ie.Si_RequestResources[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SI_RequestConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIRequestConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rach-OccasionsSI: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Rach_OccasionsSI = &struct {
				Rach_ConfigSI        RACH_ConfigGeneric
				Ssb_PerRACH_Occasion int64
			}{}
			c := ie.Rach_OccasionsSI
			{
				if err := c.Rach_ConfigSI.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeEnumerated(sIRequestConfigRachOccasionsSISsbPerRACHOccasionConstraints)
				if err != nil {
					return err
				}
				c.Ssb_PerRACH_Occasion = v
			}
		}
	}

	// 3. si-RequestPeriod: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sIRequestConfigSiRequestPeriodConstraints)
			if err != nil {
				return err
			}
			ie.Si_RequestPeriod = &idx
		}
	}

	// 4. si-RequestResources: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sIRequestConfigSiRequestResourcesConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Si_RequestResources = make([]SI_RequestResources, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Si_RequestResources[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
