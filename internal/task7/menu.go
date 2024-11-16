package task7

import (
	"DatabaseCourse/internal/menu"
	"DatabaseCourse/internal/task7/domain"
	"fmt"
	"os"
)

type Task7Menu struct {
	storage *Task7Storage
	*menu.Menu
}

type Task7MenuItem struct {
	num  int
	name string
	biz  func() (string, error)
}

func (mi Task7MenuItem) GetNum() int     { return mi.num }
func (mi Task7MenuItem) GetName() string { return mi.name }

func (mi Task7MenuItem) DoBiz() (string, error) {
	return mi.biz()
}

func NewTask7Menu(storage *Task7Storage) *Task7Menu {
	t6m := &Task7Menu{
		Menu:    menu.NewMenu(),
		storage: storage,
	}
	// t6m.scalar1()
	t6m.AddItem(Task7MenuItem{num: 1, name: "Get clients in alphabetical order", biz: t6m.object1})
	t6m.AddItem(Task7MenuItem{num: 2, name: "Get clients with loans exceeding 1 million", biz: t6m.object2})
	t6m.AddItem(Task7MenuItem{num: 3, name: "Get clients with active loans", biz: t6m.object3})
	t6m.AddItem(Task7MenuItem{num: 4, name: "Get clients and their total balances", biz: t6m.object4})
	t6m.AddItem(Task7MenuItem{num: 5, name: "Get clients with maximum number of accounts", biz: t6m.object5})
	t6m.AddItem(Task7MenuItem{num: 6, name: "Dump clients into json file", biz: t6m.jsonDump})
	t6m.AddItem(Task7MenuItem{num: 7, name: "Read clients json file", biz: t6m.jsonRead})
	t6m.AddItem(Task7MenuItem{num: 8, name: "Transform client names to upper", biz: t6m.jsonUpper})
	t6m.AddItem(Task7MenuItem{num: 9, name: "Add client to json file", biz: t6m.jsonAppend})
	t6m.AddItem(Task7MenuItem{num: 10, name: "Get card by id", biz: t6m.GetCard})
	t6m.AddItem(Task7MenuItem{num: 11, name: "Get clients card", biz: t6m.GetClientCard})
	t6m.AddItem(Task7MenuItem{num: 12, name: "Create card", biz: t6m.AddCard})
	t6m.AddItem(Task7MenuItem{num: 13, name: "Block card by id", biz: t6m.BlockCard})
	t6m.AddItem(Task7MenuItem{num: 14, name: "Delete card by id", biz: t6m.DeleteCard})
	t6m.AddItem(Task7MenuItem{num: 15, name: "Get client total balance by id", biz: t6m.GetClientTotalBalance})

	return t6m
}

func (m *Task7Menu) object1() (string, error) {
	resp, err := m.storage.GetClientsAlphabetically()
	if err != nil {
		return "", err
	}

	printResp := make([]domain.DomainType, len(resp))
	for i, c := range resp {
		printResp[i] = c
	}
	return domain.PrintDomainTable(printResp), nil
}

func (m *Task7Menu) object2() (string, error) {
	resp, err := m.storage.GetUnpaidLoansExceeding1Million()
	if err != nil {
		return "", err
	}

	printResp := make([]domain.DomainType, len(resp))
	for i, c := range resp {
		printResp[i] = c
	}
	return domain.PrintDomainTable(printResp), nil
}

func (m *Task7Menu) object3() (string, error) {
	resp, err := m.storage.GetClientsWithLoan()
	if err != nil {
		return "", err
	}

	printResp := make([]domain.DomainType, len(resp))
	for i, c := range resp {
		printResp[i] = c
	}
	return domain.PrintDomainTable(printResp), nil
}

func (m *Task7Menu) object4() (string, error) {
	resp, err := m.storage.GetClientsAccountSums()
	if err != nil {
		return "", err
	}

	printResp := make([]domain.DomainType, len(resp))
	for i, c := range resp {
		printResp[i] = c
	}
	return domain.PrintDomainTable(printResp), nil
}

func (m *Task7Menu) object5() (string, error) {
	resp, err := m.storage.GetClientsWithMostAccounts()
	if err != nil {
		return "", err
	}

	printResp := make([]domain.DomainType, len(resp))
	for i, c := range resp {
		printResp[i] = c
	}
	return domain.PrintDomainTable(printResp), nil
}

func (m *Task7Menu) jsonDump() (string, error) {
	clients, err := m.storage.GetShortClients()
	if err != nil {
		return "", err
	}
	var fname string
	fmt.Print("Enter output file name: ")
	fmt.Scanf("%s", &fname)
	f, err := os.Create(fname)
	if err != nil {
		return "", err
	}
	defer f.Close()
	err = DumpClientsToJson(clients, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Clients have been dumped to JSON file: %s", fname), nil
}

func (m *Task7Menu) jsonRead() (string, error) {
	var fname string
	fmt.Print("Enter input file name: ")
	fmt.Scanf("%s", &fname)
	f, err := os.Open(fname)
	if err != nil {
		return "", err
	}
	defer f.Close()

	clients, err := ReadJsonClients(f)
	if err != nil {
		return "", err
	}
	printResp := make([]domain.DomainType, len(clients))
	for i, c := range clients {
		printResp[i] = c
	}
	return domain.PrintDomainTable(printResp), nil
}

func (m *Task7Menu) jsonUpper() (string, error) {
	var fname string
	fmt.Print("Enter input file name: ")
	fmt.Scanf("%s", &fname)
	f, err := os.Open(fname)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var foutname string
	fmt.Print("Enter output file name: ")
	fmt.Scanf("%s", &foutname)
	fout, err := os.Create(foutname)
	if err != nil {
		return "", err
	}
	defer fout.Close()

	err = SetAllNamesUpperJson(f, fout)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Clients have been converted to uppercase in JSON file: %s", foutname), nil
}

func (m *Task7Menu) jsonAppend() (string, error) {
	var fname string
	fmt.Print("Enter input file name: ")
	fmt.Scanf("%s", &fname)
	f, err := os.Open(fname)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var foutname string
	fmt.Print("Enter output file name: ")
	fmt.Scanf("%s", &foutname)
	fout, err := os.Create(foutname)
	if err != nil {
		return "", err
	}
	defer fout.Close()

	var client domain.ShortClient
	fmt.Print("Enter first name: ")
	_, err = fmt.Scanf("%s", &client.FirstName)
	if err != nil {
		return "", err
	}

	fmt.Print("Enter last name: ")
	_, err = fmt.Scanf("%s", &client.LastName)
	if err != nil {
		return "", err
	}

	fmt.Print("Enter email: ")
	_, err = fmt.Scanf("%s", &client.Email)
	if err != nil {
		return "", err
	}

	fmt.Print("Enter phone number: ")
	_, err = fmt.Scanf("%s", &client.PhoneNumber)
	if err != nil {
		return "", err
	}
	err = AddShortClientJson(f, fout, &client)
	if err != nil {
		return "", err
	}

	fin, err := os.Open(foutname)
	if err != nil {
		return "", err
	}
	defer fin.Close()

	clients, err := ReadJsonClients(fin)
	if err != nil {
		return "", err
	}
	printResp := make([]domain.DomainType, len(clients))
	for i, c := range clients {
		printResp[i] = c
	}
	return domain.PrintDomainTable(printResp), nil
}

func (m *Task7Menu) GetCard() (string, error) {
	var cardId int64
	fmt.Print("Enter card ID: ")
	_, err := fmt.Scanf("%d", &cardId)
	if err != nil {
		return "", err
	}

	card, err := m.storage.GetCardById(cardId)
	if err != nil {
		return "", err
	}

	return domain.PrintDomainRow(card), nil
}

func (m *Task7Menu) GetClientCard() (string, error) {
	var clientId int64
	fmt.Print("Enter client ID: ")
	_, err := fmt.Scanf("%d", &clientId)
	if err != nil {
		return "", err
	}

	cards, err := m.storage.GetCardsForClient(clientId)
	if err != nil {
		return "", err
	}

	printResp := make([]domain.DomainType, len(cards))
	for i, c := range cards {
		printResp[i] = c
	}
	return domain.PrintDomainTable(printResp), nil
}

func (m *Task7Menu) AddCard() (string, error) {
	var card domain.Card
	fmt.Print("Enter card number: ")
	_, err := fmt.Scanf("%s", &card.CardNumber)
	if err != nil {
		return "", err
	}

	fmt.Print("Enter cvv: ")
	_, err = fmt.Scanf("%s", &card.CVV)
	if err != nil {
		return "", err
	}

	fmt.Print("Enter account id: ")
	_, err = fmt.Scanf("%d", &card.AccountId)
	if err != nil {
		return "", err
	}
	err = m.storage.CreateCard(&card)
	if err != nil {
		return "", err
	}
	return domain.PrintDomainRow(card), nil
}

func (m *Task7Menu) BlockCard() (string, error) {
	var cardId int64
	fmt.Print("Enter card ID: ")
	_, err := fmt.Scanf("%d", &cardId)
	if err != nil {
		return "", err
	}

	card, err := m.storage.BlockCard(cardId)
	if err != nil {
		return "", err
	}

	return domain.PrintDomainRow(card), nil
}

func (m *Task7Menu) DeleteCard() (string, error) {
	var cardId int64
	fmt.Print("Enter card ID: ")
	_, err := fmt.Scanf("%d", &cardId)
	if err != nil {
		return "", err
	}

	card, err := m.storage.DeleteCard(cardId)
	if err != nil {
		return "", err
	}
	return domain.PrintDomainRow(card), nil
}

func (m *Task7Menu) GetClientTotalBalance() (string, error) {
	var clientId int64
	fmt.Print("Enter client ID: ")
	_, err := fmt.Scanf("%d", &clientId)
	if err != nil {
		return "", err
	}

	balance, err := m.storage.GetClientTotalBalance(clientId)
	if err != nil {
		return "", err
	}

	return domain.PrintDomainRow(balance), nil
}
