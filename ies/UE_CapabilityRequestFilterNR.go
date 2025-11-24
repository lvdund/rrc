package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterNR struct {
	FrequencyBandListFilter *FreqBandList                       `optional`
	NonCriticalExtension    *UE_CapabilityRequestFilterNR_v1540 `optional`
}

func (ie *UE_CapabilityRequestFilterNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.FrequencyBandListFilter != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FrequencyBandListFilter != nil {
		if err = ie.FrequencyBandListFilter.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyBandListFilter", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterNR) Decode(r *uper.UperReader) error {
	var err error
	var FrequencyBandListFilterPresent bool
	if FrequencyBandListFilterPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FrequencyBandListFilterPresent {
		ie.FrequencyBandListFilter = new(FreqBandList)
		if err = ie.FrequencyBandListFilter.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyBandListFilter", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_CapabilityRequestFilterNR_v1540)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
