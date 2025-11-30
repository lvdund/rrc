package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TCI_UL_State_r17_referenceSignal_r17_Choice_nothing uint64 = iota
	TCI_UL_State_r17_referenceSignal_r17_Choice_Ssb_Index_r17
	TCI_UL_State_r17_referenceSignal_r17_Choice_Csi_RS_Index_r17
	TCI_UL_State_r17_referenceSignal_r17_Choice_Srs_r17
)

type TCI_UL_State_r17_referenceSignal_r17 struct {
	Choice           uint64
	Ssb_Index_r17    *SSB_Index
	Csi_RS_Index_r17 *NZP_CSI_RS_ResourceId
	Srs_r17          *SRS_ResourceId
}

func (ie *TCI_UL_State_r17_referenceSignal_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TCI_UL_State_r17_referenceSignal_r17_Choice_Ssb_Index_r17:
		if err = ie.Ssb_Index_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_Index_r17", err)
		}
	case TCI_UL_State_r17_referenceSignal_r17_Choice_Csi_RS_Index_r17:
		if err = ie.Csi_RS_Index_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RS_Index_r17", err)
		}
	case TCI_UL_State_r17_referenceSignal_r17_Choice_Srs_r17:
		if err = ie.Srs_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Srs_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TCI_UL_State_r17_referenceSignal_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TCI_UL_State_r17_referenceSignal_r17_Choice_Ssb_Index_r17:
		ie.Ssb_Index_r17 = new(SSB_Index)
		if err = ie.Ssb_Index_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Index_r17", err)
		}
	case TCI_UL_State_r17_referenceSignal_r17_Choice_Csi_RS_Index_r17:
		ie.Csi_RS_Index_r17 = new(NZP_CSI_RS_ResourceId)
		if err = ie.Csi_RS_Index_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_Index_r17", err)
		}
	case TCI_UL_State_r17_referenceSignal_r17_Choice_Srs_r17:
		ie.Srs_r17 = new(SRS_ResourceId)
		if err = ie.Srs_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
