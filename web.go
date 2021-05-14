package main

type CycleData struct {
	Cycling             bool `json:","`
	PowerWatt           int  `json:","`
	RPM                 int  `json:","`
	SpeedKMH            int  `json:","`
	DistanceMeters      int  `json:","`
	PedalingTimeSeconds int  `json:","`
	EnergyJoules        int  `json:","`
	Pulse               int  `json:","`
	Gear                int  `json:","`
	EnergyJoulesReal    int  `json:","`
}
