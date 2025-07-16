package errs

import (
	"slices"
	"testing"
)

func TestErrs_Add(t *testing.T) {
	var err = Errors{}

	err.Add("key")
	if len(err) != 0 {
		t.Errorf("len(err) = %d; want 0", len(err))
	}

	err.Add("key", "value", "test")
	if len(err) != 1 {
		t.Errorf("len(err) = %d; want 1", len(err))
	}
	if len(err["key"]) != 2 {
		t.Errorf(`len(err["key"]) = %d; want 2`, len(err["key"]))
	}
	if !slices.Equal(err["key"], []string{"value", "test"}) {
		t.Errorf(`len(err["key"]) = %v; want %v`, err["key"], []string{"value", "test"})
	}

	err.Add("q", "r")
	if len(err) != 2 {
		t.Errorf("len(err) = %d; want 2", len(err))
	}
	if len(err["q"]) != 1 {
		t.Errorf(`len(err["q"]) = %d; want 1`, len(err["q"]))
	}
	if !slices.Equal(err["q"], []string{"r"}) {
		t.Errorf(`len(err["q"]) = %v; want %v`, err["q"], []string{"r"})
	}

	err.Add("key", "top", "down")
	if len(err) != 2 {
		t.Errorf("len(err) = %d; want 2", len(err))
	}
	if len(err["key"]) != 4 {
		t.Errorf("len(err['key']) = %d; want 4", len(err))
	}
	if len(err["q"]) != 1 {
		t.Errorf("len(err['q']) = %d; want 1", len(err))
	}
	if !slices.Equal(err["key"], []string{"value", "test", "top", "down"}) {
		t.Errorf("err['key'] = %v; want %v", err["key"], []string{"value", "test", "top", "down"})
	}
	if !slices.Equal(err["q"], []string{"r"}) {
		t.Errorf("err['q'] = %v; want %v", err["q"], []string{"r"})
	}
}

func TestErrs_AddErrors(t *testing.T) {
	var err1 = Errors{}
	var err2 = Errors{}

	err1.AddErrors("xxx", err2)
	if len(err1) != 0 {
		t.Errorf("len(err1) = %d; want 0", len(err1))
	}

	err2.Add("zzz")
	err1.AddErrors("xxx", err2)
	if len(err1) != 0 {
		t.Errorf("len(err1) = %d; want 0", len(err1))
	}

	err2.Add("p", "1", "2")
	err2.Add("q", "3")
	err1.AddErrors("i.", err2)
	err1.AddErrors("j.", err2)
	err1.AddErrors("i.", err2)
	if len(err1) != 4 {
		t.Errorf("len(err1) = %d; want 4", len(err1))
	}
	if !slices.Equal(err1["i.p"], []string{"1", "2", "1", "2"}) {
		t.Errorf("err['i.p'] = %v; want %v", err1["i.p"], []string{"1", "2", "1", "2"})
	}
	if !slices.Equal(err1["i.q"], []string{"3", "3"}) {
		t.Errorf("err['i.q'] = %v; want %v", err1["i.q"], []string{"3", "3"})
	}
	if !slices.Equal(err1["j.p"], []string{"1", "2"}) {
		t.Errorf("err['j.p'] = %v; want %v", err1["j.p"], []string{"1", "2"})
	}
	if !slices.Equal(err1["j.q"], []string{"3"}) {
		t.Errorf("err['j.q'] = %v; want %v", err1["j.q"], []string{"3"})
	}
}
