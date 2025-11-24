package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_CandidateList_CriticalExtensions_C1_Choice_nothing uint64 = iota
	CG_CandidateList_CriticalExtensions_C1_Choice_Cg_CandidateList_r17
	CG_CandidateList_CriticalExtensions_C1_Choice_Spare3
	CG_CandidateList_CriticalExtensions_C1_Choice_Spare2
	CG_CandidateList_CriticalExtensions_C1_Choice_Spare1
)

type CG_CandidateList_CriticalExtensions_C1 struct {
	Choice               uint64
	Cg_CandidateList_r17 *CG_CandidateList_r17_IEs
	Spare3               uper.NULL `madatory`
	Spare2               uper.NULL `madatory`
	Spare1               uper.NULL `madatory`
}

func (ie *CG_CandidateList_CriticalExtensions_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_CandidateList_CriticalExtensions_C1_Choice_Cg_CandidateList_r17:
		if err = ie.Cg_CandidateList_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Cg_CandidateList_r17", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_CandidateList_CriticalExtensions_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_CandidateList_CriticalExtensions_C1_Choice_Cg_CandidateList_r17:
		ie.Cg_CandidateList_r17 = new(CG_CandidateList_r17_IEs)
		if err = ie.Cg_CandidateList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_CandidateList_r17", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
