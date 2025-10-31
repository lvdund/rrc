package ies

// PositionVelocity-r17 ::= SEQUENCE
type PositionvelocityR17 struct {
	PositionxR17  PositionstatevectorR17
	PositionyR17  PositionstatevectorR17
	PositionzR17  PositionstatevectorR17
	VelocityvxR17 VelocitystatevectorR17
	VelocityvyR17 VelocitystatevectorR17
	VelocityvzR17 VelocitystatevectorR17
}
