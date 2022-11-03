package resources

import (
	"reflect"
	"testing"
)

func Test_GetArgs(t *testing.T) {
	t.Run("Return string type value.", func(t *testing.T) {
		args := GetArgs(struct {
			StringValue string `thanos:"arg=%s"`
		}{
			StringValue: "val",
		})

		wanted := []string{"arg=val"}
		if !reflect.DeepEqual(args, wanted) {
			t.Fatalf("GetArgs != %v, wanted %v", args, wanted)
		}

	})

	t.Run("Return int type value.", func(t *testing.T) {
		args := GetArgs(struct {
			IntValue int `thanos:"arg=%d"`
		}{
			IntValue: 1,
		})

		wanted := []string{"arg=1"}
		if !reflect.DeepEqual(args, wanted) {
			t.Fatalf("GetArgs != %v, wanted %v", args, wanted)
		}

	})

	t.Run("Return bool type value.", func(t *testing.T) {
		args := GetArgs(struct {
			BoolValue bool `thanos:"arg"`
		}{
			BoolValue: true,
		})

		wanted := []string{"arg"}
		if !reflect.DeepEqual(args, wanted) {
			t.Fatalf("GetArgs != %v, wanted %v", args, wanted)
		}
	})

	t.Run("Return bool pointer type value.", func(t *testing.T) {
		b := true
		args := GetArgs(struct {
			BoolValue *bool `thanos:"arg"`
		}{
			BoolValue: &b,
		})

		wanted := []string{"arg"}
		if !reflect.DeepEqual(args, wanted) {
			t.Fatalf("GetArgs != %v, wanted %v", args, wanted)
		}
	})
}
