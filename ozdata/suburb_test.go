package ozdata

import "testing"

var validInput = "../data/suburbs.json"

var fileTests = []struct {
	in  string
	out bool
}{
	{"/tmp/sdfsddfsdf.json", false},
	{"../data/test.json", false},
	{validInput, true},
}

func Test_NewSuburbDataLoadsFromFile(t *testing.T) {
	for _, jt := range fileTests {
		output, err := NewSuburbs(jt.in)
		if err != nil && jt.out == true {
			t.Errorf("Input %q: Error: %q", jt.in, err)
		}

		if err != nil && jt.out == false {
			if len(output.Suburbs) > 0 {
				t.Errorf("Input %q: Error: %q", jt.in, err)
			}
		}
	}
}

var suburbTests = []struct {
	in  int64
	out bool
}{
	{2016, true},
	{2000, true},
	{2000000, false},
	{0000, false},
}

func Test_GettingSuburbByCode(t *testing.T) {
	suburbs, err := LoadSuburbs()
	if err != nil {
		t.Errorf("Input %q: Error: %q", validInput, err)
	}

	for _, st := range suburbTests {
		_, err = suburbs.GetSuburbsByPostCode(st.in)
		if err != nil && st.out == true {
			t.Errorf("Input %d: Error: %q", st.in, err)
		}
	}
}
