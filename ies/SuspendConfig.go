package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuspendConfig struct {
	FullI_RNTI                  I_RNTI_Value              `madatory`
	ShortI_RNTI                 ShortI_RNTI_Value         `madatory`
	Ran_PagingCycle             PagingCycle               `madatory`
	Ran_NotificationAreaInfo    *RAN_NotificationAreaInfo `optional`
	T380                        *PeriodicRNAU_TimerValue  `optional`
	NextHopChainingCount        NextHopChainingCount      `madatory`
	Sl_UEIdentityRemote_r17     *RNTI_Value               `optional,ext-1`
	Sdt_Config_r17              *SDT_Config_r17           `optional,ext-1,setuprelease`
	Srs_PosRRC_Inactive_r17     *SRS_PosRRC_Inactive_r17  `optional,ext-1,setuprelease`
	Ran_ExtendedPagingCycle_r17 *ExtendedPagingCycle_r17  `optional,ext-1`
}

func (ie *SuspendConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Sl_UEIdentityRemote_r17 != nil || ie.Sdt_Config_r17 != nil || ie.Srs_PosRRC_Inactive_r17 != nil || ie.Ran_ExtendedPagingCycle_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Ran_NotificationAreaInfo != nil, ie.T380 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.FullI_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode FullI_RNTI", err)
	}
	if err = ie.ShortI_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode ShortI_RNTI", err)
	}
	if err = ie.Ran_PagingCycle.Encode(w); err != nil {
		return utils.WrapError("Encode Ran_PagingCycle", err)
	}
	if ie.Ran_NotificationAreaInfo != nil {
		if err = ie.Ran_NotificationAreaInfo.Encode(w); err != nil {
			return utils.WrapError("Encode Ran_NotificationAreaInfo", err)
		}
	}
	if ie.T380 != nil {
		if err = ie.T380.Encode(w); err != nil {
			return utils.WrapError("Encode T380", err)
		}
	}
	if err = ie.NextHopChainingCount.Encode(w); err != nil {
		return utils.WrapError("Encode NextHopChainingCount", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sl_UEIdentityRemote_r17 != nil || ie.Sdt_Config_r17 != nil || ie.Srs_PosRRC_Inactive_r17 != nil || ie.Ran_ExtendedPagingCycle_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SuspendConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_UEIdentityRemote_r17 != nil, ie.Sdt_Config_r17 != nil, ie.Srs_PosRRC_Inactive_r17 != nil, ie.Ran_ExtendedPagingCycle_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_UEIdentityRemote_r17 optional
			if ie.Sl_UEIdentityRemote_r17 != nil {
				if err = ie.Sl_UEIdentityRemote_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_UEIdentityRemote_r17", err)
				}
			}
			// encode Sdt_Config_r17 optional
			if ie.Sdt_Config_r17 != nil {
				tmp_Sdt_Config_r17 := utils.SetupRelease[*SDT_Config_r17]{
					Setup: ie.Sdt_Config_r17,
				}
				if err = tmp_Sdt_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sdt_Config_r17", err)
				}
			}
			// encode Srs_PosRRC_Inactive_r17 optional
			if ie.Srs_PosRRC_Inactive_r17 != nil {
				tmp_Srs_PosRRC_Inactive_r17 := utils.SetupRelease[*SRS_PosRRC_Inactive_r17]{
					Setup: ie.Srs_PosRRC_Inactive_r17,
				}
				if err = tmp_Srs_PosRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_PosRRC_Inactive_r17", err)
				}
			}
			// encode Ran_ExtendedPagingCycle_r17 optional
			if ie.Ran_ExtendedPagingCycle_r17 != nil {
				if err = ie.Ran_ExtendedPagingCycle_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ran_ExtendedPagingCycle_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SuspendConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ran_NotificationAreaInfoPresent bool
	if Ran_NotificationAreaInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var T380Present bool
	if T380Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.FullI_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode FullI_RNTI", err)
	}
	if err = ie.ShortI_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode ShortI_RNTI", err)
	}
	if err = ie.Ran_PagingCycle.Decode(r); err != nil {
		return utils.WrapError("Decode Ran_PagingCycle", err)
	}
	if Ran_NotificationAreaInfoPresent {
		ie.Ran_NotificationAreaInfo = new(RAN_NotificationAreaInfo)
		if err = ie.Ran_NotificationAreaInfo.Decode(r); err != nil {
			return utils.WrapError("Decode Ran_NotificationAreaInfo", err)
		}
	}
	if T380Present {
		ie.T380 = new(PeriodicRNAU_TimerValue)
		if err = ie.T380.Decode(r); err != nil {
			return utils.WrapError("Decode T380", err)
		}
	}
	if err = ie.NextHopChainingCount.Decode(r); err != nil {
		return utils.WrapError("Decode NextHopChainingCount", err)
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Sl_UEIdentityRemote_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sdt_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_PosRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ran_ExtendedPagingCycle_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_UEIdentityRemote_r17 optional
			if Sl_UEIdentityRemote_r17Present {
				ie.Sl_UEIdentityRemote_r17 = new(RNTI_Value)
				if err = ie.Sl_UEIdentityRemote_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_UEIdentityRemote_r17", err)
				}
			}
			// decode Sdt_Config_r17 optional
			if Sdt_Config_r17Present {
				tmp_Sdt_Config_r17 := utils.SetupRelease[*SDT_Config_r17]{}
				if err = tmp_Sdt_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sdt_Config_r17", err)
				}
				ie.Sdt_Config_r17 = tmp_Sdt_Config_r17.Setup
			}
			// decode Srs_PosRRC_Inactive_r17 optional
			if Srs_PosRRC_Inactive_r17Present {
				tmp_Srs_PosRRC_Inactive_r17 := utils.SetupRelease[*SRS_PosRRC_Inactive_r17]{}
				if err = tmp_Srs_PosRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_PosRRC_Inactive_r17", err)
				}
				ie.Srs_PosRRC_Inactive_r17 = tmp_Srs_PosRRC_Inactive_r17.Setup
			}
			// decode Ran_ExtendedPagingCycle_r17 optional
			if Ran_ExtendedPagingCycle_r17Present {
				ie.Ran_ExtendedPagingCycle_r17 = new(ExtendedPagingCycle_r17)
				if err = ie.Ran_ExtendedPagingCycle_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ran_ExtendedPagingCycle_r17", err)
				}
			}
		}
	}
	return nil
}
