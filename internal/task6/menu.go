package task6

import (
	"DatabaseCourse/internal/menu"
	"fmt"
)

type Task6Menu struct {
	storage *Task6Storage
	*menu.Menu
}

type Task6MenuItem struct {
	num  int
	name string
	biz  func() (string, error)
}

func (mi Task6MenuItem) GetNum() int     { return mi.num }
func (mi Task6MenuItem) GetName() string { return mi.name }

func (mi Task6MenuItem) DoBiz() (string, error) {
	return mi.biz()
}

func NewTask6Menu(storage *Task6Storage) *Task6Menu {
	t6m := &Task6Menu{
		Menu:    menu.NewMenu(),
		storage: storage,
	}
	// t6m.scalar1()
	t6m.AddItem(Task6MenuItem{num: 1, name: "Get number of clients", biz: t6m.scalar1})
	t6m.AddItem(Task6MenuItem{num: 2, name: "Get all cards with their owners", biz: t6m.joins2})
	t6m.AddItem(Task6MenuItem{num: 3, name: "Get account statistics by type", biz: t6m.otv3})
	t6m.AddItem(Task6MenuItem{num: 4, name: "Get attributes names and types of specified table", biz: t6m.meta4})
	t6m.AddItem(Task6MenuItem{num: 5, name: "Get clients balance by id", biz: t6m.scalarfunc5})
	t6m.AddItem(Task6MenuItem{num: 6, name: "Get clients name and numbers with blocked cards", biz: t6m.tablefunc6})
	t6m.AddItem(Task6MenuItem{num: 7, name: "Default random person's loan", biz: t6m.proc7})
	t6m.AddItem(Task6MenuItem{num: 8, name: "Get database user name", biz: t6m.sysproc8})
	t6m.AddItem(Task6MenuItem{num: 9, name: "Create stocks table", biz: t6m.create9})
	t6m.AddItem(Task6MenuItem{num: 10, name: "Insert stocks in table", biz: t6m.ins10})

	return t6m
}

func (m *Task6Menu) scalar1() (string, error) {
	resp, err := m.storage.GetClientsCount()
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (m *Task6Menu) joins2() (string, error) {
	resp, err := m.storage.GetClientsAndCards()
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (m *Task6Menu) otv3() (string, error) {
	resp, err := m.storage.GetAccountTypesAndStats()
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (m *Task6Menu) meta4() (string, error) {
	var tname string
	fmt.Print("Input table name: ")
	fmt.Scanf("%s", &tname)
	resp, err := m.storage.GetTableAttributes(tname)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (m *Task6Menu) scalarfunc5() (string, error) {
	var cid int64
	fmt.Print("Input client ID: ")
	fmt.Scanf("%d", &cid)
	resp, err := m.storage.GetSumBalance(cid)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (m *Task6Menu) tablefunc6() (string, error) {
	resp, err := m.storage.GetBlockedClients()
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (m *Task6Menu) proc7() (string, error) {
	resp, err := m.storage.CallPromoRandom()
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (m *Task6Menu) sysproc8() (string, error) {
	resp, err := m.storage.GetCurrentUser()
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (m *Task6Menu) create9() (string, error) {
	err := m.storage.CreateStockTable()
	if err != nil {
		return "", err
	}
	return "Stock table created successfully", nil
}

func (m *Task6Menu) ins10() (string, error) {
	var company string
	var price float64
	var quantity int
	var dividend float64
	fmt.Print("Input company name: ")
	_, err := fmt.Scanf("%s", &company)
	if err != nil {
		return "", err
	}
	fmt.Println()
	fmt.Print("Input price: ")
	_, err = fmt.Scanf("%f", &price)
	if err != nil {
		return "", err
	}

	fmt.Println()
	fmt.Print("Input quantity: ")
	_, err = fmt.Scanf("%d", &quantity)
	if err != nil {
		return "", err
	}

	fmt.Println()
	fmt.Print("Input dividend: ")
	_, err = fmt.Scanf("%f", &dividend)
	if err != nil {
		return "", err
	}
	stock := Stocks{
		Company:  company,
		Price:    price,
		Quantity: quantity,
		Dividend: dividend,
	}
	err = m.storage.InsertStock(stock)
	if err != nil {
		return "", err
	}
	return "Stock inserted successfully", nil
}
