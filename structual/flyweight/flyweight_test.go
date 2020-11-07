package structual

import (
	"fmt"
	"testing"
)

func TestTeamFlyweightFactoryGetTeam(t *testing.T) {
	factory := NewTeamFactory()

	teamA1 := factory.GetTeam(TEAMA)
	if teamA1 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(TEAMA)
	if teamA2 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	if teamA1 != teamA2 {
		t.Error("TEAM_A objects weren't the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1: %d", factory.GetNumberOfObjects())
	}
}

func TestHighVolume(t *testing.T) {
	factory := NewTeamFactory()

	teams := make([]*Team, 1_000_000)
	for i := 0; i < 500_000; i++ {
		teams[i] = factory.GetTeam(TEAMA)
	}
	for i := 500_000; i < 1_000_000; i++ {
		teams[i] = factory.GetTeam(TEAMB)
	}

	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("The number of objects created was not 2: %d", factory.GetNumberOfObjects())
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
}
