package jot

import (
	"reflect"
	"testing"
)

type TestPrinter struct {
	printCalled   bool
	printValues   []interface{}
	printfCalled  bool
	printfFormat  string
	printfValues  []interface{}
	printlnCalled bool
	printlnValues []interface{}
}

func (t *TestPrinter) Print(v ...interface{}) {
	t.printValues = v
	t.printCalled = true
}

func (t *TestPrinter) Printf(format string, v ...interface{}) {
	t.printfFormat = format
	t.printfValues = v
	t.printfCalled = true
}

func (t *TestPrinter) Println(v ...interface{}) {
	t.printlnValues = v
	t.printlnCalled = true
}

func TestPrintEnabled(t *testing.T) {
	p := &TestPrinter{}
	j := New(p)
	j.Enable()

	expected := []interface{}{"Some", "values", "passed", "in"}
	j.Print(expected...)

	if !reflect.DeepEqual(expected, p.printValues) {
		t.Fatal("Passed values do not match", expected, p.printValues)
	}
}

func TestPrintfEnabled(t *testing.T) {
	p := &TestPrinter{}
	j := New(p)
	j.Enable()

	format := "format %s %s %s %s"
	expected := []interface{}{"Some", "values", "passed", "in"}
	j.Printf(format, expected...)

	if format != p.printfFormat {
		t.Fatal("Passed format does not match", format, p.printfFormat)
	}

	if !reflect.DeepEqual(expected, p.printfValues) {
		t.Fatal("Passed values do not match", expected, p.printfValues)
	}
}

func TestPrintlnEnabled(t *testing.T) {
	p := &TestPrinter{}
	j := New(p)
	j.Enable()

	expected := []interface{}{"Some", "values", "passed", "in"}
	j.Println(expected...)

	if !reflect.DeepEqual(expected, p.printlnValues) {
		t.Fatal("Passed values do not match", expected, p.printlnValues)
	}
}

func TestPrintDisabled(t *testing.T) {
	p := &TestPrinter{}
	j := New(p)

	j.Print("Some", "values", "passed", "in")

	if p.printCalled {
		t.Fatal("Expected printer to not be called")
	}
}

func TestPrintfDisabled(t *testing.T) {
	p := &TestPrinter{}
	j := New(p)

	j.Printf("format %s %s %s %s", "Some", "values", "passed", "in")

	if p.printfCalled {
		t.Fatal("Expected printer to not be called")
	}
}

func TestPrintlnDisabled(t *testing.T) {
	p := &TestPrinter{}
	j := New(p)

	j.Println("Some", "values", "passed", "in")

	if p.printlnCalled {
		t.Fatal("Expected printer to not be called")
	}
}
