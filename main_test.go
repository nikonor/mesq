package mesq

import "testing"

func TestAddMessageToQ(t *testing.T) {
	cases := []struct {
		in    string
		cmd   string
		count int
	}{
		{`{"command":"UNSUNSCRIBE","citizen_id":222,"channel":2}`, "U", 1},
		{`{"command":"subscribe","citizen_id":222,"channel":1}`, "S", 2},
		{`{"command":"subscribe","citizen_id":222,"channel":2}`, "S", 1},
		{`{"command":"subscribe","citizen_id":222,"channel":3}`, "S", 2},
		{`{"command":"unsubscribe","citizen_id":222,"channel":1}`, "S", 1},
	}

	for _, c := range cases {
		if err := AddMessageToQ([]byte(c.in)); err != nil {
			t.Error(err)
		}
		if len(Heap.Dict) != c.count {
			t.Error("ошибка подсчета")
		}
	}

	showQ()
}

func TestMakeKey(t *testing.T) {
	cases := []struct {
		cmd       string
		citID     int64
		channel   int64
		outDirect string
		outBack   string
	}{
		{"subscribe", 1, 1, "s-1-1", "u-1-1"},
		{"SUBSCRIBE", 1, 1, "s-1-1", "u-1-1"},
		{"unsubscribe", 1, 1, "u-1-1", "s-1-1"},
		{"UBSUBSCRIBE", 1, 1, "u-1-1", "s-1-1"},
	}

	for _, c := range cases {
		s1, s2 := MakeKey(c.cmd, c.citID, c.channel)
		if s1 != c.outDirect || s2 != c.outBack {
			t.Errorf("%s != %s || %s != %s\n", s1, c.outDirect, s2, c.outBack)
		}
	}
}
