package dataconversion_test

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestInt2Float(t *testing.T) {

	// Test case 1
	input := 10
	expectedOutput := 10.0
	output := float64(input)
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}

	// Test case 2
	input = -10
	expectedOutput = -10.0
	output = float64(input)
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}
}

func TestFloat2Int(t *testing.T) {

	// Test case 1
	input := 10.0
	expectedOutput := 10
	output := int(input)
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}

	// Test case 2
	input = -10.0
	expectedOutput = -10
	output = int(input)
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}
}

func TestInt82Float32(t *testing.T) {
	// Test case 1
	input := int8(10)
	expectedOutput := float32(10)
	output := float32(input)
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}

	// Test case 2
	input = int8(-10)
	expectedOutput = float32(-10)
	output = float32(input)
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}
}

func TestFloat322Float64(t *testing.T) {
	// Test case 1
	input := float32(10)
	expectedOutput := float64(10)
	output := float64(input)
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}

	// Test case 2
	input = float32(-10)
	expectedOutput = float64(-10)
	output = float64(input)
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}

}

func TestInt2String(t *testing.T) {
	input := 10
	expectedOutput := "10"
	output := strconv.Itoa(input)
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}
}

func TestString2Int(t *testing.T) {
	intput := "10"
	expectedOutput := 10
	output, err := strconv.Atoi(intput)
	if err != nil {
		t.Errorf("Error converting string to int: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}
}

func TestString2Float32(t *testing.T) {

	intput := "10.5"
	expectedOutput := float32(10.5)
	output, err := strconv.ParseFloat(intput, 32)
	if err != nil {
		t.Errorf("Error converting string to float32: %v", err)
	}
	if float32(output) != expectedOutput {
		t.Errorf("Expected output %v, got %v", expectedOutput, output)
	}
}

type Person struct {
	Name string `json:"name"`
}

func TestEncodingAndDecoding(t *testing.T) {

	person := Person{
		Name: "John",
	}

	// Encode the person struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		t.Errorf("Error encoding person struct to JSON: %v", err)
	}
	t.Logf("JSON data: %s", jsonData)

	// Decode the JSON data back to a person struct
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		t.Errorf("Error decoding JSON data to person struct: %v", err)
	}
	t.Logf("Decoded person name: %v", decodedPerson.Name)
}

type Share interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}
func TestInterfaceConversion(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 3}
	s := []Share{r, c}
	for _, item := range s {
		area := item.Area()
		t.Logf("Area: %f\n", area)
	}
}

func TestArithmeticOperations(t *testing.T) {
	var length int = 10
	var Width float32 = 5.5
	result := float32(length) * Width
	t.Logf("Result: %f\n", result)
}

func TestAssignValue(t *testing.T) {
	var a int = 10
	result := strconv.Itoa(a)
	t.Logf("Result: %s\n", result)
}

func processFloat64(f float64) {
	// do something with f
}

func TestFunctionCall(t *testing.T) {
	var input int = 10
	processFloat64(float64(input))
}
