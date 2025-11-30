package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FreqPriorityDedicatedSlicing_r17 struct {
	Dl_ExplicitCarrierFreq_r17 ARFCN_ValueNR               `madatory`
	SliceInfoListDedicated_r17 *SliceInfoListDedicated_r17 `optional`
}

func (ie *FreqPriorityDedicatedSlicing_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SliceInfoListDedicated_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Dl_ExplicitCarrierFreq_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_ExplicitCarrierFreq_r17", err)
	}
	if ie.SliceInfoListDedicated_r17 != nil {
		if err = ie.SliceInfoListDedicated_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SliceInfoListDedicated_r17", err)
		}
	}
	return nil
}

func (ie *FreqPriorityDedicatedSlicing_r17) Decode(r *aper.AperReader) error {
	var err error
	var SliceInfoListDedicated_r17Present bool
	if SliceInfoListDedicated_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Dl_ExplicitCarrierFreq_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_ExplicitCarrierFreq_r17", err)
	}
	if SliceInfoListDedicated_r17Present {
		ie.SliceInfoListDedicated_r17 = new(SliceInfoListDedicated_r17)
		if err = ie.SliceInfoListDedicated_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SliceInfoListDedicated_r17", err)
		}
	}
	return nil
}
