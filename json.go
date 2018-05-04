package main

import (
    "encoding/json"
    "fmt"
)

type Shop struct {
    OneItem []Product
    SecondItem []Fruit
}

type Product struct {
    Name      string
    ProductID int64
	Number    int
	Price     float64
    IsOnSale  bool
}

type Fruit struct {
    Name      string
    FruitID int64
	Number    int
	Price     float64
    IsOnSale  bool
}


func makeJson() {
    var p Shop
    p.OneItem = append(p.OneItem, Product{Name: "小米7", ProductID: 1, Number: 1000, IsOnSale: true})
    p.OneItem = append(p.OneItem, Product{Name: "Apple x", ProductID: 2, Number: 100, IsOnSale: true})
    p.SecondItem = append(p.SecondItem, Fruit{Name: "watermelon", FruitID: 3, Number: 20, IsOnSale: false})
	jsonData, _ := json.Marshal(p)
    fmt.Printf("json: %s\n", jsonData)
}

func decodeJson() error {
    var p Shop
    data := []byte(`{"OneItem":[{"Name":"小米7","ProductID":1,"Number":1000,"Price":0,"IsOnSale":true},{"Name":"Apple x","ProductID":2,"Number":100,"Price":0,"IsOnSale":true}],"SecondItem":[{"Name":"watermelon","FruitID":3,"Number":20,"Price":0,"IsOnSale":false}]}`)
    err := json.Unmarshal(data, &p)
    if (err != nil ) {
        return err
    }
    fmt.Printf("%#v\n", p)
    return nil
}

func main() {
    makeJson()
    err := decodeJson()
    if (err != nil){
        fmt.Println(err)
    }
}
