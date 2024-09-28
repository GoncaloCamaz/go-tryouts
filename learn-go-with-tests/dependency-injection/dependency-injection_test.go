package dependencyinjection

import "testing"

func TestDependencyInjection(t *testing.T) {
	t.Run("testing machine gun", func(t *testing.T) {
		ws := newWarShip(MachineGun{rounds: 10000, reloadTime: 100})
		ws.gs.fireGun()

	})

	t.Run("testing canon gun", func(t *testing.T) {
		ws := newWarShip(CanonGun{rounds: 10, reloadTime: 1000000})
		ws.gs.fireGun()
	})

	t.Run("testing anti submarine gun", func(t *testing.T) {
		ws := newWarShip(AntiSubmarineGun{rounds: 3, reloadTime: 10000})
		ws.gs.fireGun()
	})
}
