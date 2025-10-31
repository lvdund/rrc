package ies

// PDCP-ParametersNR-r15 ::= SEQUENCE
type PdcpParametersnrR15 struct {
	RohcProfilesR15                RohcProfilesupportlistR15
	RohcContextmaxsessionsR15      PdcpParametersnrR15RohcContextmaxsessionsR15
	RohcProfilesulOnlyR15          PdcpParametersnrR15RohcProfilesulOnlyR15
	RohcContextcontinueR15         *PdcpParametersnrR15RohcContextcontinueR15
	OutoforderdeliveryR15          *PdcpParametersnrR15OutoforderdeliveryR15
	SnSizeloR15                    *PdcpParametersnrR15SnSizeloR15
	ImsVoiceovernrPdcpMcgBearerR15 *PdcpParametersnrR15ImsVoiceovernrPdcpMcgBearerR15
	ImsVoiceovernrPdcpScgBearerR15 *PdcpParametersnrR15ImsVoiceovernrPdcpScgBearerR15
}
