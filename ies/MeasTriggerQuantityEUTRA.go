package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasTriggerQuantityEUTRA_Choice_nothing uint64 = iota
	MeasTriggerQuantityEUTRA_Choice_Rsrp
	MeasTriggerQuantityEUTRA_Choice_Rsrq
	MeasTriggerQuantityEUTRA_Choice_Sinr
)

type MeasTriggerQuantityEUTRA struct {
	Choice uint64
	Rsrp   *RSRP_RangeEUTRA
	Rsrq   *RSRQ_RangeEUTRA
	Sinr   *SINR_RangeEUTRA
}

func (ie *MeasTriggerQuantityEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityEUTRA_Choice_Rsrp:
		if err = ie.Rsrp.Encode(w); err != nil {
			err = utils.WrapError("Encode Rsrp", err)
		}
	case MeasTriggerQuantityEUTRA_Choice_Rsrq:
		if err = ie.Rsrq.Encode(w); err != nil {
			err = utils.WrapError("Encode Rsrq", err)
		}
	case MeasTriggerQuantityEUTRA_Choice_Sinr:
		if err = ie.Sinr.Encode(w); err != nil {
			err = utils.WrapError("Encode Sinr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasTriggerQuantityEUTRA) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityEUTRA_Choice_Rsrp:
		ie.Rsrp = new(RSRP_RangeEUTRA)
		if err = ie.Rsrp.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp", err)
		}
	case MeasTriggerQuantityEUTRA_Choice_Rsrq:
		ie.Rsrq = new(RSRQ_RangeEUTRA)
		if err = ie.Rsrq.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrq", err)
		}
	case MeasTriggerQuantityEUTRA_Choice_Sinr:
		ie.Sinr = new(SINR_RangeEUTRA)
		if err = ie.Sinr.Decode(r); err != nil {
			return utils.WrapError("Decode Sinr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
