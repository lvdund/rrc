package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DedicatedInfoF1c_r17 struct {
	Value []byte `madatory`
}

func (ie *DedicatedInfoF1c_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, nil, false); err != nil {
		return utils.WrapError("Encode DedicatedInfoF1c_r17", err)
	}
	return nil
}

func (ie *DedicatedInfoF1c_r17) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(nil, false); err != nil {
		return utils.WrapError("Decode DedicatedInfoF1c_r17", err)
	}
	ie.Value = v
	return nil
}
