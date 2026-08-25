package model

import "testing"

func TestTransitionValidation(t *testing.T) {
	if ValidateTransition(Transition{RecordID: "x", From: "draft", To: "approved", Actor: "u"}) == nil {
		t.Fatal("invalid transition accepted")
	}
	if NormalizeStatus("") != "draft" {
		t.Fatal("default")
	}
}
