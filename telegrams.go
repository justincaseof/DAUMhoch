package main

type RunDaten struct {
	Program             int
	Person              int
	Cycling             bool
	PowerWatt           int
	RPM                 int
	SpeedKMH            int
	DistanceMeters      int
	PedalingTimeSeconds int
	EnergyJoules        int
	Pulse               int
	Gear                int
	EnergyJoulesReal    int
}

type RunDatenResponsePayload []byte

func (data RunDatenResponsePayload) parse(responsebytes []byte) RunDaten {
	var result RunDaten

	result.Program = int(responsebytes[0])
	result.Person = int(responsebytes[1])

	cycleState := responsebytes[2]
	result.Cycling = cycleState > 0
	if result.Cycling {
		// anything?
	}

	// check: OK
	result.PowerWatt = 0
	powerRaw := responsebytes[3]
	if powerRaw >= 5 && powerRaw <= 80 || powerRaw >= 10 && powerRaw <= 160 { // validation
		result.PowerWatt = int(powerRaw) * 5
	}

	// check: OK
	result.RPM = int(responsebytes[4])

	// check: OK
	result.SpeedKMH = int(responsebytes[5])

	// wrong somehow
	result.DistanceMeters = (int(responsebytes[6])<<8 + int(responsebytes[7])) * 100

	// wrong somehow
	result.PedalingTimeSeconds = int(responsebytes[8])<<8 + int(responsebytes[9])

	// wrong somehow
	result.EnergyJoules = int(responsebytes[10])<<8 + int(responsebytes[11])

	// check: OK
	result.Pulse = int(responsebytes[12])
	//result.PulseState = int(responsebytes[13])

	result.Gear = int(responsebytes[14])

	result.EnergyJoulesReal = int(responsebytes[15]) + int(responsebytes[16])*100

	return result
}
