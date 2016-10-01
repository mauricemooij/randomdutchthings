package randomdutchthings

import (
    "strings"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().Unix())
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var numberRunesWithoutZero = []rune("123456789")

var numberRunes = []rune("0123456789")

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func random(min, max int) int {
    return rand.Intn(max - min) + min
}

func randStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}


func randNumberRunes(n int, includeZeros bool) string {
    r := numberRunes
    if includeZeros == false {
        r = numberRunesWithoutZero
    }

    b := make([]rune, n)
    for i := range b {
        b[i] = r[rand.Intn(len(r))]
    }

    return string(b)
}

func RandomFirstName() string {   
    firstNameEntry := firstNameList[rand.Intn(len(firstNameList))]
    firstName := firstNameEntry.firstName
    
    return firstName
}

func RandomLastName() string {
    lastNameEntry := lastNameList[rand.Intn(len(lastNameList))]
    prefix := lastNameEntry.prefix
    lastName := lastNameEntry.lastName
  
    return strings.TrimSpace(prefix + " " + lastName)
}

func RandomFullName() string {
    firstName := RandomFirstName()
    lastName := RandomLastName()
    
    return firstName + " " + lastName
}

func RandomPostalCode() string {    
    firstNumber := randNumberRunes(1, false)
    otherNumbers := randNumberRunes(3, true)
    letters := randStringRunes(2)
    for letters == "sa" || letters == "sd" || letters == "ss" {
        letters = randStringRunes(2)
    }
    return firstNumber + otherNumbers + " " + letters
}

func RandomCity() string { 
    cityEntry := cityList[rand.Intn(len(cityList))]
    city := cityEntry.city
    
    return city
}

func RandomPostalCodeCity() string {
    postalCode := RandomPostalCode()
    city := RandomCity()
    return postalCode + " " + city
}

func RandomStreet() string {
    streetEntry := streetList[rand.Intn(len(streetList))]
    street := streetEntry.street
    firstNumber := randNumberRunes(1, false)
    secondNumber := ""
    if rand.Intn(2) == 1 {
        secondNumber = randNumberRunes(1, true) 
    } 

    thirdNumber := ""
    if rand.Intn(4) == 1 {
        thirdNumber = randNumberRunes(1, true) 
    }

    fourthNumber := ""
    if rand.Intn(500) == 1 {
        fourthNumber = randNumberRunes(1, true) 
    }    

    return strings.TrimSpace(street + " " + firstNumber + secondNumber + thirdNumber + fourthNumber)
}

