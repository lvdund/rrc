package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EthernetHeaderCompression_r16 struct {
	Ehc_Common_r16   EthernetHeaderCompression_r16_ehc_Common_r16    `madatory`
	Ehc_Downlink_r16 *EthernetHeaderCompression_r16_ehc_Downlink_r16 `optional,ext`
	Ehc_Uplink_r16   *EthernetHeaderCompression_r16_ehc_Uplink_r16   `lb:1,ub:32767,optional,ext`
}

func (ie *EthernetHeaderCompression_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Ehc_Common_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ehc_Common_r16", err)
	}
	return nil
}

func (ie *EthernetHeaderCompression_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Ehc_Common_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ehc_Common_r16", err)
	}
	return nil
}
