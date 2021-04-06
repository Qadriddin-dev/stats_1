package stats

import (
	"github.com/Qadriddin-dev/bank_1/pkg/types"
	"fmt"
)
func ExampleAvg() {
	payments := []types.Payment{
		{1, 10, "T", types.StatusOk},
		{1, 11, "U", types.StatusFail},
		{1, 12, "R", types.StatusInProgress},
	}
	fmt.Println(Avg(payments))
	//Output:
	// 11
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
	  {
		ID:2,
		Amount:53_00,
		Category: "Cafe",
		Status: types.StatusOk,
	  },
	  {
		ID:1,
		Amount:51_00,
		Category: "Cafe",
		Status: types.StatusFail,
	  },
	  {
		ID:3,
		Amount:52_00,
		Category: "Restaurant",
		Status: types.StatusOk,
	  },
	}
  
	fmt.Println(TotalInCategory(payments, "Cafe"))
  
	//Output: 5300
  }