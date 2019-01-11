package nest

import "testing"

func TestGetSmokeCOAlarms(t *testing.T) {
	t.Run("One Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		alarms, err := n.GetSmokeCOAlarms()
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := 1
			if len(alarms) != expected {
				t.Fatalf("Expected %d smoke/co alarm(s), got %d", expected, len(alarms))
			}
		}

		{
			expected := "en-US"
			if alarms[0].Locale != expected {
				t.Fatalf("Expected Locale to equal %s, got %s", expected, alarms[0].Locale)
			}
		}

		{
			expected := "abc"
			if alarms[0].DeviceID != expected {
				t.Fatalf("Expected DeviceID to equal %s, got %s", expected, alarms[0].DeviceID)
			}
		}

		{
			expected := "test smoke alarm"
			if alarms[0].Name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, alarms[0].Name)
			}
		}

		{
			expected := true
			if alarms[0].IsOnline != expected {
				t.Fatalf("Expected IsOnline to equal %v, got %v", expected, alarms[0].IsOnline)
			}
		}
	})

	t.Run("No Smoke/CO Alarms found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		alarms, err := n.GetSmokeCOAlarms()
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := 0
			if len(alarms) != expected {
				t.Fatalf("Expected %d smoke/co alarm(s), got %d", expected, len(alarms))
			}
		}
	})
}

func TestGetSmokeCOAlarm(t *testing.T) {
	t.Run("Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		alarm, err := n.GetSmokeCOAlarm("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "en-US"
			if alarm.Locale != expected {
				t.Fatalf("Expected Locale to equal %s, got %s", expected, alarm.Locale)
			}
		}

		{
			expected := "abc"
			if alarm.DeviceID != expected {
				t.Fatalf("Expected DeviceID to equal %s, got %s", expected, alarm.DeviceID)
			}
		}

		{
			expected := "test smoke alarm"
			if alarm.Name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, alarm.Name)
			}
		}

		{
			expected := true
			if alarm.IsOnline != expected {
				t.Fatalf("Expected IsOnline to equal %v, got %v", expected, alarm.IsOnline)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetSmokeCOAlarm("")
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

	t.Run("Smoke/CO Alarm not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetSmokeCOAlarm("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Smoke/CO Alarm def not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetSmokeCOAlarmLocale(t *testing.T) {
	t.Run("Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		locale, err := n.GetSmokeCOAlarmLocale("abc")
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

		_, err := n.GetSmokeCOAlarmLocale("")
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

	t.Run("Smoke/CO Alarm not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetSmokeCOAlarmLocale("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Smoke Co Alarm locale not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetSmokeCOAlarmSoftwareVersion(t *testing.T) {
	t.Run("Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		version, err := n.GetSmokeCOAlarmSoftwareVersion("abc")
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

		_, err := n.GetSmokeCOAlarmSoftwareVersion("")
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

	t.Run("Smoke/CO Alarm not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetSmokeCOAlarmSoftwareVersion("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Smoke Co Alarm software_version not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetSmokeCOAlarmName(t *testing.T) {
	t.Run("Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		name, err := n.GetSmokeCOAlarmName("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "test smoke alarm"
			if name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, name)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetSmokeCOAlarmName("")
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

	t.Run("Smoke/CO Alarm not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetSmokeCOAlarmName("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Smoke Co Alarm name not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetSmokeCOAlarmLastConnection(t *testing.T) {
	t.Run("Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		lastConn, err := n.GetSmokeCOAlarmLastConnection("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "2019-01-02T14:27:53.729Z"
			if lastConn != expected {
				t.Fatalf("Expected Last Connection to equal %s, got %s", expected, lastConn)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetSmokeCOAlarmLastConnection("")
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

	t.Run("Smoke/CO Alarm not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetSmokeCOAlarmLastConnection("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Smoke Co Alarm last_connection not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestIsSmokeCOAlarmOnline(t *testing.T) {
	t.Run("Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		isOnline, err := n.IsSmokeCOAlarmOnline("abc")
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

		_, err := n.IsSmokeCOAlarmOnline("")
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

	t.Run("Smoke/CO Alarm not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.IsSmokeCOAlarmOnline("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Smoke Co Alarm is_online not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetSmokeCOAlarmBatteryHealth(t *testing.T) {
	t.Run("Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		batteryHealth, err := n.GetSmokeCOAlarmBatteryHealth("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "ok"
			if batteryHealth != expected {
				t.Fatalf("Expected Battery Health to equal %s, got %s", expected, batteryHealth)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetSmokeCOAlarmBatteryHealth("")
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

	t.Run("Smoke/CO Alarm not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetSmokeCOAlarmBatteryHealth("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Smoke Co Alarm battery_health not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetCOAlarmState(t *testing.T) {
	t.Run("Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		alarmState, err := n.GetCOAlarmState("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "ok"
			if alarmState != expected {
				t.Fatalf("Expected CO Alarm State to equal %s, got %s", expected, alarmState)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetCOAlarmState("")
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

	t.Run("Smoke/CO Alarm not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetCOAlarmState("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Smoke Co Alarm co_alarm_state not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetSmokeAlarmState(t *testing.T) {
	t.Run("Smoke/CO Alarm found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		alarmState, err := n.GetSmokeAlarmState("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "ok"
			if alarmState != expected {
				t.Fatalf("Expected Smoke Alarm State to equal %s, got %s", expected, alarmState)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetSmokeAlarmState("")
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

	t.Run("Smoke/CO Alarm not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetSmokeAlarmState("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Smoke Co Alarm smoke_alarm_state not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}
