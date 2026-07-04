// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ConfiguredGrantConfigDedicatedSL-PRS-RP-r18 (line 27095).

var sLConfiguredGrantConfigDedicatedSLPRSRPR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-ConfigIndexCG-r18"},
		{Name: "sl-PRS-PeriodCG-r18", Optional: true},
		{Name: "sl-PRS-ResourcePoolID-r18", Optional: true},
		{Name: "rrc-ConfiguredSidelinkGrantDedicatedSL-PRS-RP-r18"},
	},
}

var sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TimeOffsetCG-Type1-r18", Optional: true},
		{Name: "sl-TimeReferenceSFN-Type1-r18", Optional: true},
		{Name: "sl-TimeResourceCG-Type1-r18", Optional: true},
		{Name: "sl-PRS-ResourceIndicationFirstType1-r18", Optional: true},
		{Name: "sl-PRS-ResourceIndicationFutureType1-r18", Optional: true},
	},
}

const (
	SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_r18_Rrc_ConfiguredSidelinkGrantDedicatedSL_PRS_RP_r18_Sl_TimeReferenceSFN_Type1_r18_Sfn512 = 0
)

var sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18SlTimeReferenceSFNType1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_r18_Rrc_ConfiguredSidelinkGrantDedicatedSL_PRS_RP_r18_Sl_TimeReferenceSFN_Type1_r18_Sfn512},
}

type SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_r18 struct {
	Sl_PRS_ConfigIndexCG_r18                          SL_ConfigIndexCG_r16
	Sl_PRS_PeriodCG_r18                               *SL_PeriodCG_r16
	Sl_PRS_ResourcePoolID_r18                         *SL_ResourcePoolID_r16
	Rrc_ConfiguredSidelinkGrantDedicatedSL_PRS_RP_r18 struct {
		Sl_TimeOffsetCG_Type1_r18                *int64
		Sl_TimeReferenceSFN_Type1_r18            *int64
		Sl_TimeResourceCG_Type1_r18              *int64
		Sl_PRS_ResourceIndicationFirstType1_r18  *int64
		Sl_PRS_ResourceIndicationFutureType1_r18 *int64
	}
}

func (ie *SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLConfiguredGrantConfigDedicatedSLPRSRPR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_PeriodCG_r18 != nil, ie.Sl_PRS_ResourcePoolID_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PRS-ConfigIndexCG-r18: ref
	{
		if err := ie.Sl_PRS_ConfigIndexCG_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-PRS-PeriodCG-r18: ref
	{
		if ie.Sl_PRS_PeriodCG_r18 != nil {
			if err := ie.Sl_PRS_PeriodCG_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-PRS-ResourcePoolID-r18: ref
	{
		if ie.Sl_PRS_ResourcePoolID_r18 != nil {
			if err := ie.Sl_PRS_ResourcePoolID_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. rrc-ConfiguredSidelinkGrantDedicatedSL-PRS-RP-r18: sequence
	{
		{
			c := &ie.Rrc_ConfiguredSidelinkGrantDedicatedSL_PRS_RP_r18
			sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Seq := e.NewSequenceEncoder(sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Constraints)
			if err := sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Seq.EncodePreamble([]bool{c.Sl_TimeOffsetCG_Type1_r18 != nil, c.Sl_TimeReferenceSFN_Type1_r18 != nil, c.Sl_TimeResourceCG_Type1_r18 != nil, c.Sl_PRS_ResourceIndicationFirstType1_r18 != nil, c.Sl_PRS_ResourceIndicationFutureType1_r18 != nil}); err != nil {
				return err
			}
			if c.Sl_TimeOffsetCG_Type1_r18 != nil {
				if err := e.EncodeInteger((*c.Sl_TimeOffsetCG_Type1_r18), per.Constrained(0, 7999)); err != nil {
					return err
				}
			}
			if c.Sl_TimeReferenceSFN_Type1_r18 != nil {
				if err := e.EncodeEnumerated((*c.Sl_TimeReferenceSFN_Type1_r18), sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18SlTimeReferenceSFNType1R18Constraints); err != nil {
					return err
				}
			}
			if c.Sl_TimeResourceCG_Type1_r18 != nil {
				if err := e.EncodeInteger((*c.Sl_TimeResourceCG_Type1_r18), per.Constrained(0, 496)); err != nil {
					return err
				}
			}
			if c.Sl_PRS_ResourceIndicationFirstType1_r18 != nil {
				if err := e.EncodeInteger((*c.Sl_PRS_ResourceIndicationFirstType1_r18), per.Constrained(0, 11)); err != nil {
					return err
				}
			}
			if c.Sl_PRS_ResourceIndicationFutureType1_r18 != nil {
				if err := e.EncodeInteger((*c.Sl_PRS_ResourceIndicationFutureType1_r18), per.Constrained(0, 143)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLConfiguredGrantConfigDedicatedSLPRSRPR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PRS-ConfigIndexCG-r18: ref
	{
		if err := ie.Sl_PRS_ConfigIndexCG_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-PRS-PeriodCG-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_PRS_PeriodCG_r18 = new(SL_PeriodCG_r16)
			if err := ie.Sl_PRS_PeriodCG_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-PRS-ResourcePoolID-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_PRS_ResourcePoolID_r18 = new(SL_ResourcePoolID_r16)
			if err := ie.Sl_PRS_ResourcePoolID_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. rrc-ConfiguredSidelinkGrantDedicatedSL-PRS-RP-r18: sequence
	{
		{
			c := &ie.Rrc_ConfiguredSidelinkGrantDedicatedSL_PRS_RP_r18
			sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Seq := d.NewSequenceDecoder(sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Constraints)
			if err := sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Seq.IsComponentPresent(0) {
				c.Sl_TimeOffsetCG_Type1_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 7999))
				if err != nil {
					return err
				}
				(*c.Sl_TimeOffsetCG_Type1_r18) = v
			}
			if sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Seq.IsComponentPresent(1) {
				c.Sl_TimeReferenceSFN_Type1_r18 = new(int64)
				v, err := d.DecodeEnumerated(sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18SlTimeReferenceSFNType1R18Constraints)
				if err != nil {
					return err
				}
				(*c.Sl_TimeReferenceSFN_Type1_r18) = v
			}
			if sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Seq.IsComponentPresent(2) {
				c.Sl_TimeResourceCG_Type1_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 496))
				if err != nil {
					return err
				}
				(*c.Sl_TimeResourceCG_Type1_r18) = v
			}
			if sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Seq.IsComponentPresent(3) {
				c.Sl_PRS_ResourceIndicationFirstType1_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 11))
				if err != nil {
					return err
				}
				(*c.Sl_PRS_ResourceIndicationFirstType1_r18) = v
			}
			if sLConfiguredGrantConfigDedicatedSLPRSRPR18RrcConfiguredSidelinkGrantDedicatedSLPRSRPR18Seq.IsComponentPresent(4) {
				c.Sl_PRS_ResourceIndicationFutureType1_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 143))
				if err != nil {
					return err
				}
				(*c.Sl_PRS_ResourceIndicationFutureType1_r18) = v
			}
		}
	}

	return nil
}
