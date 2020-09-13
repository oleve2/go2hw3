package main

import (
	"fmt"

	"github.com/wool/go2hw3/pkg/card"
)

func main() {
	// ========================================================
	fmt.Println("==================================")
	card1 := &card.Card{
		ID: 1, Type: "Master", CardNumber: "1111 1111 1111 1111", CardDueDate: "2030-01-01",
		Transactions: []*card.Transaction{
			&card.Transaction{
				ID: 1, TranType: "purchase", TranSum: 735_55,
				TranDate: 1597293926, MccCode: "5411", Status: "done Супермаркеты"},
			&card.Transaction{
				ID: 2, TranType: "refill", TranSum: 2000_00,
				TranDate: 1597293926, MccCode: "", Status: "done"},
			&card.Transaction{
				ID: 3, TranType: "purchase", TranSum: 1203_91,
				TranDate: 1597293926, MccCode: "5812", Status: "done Рестораны"},
		},
	}
	fmt.Println("Данные карты:", card1)
	fmt.Println("Список транзакций:")
	for _, v := range card1.Transactions {
		fmt.Println(v)
	}

	// ========================================================
	fmt.Println("==================================")
	tr4 := &card.Transaction{
		ID: 4, TranType: "purchase", TranSum: 1_000_000_00,
		TranDate: 1597293926, MccCode: "Пополнение", Status: "done"}
	fmt.Println(tr4)
	card.AddTransaction(card1, tr4)
	fmt.Println("Список транзакций:")
	for _, v := range card1.Transactions {
		fmt.Println(v)
	}

	// ========================================================
	fmt.Println("==================================")
	codes := []string{"5411", "Пополнение"}
	fmt.Println("codes=", codes)
	//
	resSum := card.SumByMCC(card1.Transactions, codes)
	fmt.Println("mcc tran sum=", resSum)

}
