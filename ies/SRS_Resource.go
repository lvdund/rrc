package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource struct {
	Srs_ResourceId               SRS_ResourceId                             `madatory`
	NrofSRS_Ports                SRS_Resource_nrofSRS_Ports                 `madatory`
	Ptrs_PortIndex               *SRS_Resource_ptrs_PortIndex               `optional`
	TransmissionComb             SRS_Resource_transmissionComb              `lb:0,ub:1,madatory`
	ResourceMapping              SRS_Resource_resourceMapping               `lb:0,ub:5,madatory`
	FreqDomainPosition           int64                                      `lb:0,ub:67,madatory`
	FreqDomainShift              int64                                      `lb:0,ub:268,madatory`
	FreqHopping                  SRS_Resource_freqHopping                   `lb:0,ub:63,madatory`
	GroupOrSequenceHopping       SRS_Resource_groupOrSequenceHopping        `madatory`
	ResourceType                 SRS_Resource_resourceType                  `madatory`
	SequenceId                   int64                                      `lb:0,ub:1023,madatory,ext`
	SpatialRelationInfo          *SRS_SpatialRelationInfo                   `optional,ext`
	ResourceMapping_r16          *SRS_Resource_resourceMapping_r16          `lb:0,ub:13,optional,ext-1`
	SpatialRelationInfo_PDC_r17  *SpatialRelationInfo_PDC_r17               `optional,ext-2,setuprelease`
	ResourceMapping_r17          *SRS_Resource_resourceMapping_r17          `lb:0,ub:13,optional,ext-2`
	PartialFreqSounding_r17      *SRS_Resource_partialFreqSounding_r17      `lb:0,ub:1,optional,ext-2`
	TransmissionComb_n8_r17      *SRS_Resource_transmissionComb_n8_r17      `lb:0,ub:7,optional,ext-2`
	Srs_TCI_State_r17            *SRS_Resource_srs_TCI_State_r17            `optional,ext-2`
	RepetitionFactor_v1730       *SRS_Resource_repetitionFactor_v1730       `optional,ext-3`
	Srs_DLorJointTCI_State_v1730 *SRS_Resource_srs_DLorJointTCI_State_v1730 `optional,ext-3`
}

func (ie *SRS_Resource) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.ResourceMapping_r16 != nil || ie.SpatialRelationInfo_PDC_r17 != nil || ie.ResourceMapping_r17 != nil || ie.PartialFreqSounding_r17 != nil || ie.TransmissionComb_n8_r17 != nil || ie.Srs_TCI_State_r17 != nil || ie.RepetitionFactor_v1730 != nil || ie.Srs_DLorJointTCI_State_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.Ptrs_PortIndex != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Srs_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_ResourceId", err)
	}
	if err = ie.NrofSRS_Ports.Encode(w); err != nil {
		return utils.WrapError("Encode NrofSRS_Ports", err)
	}
	if ie.Ptrs_PortIndex != nil {
		if err = ie.Ptrs_PortIndex.Encode(w); err != nil {
			return utils.WrapError("Encode Ptrs_PortIndex", err)
		}
	}
	if err = ie.TransmissionComb.Encode(w); err != nil {
		return utils.WrapError("Encode TransmissionComb", err)
	}
	if err = ie.ResourceMapping.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceMapping", err)
	}
	if err = w.WriteInteger(ie.FreqDomainPosition, &aper.Constraint{Lb: 0, Ub: 67}, false); err != nil {
		return utils.WrapError("WriteInteger FreqDomainPosition", err)
	}
	if err = w.WriteInteger(ie.FreqDomainShift, &aper.Constraint{Lb: 0, Ub: 268}, false); err != nil {
		return utils.WrapError("WriteInteger FreqDomainShift", err)
	}
	if err = ie.FreqHopping.Encode(w); err != nil {
		return utils.WrapError("Encode FreqHopping", err)
	}
	if err = ie.GroupOrSequenceHopping.Encode(w); err != nil {
		return utils.WrapError("Encode GroupOrSequenceHopping", err)
	}
	if err = ie.ResourceType.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceType", err)
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.ResourceMapping_r16 != nil, ie.SpatialRelationInfo_PDC_r17 != nil || ie.ResourceMapping_r17 != nil || ie.PartialFreqSounding_r17 != nil || ie.TransmissionComb_n8_r17 != nil || ie.Srs_TCI_State_r17 != nil, ie.RepetitionFactor_v1730 != nil || ie.Srs_DLorJointTCI_State_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SRS_Resource", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ResourceMapping_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ResourceMapping_r16 optional
			if ie.ResourceMapping_r16 != nil {
				if err = ie.ResourceMapping_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourceMapping_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.SpatialRelationInfo_PDC_r17 != nil, ie.ResourceMapping_r17 != nil, ie.PartialFreqSounding_r17 != nil, ie.TransmissionComb_n8_r17 != nil, ie.Srs_TCI_State_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SpatialRelationInfo_PDC_r17 optional
			if ie.SpatialRelationInfo_PDC_r17 != nil {
				tmp_SpatialRelationInfo_PDC_r17 := utils.SetupRelease[*SpatialRelationInfo_PDC_r17]{
					Setup: ie.SpatialRelationInfo_PDC_r17,
				}
				if err = tmp_SpatialRelationInfo_PDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelationInfo_PDC_r17", err)
				}
			}
			// encode ResourceMapping_r17 optional
			if ie.ResourceMapping_r17 != nil {
				if err = ie.ResourceMapping_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourceMapping_r17", err)
				}
			}
			// encode PartialFreqSounding_r17 optional
			if ie.PartialFreqSounding_r17 != nil {
				if err = ie.PartialFreqSounding_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PartialFreqSounding_r17", err)
				}
			}
			// encode TransmissionComb_n8_r17 optional
			if ie.TransmissionComb_n8_r17 != nil {
				if err = ie.TransmissionComb_n8_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TransmissionComb_n8_r17", err)
				}
			}
			// encode Srs_TCI_State_r17 optional
			if ie.Srs_TCI_State_r17 != nil {
				if err = ie.Srs_TCI_State_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_TCI_State_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.RepetitionFactor_v1730 != nil, ie.Srs_DLorJointTCI_State_v1730 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RepetitionFactor_v1730 optional
			if ie.RepetitionFactor_v1730 != nil {
				if err = ie.RepetitionFactor_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RepetitionFactor_v1730", err)
				}
			}
			// encode Srs_DLorJointTCI_State_v1730 optional
			if ie.Srs_DLorJointTCI_State_v1730 != nil {
				if err = ie.Srs_DLorJointTCI_State_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_DLorJointTCI_State_v1730", err)
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

func (ie *SRS_Resource) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ptrs_PortIndexPresent bool
	if Ptrs_PortIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Srs_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_ResourceId", err)
	}
	if err = ie.NrofSRS_Ports.Decode(r); err != nil {
		return utils.WrapError("Decode NrofSRS_Ports", err)
	}
	if Ptrs_PortIndexPresent {
		ie.Ptrs_PortIndex = new(SRS_Resource_ptrs_PortIndex)
		if err = ie.Ptrs_PortIndex.Decode(r); err != nil {
			return utils.WrapError("Decode Ptrs_PortIndex", err)
		}
	}
	if err = ie.TransmissionComb.Decode(r); err != nil {
		return utils.WrapError("Decode TransmissionComb", err)
	}
	if err = ie.ResourceMapping.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceMapping", err)
	}
	var tmp_int_FreqDomainPosition int64
	if tmp_int_FreqDomainPosition, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 67}, false); err != nil {
		return utils.WrapError("ReadInteger FreqDomainPosition", err)
	}
	ie.FreqDomainPosition = tmp_int_FreqDomainPosition
	var tmp_int_FreqDomainShift int64
	if tmp_int_FreqDomainShift, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 268}, false); err != nil {
		return utils.WrapError("ReadInteger FreqDomainShift", err)
	}
	ie.FreqDomainShift = tmp_int_FreqDomainShift
	if err = ie.FreqHopping.Decode(r); err != nil {
		return utils.WrapError("Decode FreqHopping", err)
	}
	if err = ie.GroupOrSequenceHopping.Decode(r); err != nil {
		return utils.WrapError("Decode GroupOrSequenceHopping", err)
	}
	if err = ie.ResourceType.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceType", err)
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			ResourceMapping_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ResourceMapping_r16 optional
			if ResourceMapping_r16Present {
				ie.ResourceMapping_r16 = new(SRS_Resource_resourceMapping_r16)
				if err = ie.ResourceMapping_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ResourceMapping_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SpatialRelationInfo_PDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ResourceMapping_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PartialFreqSounding_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TransmissionComb_n8_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_TCI_State_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SpatialRelationInfo_PDC_r17 optional
			if SpatialRelationInfo_PDC_r17Present {
				tmp_SpatialRelationInfo_PDC_r17 := utils.SetupRelease[*SpatialRelationInfo_PDC_r17]{}
				if err = tmp_SpatialRelationInfo_PDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpatialRelationInfo_PDC_r17", err)
				}
				ie.SpatialRelationInfo_PDC_r17 = tmp_SpatialRelationInfo_PDC_r17.Setup
			}
			// decode ResourceMapping_r17 optional
			if ResourceMapping_r17Present {
				ie.ResourceMapping_r17 = new(SRS_Resource_resourceMapping_r17)
				if err = ie.ResourceMapping_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ResourceMapping_r17", err)
				}
			}
			// decode PartialFreqSounding_r17 optional
			if PartialFreqSounding_r17Present {
				ie.PartialFreqSounding_r17 = new(SRS_Resource_partialFreqSounding_r17)
				if err = ie.PartialFreqSounding_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PartialFreqSounding_r17", err)
				}
			}
			// decode TransmissionComb_n8_r17 optional
			if TransmissionComb_n8_r17Present {
				ie.TransmissionComb_n8_r17 = new(SRS_Resource_transmissionComb_n8_r17)
				if err = ie.TransmissionComb_n8_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TransmissionComb_n8_r17", err)
				}
			}
			// decode Srs_TCI_State_r17 optional
			if Srs_TCI_State_r17Present {
				ie.Srs_TCI_State_r17 = new(SRS_Resource_srs_TCI_State_r17)
				if err = ie.Srs_TCI_State_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_TCI_State_r17", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			RepetitionFactor_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_DLorJointTCI_State_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RepetitionFactor_v1730 optional
			if RepetitionFactor_v1730Present {
				ie.RepetitionFactor_v1730 = new(SRS_Resource_repetitionFactor_v1730)
				if err = ie.RepetitionFactor_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode RepetitionFactor_v1730", err)
				}
			}
			// decode Srs_DLorJointTCI_State_v1730 optional
			if Srs_DLorJointTCI_State_v1730Present {
				ie.Srs_DLorJointTCI_State_v1730 = new(SRS_Resource_srs_DLorJointTCI_State_v1730)
				if err = ie.Srs_DLorJointTCI_State_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_DLorJointTCI_State_v1730", err)
				}
			}
		}
	}
	return nil
}
