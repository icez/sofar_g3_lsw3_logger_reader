package mosquitto

import "testing"

func TestNormalizeUnit(t *testing.T) {
	cases := map[string]string{
		"℃":   "°C",
		"kWh": "kWh",
		"":    "",
	}

	for unit, expected := range cases {
		if got := normalizeUnit(unit); got != expected {
			t.Errorf("normalizeUnit(%q) = %q, want %q", unit, got, expected)
		}
	}
}

func TestUnit2DeviceClass(t *testing.T) {
	// units as they reach Home Assistant, i.e. already normalized
	cases := map[string]string{
		"kWh": "energy",
		"kW":  "power",
		"Hz":  "frequency",
		"V":   "voltage",
		"A":   "current",
		"°C":  "temperature",
		"min": "duration",
		// no device class matches these, and a wrong one makes HA reject the entity
		"kΩ":   "",
		"p.u.": "",
		"%":    "",
		"":     "",
	}

	for unit, expected := range cases {
		if got := unit2DeviceClass(unit); got != expected {
			t.Errorf("unit2DeviceClass(%q) = %q, want %q", unit, got, expected)
		}
	}
}

func TestValueTemplate(t *testing.T) {
	if got, expected := valueTemplate("PV_Power", "0.1"), "{{ value_json.PV_Power|int * 0.1 }}"; got != expected {
		t.Errorf("valueTemplate with factor = %q, want %q", got, expected)
	}

	if got, expected := valueTemplate("SysState", ""), "{{ value_json.SysState|int }}"; got != expected {
		t.Errorf("valueTemplate without factor = %q, want %q", got, expected)
	}
}
