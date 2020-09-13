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
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// SumByMCC - функция сумм по коду mmc
func SumByMCC(transactions []*Transaction, mcc []string) int64 {
	var totalMcc int64
	for _, v := range transactions {
		if valInSlice(v.MccCode, mcc) == true {
			totalMcc += v.TranSum
		}
	}
	return totalMcc
}
