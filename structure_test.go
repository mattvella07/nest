package nest

import "testing"

func TestGetStructures(t *testing.T) {
	t.Run("One structure found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		structures, err := n.GetStructures()
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := 1
			if len(structures) != expected {
				t.Fatalf("Expected %d thermostat(s), got %d", expected, len(structures))
			}
		}

		{
			expected := "abc123"
			if structures[0].StructureID != expected {
				t.Fatalf("Expected StructureID to equal %s, got %s", expected, structures[0].StructureID)
			}
		}

		{
			expected := "home"
			if structures[0].Away != expected {
				t.Fatalf("Expected Away to equal %s, got %s", expected, structures[0].Away)
			}
		}

		{
			expected := "test structure"
			if structures[0].Name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, structures[0].Name)
			}
		}

		{
			expected := "America/Chicago"
			if structures[0].TimeZone != expected {
				t.Fatalf("Expected TimeZone to equal %s, got %s", expected, structures[0].TimeZone)
			}
		}
	})

	t.Run("No structures found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		structures, err := n.GetStructures()
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := 0
			if len(structures) != expected {
				t.Fatalf("Expected %d structure(s), got %d", expected, len(structures))
			}
		}
	})
}

func TestGetStructure(t *testing.T) {
	t.Run("Structure found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		structure, err := n.GetStructure("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "abc123"
			if structure.StructureID != expected {
				t.Fatalf("Expected StructureID to equal %s, got %s", expected, structure.StructureID)
			}
		}

		{
			expected := "home"
			if structure.Away != expected {
				t.Fatalf("Expected Away to equal %s, got %s", expected, structure.Away)
			}
		}

		{
			expected := "test structure"
			if structure.Name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, structure.Name)
			}
		}

		{
			expected := "America/Chicago"
			if structure.TimeZone != expected {
				t.Fatalf("Expected TimeZone to equal %s, got %s", expected, structure.TimeZone)
			}
		}
	})

	t.Run("Invalid structure id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetStructure("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Structure not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetStructure("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure def not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetStructureThermostats(t *testing.T) {
	t.Run("Structure found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		thermostats, err := n.GetStructureThermostats("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "[\"123\"]"
			if thermostats != expected {
				t.Fatalf("Expected Thermostats to equal %s, got %s", expected, thermostats)
			}
		}
	})

	t.Run("Invalid structure id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetStructureThermostats("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Structure not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetStructureThermostats("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure thermostats not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetStructureSmokeCOAlarms(t *testing.T) {
	t.Run("Structure found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		alarms, err := n.GetStructureSmokeCOAlarms("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "[\"456\"]"
			if alarms != expected {
				t.Fatalf("Expected Smoke/CO Alarms to equal %s, got %s", expected, alarms)
			}
		}
	})

	t.Run("Invalid structure id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetStructureSmokeCOAlarms("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Structure not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetStructureSmokeCOAlarms("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure smoke_co_alarms not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetStructureCameras(t *testing.T) {
	t.Run("Structure found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		cameras, err := n.GetStructureCameras("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "[\"789\"]"
			if cameras != expected {
				t.Fatalf("Expected Cameras to equal %s, got %s", expected, cameras)
			}
		}
	})

	t.Run("Invalid structure id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetStructureCameras("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Structure not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetStructureCameras("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure cameras not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetStructureOccupancyState(t *testing.T) {
	t.Run("Structure found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		state, err := n.GetStructureOccupancyState("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "home"
			if state != expected {
				t.Fatalf("Expected Occupancy State to equal %s, got %s", expected, state)
			}
		}
	})

	t.Run("Invalid structure id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetStructureOccupancyState("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Structure not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetStructureOccupancyState("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure away not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}

func TestGetStructureName(t *testing.T) {
	t.Run("Structure found", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		name, err := n.GetStructureName("abc")
		if err != nil {
			t.Fatal(err)
		}

		{
			expected := "test structure"
			if name != expected {
				t.Fatalf("Expected Name to equal %s, got %s", expected, name)
			}
		}
	})

	t.Run("Invalid structure id", func(t *testing.T) {
		n, server := createTestConnection(1)
		defer server.Close()

		_, err := n.GetStructureName("")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure ID must not be empty"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})

	t.Run("Structure not found", func(t *testing.T) {
		n, server := createTestConnection(2)
		defer server.Close()

		_, err := n.GetStructureName("def")
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		{
			expected := "Structure name not found"
			if err.Error() != expected {
				t.Fatalf("Expected error message to equal %s, got %s", expected, err.Error())
			}
		}
	})
}
