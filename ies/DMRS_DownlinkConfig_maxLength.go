package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DMRS_DownlinkConfig_maxLength_Enum_len2 aper.Enumerated = 0
)

type DMRS_DownlinkConfig_maxLength struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *DMRS_DownlinkConfig_maxLength) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode DMRS_DownlinkConfig_maxLength", err)
	}
	return nil
}

func (ie *DMRS_DownlinkConfig_maxLength) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode DMRS_DownlinkConfig_maxLength", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
