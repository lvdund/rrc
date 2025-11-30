package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RadioLinkMonitoringRS_detectionResource_Choice_nothing uint64 = iota
	RadioLinkMonitoringRS_detectionResource_Choice_Ssb_Index
	RadioLinkMonitoringRS_detectionResource_Choice_Csi_RS_Index
)

type RadioLinkMonitoringRS_detectionResource struct {
	Choice       uint64
	Ssb_Index    *SSB_Index
	Csi_RS_Index *NZP_CSI_RS_ResourceId
}

func (ie *RadioLinkMonitoringRS_detectionResource) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RadioLinkMonitoringRS_detectionResource_Choice_Ssb_Index:
		if err = ie.Ssb_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_Index", err)
		}
	case RadioLinkMonitoringRS_detectionResource_Choice_Csi_RS_Index:
		if err = ie.Csi_RS_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RS_Index", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RadioLinkMonitoringRS_detectionResource) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RadioLinkMonitoringRS_detectionResource_Choice_Ssb_Index:
		ie.Ssb_Index = new(SSB_Index)
		if err = ie.Ssb_Index.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Index", err)
		}
	case RadioLinkMonitoringRS_detectionResource_Choice_Csi_RS_Index:
		ie.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
		if err = ie.Csi_RS_Index.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_Index", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
