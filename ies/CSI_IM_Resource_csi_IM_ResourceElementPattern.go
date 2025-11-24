package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_nothing uint64 = iota
	CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_Pattern0
	CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_Pattern1
)

type CSI_IM_Resource_csi_IM_ResourceElementPattern struct {
	Choice   uint64
	Pattern0 *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0
	Pattern1 *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_Pattern0:
		if err = ie.Pattern0.Encode(w); err != nil {
			err = utils.WrapError("Encode Pattern0", err)
		}
	case CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_Pattern1:
		if err = ie.Pattern1.Encode(w); err != nil {
			err = utils.WrapError("Encode Pattern1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_Pattern0:
		ie.Pattern0 = new(CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0)
		if err = ie.Pattern0.Decode(r); err != nil {
			return utils.WrapError("Decode Pattern0", err)
		}
	case CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_Pattern1:
		ie.Pattern1 = new(CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1)
		if err = ie.Pattern1.Decode(r); err != nil {
			return utils.WrapError("Decode Pattern1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
