// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ReservationPeriodAllowedDedicatedSL-PRS-RP-r18 (line 27658).

var sLReservationPeriodAllowedDedicatedSLPRSRPR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-ResourceReservePeriod1-r18"},
		{Name: "sl-ResourceReservePeriod2-r18"},
	},
}

const (
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18 = 0
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod2_r18 = 1
)

const (
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms0     = 0
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms100   = 1
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms160   = 2
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms200   = 3
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms300   = 4
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms320   = 5
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms400   = 6
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms500   = 7
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms600   = 8
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms640   = 9
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms700   = 10
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms800   = 11
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms900   = 12
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms1000  = 13
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms1280  = 14
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms2560  = 15
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms5120  = 16
	SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms10240 = 17
)

var sLReservationPeriodAllowedDedicatedSLPRSRPR18SlResourceReservePeriod1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms0, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms100, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms160, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms200, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms300, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms320, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms400, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms500, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms600, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms640, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms700, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms800, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms900, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms1000, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms1280, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms2560, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms5120, SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18_Ms10240},
}

type SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18 struct {
	Choice                        int
	Sl_ResourceReservePeriod1_r18 *int64
	Sl_ResourceReservePeriod2_r18 *int64
}

func (ie *SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLReservationPeriodAllowedDedicatedSLPRSRPR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18:
		if err := e.EncodeEnumerated((*ie.Sl_ResourceReservePeriod1_r18), sLReservationPeriodAllowedDedicatedSLPRSRPR18SlResourceReservePeriod1R18Constraints); err != nil {
			return err
		}
	case SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod2_r18:
		if err := e.EncodeInteger((*ie.Sl_ResourceReservePeriod2_r18), per.Constrained(1, 99)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SL-ReservationPeriodAllowedDedicatedSL-PRS-RP-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLReservationPeriodAllowedDedicatedSLPRSRPR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-ReservationPeriodAllowedDedicatedSL-PRS-RP-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod1_r18:
		ie.Sl_ResourceReservePeriod1_r18 = new(int64)
		v, err := d.DecodeEnumerated(sLReservationPeriodAllowedDedicatedSLPRSRPR18SlResourceReservePeriod1R18Constraints)
		if err != nil {
			return err
		}
		(*ie.Sl_ResourceReservePeriod1_r18) = v
	case SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18_Sl_ResourceReservePeriod2_r18:
		ie.Sl_ResourceReservePeriod2_r18 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(1, 99))
		if err != nil {
			return err
		}
		(*ie.Sl_ResourceReservePeriod2_r18) = v
	default:
		return &per.DecodeError{TypeName: "SL-ReservationPeriodAllowedDedicatedSL-PRS-RP-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
