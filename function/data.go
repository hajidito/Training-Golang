package function

import "assignment/models"

func DataTeman() []models.Teman{
	temans := []models.Teman{
		{"Wasis", "Jakarta", "Developer", "Suka bahasa pemrograman Go"},
		{"Dio", "Surabaya", "Vokalis", "Ingin belajar bahasa pemrograman baru"},
		{"Juan", "Jakarta", "QA", "Suka bahasa pemrograman Go"},
	}
	return temans
}
