package stats

import (
	"reflect"
	"testing"

	"github.com/Qadriddin-dev/bank_1/pkg/types"
)


func TestCategoriesAvg_multiple(t *testing.T) {
	payments := []types.Payment{
	  {
		ID:       1,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       2,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       3,
		Category: "restaurant",
		Amount:   10_00,
	  },
	  {
		ID:       2,
		Category: "restaurant",
		Amount:   100_00,
	  },
	}
	  expected := map[types.Category]types.Money{
	  "cafe":100_00,
	  "restaurant":55_00,
	}
  
  
	res := CategoriesAvg(payments)
  
  
	if !reflect.DeepEqual(expected, res) {
		t.Errorf(" got > %v want > %v", res, expected)
	  }
  
  }
  func TestPeriodsDynamic_empty(t *testing.T) {
	first, second := map[types.Category]types.Money{}, map[types.Category]types.Money{}
	want := map[types.Category]types.Money{}
	got := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
func TestPeriodsDynamic_oneParam(t *testing.T) {
	first, second := map[types.Category]types.Money{
		"car": 10,
	}, map[types.Category]types.Money{
		"food": 5,
	}
	want := map[types.Category]types.Money{
		"car":  -10,
		"food": 5,
	}
	got := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestPeriodsDynamic_manyParam(t *testing.T) {
	first, second := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}, map[types.Category]types.Money{
		"auto":   10,
		"food":   25,
		"mobile": 5,
	}
	want := map[types.Category]types.Money{
		"auto":   0,
		"food":   5,
		"mobile": 5,
	}
	got := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}