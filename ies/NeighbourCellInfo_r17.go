package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NeighbourCellInfo_r17 struct {
	EpochTime_r17     EpochTime_r17     `madatory`
	EphemerisInfo_r17 EphemerisInfo_r17 `madatory`
}

func (ie *NeighbourCellInfo_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.EpochTime_r17.Encode(w); err != nil {
		return utils.WrapError("Encode EpochTime_r17", err)
	}
	if err = ie.EphemerisInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode EphemerisInfo_r17", err)
	}
	return nil
}

func (ie *NeighbourCellInfo_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.EpochTime_r17.Decode(r); err != nil {
		return utils.WrapError("Decode EpochTime_r17", err)
	}
	if err = ie.EphemerisInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode EphemerisInfo_r17", err)
	}
	return nil
}
