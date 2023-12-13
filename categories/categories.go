package categories

func GetCategoriesMap() map[string]string {
	return map[string]string{
		"amusement_park": "Парк развлечений",
		"aquarium":       "Океанариум",
		"art_gallery":    "Галерея",
		"bar":            "Бар",
		"mosque":         "Мечеть",
		"movie_theater":  "Кинотеатр",
		"bowling_alley":  "Боулинг",
		"museum":         "Музей",
		"cafe":           "Кафе",
		"night_club":     "Клуб",
		"park":           "Парк",
		"casino":         "Казино",
		"church":         "Церковь",
		"city_hall":      "Мэрия",
		"restaurant":     "Ресторан",
		"shopping_mall":  "Торговый центр",
		"stadium":        "Стадион",
		"synagogue":      "Синагога",
		"hindu_temple":   "Храм Хинду",
		"zoo":            "Зоопарк",
		//"tourist_attraction": "Достопримечательность",
	}
}

func GetReversedCategoryMap() (result map[string]string) {
	result = make(map[string]string, 0)
	categories := GetCategoriesMap()
	for k, v := range categories {
		result[v] = k
	}
	return
}
