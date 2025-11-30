package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_nothing uint64 = iota
	MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_Csi_RS_Resource_r17
	MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_Ssb_Resource_r17
)

type MPE_Resource_r17_mpe_ReferenceSignal_r17 struct {
	Choice              uint64
	Csi_RS_Resource_r17 *NZP_CSI_RS_ResourceId
	Ssb_Resource_r17    *SSB_Index
}

func (ie *MPE_Resource_r17_mpe_ReferenceSignal_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_Csi_RS_Resource_r17:
		if err = ie.Csi_RS_Resource_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RS_Resource_r17", err)
		}
	case MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_Ssb_Resource_r17:
		if err = ie.Ssb_Resource_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_Resource_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MPE_Resource_r17_mpe_ReferenceSignal_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_Csi_RS_Resource_r17:
		ie.Csi_RS_Resource_r17 = new(NZP_CSI_RS_ResourceId)
		if err = ie.Csi_RS_Resource_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_Resource_r17", err)
		}
	case MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_Ssb_Resource_r17:
		ie.Ssb_Resource_r17 = new(SSB_Index)
		if err = ie.Ssb_Resource_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Resource_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
