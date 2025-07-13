package errs

import (
	"slices"
	"testing"
)

func TestErrs_initIfNil(t *testing.T) {
	var err Errors
	err.initIfNil()
	if err == nil {
		t.Errorf("err.initIfNil(): got %v; want not nil", err)
	}
	tmp := err
	err.initIfNil()
	err["a"] = []string{"b"}
	if v, ok := tmp["a"]; !ok || v[0] != "b" {
		t.Errorf("err.initIfNil(): got %v; want %v", tmp, err)
	}
}

func TestErrs_Add(t *testing.T) {
	var err Errors
	err.Add("key", "value")
	if len(err) != 1 {
		t.Errorf("len(err) = %d; want 1", len(err))
	}
	err.Add("p", "q")
	if len(err) != 2 {
		t.Errorf("len(err) = %d; want 2", len(err))
	}
	err.Add("key", "test")
	if len(err) != 2 {
		t.Errorf("len(err) = %d; want 2", len(err))
	}
	if !slices.Equal(err["p"], []string{"q"}) {
		t.Errorf("err['p'] = %v; want %v", err["p"], []string{"q"})
	}
	if !slices.Equal(err["key"], []string{"value", "test"}) {
		t.Errorf("err['key'] = %v; want %v", err["key"], []string{"value", "test"})
	}
}

func TestErrs_AddMessages(t *testing.T) {
	var err Errors

	err.AddMessages("key", []string{})
	if err == nil {
		t.Errorf("err.AddMessages(): got %v; want not nil", err)
	}
	if len(err) != 0 {
		t.Errorf("len(err) = %d; want 0", len(err))
	}

	err.AddMessages("key", []string{"value", "test"})
	if len(err) != 1 {
		t.Errorf("len(err) = %d; want 1", len(err))
	}

	err.AddMessages("q", []string{"r"})
	if len(err) != 2 {
		t.Errorf("len(err) = %d; want 2", len(err))
	}

	err.addMessages("key", []string{"top", "down"})
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
	var err1, err2 Errors

	err1.AddErrors("xxx", err2)
	if len(err1) != 0 {
		t.Errorf("len(err1) = %d; want 0", len(err1))
	}

	err2.AddMessages("zzz", []string{})
	err1.AddErrors("xxx", err2)
	if len(err1) != 0 {
		t.Errorf("len(err1) = %d; want 0", len(err1))
	}

	err2.AddMessages("p", []string{"1", "2"})
	err2.addMessages("q", []string{"3"})
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
