package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EthernetHeaderCompression_r16_ehc_Uplink_r16 struct {
	MaxCID_EHC_UL_r16      int64                                                                `lb:1,ub:32767,madatory`
	Drb_ContinueEHC_UL_r16 *EthernetHeaderCompression_r16_ehc_Uplink_r16_drb_ContinueEHC_UL_r16 `optional`
}

func (ie *EthernetHeaderCompression_r16_ehc_Uplink_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Drb_ContinueEHC_UL_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.MaxCID_EHC_UL_r16, &aper.Constraint{Lb: 1, Ub: 32767}, false); err != nil {
		return utils.WrapError("WriteInteger MaxCID_EHC_UL_r16", err)
	}
	if ie.Drb_ContinueEHC_UL_r16 != nil {
		if err = ie.Drb_ContinueEHC_UL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Drb_ContinueEHC_UL_r16", err)
		}
	}
	return nil
}

func (ie *EthernetHeaderCompression_r16_ehc_Uplink_r16) Decode(r *aper.AperReader) error {
	var err error
	var Drb_ContinueEHC_UL_r16Present bool
	if Drb_ContinueEHC_UL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_MaxCID_EHC_UL_r16 int64
	if tmp_int_MaxCID_EHC_UL_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32767}, false); err != nil {
		return utils.WrapError("ReadInteger MaxCID_EHC_UL_r16", err)
	}
	ie.MaxCID_EHC_UL_r16 = tmp_int_MaxCID_EHC_UL_r16
	if Drb_ContinueEHC_UL_r16Present {
		ie.Drb_ContinueEHC_UL_r16 = new(EthernetHeaderCompression_r16_ehc_Uplink_r16_drb_ContinueEHC_UL_r16)
		if err = ie.Drb_ContinueEHC_UL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Drb_ContinueEHC_UL_r16", err)
		}
	}
	return nil
}
