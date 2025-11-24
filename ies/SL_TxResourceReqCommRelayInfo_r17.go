package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReqCommRelayInfo_r17 struct {
	Sl_RelayDRXConfig_r17         *SL_TxResourceReq_v1700       `optional`
	Sl_TxResourceReqCommRelay_r17 SL_TxResourceReqCommRelay_r17 `madatory`
}

func (ie *SL_TxResourceReqCommRelayInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_RelayDRXConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_RelayDRXConfig_r17 != nil {
		if err = ie.Sl_RelayDRXConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RelayDRXConfig_r17", err)
		}
	}
	if err = ie.Sl_TxResourceReqCommRelay_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_TxResourceReqCommRelay_r17", err)
	}
	return nil
}

func (ie *SL_TxResourceReqCommRelayInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sl_RelayDRXConfig_r17Present bool
	if Sl_RelayDRXConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_RelayDRXConfig_r17Present {
		ie.Sl_RelayDRXConfig_r17 = new(SL_TxResourceReq_v1700)
		if err = ie.Sl_RelayDRXConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RelayDRXConfig_r17", err)
		}
	}
	if err = ie.Sl_TxResourceReqCommRelay_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_TxResourceReqCommRelay_r17", err)
	}
	return nil
}
