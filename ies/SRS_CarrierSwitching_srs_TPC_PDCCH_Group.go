package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_nothing uint64 = iota
	SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_TypeA
	SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_TypeB
)

type SRS_CarrierSwitching_srs_TPC_PDCCH_Group struct {
	Choice uint64
	TypeA  []SRS_TPC_PDCCH_Config `lb:1,ub:32,madatory`
	TypeB  *SRS_TPC_PDCCH_Config
}

func (ie *SRS_CarrierSwitching_srs_TPC_PDCCH_Group) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_TypeA:
		tmp := utils.NewSequence[*SRS_TPC_PDCCH_Config]([]*SRS_TPC_PDCCH_Config{}, uper.Constraint{Lb: 1, Ub: 32}, false)
		for _, i := range ie.TypeA {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode TypeA", err)
		}
	case SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_TypeB:
		if err = ie.TypeB.Encode(w); err != nil {
			err = utils.WrapError("Encode TypeB", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_CarrierSwitching_srs_TPC_PDCCH_Group) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_TypeA:
		tmp := utils.NewSequence[*SRS_TPC_PDCCH_Config]([]*SRS_TPC_PDCCH_Config{}, uper.Constraint{Lb: 1, Ub: 32}, false)
		fn := func() *SRS_TPC_PDCCH_Config {
			return new(SRS_TPC_PDCCH_Config)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode TypeA", err)
		}
		ie.TypeA = []SRS_TPC_PDCCH_Config{}
		for _, i := range tmp.Value {
			ie.TypeA = append(ie.TypeA, *i)
		}
	case SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_TypeB:
		ie.TypeB = new(SRS_TPC_PDCCH_Config)
		if err = ie.TypeB.Decode(r); err != nil {
			return utils.WrapError("Decode TypeB", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
