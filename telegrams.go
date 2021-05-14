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

	result.Program = int(responsebytes[2])
	result.Person = int(responsebytes[3])

	cycleState := responsebytes[4]
	result.Cycling = cycleState > 0
	if result.Cycling {
		// anything?
	}

	result.PowerWatt = 0
	powerRaw := responsebytes[5]
	if powerRaw >= 5 && powerRaw <= 80 || powerRaw >= 10 && powerRaw <= 160 { // validation
		result.PowerWatt = int(powerRaw) * 5
	}

	result.RPM = int(responsebytes[6])

	result.DistanceMeters = int(responsebytes[7]) + int(responsebytes[8])*100

	result.PedalingTimeSeconds = int(responsebytes[9]) + int(responsebytes[10])

	result.EnergyJoules = int(responsebytes[11]) + int(responsebytes[12])*100

	result.Pulse = int(responsebytes[13])

	result.Gear = int(responsebytes[15])

	result.EnergyJoulesReal = int(responsebytes[16]) + int(responsebytes[17])*100

	return result
}
