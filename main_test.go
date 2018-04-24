package mesq

import "testing"

func TestAddMessageToQ(t *testing.T){
    cases := []struct{
        in string
        cmd string
    }{
        {`{"command":"subscribe","citizen_id":222,"channels":[1,3,5]}`,"S"},
    }
    
    for _, c := range  cases {
        if err := AddMessageToQ([]byte(c.in)); err != nil {
            t.Error(err)
        }
    }
}

func TestMakeKey (t *testing.T) {
    cases := []struct {
        cmd string
        citID int64
        channel int64
        out string
    }{
        {"subscribe",1,1,"s-1-1"},
        {"SUBSCRIBE",1,1,"s-1-1"},
    }
    
    for _, c := range cases {
        s := MakeKey(c.cmd, c.citID, c.channel)
        if s != c.out {
            t.Errorf("%s != %s\n", s, c.out)
        }
    }
}