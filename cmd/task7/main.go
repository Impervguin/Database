package main

import (
	"DatabaseCourse/internal/task7"
)

func main() {
	t7s, err := task7.NewTask7Storage()
	if err != nil {
		panic(err)
	}

	menu := task7.NewTask7Menu(t7s)
	menu.Serve()

	// res, err := t7s.GetCardById(3000)
	// if err != nil {
	// 	panic(err)
	// }

	// newCard := &domain.Card{
	// 	CardNumber: "4567890323456789",
	// 	AccountId:  125,
	// 	CVV:        "123",
	// }

	// res, err := t7s.GetClientTotalBalance(5)
	// if err != nil {
	// 	panic(err)
	// }

	// f, err := os.Create("t.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// err = task7.DumpClientsToJson(res, f)
	// if err != nil {
	// 	panic(err)
	// }

	// f, err = os.Create("tu.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// rf, err := os.Open("t.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rf.Close()

	// err = task7.SetAllNamesUpperJson(rf, f)

	// rf, err = os.Open("tu.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rf.Close()

	// clients, err := task7.ReadJsonClients(rf)
	// if err != nil {
	// 	panic(err)
	// }

	// printRes := make([]domain.DomainType, len(res))
	// for i, client := range res {
	// 	printRes[i] = client
	// }

	// fmt.Println(domain.PrintDomainRow(res))
}
