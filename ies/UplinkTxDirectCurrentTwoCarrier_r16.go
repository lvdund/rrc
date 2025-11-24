package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentTwoCarrier_r16 struct {
	CarrierOneInfo_r16           UplinkTxDirectCurrentCarrierInfo_r16     `madatory`
	CarrierTwoInfo_r16           UplinkTxDirectCurrentCarrierInfo_r16     `madatory`
	SinglePA_TxDirectCurrent_r16 UplinkTxDirectCurrentTwoCarrierInfo_r16  `madatory`
	SecondPA_TxDirectCurrent_r16 *UplinkTxDirectCurrentTwoCarrierInfo_r16 `optional`
}

func (ie *UplinkTxDirectCurrentTwoCarrier_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SecondPA_TxDirectCurrent_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierOneInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierOneInfo_r16", err)
	}
	if err = ie.CarrierTwoInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierTwoInfo_r16", err)
	}
	if err = ie.SinglePA_TxDirectCurrent_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SinglePA_TxDirectCurrent_r16", err)
	}
	if ie.SecondPA_TxDirectCurrent_r16 != nil {
		if err = ie.SecondPA_TxDirectCurrent_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SecondPA_TxDirectCurrent_r16", err)
		}
	}
	return nil
}

func (ie *UplinkTxDirectCurrentTwoCarrier_r16) Decode(r *uper.UperReader) error {
	var err error
	var SecondPA_TxDirectCurrent_r16Present bool
	if SecondPA_TxDirectCurrent_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierOneInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierOneInfo_r16", err)
	}
	if err = ie.CarrierTwoInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierTwoInfo_r16", err)
	}
	if err = ie.SinglePA_TxDirectCurrent_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SinglePA_TxDirectCurrent_r16", err)
	}
	if SecondPA_TxDirectCurrent_r16Present {
		ie.SecondPA_TxDirectCurrent_r16 = new(UplinkTxDirectCurrentTwoCarrierInfo_r16)
		if err = ie.SecondPA_TxDirectCurrent_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SecondPA_TxDirectCurrent_r16", err)
		}
	}
	return nil
}
