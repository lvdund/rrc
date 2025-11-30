package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasTriggerQuantity_Choice_nothing uint64 = iota
	MeasTriggerQuantity_Choice_Rsrp
	MeasTriggerQuantity_Choice_Rsrq
	MeasTriggerQuantity_Choice_Sinr
)

type MeasTriggerQuantity struct {
	Choice uint64
	Rsrp   *RSRP_Range
	Rsrq   *RSRQ_Range
	Sinr   *SINR_Range
}

func (ie *MeasTriggerQuantity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantity_Choice_Rsrp:
		if err = ie.Rsrp.Encode(w); err != nil {
			err = utils.WrapError("Encode Rsrp", err)
		}
	case MeasTriggerQuantity_Choice_Rsrq:
		if err = ie.Rsrq.Encode(w); err != nil {
			err = utils.WrapError("Encode Rsrq", err)
		}
	case MeasTriggerQuantity_Choice_Sinr:
		if err = ie.Sinr.Encode(w); err != nil {
			err = utils.WrapError("Encode Sinr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasTriggerQuantity) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantity_Choice_Rsrp:
		ie.Rsrp = new(RSRP_Range)
		if err = ie.Rsrp.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp", err)
		}
	case MeasTriggerQuantity_Choice_Rsrq:
		ie.Rsrq = new(RSRQ_Range)
		if err = ie.Rsrq.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrq", err)
		}
	case MeasTriggerQuantity_Choice_Sinr:
		ie.Sinr = new(SINR_Range)
		if err = ie.Sinr.Decode(r); err != nil {
			return utils.WrapError("Decode Sinr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
