package sysfs

import (
	"testing"
)

func TestCPU(t *testing.T) {

	online := Devices.Object("system").SubObject("cpu").Attribute("online")
	s, err := online.Read()

	// cpu0 := Devices.Object("system").SubObject("cpu").SubObject("cpu0")
	// if !cpu0.Exists() {
	// 	t.Fail()
	// }

	// bios_limit := cpu0.SubObject("power").Attribute("control")
	// if !bios_limit.Exists() {
	// 	t.Fail()
	// }

	// s, err := bios_limit.Read()
	if err != nil {
		t.Fatal(err)
	}
	if s == "" {
		t.Fatal("No value in")
	}
}
