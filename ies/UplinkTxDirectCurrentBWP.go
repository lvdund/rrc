package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentBWP struct {
	Bwp_Id                  BWP_Id `madatory`
	Shift7dot5kHz           bool   `madatory`
	TxDirectCurrentLocation int64  `lb:0,ub:3301,madatory`
}

func (ie *UplinkTxDirectCurrentBWP) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Bwp_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Bwp_Id", err)
	}
	if err = w.WriteBoolean(ie.Shift7dot5kHz); err != nil {
		return utils.WrapError("WriteBoolean Shift7dot5kHz", err)
	}
	if err = w.WriteInteger(ie.TxDirectCurrentLocation, &aper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
		return utils.WrapError("WriteInteger TxDirectCurrentLocation", err)
	}
	return nil
}

func (ie *UplinkTxDirectCurrentBWP) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Bwp_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Bwp_Id", err)
	}
	var tmp_bool_Shift7dot5kHz bool
	if tmp_bool_Shift7dot5kHz, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Shift7dot5kHz", err)
	}
	ie.Shift7dot5kHz = tmp_bool_Shift7dot5kHz
	var tmp_int_TxDirectCurrentLocation int64
	if tmp_int_TxDirectCurrentLocation, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
		return utils.WrapError("ReadInteger TxDirectCurrentLocation", err)
	}
	ie.TxDirectCurrentLocation = tmp_int_TxDirectCurrentLocation
	return nil
}
