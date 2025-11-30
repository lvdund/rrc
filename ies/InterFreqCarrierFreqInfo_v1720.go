package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo_v1720 struct {
	Smtc4list_r17 *SSB_MTC4List_r17 `optional`
}

func (ie *InterFreqCarrierFreqInfo_v1720) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Smtc4list_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Smtc4list_r17 != nil {
		if err = ie.Smtc4list_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc4list_r17", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1720) Decode(r *aper.AperReader) error {
	var err error
	var Smtc4list_r17Present bool
	if Smtc4list_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Smtc4list_r17Present {
		ie.Smtc4list_r17 = new(SSB_MTC4List_r17)
		if err = ie.Smtc4list_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc4list_r17", err)
		}
	}
	return nil
}
