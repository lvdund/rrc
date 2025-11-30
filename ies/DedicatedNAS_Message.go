package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DedicatedNAS_Message struct {
	Value []byte `madatory`
}

func (ie *DedicatedNAS_Message) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, nil, false); err != nil {
		return utils.WrapError("Encode DedicatedNAS_Message", err)
	}
	return nil
}

func (ie *DedicatedNAS_Message) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(nil, false); err != nil {
		return utils.WrapError("Decode DedicatedNAS_Message", err)
	}
	ie.Value = v
	return nil
}
