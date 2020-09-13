package card

// Card - Card
type Card struct {
	ID           int64
	Type         string
	CardNumber   string
	CardDueDate  string
	Transactions []*Transaction
}

// Transaction - Transaction
type Transaction struct {
	ID       int64
	TranType string
	TranSum  int64
	TranDate int64 // unix timestamp
	MccCode  string
	Status   string
}

// AddTransaction - добавление транцакции
func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

// функция проверки вхождения
func valInSlice(val string, arr []string) bool {
	var flg bool = false
	for _, v := range arr {
		//fmt.Println(reflect.TypeOf(val), reflect.TypeOf(v))
		if v == val {
			flg = true
		}
	}
	return flg
}

// SumByMCC - функция сумм по коду mmc
func SumByMCC(transactions []*Transaction, mcc []string) int64 {
	// итоговая сумма
	var totalMcc int64 = 0
	// основной цикл
	for _, v := range transactions {
		//fmt.Println( v.MccCode, valInSlice(v.MccCode, mcc) )
		if valInSlice(v.MccCode, mcc) == true {
			//fmt.Println(i,v, v.MccCode)
			totalMcc += v.TranSum
		}
	}
	//fmt.Println(totalMcc)
	return totalMcc
}
