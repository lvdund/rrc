package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_Choice_nothing uint64 = iota
	PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_Choice_StaticBundling_r16
	PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_Choice_DynamicBundling_r16
)

type PDSCH_Config_prb_BundlingTypeDCI_1_2_r16 struct {
	Choice              uint64
	StaticBundling_r16  *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16
	DynamicBundling_r16 *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_Choice_StaticBundling_r16:
		if err = ie.StaticBundling_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode StaticBundling_r16", err)
		}
	case PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_Choice_DynamicBundling_r16:
		if err = ie.DynamicBundling_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode DynamicBundling_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_Choice_StaticBundling_r16:
		ie.StaticBundling_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16)
		if err = ie.StaticBundling_r16.Decode(r); err != nil {
			return utils.WrapError("Decode StaticBundling_r16", err)
		}
	case PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_Choice_DynamicBundling_r16:
		ie.DynamicBundling_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16)
		if err = ie.DynamicBundling_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicBundling_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
