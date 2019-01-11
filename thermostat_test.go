package nest

import (
	"testing"
)

func TestGetThermostats(t *testing.T) {
	t.Run("One thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		thermostats, err := n.GetThermostats()
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := 1
			if len(thermostats) != expected {
				t.Fatalf("Expected %d thermostat(s), got %d", expected, len(thermostats))
			}
		}

		{
			expected := 30
			if thermostats[0].Humidity != expected {
				t.Fatalf("Expected Humidity to equal %d, got %d", expected, thermostats[0].Humidity)
			}
		}

		{
			expected := "abc"
			if thermostats[0].DeviceID != expected {
				t.Fatalf("Expected DeviceID to equal %s, got %s", expected, thermostats[0].DeviceID)
			}
		}

		{
			expected := "test thermostat"
			if thermostats[0].Name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, thermostats[0].Name)
			}
		}

		{
			expected := 68
			if thermostats[0].TargetTemperatureF != expected {
				t.Fatalf("Expected TargetTemperatureF to equal %d, got %d", expected, thermostats[0].TargetTemperatureF)
			}
		}
	})

	t.Run("No thermostats found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		thermostats, err := n.GetThermostats()
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := 0
			if len(thermostats) != expected {
				t.Fatalf("Expected %d thermostat(s), got %d", expected, len(thermostats))
			}
		}
	})
}

func TestGetThermostat(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		thermostat, err := n.GetThermostat("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := 30
			if thermostat.Humidity != expected {
				t.Fatalf("Expected Humidity to equal %d, got %d", expected, thermostat.Humidity)
			}
		}

		{
			expected := "abc"
			if thermostat.DeviceID != expected {
				t.Fatalf("Expected DeviceID to equal %s, got %s", expected, thermostat.DeviceID)
			}
		}

		{
			expected := "test thermostat"
			if thermostat.Name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, thermostat.Name)
			}
		}

		{
			expected := 68
			if thermostat.TargetTemperatureF != expected {
				t.Fatalf("Expected TargetTemperatureF to equal %d, got %d", expected, thermostat.TargetTemperatureF)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetThermostat("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetThermostat("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat def not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetThermostatLocale(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		locale, err := n.GetThermostatLocale("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "en-US"
			if locale != expected {
				t.Fatalf("Expected Locale to equal %s, got %s", expected, locale)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetThermostatLocale("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetThermostatLocale("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat locale not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetThermostatSoftwareVersion(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		version, err := n.GetThermostatSoftwareVersion("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "1.0"
			if version != expected {
				t.Fatalf("Expected Software Version to equal %s, got %s", expected, version)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetThermostatSoftwareVersion("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetThermostatSoftwareVersion("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat software_version not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetThermostatName(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		name, err := n.GetThermostatName("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "test thermostat"
			if name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, name)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetThermostatName("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetThermostatName("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat name not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetThermostatLastConnection(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		lastConn, err := n.GetThermostatLastConnection("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "2016-12-31T23:59:59.000Z"
			if lastConn != expected {
				t.Fatalf("Expected Last Connection to equal %s, got %s", expected, lastConn)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetThermostatLastConnection("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetThermostatLastConnection("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat last_connection not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestIsThermostatOnline(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		isOnline, err := n.IsThermostatOnline("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "true"
			if isOnline != expected {
				t.Fatalf("Expected Is Online to equal %s, got %s", expected, isOnline)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.IsThermostatOnline("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.IsThermostatOnline("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat is_online not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetTemperatureScale(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		tempScale, err := n.GetTemperatureScale("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "F"
			if tempScale != expected {
				t.Fatalf("Expected Temperature Scale to equal %s, got %s", expected, tempScale)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetTemperatureScale("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetTemperatureScale("ghi")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat temperature_scale not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetTargetTemperature(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		temp, err := n.GetTargetTemperature("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "68"
			if temp != expected {
				t.Fatalf("Expected Target Temperature to equal %s, got %s", expected, temp)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetTargetTemperature("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetTargetTemperature("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat target_temperature_f not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetTargetHighLowTemperature(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		high, low, err := n.GetTargetHighLowTemperature("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "72"
			if high != expected {
				t.Fatalf("Expected Target Temperature High to equal %s, got %s", expected, high)
			}
		}

		{
			expected := "70"
			if low != expected {
				t.Fatalf("Expected Target Temperature Low to equal %s, got %s", expected, low)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, _, err := n.GetTargetHighLowTemperature("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, _, err := n.GetTargetHighLowTemperature("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat target_temperature_high_f not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetHVACMode(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		mode, err := n.GetHVACMode("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "heat-cool"
			if mode != expected {
				t.Fatalf("Expected HVAC Mode to equal %s, got %s", expected, mode)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetHVACMode("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetHVACMode("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat hvac_mode not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetThermostatLabel(t *testing.T) {
	t.Run("Thermostat found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		label, err := n.GetThermostatLabel("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "test thermostat label"
			if label != expected {
				t.Fatalf("Expected Label to equal %s, got %s", expected, label)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetThermostatLabel("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Thermostat not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetThermostatLabel("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Thermostat label not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestSetTemperatureScale(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Success", func(t *testing.T) {
		err := n.SetTemperatureScale("abc", "F")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.SetTemperatureScale("", "F")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid scale 1", func(t *testing.T) {
		err := n.SetTemperatureScale("abc", "")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Scale must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid scale 2", func(t *testing.T) {
		err := n.SetTemperatureScale("abc", "A")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Temperature Scale must be one of the following: [F C]"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestSetTargetTemperatureF(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Success", func(t *testing.T) {
		err := n.SetTargetTemperatureF("abc", 70)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.SetTargetTemperatureF("", 70)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid temperature", func(t *testing.T) {
		err := n.SetTargetTemperatureF("abc", 100)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Target Temperature must be in the range of 50 - 90"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestSetTargetTemperatureC(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.SetTargetTemperatureC("", 25)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid temperature", func(t *testing.T) {
		err := n.SetTargetTemperatureC("abc", 50)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Target Temperature must be in the range of 9 - 32"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid scale", func(t *testing.T) {
		err := n.SetTargetTemperatureC("abc", 30)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Temperature Scale must be set to C"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestSetTargetHighLowTemperatureF(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Success", func(t *testing.T) {
		err := n.SetTargetHighLowTemperatureF("abc", 70, 72)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.SetTargetHighLowTemperatureF("", 70, 72)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid high temperature", func(t *testing.T) {
		err := n.SetTargetHighLowTemperatureF("abc", 100, 72)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Target High Temperature must be in the range of 50 - 90"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid low temperature", func(t *testing.T) {
		err := n.SetTargetHighLowTemperatureF("abc", 70, 100)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Target Low Temperature must be in the range of 50 - 90"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestSetTargetHighLowTemperatureC(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.SetTargetHighLowTemperatureC("", 30, 32)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid high temperature", func(t *testing.T) {
		err := n.SetTargetHighLowTemperatureC("abc", 50, 32)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Target High Temperature must be in the range of 9 - 32"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid low temperature", func(t *testing.T) {
		err := n.SetTargetHighLowTemperatureC("abc", 30, 50)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Target Low Temperature must be in the range of 9 - 32"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid scale", func(t *testing.T) {
		err := n.SetTargetHighLowTemperatureC("abc", 30, 32)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Temperature Scale must be set to C"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestSetHVACMode(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Success", func(t *testing.T) {
		err := n.SetHVACMode("abc", "heat-cool")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.SetHVACMode("", "heat-cool")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid mode 1", func(t *testing.T) {
		err := n.SetHVACMode("abc", "")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "HVAC Mode must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid mode 2", func(t *testing.T) {
		err := n.SetHVACMode("abc", "invalid")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "HVAC Mode must be one of the following: [heat cool heat-cool eco off]"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestSetThermostatLabel(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Success", func(t *testing.T) {
		err := n.SetThermostatLabel("abc", "New Label")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.SetThermostatLabel("", "New label")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid label", func(t *testing.T) {
		err := n.SetThermostatLabel("abc", "")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Label must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestTurnOnFanTimer(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Success", func(t *testing.T) {
		err := n.TurnOnFanTimer("abc", 15)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.TurnOnFanTimer("", 15)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Invalid fan timer", func(t *testing.T) {
		err := n.TurnOnFanTimer("abc", 5)
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Fan Timer Duration must be one of the following: [15 30 45 60 120 240 480 720]"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestTurnOffFanTimer(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Success", func(t *testing.T) {
		err := n.TurnOffFanTimer("abc")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.TurnOffFanTimer("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Device ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}
