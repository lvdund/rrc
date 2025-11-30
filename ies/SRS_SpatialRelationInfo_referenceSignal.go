package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_SpatialRelationInfo_referenceSignal_Choice_nothing uint64 = iota
	SRS_SpatialRelationInfo_referenceSignal_Choice_Ssb_Index
	SRS_SpatialRelationInfo_referenceSignal_Choice_Csi_RS_Index
	SRS_SpatialRelationInfo_referenceSignal_Choice_Srs
)

type SRS_SpatialRelationInfo_referenceSignal struct {
	Choice       uint64
	Ssb_Index    *SSB_Index
	Csi_RS_Index *NZP_CSI_RS_ResourceId
	Srs          *SRS_SpatialRelationInfo_referenceSignal_srs
}

func (ie *SRS_SpatialRelationInfo_referenceSignal) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfo_referenceSignal_Choice_Ssb_Index:
		if err = ie.Ssb_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_Index", err)
		}
	case SRS_SpatialRelationInfo_referenceSignal_Choice_Csi_RS_Index:
		if err = ie.Csi_RS_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RS_Index", err)
		}
	case SRS_SpatialRelationInfo_referenceSignal_Choice_Srs:
		if err = ie.Srs.Encode(w); err != nil {
			err = utils.WrapError("Encode Srs", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_SpatialRelationInfo_referenceSignal) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfo_referenceSignal_Choice_Ssb_Index:
		ie.Ssb_Index = new(SSB_Index)
		if err = ie.Ssb_Index.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Index", err)
		}
	case SRS_SpatialRelationInfo_referenceSignal_Choice_Csi_RS_Index:
		ie.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
		if err = ie.Csi_RS_Index.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_Index", err)
		}
	case SRS_SpatialRelationInfo_referenceSignal_Choice_Srs:
		ie.Srs = new(SRS_SpatialRelationInfo_referenceSignal_srs)
		if err = ie.Srs.Decode(r); err != nil {
			return utils.WrapError("Decode Srs", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
