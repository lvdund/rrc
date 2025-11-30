package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ULInformationTransfer_v1700_IEs struct {
	DedicatedInfoF1c_r17 *DedicatedInfoF1c_r17 `optional`
	NonCriticalExtension interface{}           `optional`
}

func (ie *ULInformationTransfer_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DedicatedInfoF1c_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DedicatedInfoF1c_r17 != nil {
		if err = ie.DedicatedInfoF1c_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DedicatedInfoF1c_r17", err)
		}
	}
	return nil
}

func (ie *ULInformationTransfer_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var DedicatedInfoF1c_r17Present bool
	if DedicatedInfoF1c_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if DedicatedInfoF1c_r17Present {
		ie.DedicatedInfoF1c_r17 = new(DedicatedInfoF1c_r17)
		if err = ie.DedicatedInfoF1c_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DedicatedInfoF1c_r17", err)
		}
	}
	return nil
}
