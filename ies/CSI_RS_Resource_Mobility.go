package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_Resource_Mobility struct {
	Csi_RS_Index                CSI_RS_Index                                       `madatory`
	SlotConfig                  CSI_RS_Resource_Mobility_slotConfig                `lb:0,ub:31,madatory`
	AssociatedSSB               *CSI_RS_Resource_Mobility_associatedSSB            `optional`
	FrequencyDomainAllocation   CSI_RS_Resource_Mobility_frequencyDomainAllocation `lb:4,ub:4,madatory`
	FirstOFDMSymbolInTimeDomain int64                                              `lb:0,ub:13,madatory`
	SequenceGenerationConfig    int64                                              `lb:0,ub:1023,madatory`
	SlotConfig_r17              *CSI_RS_Resource_Mobility_slotConfig_r17           `lb:0,ub:255,optional,ext-1`
}

func (ie *CSI_RS_Resource_Mobility) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.SlotConfig_r17 != nil
	preambleBits := []bool{hasExtensions, ie.AssociatedSSB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Csi_RS_Index.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_RS_Index", err)
	}
	if err = ie.SlotConfig.Encode(w); err != nil {
		return utils.WrapError("Encode SlotConfig", err)
	}
	if ie.AssociatedSSB != nil {
		if err = ie.AssociatedSSB.Encode(w); err != nil {
			return utils.WrapError("Encode AssociatedSSB", err)
		}
	}
	if err = ie.FrequencyDomainAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode FrequencyDomainAllocation", err)
	}
	if err = w.WriteInteger(ie.FirstOFDMSymbolInTimeDomain, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger FirstOFDMSymbolInTimeDomain", err)
	}
	if err = w.WriteInteger(ie.SequenceGenerationConfig, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger SequenceGenerationConfig", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.SlotConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_RS_Resource_Mobility", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.SlotConfig_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SlotConfig_r17 optional
			if ie.SlotConfig_r17 != nil {
				if err = ie.SlotConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SlotConfig_r17", err)
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

func (ie *CSI_RS_Resource_Mobility) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var AssociatedSSBPresent bool
	if AssociatedSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Csi_RS_Index.Decode(r); err != nil {
		return utils.WrapError("Decode Csi_RS_Index", err)
	}
	if err = ie.SlotConfig.Decode(r); err != nil {
		return utils.WrapError("Decode SlotConfig", err)
	}
	if AssociatedSSBPresent {
		ie.AssociatedSSB = new(CSI_RS_Resource_Mobility_associatedSSB)
		if err = ie.AssociatedSSB.Decode(r); err != nil {
			return utils.WrapError("Decode AssociatedSSB", err)
		}
	}
	if err = ie.FrequencyDomainAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode FrequencyDomainAllocation", err)
	}
	var tmp_int_FirstOFDMSymbolInTimeDomain int64
	if tmp_int_FirstOFDMSymbolInTimeDomain, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger FirstOFDMSymbolInTimeDomain", err)
	}
	ie.FirstOFDMSymbolInTimeDomain = tmp_int_FirstOFDMSymbolInTimeDomain
	var tmp_int_SequenceGenerationConfig int64
	if tmp_int_SequenceGenerationConfig, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger SequenceGenerationConfig", err)
	}
	ie.SequenceGenerationConfig = tmp_int_SequenceGenerationConfig

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

			SlotConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SlotConfig_r17 optional
			if SlotConfig_r17Present {
				ie.SlotConfig_r17 = new(CSI_RS_Resource_Mobility_slotConfig_r17)
				if err = ie.SlotConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SlotConfig_r17", err)
				}
			}
		}
	}
	return nil
}
