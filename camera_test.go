package nest

import "testing"

func TestGetCameras(t *testing.T) {
	t.Run("One camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		cameras, err := n.GetCameras()
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := 1
			if len(cameras) != expected {
				t.Fatalf("Expected %d camera(s), got %d", expected, len(cameras))
			}
		}

		{
			expected := true
			if cameras[0].IsOnline != expected {
				t.Fatalf("Expected IsOnline to equal %v, got %v", expected, cameras[0].IsOnline)
			}
		}

		{
			expected := "abc"
			if cameras[0].DeviceID != expected {
				t.Fatalf("Expected DeviceID to equal %s, got %s", expected, cameras[0].DeviceID)
			}
		}

		{
			expected := "test camera"
			if cameras[0].Name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, cameras[0].Name)
			}
		}

		{
			expected := "https://home.nest.com/cameras/abc?auth=camera_token"
			if cameras[0].WebURL != expected {
				t.Fatalf("Expected WebURL to equal %s, got %s", expected, cameras[0].WebURL)
			}
		}
	})

	t.Run("No cameras found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		cameras, err := n.GetCameras()
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := 0
			if len(cameras) != expected {
				t.Fatalf("Expected %d camera(s), got %d", expected, len(cameras))
			}
		}
	})
}

func TestGetCamera(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		cameras, err := n.GetCamera("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := true
			if cameras.IsOnline != expected {
				t.Fatalf("Expected IsOnline to equal %v, got %v", expected, cameras.IsOnline)
			}
		}

		{
			expected := "abc"
			if cameras.DeviceID != expected {
				t.Fatalf("Expected DeviceID to equal %s, got %s", expected, cameras.DeviceID)
			}
		}

		{
			expected := "test camera"
			if cameras.Name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, cameras.Name)
			}
		}

		{
			expected := "https://home.nest.com/cameras/abc?auth=camera_token"
			if cameras.WebURL != expected {
				t.Fatalf("Expected WebURL to equal %s,  got %s", expected, cameras.WebURL)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetCamera("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetCamera("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera def not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetCameraSoftwareVersion(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		version, err := n.GetCameraSoftwareVersion("abc")
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

		_, err := n.GetCameraSoftwareVersion("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetCameraSoftwareVersion("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera software_version not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetCameraName(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		name, err := n.GetCameraName("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "test camera"
			if name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, name)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetCameraName("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetCameraName("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera name not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestIsCameraOnline(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		isOnline, err := n.IsCameraOnline("abc")
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

		_, err := n.IsCameraOnline("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.IsCameraOnline("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera is_online not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestIsCameraStreaming(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		isStreaming, err := n.IsCameraStreaming("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "true"
			if isStreaming != expected {
				t.Fatalf("Expected Is Streaming to equal %s, got %s", expected, isStreaming)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.IsCameraStreaming("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.IsCameraStreaming("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera is_streaming not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestIsCameraAudioInputEnabled(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		isAudioEnabled, err := n.IsCameraAudioInputEnabled("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "true"
			if isAudioEnabled != expected {
				t.Fatalf("Expected Is Audio Input Enabled to equal %s, got %s", expected, isAudioEnabled)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.IsCameraAudioInputEnabled("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.IsCameraAudioInputEnabled("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera is_audio_input_enabled not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestIsCameraVideoHistoryEnabled(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		isVideoHistoryEnabled, err := n.IsCameraVideoHistoryEnabled("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "true"
			if isVideoHistoryEnabled != expected {
				t.Fatalf("Expected Is Video History Enabled to equal %s, got %s", expected, isVideoHistoryEnabled)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.IsCameraVideoHistoryEnabled("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.IsCameraVideoHistoryEnabled("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera is_video_history_enabled not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetCameraWebURL(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		webURL, err := n.GetCameraWebURL("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "https://home.nest.com/cameras/abc?auth=camera_token"
			if webURL != expected {
				t.Fatalf("Expected Web URL to equal %s, got %s", expected, webURL)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetCameraWebURL("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetCameraWebURL("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera web_url not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetCameraAppURL(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		appURL, err := n.GetCameraAppURL("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "nestmobile://cameras/abc?auth=camera_token"
			if appURL != expected {
				t.Fatalf("Expected Web URL to equal %s, got %s", expected, appURL)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetCameraAppURL("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetCameraAppURL("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera app_url not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetCameraLastEvent(t *testing.T) {
	t.Run("Camera found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		lastEvent, err := n.GetCameraLastEvent("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "event"
			if lastEvent != expected {
				t.Fatalf("Expected Web URL to equal %s, got %s", expected, lastEvent)
			}
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetCameraLastEvent("")
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

	t.Run("Camera not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetCameraLastEvent("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Camera last_event not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestTurnOnStreaming(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Success", func(t *testing.T) {
		err := n.TurnOnStreaming("abc")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.TurnOnStreaming("")
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

func TestTurnOffStreaming(t *testing.T) {
	n, server := createTestConnection(1)
	defer server.Close()

	t.Run("Success", func(t *testing.T) {
		err := n.TurnOffStreaming("abc")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Invalid device id", func(t *testing.T) {
		err := n.TurnOffStreaming("")
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
