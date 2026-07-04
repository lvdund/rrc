// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBS-RNTI-SpecificConfig-r17 (line 9057).

var mBSRNTISpecificConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-RNTI-SpecificConfigId-r17"},
		{Name: "groupCommon-RNTI-r17"},
		{Name: "drx-ConfigPTM-r17", Optional: true},
		{Name: "harq-FeedbackEnablerMulticast-r17", Optional: true},
		{Name: "harq-FeedbackOptionMulticast-r17", Optional: true},
		{Name: "pdsch-AggregationFactor-r17", Optional: true},
	},
}

var mBS_RNTI_SpecificConfig_r17GroupCommonRNTIR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "g-RNTI"},
		{Name: "g-CS-RNTI"},
	},
}

const (
	MBS_RNTI_SpecificConfig_r17_GroupCommon_RNTI_r17_G_RNTI    = 0
	MBS_RNTI_SpecificConfig_r17_GroupCommon_RNTI_r17_G_CS_RNTI = 1
)

var mBS_RNTI_SpecificConfig_r17DrxConfigPTMR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MBS_RNTI_SpecificConfig_r17_Drx_ConfigPTM_r17_Release = 0
	MBS_RNTI_SpecificConfig_r17_Drx_ConfigPTM_r17_Setup   = 1
)

const (
	MBS_RNTI_SpecificConfig_r17_Harq_FeedbackEnablerMulticast_r17_Dci_Enabler = 0
	MBS_RNTI_SpecificConfig_r17_Harq_FeedbackEnablerMulticast_r17_Enabled     = 1
)

var mBSRNTISpecificConfigR17HarqFeedbackEnablerMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MBS_RNTI_SpecificConfig_r17_Harq_FeedbackEnablerMulticast_r17_Dci_Enabler, MBS_RNTI_SpecificConfig_r17_Harq_FeedbackEnablerMulticast_r17_Enabled},
}

const (
	MBS_RNTI_SpecificConfig_r17_Harq_FeedbackOptionMulticast_r17_Ack_Nack  = 0
	MBS_RNTI_SpecificConfig_r17_Harq_FeedbackOptionMulticast_r17_Nack_Only = 1
)

var mBSRNTISpecificConfigR17HarqFeedbackOptionMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MBS_RNTI_SpecificConfig_r17_Harq_FeedbackOptionMulticast_r17_Ack_Nack, MBS_RNTI_SpecificConfig_r17_Harq_FeedbackOptionMulticast_r17_Nack_Only},
}

const (
	MBS_RNTI_SpecificConfig_r17_Pdsch_AggregationFactor_r17_N2 = 0
	MBS_RNTI_SpecificConfig_r17_Pdsch_AggregationFactor_r17_N4 = 1
	MBS_RNTI_SpecificConfig_r17_Pdsch_AggregationFactor_r17_N8 = 2
)

var mBSRNTISpecificConfigR17PdschAggregationFactorR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MBS_RNTI_SpecificConfig_r17_Pdsch_AggregationFactor_r17_N2, MBS_RNTI_SpecificConfig_r17_Pdsch_AggregationFactor_r17_N4, MBS_RNTI_SpecificConfig_r17_Pdsch_AggregationFactor_r17_N8},
}

type MBS_RNTI_SpecificConfig_r17 struct {
	Mbs_RNTI_SpecificConfigId_r17 MBS_RNTI_SpecificConfigId_r17
	GroupCommon_RNTI_r17          struct {
		Choice    int
		G_RNTI    *RNTI_Value
		G_CS_RNTI *RNTI_Value
	}
	Drx_ConfigPTM_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *DRX_ConfigPTM_r17
	}
	Harq_FeedbackEnablerMulticast_r17 *int64
	Harq_FeedbackOptionMulticast_r17  *int64
	Pdsch_AggregationFactor_r17       *int64
}

func (ie *MBS_RNTI_SpecificConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSRNTISpecificConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Drx_ConfigPTM_r17 != nil, ie.Harq_FeedbackEnablerMulticast_r17 != nil, ie.Harq_FeedbackOptionMulticast_r17 != nil, ie.Pdsch_AggregationFactor_r17 != nil}); err != nil {
		return err
	}

	// 2. mbs-RNTI-SpecificConfigId-r17: ref
	{
		if err := ie.Mbs_RNTI_SpecificConfigId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. groupCommon-RNTI-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(mBS_RNTI_SpecificConfig_r17GroupCommonRNTIR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.GroupCommon_RNTI_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.GroupCommon_RNTI_r17.Choice {
		case MBS_RNTI_SpecificConfig_r17_GroupCommon_RNTI_r17_G_RNTI:
			if err := ie.GroupCommon_RNTI_r17.G_RNTI.Encode(e); err != nil {
				return err
			}
		case MBS_RNTI_SpecificConfig_r17_GroupCommon_RNTI_r17_G_CS_RNTI:
			if err := ie.GroupCommon_RNTI_r17.G_CS_RNTI.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.GroupCommon_RNTI_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 4. drx-ConfigPTM-r17: choice
	{
		if ie.Drx_ConfigPTM_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(mBS_RNTI_SpecificConfig_r17DrxConfigPTMR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Drx_ConfigPTM_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Drx_ConfigPTM_r17).Choice {
			case MBS_RNTI_SpecificConfig_r17_Drx_ConfigPTM_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case MBS_RNTI_SpecificConfig_r17_Drx_ConfigPTM_r17_Setup:
				if err := (*ie.Drx_ConfigPTM_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Drx_ConfigPTM_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. harq-FeedbackEnablerMulticast-r17: enumerated
	{
		if ie.Harq_FeedbackEnablerMulticast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Harq_FeedbackEnablerMulticast_r17, mBSRNTISpecificConfigR17HarqFeedbackEnablerMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. harq-FeedbackOptionMulticast-r17: enumerated
	{
		if ie.Harq_FeedbackOptionMulticast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Harq_FeedbackOptionMulticast_r17, mBSRNTISpecificConfigR17HarqFeedbackOptionMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. pdsch-AggregationFactor-r17: enumerated
	{
		if ie.Pdsch_AggregationFactor_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_AggregationFactor_r17, mBSRNTISpecificConfigR17PdschAggregationFactorR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MBS_RNTI_SpecificConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSRNTISpecificConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mbs-RNTI-SpecificConfigId-r17: ref
	{
		if err := ie.Mbs_RNTI_SpecificConfigId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. groupCommon-RNTI-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(mBS_RNTI_SpecificConfig_r17GroupCommonRNTIR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.GroupCommon_RNTI_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MBS_RNTI_SpecificConfig_r17_GroupCommon_RNTI_r17_G_RNTI:
			ie.GroupCommon_RNTI_r17.G_RNTI = new(RNTI_Value)
			if err := ie.GroupCommon_RNTI_r17.G_RNTI.Decode(d); err != nil {
				return err
			}
		case MBS_RNTI_SpecificConfig_r17_GroupCommon_RNTI_r17_G_CS_RNTI:
			ie.GroupCommon_RNTI_r17.G_CS_RNTI = new(RNTI_Value)
			if err := ie.GroupCommon_RNTI_r17.G_CS_RNTI.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. drx-ConfigPTM-r17: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Drx_ConfigPTM_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DRX_ConfigPTM_r17
			}{}
			choiceDec := d.NewChoiceDecoder(mBS_RNTI_SpecificConfig_r17DrxConfigPTMR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Drx_ConfigPTM_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MBS_RNTI_SpecificConfig_r17_Drx_ConfigPTM_r17_Release:
				(*ie.Drx_ConfigPTM_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case MBS_RNTI_SpecificConfig_r17_Drx_ConfigPTM_r17_Setup:
				(*ie.Drx_ConfigPTM_r17).Setup = new(DRX_ConfigPTM_r17)
				if err := (*ie.Drx_ConfigPTM_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. harq-FeedbackEnablerMulticast-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(mBSRNTISpecificConfigR17HarqFeedbackEnablerMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Harq_FeedbackEnablerMulticast_r17 = &idx
		}
	}

	// 6. harq-FeedbackOptionMulticast-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(mBSRNTISpecificConfigR17HarqFeedbackOptionMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Harq_FeedbackOptionMulticast_r17 = &idx
		}
	}

	// 7. pdsch-AggregationFactor-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(mBSRNTISpecificConfigR17PdschAggregationFactorR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_AggregationFactor_r17 = &idx
		}
	}

	return nil
}
