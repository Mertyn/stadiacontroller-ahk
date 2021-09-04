package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/71/stadiacontroller"
)

func main() {
	flag.Parse()

	err := run()

	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	controller := stadiacontroller.NewStadiaController()

	defer controller.Close()

	emulator, err := stadiacontroller.NewEmulator(func(vibration stadiacontroller.Vibration) {
		controller.Vibrate(vibration.LargeMotor, vibration.SmallMotor)
	})

	if err != nil {
		return fmt.Errorf("unable to start ViGEm client: %w", err)
	}

	defer emulator.Close()

	x360, err := emulator.CreateXbox360Controller()

	if err != nil {
		return fmt.Errorf("unable to create emulated Xbox 360 controller: %w", err)
	}

	defer x360.Close()

	if err = x360.Connect(); err != nil {
		return fmt.Errorf("unable to connect to emulated Xbox 360 controller: %w", err)
	}

	assistantPressed, capturePressed := false, false

	stadiacontroller.InitAHK()

	for {
		report, err := controller.GetReport()

		if err != nil {
			if errors.Is(err, stadiacontroller.RetryError) {
				time.Sleep(1 * time.Second)
				continue
			}
			return err
		}

		err = x360.Send(&report)

		if err != nil {
			return err
		}

		if report.Assistant != assistantPressed {
			assistantPressed = report.Assistant

			if assistantPressed {
				stadiacontroller.CallAHKFunction("Assistant")
			}
		}

		if report.Capture != capturePressed {
			capturePressed = report.Capture

			if capturePressed {
				stadiacontroller.CallAHKFunction("Capture")
			}
		}
	}
}
