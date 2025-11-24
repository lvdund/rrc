package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_nothing uint64 = iota
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Ssb_Index
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Csi_RS_Index
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Dl_PRS_PDC
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Srs
)

type SpatialRelationInfo_PDC_r17_referenceSignal struct {
	Choice       uint64
	Ssb_Index    *SSB_Index
	Csi_RS_Index *NZP_CSI_RS_ResourceId
	Dl_PRS_PDC   *NR_DL_PRS_ResourceID_r17
	Srs          *SpatialRelationInfo_PDC_r17_referenceSignal_srs
}

func (ie *SpatialRelationInfo_PDC_r17_referenceSignal) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Ssb_Index:
		if err = ie.Ssb_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_Index", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Csi_RS_Index:
		if err = ie.Csi_RS_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RS_Index", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Dl_PRS_PDC:
		if err = ie.Dl_PRS_PDC.Encode(w); err != nil {
			err = utils.WrapError("Encode Dl_PRS_PDC", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Srs:
		if err = ie.Srs.Encode(w); err != nil {
			err = utils.WrapError("Encode Srs", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SpatialRelationInfo_PDC_r17_referenceSignal) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Ssb_Index:
		ie.Ssb_Index = new(SSB_Index)
		if err = ie.Ssb_Index.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Index", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Csi_RS_Index:
		ie.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
		if err = ie.Csi_RS_Index.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_Index", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Dl_PRS_PDC:
		ie.Dl_PRS_PDC = new(NR_DL_PRS_ResourceID_r17)
		if err = ie.Dl_PRS_PDC.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_PRS_PDC", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_Srs:
		ie.Srs = new(SpatialRelationInfo_PDC_r17_referenceSignal_srs)
		if err = ie.Srs.Decode(r); err != nil {
			return utils.WrapError("Decode Srs", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
