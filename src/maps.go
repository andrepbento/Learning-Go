package main

import (
    "fmt"
)

func main() {
    //      declare a variable, by default map will be nil
    // var map_variable map[key_data_type]value_data_type
    //      define the map as nil map can not be assigned any value
    // map_variable = make(map[key_data_type]value_data_type)

    var countryCapitalMap map[string]string
    //  create a map
    countryCapitalMap = make(map[string]string)

    //  insert key-value pairs in the map
    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Rome"
    countryCapitalMap["Japan"] = "Tokyo"
    countryCapitalMap["India"] = "New Delhi"

    // print map using keys
    for country := range countryCapitalMap {
        fmt.Println("Capital of",country,"is",countryCapitalMap[country])
    }

    // test if entry is present in the map or not
    capital, ok := countryCapitalMap["United States"]

    if(ok) {
        fmt.Println("Capital of United States is", capital)
    } else {
        fmt.Println("Capital of United States is not present")
    }

    //      delete() function
    fmt.Println("\nOriginal map")

    // print map
    for country := range countryCapitalMap {
        fmt.Println("Capital of",country,"is",countryCapitalMap[country])
    }

    // delete an entry
    delete(countryCapitalMap, "France")
    fmt.Println("\nEntry for France is deleted")

    fmt.Println("\nUpdated map")

    // print map
    for country := range countryCapitalMap {
        fmt.Println("Capital of",country,"is",countryCapitalMap[country])
    }
}