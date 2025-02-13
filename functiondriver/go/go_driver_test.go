package godriver

import (
	"context"
	"os"
	"testing"

	"github.com/cofunclabs/cofunc/service/resource"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	dr := New("print", "print", "")
	if dr == nil {
		t.FailNow()
	}
	assert.Equal(t, "print", dr.fname)
	assert.Equal(t, "print", dr.path)
	err := dr.Load(context.Background(), resource.Resources{
		Logwriter: os.Stdout,
	})
	args := dr.mergeArgs(map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	})
	assert.NoError(t, err)
	assert.Len(t, dr.manifest.Args, 0)
	assert.Len(t, args, 3)
}

func TestRun(t *testing.T) {
	dr := New("print", "print", "")
	if dr == nil {
		t.FailNow()
	}
	err := dr.Load(context.Background(), resource.Resources{
		Logwriter: os.Stdout,
	})
	assert.NoError(t, err)
	out, err := dr.Run(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, "ok", out["status"])
}
