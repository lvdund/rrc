package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB12_IEs_r16 struct {
	Sl_ConfigCommonNR_r16             SL_ConfigCommonNR_r16                      `madatory`
	LateNonCriticalExtension          *[]byte                                    `optional`
	Sl_DRX_ConfigCommonGC_BC_r17      *SL_DRX_ConfigGC_BC_r17                    `optional,ext-1`
	Sl_DiscConfigCommon_r17           *SL_DiscConfigCommon_r17                   `optional,ext-1`
	Sl_L2U2N_Relay_r17                *SIB12_IEs_r16_sl_L2U2N_Relay_r17          `optional,ext-1`
	Sl_NonRelayDiscovery_r17          *SIB12_IEs_r16_sl_NonRelayDiscovery_r17    `optional,ext-1`
	Sl_L3U2N_RelayDiscovery_r17       *SIB12_IEs_r16_sl_L3U2N_RelayDiscovery_r17 `optional,ext-1`
	Sl_TimersAndConstantsRemoteUE_r17 *UE_TimersAndConstantsRemoteUE_r17         `optional,ext-1`
}

func (ie *SIB12_IEs_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Sl_DRX_ConfigCommonGC_BC_r17 != nil || ie.Sl_DiscConfigCommon_r17 != nil || ie.Sl_L2U2N_Relay_r17 != nil || ie.Sl_NonRelayDiscovery_r17 != nil || ie.Sl_L3U2N_RelayDiscovery_r17 != nil || ie.Sl_TimersAndConstantsRemoteUE_r17 != nil
	preambleBits := []bool{hasExtensions, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_ConfigCommonNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ConfigCommonNR_r16", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sl_DRX_ConfigCommonGC_BC_r17 != nil || ie.Sl_DiscConfigCommon_r17 != nil || ie.Sl_L2U2N_Relay_r17 != nil || ie.Sl_NonRelayDiscovery_r17 != nil || ie.Sl_L3U2N_RelayDiscovery_r17 != nil || ie.Sl_TimersAndConstantsRemoteUE_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB12_IEs_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_DRX_ConfigCommonGC_BC_r17 != nil, ie.Sl_DiscConfigCommon_r17 != nil, ie.Sl_L2U2N_Relay_r17 != nil, ie.Sl_NonRelayDiscovery_r17 != nil, ie.Sl_L3U2N_RelayDiscovery_r17 != nil, ie.Sl_TimersAndConstantsRemoteUE_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_DRX_ConfigCommonGC_BC_r17 optional
			if ie.Sl_DRX_ConfigCommonGC_BC_r17 != nil {
				if err = ie.Sl_DRX_ConfigCommonGC_BC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_DRX_ConfigCommonGC_BC_r17", err)
				}
			}
			// encode Sl_DiscConfigCommon_r17 optional
			if ie.Sl_DiscConfigCommon_r17 != nil {
				if err = ie.Sl_DiscConfigCommon_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_DiscConfigCommon_r17", err)
				}
			}
			// encode Sl_L2U2N_Relay_r17 optional
			if ie.Sl_L2U2N_Relay_r17 != nil {
				if err = ie.Sl_L2U2N_Relay_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_L2U2N_Relay_r17", err)
				}
			}
			// encode Sl_NonRelayDiscovery_r17 optional
			if ie.Sl_NonRelayDiscovery_r17 != nil {
				if err = ie.Sl_NonRelayDiscovery_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_NonRelayDiscovery_r17", err)
				}
			}
			// encode Sl_L3U2N_RelayDiscovery_r17 optional
			if ie.Sl_L3U2N_RelayDiscovery_r17 != nil {
				if err = ie.Sl_L3U2N_RelayDiscovery_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_L3U2N_RelayDiscovery_r17", err)
				}
			}
			// encode Sl_TimersAndConstantsRemoteUE_r17 optional
			if ie.Sl_TimersAndConstantsRemoteUE_r17 != nil {
				if err = ie.Sl_TimersAndConstantsRemoteUE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_TimersAndConstantsRemoteUE_r17", err)
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

func (ie *SIB12_IEs_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_ConfigCommonNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ConfigCommonNR_r16", err)
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Sl_DRX_ConfigCommonGC_BC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_DiscConfigCommon_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_L2U2N_Relay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_NonRelayDiscovery_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_L3U2N_RelayDiscovery_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_TimersAndConstantsRemoteUE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_DRX_ConfigCommonGC_BC_r17 optional
			if Sl_DRX_ConfigCommonGC_BC_r17Present {
				ie.Sl_DRX_ConfigCommonGC_BC_r17 = new(SL_DRX_ConfigGC_BC_r17)
				if err = ie.Sl_DRX_ConfigCommonGC_BC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_DRX_ConfigCommonGC_BC_r17", err)
				}
			}
			// decode Sl_DiscConfigCommon_r17 optional
			if Sl_DiscConfigCommon_r17Present {
				ie.Sl_DiscConfigCommon_r17 = new(SL_DiscConfigCommon_r17)
				if err = ie.Sl_DiscConfigCommon_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_DiscConfigCommon_r17", err)
				}
			}
			// decode Sl_L2U2N_Relay_r17 optional
			if Sl_L2U2N_Relay_r17Present {
				ie.Sl_L2U2N_Relay_r17 = new(SIB12_IEs_r16_sl_L2U2N_Relay_r17)
				if err = ie.Sl_L2U2N_Relay_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_L2U2N_Relay_r17", err)
				}
			}
			// decode Sl_NonRelayDiscovery_r17 optional
			if Sl_NonRelayDiscovery_r17Present {
				ie.Sl_NonRelayDiscovery_r17 = new(SIB12_IEs_r16_sl_NonRelayDiscovery_r17)
				if err = ie.Sl_NonRelayDiscovery_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_NonRelayDiscovery_r17", err)
				}
			}
			// decode Sl_L3U2N_RelayDiscovery_r17 optional
			if Sl_L3U2N_RelayDiscovery_r17Present {
				ie.Sl_L3U2N_RelayDiscovery_r17 = new(SIB12_IEs_r16_sl_L3U2N_RelayDiscovery_r17)
				if err = ie.Sl_L3U2N_RelayDiscovery_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_L3U2N_RelayDiscovery_r17", err)
				}
			}
			// decode Sl_TimersAndConstantsRemoteUE_r17 optional
			if Sl_TimersAndConstantsRemoteUE_r17Present {
				ie.Sl_TimersAndConstantsRemoteUE_r17 = new(UE_TimersAndConstantsRemoteUE_r17)
				if err = ie.Sl_TimersAndConstantsRemoteUE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_TimersAndConstantsRemoteUE_r17", err)
				}
			}
		}
	}
	return nil
}
