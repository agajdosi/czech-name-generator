package main

import (
	"math/rand"
	"time"
)

func seed() {
	rand.Seed(time.Now().UTC().UnixNano())

	return
}

//GenerateName generates a random name
func GenerateName(sex string) (string, string) {
	seed()
	if sex == "M" {

		name := maleNames[rand.Intn(len(maleNames))]
		surname := maleSurnames[rand.Intn(len(maleSurnames))]

		return name, surname
	}

	return "", ""
}

//GeneratePassword generates random password
func GeneratePassword(n int) string {
	seed()
	var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//GenerateBirthday generates a random birthday
func GenerateBirthday() time.Time {
	seed()
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2000, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min

	return time.Unix(sec, 0)
}

var maleSurnames = [...]string{
	"Novák",
	"Svoboda",
	"Novotný",
	"Dvořák",
	"Černý",
	"Procházka",
	"Kučera",
	"Veselý",
	"Horák",
	"Němec",
	"Marek",
	"Pospíšil",
	"Pokorný",
	"Hájek",
	"Král",
	"Jelínek",
	"Růžička",
	"Beneš",
	"Fiala",
	"Sedláček",
	"Doležal",
	"Zeman",
	"Kolář",
	"Navrátil",
	"Čermák",
	"Vaněk",
	"Urban",
	"Blažek",
	"Kříž",
	"Kovář",
	"Kratochvíl",
	"Bartoš",
	"Vlček",
	"Polák",
	"Musil",
	"Kopecký",
	"Šimek",
	"Konečný",
	"Malý",
	"Holub",
	"Čech",
	"Štěpánek",
	"Staněk",
	"Kadlec",
	"Dostál",
	"Soukup",
	"Šťastný",
	"Mareš",
	"Moravec",
	"Sýkora",
	"Tichý",
	"Valenta",
	"Vávra",
	"Matoušek",
	"Říha",
	"Bláha",
	"Bureš",
	"Ševčík",
	"Hruška",
	"Mašek",
	"Dušek",
	"Pavlík",
	"Havlíček",
	"Janda",
	"Hrubý",
	"Mach",
	"Liška",
	"Bednář",
	"Beran",
	"Vítek",
	"Macháček",
	"Müller",
	"Toman",
	"Vacek",
	"Bárta",
	"Tůma",
	"Žák",
	"Šmíd",
	"Kašpar",
	"Stejskal",
	"Sedlák",
	"Švec",
	"Horváth",
	"Jaroš",
	"Strnad",
	"Slavík",
	"Němeček",
	"Bílek",
	"Ježek",
	"Matějka",
	"Beránek",
	"Tesař",
	"Horáček",
	"Fišer",
	"Kraus",
	"Brož",
	"Kubíček",
	"Prokop",
	"Pavlíček",
	"Havel",
}

var maleNames = [...]string{
	"Jiří",
	"Jan",
	"Lukáš",
	"Petr",
	"Daniel",
	"Ondřej",
	"Adam",
	"Filip",
	"Pavel",
	"Luděk",
	"Ivan",
	"Alois",
	"Jonáš",
	"Vojtěch",
}
