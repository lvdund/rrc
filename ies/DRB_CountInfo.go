package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRB_CountInfo struct {
	Drb_Identity   DRB_Identity `madatory`
	Count_Uplink   int64        `lb:0,ub:4294967295,madatory`
	Count_Downlink int64        `lb:0,ub:4294967295,madatory`
}

func (ie *DRB_CountInfo) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Drb_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode Drb_Identity", err)
	}
	if err = w.WriteInteger(ie.Count_Uplink, &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("WriteInteger Count_Uplink", err)
	}
	if err = w.WriteInteger(ie.Count_Downlink, &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("WriteInteger Count_Downlink", err)
	}
	return nil
}

func (ie *DRB_CountInfo) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Drb_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode Drb_Identity", err)
	}
	var tmp_int_Count_Uplink int64
	if tmp_int_Count_Uplink, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("ReadInteger Count_Uplink", err)
	}
	ie.Count_Uplink = tmp_int_Count_Uplink
	var tmp_int_Count_Downlink int64
	if tmp_int_Count_Downlink, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("ReadInteger Count_Downlink", err)
	}
	ie.Count_Downlink = tmp_int_Count_Downlink
	return nil
}
