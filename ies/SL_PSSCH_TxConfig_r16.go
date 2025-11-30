package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSSCH_TxConfig_r16 struct {
	Sl_TypeTxSync_r16             *SL_TypeTxSync_r16                         `optional`
	Sl_ThresUE_Speed_r16          SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16 `madatory`
	Sl_ParametersAboveThres_r16   SL_PSSCH_TxParameters_r16                  `madatory`
	Sl_ParametersBelowThres_r16   SL_PSSCH_TxParameters_r16                  `madatory`
	Sl_ParametersAboveThres_v1650 *SL_MinMaxMCS_List_r16                     `optional,ext-1`
	Sl_ParametersBelowThres_v1650 *SL_MinMaxMCS_List_r16                     `optional,ext-1`
}

func (ie *SL_PSSCH_TxConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Sl_ParametersAboveThres_v1650 != nil || ie.Sl_ParametersBelowThres_v1650 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_TypeTxSync_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_TypeTxSync_r16 != nil {
		if err = ie.Sl_TypeTxSync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TypeTxSync_r16", err)
		}
	}
	if err = ie.Sl_ThresUE_Speed_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ThresUE_Speed_r16", err)
	}
	if err = ie.Sl_ParametersAboveThres_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ParametersAboveThres_r16", err)
	}
	if err = ie.Sl_ParametersBelowThres_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ParametersBelowThres_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sl_ParametersAboveThres_v1650 != nil || ie.Sl_ParametersBelowThres_v1650 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_PSSCH_TxConfig_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_ParametersAboveThres_v1650 != nil, ie.Sl_ParametersBelowThres_v1650 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_ParametersAboveThres_v1650 optional
			if ie.Sl_ParametersAboveThres_v1650 != nil {
				if err = ie.Sl_ParametersAboveThres_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_ParametersAboveThres_v1650", err)
				}
			}
			// encode Sl_ParametersBelowThres_v1650 optional
			if ie.Sl_ParametersBelowThres_v1650 != nil {
				if err = ie.Sl_ParametersBelowThres_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_ParametersBelowThres_v1650", err)
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

func (ie *SL_PSSCH_TxConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TypeTxSync_r16Present bool
	if Sl_TypeTxSync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_TypeTxSync_r16Present {
		ie.Sl_TypeTxSync_r16 = new(SL_TypeTxSync_r16)
		if err = ie.Sl_TypeTxSync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TypeTxSync_r16", err)
		}
	}
	if err = ie.Sl_ThresUE_Speed_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ThresUE_Speed_r16", err)
	}
	if err = ie.Sl_ParametersAboveThres_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ParametersAboveThres_r16", err)
	}
	if err = ie.Sl_ParametersBelowThres_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ParametersBelowThres_r16", err)
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

			Sl_ParametersAboveThres_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_ParametersBelowThres_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_ParametersAboveThres_v1650 optional
			if Sl_ParametersAboveThres_v1650Present {
				ie.Sl_ParametersAboveThres_v1650 = new(SL_MinMaxMCS_List_r16)
				if err = ie.Sl_ParametersAboveThres_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_ParametersAboveThres_v1650", err)
				}
			}
			// decode Sl_ParametersBelowThres_v1650 optional
			if Sl_ParametersBelowThres_v1650Present {
				ie.Sl_ParametersBelowThres_v1650 = new(SL_MinMaxMCS_List_r16)
				if err = ie.Sl_ParametersBelowThres_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_ParametersBelowThres_v1650", err)
				}
			}
		}
	}
	return nil
}
