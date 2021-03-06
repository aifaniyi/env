package loader

import (
	"os"
	"testing"
)

var (
	env = "ENV"
)

func TestLoadTrimmed(t *testing.T) {
	os.Setenv(env, "")
	if loadTrimmed(env) != "" {
		t.Fatalf("%s should be empty", env)
	}

	expected := "VALUE1"
	os.Setenv(env, expected)
	if loadTrimmed(env) != expected {
		t.Fatalf("%s should be equal to %s", env, expected)
	}

	input := "    VALUE2   "
	expected = "VALUE2"
	os.Setenv(env, input)
	if loadTrimmed(env) != expected {
		t.Fatalf("%s should be equal to %s", env, expected)
	}

	input = "    VALUE3   "
	expected = "VALUE2"
	os.Setenv(env, input)
	if loadTrimmed(env) == expected {
		t.Fatalf("%s should not be equal to %s", env, expected)
	}
}

func TestLoadInt(t *testing.T) {
	// load default
	os.Setenv(env, "")
	expectedInt := int64(20)
	if LoadInt(env, 20) != expectedInt {
		t.Fatalf("%s should return default value %d", env, expectedInt)
	}

	os.Setenv(env, "300")
	expectedInt = int64(300)
	if LoadInt(env, 20) != expectedInt {
		t.Fatalf("%s should return provided value %d", env, expectedInt)
	}

	os.Setenv(env, "A300")
	expectedInt = int64(20)
	if LoadInt(env, 20) != expectedInt {
		t.Fatalf("%s should return default value %d", env, expectedInt)
	}
}

func TestLoadBool(t *testing.T) {
	// load default
	os.Setenv(env, "")
	expectedBool := true
	if LoadBool(env, true) != expectedBool {
		t.Fatalf("%s should return default value %t", env, expectedBool)
	}

	os.Setenv(env, "true")
	expectedBool = true
	if LoadBool(env, true) != expectedBool {
		t.Fatalf("%s should return provided value %t", env, expectedBool)
	}

	os.Setenv(env, "t")
	expectedBool = true
	if LoadBool(env, true) != expectedBool {
		t.Fatalf("%s should return default value %t", env, expectedBool)
	}
}

func TestLoadString(t *testing.T) {
	// load default
	os.Setenv(env, "")
	expectedString := "default"
	if LoadString(env, "default") != expectedString {
		t.Fatalf("%s should return default value %s", env, expectedString)
	}

	os.Setenv(env, "environment_value")
	expectedString = "environment_value"
	if LoadString(env, "default") != expectedString {
		t.Fatalf("%s should return provided value %s", env, expectedString)
	}
}

func TestLoadFloat(t *testing.T) {
	// load default
	os.Setenv(env, "")
	expectedFloat := float64(20.0)
	if LoadFloat(env, 20) != expectedFloat {
		t.Fatalf("%s should return default value %f", env, expectedFloat)
	}

	os.Setenv(env, "300")
	expectedFloat = float64(300)
	if LoadFloat(env, 20) != expectedFloat {
		t.Fatalf("%s should return provided value %f", env, expectedFloat)
	}

	os.Setenv(env, "A300")
	expectedFloat = float64(20)
	if LoadFloat(env, 20) != expectedFloat {
		t.Fatalf("%s should return default value %f", env, expectedFloat)
	}
}

func TestLoadArray(t *testing.T) {
	os.Setenv(env, "hello, world")

	el0 := "hello"
	el1 := "world"

	result := LoadArray(env, ",", []string{"default", "value"})
	if result[0] != el0 {
		print(t, "Failed to load correct array environment variable for el0\n", el0, result[0])
	}

	result = LoadArray(env, ",", []string{"default", "value"})
	if result[1] != el1 {
		print(t, "Failed to load correct array environment variable for el1\n", el1, result[1])
	}

}

func print(t *testing.T, message string, expected, found interface{}) {
	t.Fatalf("\n%s\n\texpected: \t%s\n\treceived: \t%s", message, expected, found)
}
