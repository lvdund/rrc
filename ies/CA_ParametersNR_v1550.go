package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1550 struct {
	Dummy *CA_ParametersNR_v1550_dummy `optional`
}

func (ie *CA_ParametersNR_v1550) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dummy != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dummy != nil {
		if err = ie.Dummy.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1550) Decode(r *aper.AperReader) error {
	var err error
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DummyPresent {
		ie.Dummy = new(CA_ParametersNR_v1550_dummy)
		if err = ie.Dummy.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
	}
	return nil
}
