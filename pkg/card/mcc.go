package card

// TranslateMCC - TranslateMCC
func TranslateMCC(code string) string {
	// иниицализация словаря
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5912": "Аптеки",
	}
	//fmt.Println(mcc)

	// ищем значение code
	var totVal string
	if value, ok := mcc[code]; ok {
		totVal = value
	} else {
		totVal = "Категория не указана"
	}
	return totVal
}
