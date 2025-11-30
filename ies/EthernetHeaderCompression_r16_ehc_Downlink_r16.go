package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EthernetHeaderCompression_r16_ehc_Downlink_r16 struct {
	Drb_ContinueEHC_DL_r16 *EthernetHeaderCompression_r16_ehc_Downlink_r16_drb_ContinueEHC_DL_r16 `optional`
}

func (ie *EthernetHeaderCompression_r16_ehc_Downlink_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Drb_ContinueEHC_DL_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Drb_ContinueEHC_DL_r16 != nil {
		if err = ie.Drb_ContinueEHC_DL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Drb_ContinueEHC_DL_r16", err)
		}
	}
	return nil
}

func (ie *EthernetHeaderCompression_r16_ehc_Downlink_r16) Decode(r *aper.AperReader) error {
	var err error
	var Drb_ContinueEHC_DL_r16Present bool
	if Drb_ContinueEHC_DL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Drb_ContinueEHC_DL_r16Present {
		ie.Drb_ContinueEHC_DL_r16 = new(EthernetHeaderCompression_r16_ehc_Downlink_r16_drb_ContinueEHC_DL_r16)
		if err = ie.Drb_ContinueEHC_DL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Drb_ContinueEHC_DL_r16", err)
		}
	}
	return nil
}
