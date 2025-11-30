package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	QCL_Info_referenceSignal_Choice_nothing uint64 = iota
	QCL_Info_referenceSignal_Choice_Csi_rs
	QCL_Info_referenceSignal_Choice_Ssb
)

type QCL_Info_referenceSignal struct {
	Choice uint64
	Csi_rs *NZP_CSI_RS_ResourceId
	Ssb    *SSB_Index
}

func (ie *QCL_Info_referenceSignal) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case QCL_Info_referenceSignal_Choice_Csi_rs:
		if err = ie.Csi_rs.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_rs", err)
		}
	case QCL_Info_referenceSignal_Choice_Ssb:
		if err = ie.Ssb.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *QCL_Info_referenceSignal) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case QCL_Info_referenceSignal_Choice_Csi_rs:
		ie.Csi_rs = new(NZP_CSI_RS_ResourceId)
		if err = ie.Csi_rs.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_rs", err)
		}
	case QCL_Info_referenceSignal_Choice_Ssb:
		ie.Ssb = new(SSB_Index)
		if err = ie.Ssb.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
