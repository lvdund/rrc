package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityInformationSidelink_v1700_IEs struct {
	Mac_ParametersSidelink_r17                   *MAC_ParametersSidelink_r17          `optional`
	SupportedBandCombinationListSidelinkNR_v1710 *BandCombinationListSidelinkNR_v1710 `optional`
	NonCriticalExtension                         interface{}                          `optional`
}

func (ie *UECapabilityInformationSidelink_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Mac_ParametersSidelink_r17 != nil, ie.SupportedBandCombinationListSidelinkNR_v1710 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mac_ParametersSidelink_r17 != nil {
		if err = ie.Mac_ParametersSidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersSidelink_r17", err)
		}
	}
	if ie.SupportedBandCombinationListSidelinkNR_v1710 != nil {
		if err = ie.SupportedBandCombinationListSidelinkNR_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandCombinationListSidelinkNR_v1710", err)
		}
	}
	return nil
}

func (ie *UECapabilityInformationSidelink_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Mac_ParametersSidelink_r17Present bool
	if Mac_ParametersSidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandCombinationListSidelinkNR_v1710Present bool
	if SupportedBandCombinationListSidelinkNR_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Mac_ParametersSidelink_r17Present {
		ie.Mac_ParametersSidelink_r17 = new(MAC_ParametersSidelink_r17)
		if err = ie.Mac_ParametersSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersSidelink_r17", err)
		}
	}
	if SupportedBandCombinationListSidelinkNR_v1710Present {
		ie.SupportedBandCombinationListSidelinkNR_v1710 = new(BandCombinationListSidelinkNR_v1710)
		if err = ie.SupportedBandCombinationListSidelinkNR_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandCombinationListSidelinkNR_v1710", err)
		}
	}
	return nil
}
