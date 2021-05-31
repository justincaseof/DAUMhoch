package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jacobsa/go-serial/serial"
	log2 "github.com/sirupsen/logrus"
)

type Lyps struct {
	address byte
}

const (
	STATE0_SEARCHING = iota
	STATE1_INITIALIZING
	STATE2_MEASURING
)

var State uint

var lyps Lyps

func main() {
	log2.SetFormatter(&log2.JSONFormatter{})

	// Set up options.
	options := serial.OpenOptions{
		//PortName:        "/dev/tty.usbserial-A8008HlV",
		PortName:          "COM7",
		BaudRate:          9600,
		DataBits:          8,
		StopBits:          1,
		ParityMode:        serial.PARITY_NONE,
		RTSCTSFlowControl: false,

		InterCharacterTimeout: 100,
		MinimumReadSize:       0,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	State = STATE0_SEARCHING
	for {

		switch State {
		case STATE0_SEARCHING:

			fmt.Printf("\nSearching Lyps...\n")

			// Write 4 bytes to the port.
			b := []byte{0x11}
			n, err := port.Write(b)
			if err != nil {
				log.Fatalf("port.Write: %v", err)
			}
			fmt.Println(" --> Wrote", n, "bytes.")

			data := make([]byte, 32)
			readByteCount, err := port.Read(data)
			if readByteCount > 0 {
				fmt.Printf(" --> Response (%d bytes):\t %x \n", readByteCount, data)

				if readByteCount == 2 {
					if data[0] == 0x11 {
						lyps.address = data[1]
						State = STATE1_INITIALIZING
						fmt.Printf(" --> Found Lyps @ address 0x%x!\n", lyps.address)
					}
				}

			} else {
				fmt.Printf(" --> timeout\n")
				State = STATE0_SEARCHING
			}
		case STATE1_INITIALIZING:
			// Write 4 bytes to the port.
			b := []byte{0x40, lyps.address}
			n, err := port.Write(b)
			if err != nil {
				log.Fatalf("port.Write: %v", err)
			}
			fmt.Println(" --> Wrote", n, "bytes.")

			data := make([]byte, 32)
			readByteCount, err := port.Read(data)
			if readByteCount > 0 {
				fmt.Printf(" --> Response (%d bytes):\t %x \n", readByteCount, data)
				if readByteCount == 19 {
					if data[0] == 0x40 {
						fmt.Println(" --> address: ", data[1])

						var lala RunDatenResponsePayload
						data := lala.parse(data[2:])
						fmt.Println("--------------------------------")
						fmt.Println("Program: ", data.Program)
						fmt.Println("Person: ", data.Person)
						fmt.Println("Cycling: ", data.Cycling)
						fmt.Println("PowerWatt: ", data.PowerWatt)
						fmt.Println("RPM: ", data.RPM)
						fmt.Println("SpeedKMH: ", data.SpeedKMH)
						fmt.Println("DistanceMeters: ", data.DistanceMeters)
						fmt.Println("PedalingTimeSeconds: ", data.PedalingTimeSeconds)
						fmt.Println("EnergyJoules: ", data.EnergyJoules)
						fmt.Println("Pulse: ", data.Pulse)
						fmt.Println("Gear: ", data.Gear)
						fmt.Println("EnergyJoulesReal: ", data.EnergyJoulesReal)

					}
				}

			} else {
				fmt.Printf(" --> timeout\n")
				State = STATE0_SEARCHING
			}
		default:
			panic("illegal state")
		}

		time.Sleep(3 * time.Second)
	}

}

func (runDaten RunDaten) encode() CycleData {
	data := &CycleData{
		Cycling:             runDaten.Cycling,
		PowerWatt:           runDaten.PowerWatt,
		RPM:                 runDaten.RPM,
		SpeedKMH:            runDaten.SpeedKMH,
		DistanceMeters:      runDaten.DistanceMeters,
		PedalingTimeSeconds: runDaten.PedalingTimeSeconds,
		EnergyJoules:        runDaten.EnergyJoules,
		Pulse:               runDaten.Pulse,
		Gear:                runDaten.Gear,
		EnergyJoulesReal:    runDaten.EnergyJoulesReal,
	}
	return *data
}
