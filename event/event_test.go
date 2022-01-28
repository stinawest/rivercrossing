package event

import "testing"



func TestPut(t *testing.T) {
	wanted := "[kylling rev korn person ---\\ \\_kylling_/ ______________/---]"
	state := PutKyll(kylling) 
	if state != wanted {
			t.Errorf("Feil, fikk %q, ønsket %q", state, wanted)
	}
}



func TestCrossFor(t *testing.T) {
    wanted := "[kylling rev korn person ---\\____________________ \\____//---]"
    state := CrossRiverFor()
    if state != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q", state, wanted)
    }
}
