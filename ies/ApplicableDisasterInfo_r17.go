package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ApplicableDisasterInfo_r17_Choice_nothing uint64 = iota
	ApplicableDisasterInfo_r17_Choice_NoDisasterRoaming_r17
	ApplicableDisasterInfo_r17_Choice_DisasterRelatedIndication_r17
	ApplicableDisasterInfo_r17_Choice_CommonPLMNs_r17
	ApplicableDisasterInfo_r17_Choice_DedicatedPLMNs_r17
)

type ApplicableDisasterInfo_r17 struct {
	Choice                        uint64
	NoDisasterRoaming_r17         aper.NULL       `madatory`
	DisasterRelatedIndication_r17 aper.NULL       `madatory`
	CommonPLMNs_r17               aper.NULL       `madatory`
	DedicatedPLMNs_r17            []PLMN_Identity `lb:1,ub:maxPLMN,madatory`
}

func (ie *ApplicableDisasterInfo_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ApplicableDisasterInfo_r17_Choice_NoDisasterRoaming_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode NoDisasterRoaming_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_DisasterRelatedIndication_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode DisasterRelatedIndication_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_CommonPLMNs_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode CommonPLMNs_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_DedicatedPLMNs_r17:
		tmp := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.DedicatedPLMNs_r17 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode DedicatedPLMNs_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ApplicableDisasterInfo_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ApplicableDisasterInfo_r17_Choice_NoDisasterRoaming_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode NoDisasterRoaming_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_DisasterRelatedIndication_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode DisasterRelatedIndication_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_CommonPLMNs_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode CommonPLMNs_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_DedicatedPLMNs_r17:
		tmp := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn := func() *PLMN_Identity {
			return new(PLMN_Identity)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode DedicatedPLMNs_r17", err)
		}
		ie.DedicatedPLMNs_r17 = []PLMN_Identity{}
		for _, i := range tmp.Value {
			ie.DedicatedPLMNs_r17 = append(ie.DedicatedPLMNs_r17, *i)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
