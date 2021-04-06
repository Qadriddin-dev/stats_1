package stats

import (
	"github.com/Qadriddin-dev/bank_1/pkg/types"
)


func Avg(payments []types.Payment) types.Money {
  var sum types.Money
  i := 0
  for _, v := range payments {
    if v.Status == types.StatusFail {
      continue
    }
    sum += v.Amount
    i++
  }
  return sum / types.Money(i)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}

//CategoriesAvg func
func CategoriesAvg(payments []types.Payment) ( map[types.Category]types.Money) {

	mp := make(map[types.Category]types.Money)
	for _, v := range payments {
  
	  if _, er := mp[v.Category]; er {
		continue
	  }
	  filtered := FilterByCategory(payments, v.Category)
	  mp[v.Category]=Avg(filtered)
	}
  
  
	return mp
  }
  
  
  //FilterByCategory func
  func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment{
  
	var filtered []types.Payment
  
	for _, v := range payments {
	  if v.Category == category {
		filtered = append(filtered, v)
	  }
	}
	return filtered
  }

func PeriodsDynamic(first, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	for key, amount := range second {
		result[key] += amount
	}
	for key, amount := range first {
		result[key] -= amount
	}
	return result
}
