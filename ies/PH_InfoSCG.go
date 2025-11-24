package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PH_InfoSCG struct {
	ServCellIndex               ServCellIndex                           `madatory`
	Ph_Uplink                   PH_UplinkCarrierSCG                     `madatory`
	Ph_SupplementaryUplink      *PH_UplinkCarrierSCG                    `optional`
	TwoSRS_PUSCH_Repetition_r17 *PH_InfoSCG_twoSRS_PUSCH_Repetition_r17 `optional,ext-1`
}

func (ie *PH_InfoSCG) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.TwoSRS_PUSCH_Repetition_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Ph_SupplementaryUplink != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ServCellIndex.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellIndex", err)
	}
	if err = ie.Ph_Uplink.Encode(w); err != nil {
		return utils.WrapError("Encode Ph_Uplink", err)
	}
	if ie.Ph_SupplementaryUplink != nil {
		if err = ie.Ph_SupplementaryUplink.Encode(w); err != nil {
			return utils.WrapError("Encode Ph_SupplementaryUplink", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.TwoSRS_PUSCH_Repetition_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PH_InfoSCG", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.TwoSRS_PUSCH_Repetition_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode TwoSRS_PUSCH_Repetition_r17 optional
			if ie.TwoSRS_PUSCH_Repetition_r17 != nil {
				if err = ie.TwoSRS_PUSCH_Repetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TwoSRS_PUSCH_Repetition_r17", err)
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

func (ie *PH_InfoSCG) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ph_SupplementaryUplinkPresent bool
	if Ph_SupplementaryUplinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ServCellIndex.Decode(r); err != nil {
		return utils.WrapError("Decode ServCellIndex", err)
	}
	if err = ie.Ph_Uplink.Decode(r); err != nil {
		return utils.WrapError("Decode Ph_Uplink", err)
	}
	if Ph_SupplementaryUplinkPresent {
		ie.Ph_SupplementaryUplink = new(PH_UplinkCarrierSCG)
		if err = ie.Ph_SupplementaryUplink.Decode(r); err != nil {
			return utils.WrapError("Decode Ph_SupplementaryUplink", err)
		}
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

			TwoSRS_PUSCH_Repetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode TwoSRS_PUSCH_Repetition_r17 optional
			if TwoSRS_PUSCH_Repetition_r17Present {
				ie.TwoSRS_PUSCH_Repetition_r17 = new(PH_InfoSCG_twoSRS_PUSCH_Repetition_r17)
				if err = ie.TwoSRS_PUSCH_Repetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TwoSRS_PUSCH_Repetition_r17", err)
				}
			}
		}
	}
	return nil
}
