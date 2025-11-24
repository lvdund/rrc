package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_prb_BundlingType_Choice_nothing uint64 = iota
	PDSCH_Config_prb_BundlingType_Choice_StaticBundling
	PDSCH_Config_prb_BundlingType_Choice_DynamicBundling
)

type PDSCH_Config_prb_BundlingType struct {
	Choice          uint64
	StaticBundling  *PDSCH_Config_prb_BundlingType_staticBundling
	DynamicBundling *PDSCH_Config_prb_BundlingType_dynamicBundling
}

func (ie *PDSCH_Config_prb_BundlingType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDSCH_Config_prb_BundlingType_Choice_StaticBundling:
		if err = ie.StaticBundling.Encode(w); err != nil {
			err = utils.WrapError("Encode StaticBundling", err)
		}
	case PDSCH_Config_prb_BundlingType_Choice_DynamicBundling:
		if err = ie.DynamicBundling.Encode(w); err != nil {
			err = utils.WrapError("Encode DynamicBundling", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PDSCH_Config_prb_BundlingType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDSCH_Config_prb_BundlingType_Choice_StaticBundling:
		ie.StaticBundling = new(PDSCH_Config_prb_BundlingType_staticBundling)
		if err = ie.StaticBundling.Decode(r); err != nil {
			return utils.WrapError("Decode StaticBundling", err)
		}
	case PDSCH_Config_prb_BundlingType_Choice_DynamicBundling:
		ie.DynamicBundling = new(PDSCH_Config_prb_BundlingType_dynamicBundling)
		if err = ie.DynamicBundling.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicBundling", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
