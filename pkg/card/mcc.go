package card

const categoryNotFound = "Категория не найдена"

// TranslateMCC - TranslateMCC
func TranslateMCC(code string) string {
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5912": "Аптеки",
	}
	if value, ok := mcc[code]; ok {
		return value
	}
	return categoryNotFound
}
