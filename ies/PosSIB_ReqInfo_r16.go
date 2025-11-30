package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PosSIB_ReqInfo_r16 struct {
	Gnss_id_r16    *GNSS_ID_r16                      `optional`
	Sbas_id_r16    *SBAS_ID_r16                      `optional`
	PosSibType_r16 PosSIB_ReqInfo_r16_posSibType_r16 `madatory`
}

func (ie *PosSIB_ReqInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Gnss_id_r16 != nil, ie.Sbas_id_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Gnss_id_r16 != nil {
		if err = ie.Gnss_id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Gnss_id_r16", err)
		}
	}
	if ie.Sbas_id_r16 != nil {
		if err = ie.Sbas_id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sbas_id_r16", err)
		}
	}
	if err = ie.PosSibType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PosSibType_r16", err)
	}
	return nil
}

func (ie *PosSIB_ReqInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	var Gnss_id_r16Present bool
	if Gnss_id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sbas_id_r16Present bool
	if Sbas_id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Gnss_id_r16Present {
		ie.Gnss_id_r16 = new(GNSS_ID_r16)
		if err = ie.Gnss_id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Gnss_id_r16", err)
		}
	}
	if Sbas_id_r16Present {
		ie.Sbas_id_r16 = new(SBAS_ID_r16)
		if err = ie.Sbas_id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sbas_id_r16", err)
		}
	}
	if err = ie.PosSibType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PosSibType_r16", err)
	}
	return nil
}
