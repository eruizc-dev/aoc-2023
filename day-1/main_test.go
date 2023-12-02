package main

import "testing"

func Test_calibrate_document(t *testing.T) {
	result := calibrate_document([]string {
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	})
	if result != 142 {
		t.Logf("Expected %d, got %d", 142, result)
		t.Fail()
	}
}

func Test_calibrate_line_num(t *testing.T) {
	test := []struct { line string; result int } {
		{ "1abc2", 12 },
		{ "pqr3stu8vwx", 38 },
		{ "a1b2c3d4e5f", 15 },
		{ "treb7uchet", 77 },
		{ "gsjgklneight6zqfz", 66 },
	}

	for _, expected := range test {
		result, err := calibrate_line(expected.line)

		if err != nil {
			t.Log(err)
			t.Fail()
		} else if result != expected.result {
			t.Logf("Expected %d, got %d", expected.result, result)
			t.Fail()
		}
	}
}

func Test_calibrate_line_str(t *testing.T) {
	test := []struct { line string; result int } {
		{ "1abc2", 12 },
		{ "pqr3stu8vwx", 38 },
		{ "a1b2c3d4e5f", 15 },
		{ "treb7uchet", 77 },
		{ "gsjgklneight6zqfz", 66 },
	}

	for _, expected := range test {
		result, err := calibrate_line(expected.line)

		if err != nil {
			t.Log(err)
			t.Fail()
		} else if result != expected.result {
			t.Logf("Expected %d, got %d", expected.result, result)
			t.Fail()
		}
	}
}
