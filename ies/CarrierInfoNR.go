package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CarrierInfoNR struct {
	CarrierFreq          ARFCN_ValueNR     `madatory`
	SsbSubcarrierSpacing SubcarrierSpacing `madatory`
	Smtc                 *SSB_MTC          `optional`
}

func (ie *CarrierInfoNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Smtc != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq", err)
	}
	if err = ie.SsbSubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode SsbSubcarrierSpacing", err)
	}
	if ie.Smtc != nil {
		if err = ie.Smtc.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc", err)
		}
	}
	return nil
}

func (ie *CarrierInfoNR) Decode(r *uper.UperReader) error {
	var err error
	var SmtcPresent bool
	if SmtcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq", err)
	}
	if err = ie.SsbSubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode SsbSubcarrierSpacing", err)
	}
	if SmtcPresent {
		ie.Smtc = new(SSB_MTC)
		if err = ie.Smtc.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc", err)
		}
	}
	return nil
}
