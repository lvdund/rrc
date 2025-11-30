package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRB_CountMSB_Info struct {
	Drb_Identity      DRB_Identity `madatory`
	CountMSB_Uplink   int64        `lb:0,ub:33554431,madatory`
	CountMSB_Downlink int64        `lb:0,ub:33554431,madatory`
}

func (ie *DRB_CountMSB_Info) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Drb_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode Drb_Identity", err)
	}
	if err = w.WriteInteger(ie.CountMSB_Uplink, &aper.Constraint{Lb: 0, Ub: 33554431}, false); err != nil {
		return utils.WrapError("WriteInteger CountMSB_Uplink", err)
	}
	if err = w.WriteInteger(ie.CountMSB_Downlink, &aper.Constraint{Lb: 0, Ub: 33554431}, false); err != nil {
		return utils.WrapError("WriteInteger CountMSB_Downlink", err)
	}
	return nil
}

func (ie *DRB_CountMSB_Info) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Drb_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode Drb_Identity", err)
	}
	var tmp_int_CountMSB_Uplink int64
	if tmp_int_CountMSB_Uplink, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 33554431}, false); err != nil {
		return utils.WrapError("ReadInteger CountMSB_Uplink", err)
	}
	ie.CountMSB_Uplink = tmp_int_CountMSB_Uplink
	var tmp_int_CountMSB_Downlink int64
	if tmp_int_CountMSB_Downlink, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 33554431}, false); err != nil {
		return utils.WrapError("ReadInteger CountMSB_Downlink", err)
	}
	ie.CountMSB_Downlink = tmp_int_CountMSB_Downlink
	return nil
}
