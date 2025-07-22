package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrs_Add(t *testing.T) {
	var err Errors

	err.Add("key")
	assert.Nil(t, err)

	err.Add("key", "value", "test")
	assert.Equal(t, err, Errors{"key": {"value", "test"}})

	err.Add("q", "r").Add("key", "top", "down")
	assert.Equal(t, err, Errors{"key": {"value", "test", "top", "down"}, "q": {"r"}})
}

func TestErrs_AddErrors(t *testing.T) {
	var err1 Errors
	var err2 Errors

	err1.AddErrors("xxx", err2)
	assert.Nil(t, err1)

	err1.AddErrors("qqq", *(err2.Add("", "aaa").Add("3", "bbb").Add("yyy", "ccc")))
	assert.Equal(t, err1, Errors{"qqq": {"aaa"}, "qqq.3": {"bbb"}, "qqq.yyy": {"ccc"}})
}

func TestSetDelimiter(t *testing.T) {
	SetDelimiter("~^~")

	var err1 Errors
	var err2 Errors

	err1.AddErrors("qqq", *(err2.Add("", "aaa").Add("3", "bbb").Add("yyy", "ccc")))
	assert.Equal(t, err1, Errors{"qqq": {"aaa"}, "qqq~^~3": {"bbb"}, "qqq~^~yyy": {"ccc"}})
}
