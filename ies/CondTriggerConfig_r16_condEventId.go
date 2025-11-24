package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CondTriggerConfig_r16_condEventId_Choice_nothing uint64 = iota
	CondTriggerConfig_r16_condEventId_Choice_CondEventA3
	CondTriggerConfig_r16_condEventId_Choice_CondEventA5
	CondTriggerConfig_r16_condEventId_Choice_CondEventA4_r17
	CondTriggerConfig_r16_condEventId_Choice_CondEventD1_r17
	CondTriggerConfig_r16_condEventId_Choice_CondEventT1_r17
)

type CondTriggerConfig_r16_condEventId struct {
	Choice          uint64
	CondEventA3     *CondTriggerConfig_r16_condEventId_condEventA3
	CondEventA5     *CondTriggerConfig_r16_condEventId_condEventA5
	CondEventA4_r17 *CondTriggerConfig_r16_condEventId_condEventA4_r17
	CondEventD1_r17 *CondTriggerConfig_r16_condEventId_condEventD1_r17
	CondEventT1_r17 *CondTriggerConfig_r16_condEventId_condEventT1_r17
}

func (ie *CondTriggerConfig_r16_condEventId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CondTriggerConfig_r16_condEventId_Choice_CondEventA3:
		if err = ie.CondEventA3.Encode(w); err != nil {
			err = utils.WrapError("Encode CondEventA3", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_CondEventA5:
		if err = ie.CondEventA5.Encode(w); err != nil {
			err = utils.WrapError("Encode CondEventA5", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_CondEventA4_r17:
		if err = ie.CondEventA4_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode CondEventA4_r17", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_CondEventD1_r17:
		if err = ie.CondEventD1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode CondEventD1_r17", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_CondEventT1_r17:
		if err = ie.CondEventT1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode CondEventT1_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CondTriggerConfig_r16_condEventId) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CondTriggerConfig_r16_condEventId_Choice_CondEventA3:
		ie.CondEventA3 = new(CondTriggerConfig_r16_condEventId_condEventA3)
		if err = ie.CondEventA3.Decode(r); err != nil {
			return utils.WrapError("Decode CondEventA3", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_CondEventA5:
		ie.CondEventA5 = new(CondTriggerConfig_r16_condEventId_condEventA5)
		if err = ie.CondEventA5.Decode(r); err != nil {
			return utils.WrapError("Decode CondEventA5", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_CondEventA4_r17:
		ie.CondEventA4_r17 = new(CondTriggerConfig_r16_condEventId_condEventA4_r17)
		if err = ie.CondEventA4_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CondEventA4_r17", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_CondEventD1_r17:
		ie.CondEventD1_r17 = new(CondTriggerConfig_r16_condEventId_condEventD1_r17)
		if err = ie.CondEventD1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CondEventD1_r17", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_CondEventT1_r17:
		ie.CondEventT1_r17 = new(CondTriggerConfig_r16_condEventId_condEventT1_r17)
		if err = ie.CondEventT1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CondEventT1_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
