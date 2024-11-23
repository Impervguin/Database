package task9

type Task9Storage struct {
	pgs *PgStorage
	red *RedisStorage
}

const TopAmount = 10

func NewTask9Storage(pgs *PgStorage, red *RedisStorage) *Task9Storage {
	return &Task9Storage{pgs: pgs, red: red}
}

func (ts *Task9Storage) GetTopWealthyClients() ([]WealthyClient, error) {
	clients, err := ts.red.GetTopWealthyClients()
	if err != nil {
		return nil, err
	}
	if clients != nil {
		return clients, nil
	}

	clients, err = ts.pgs.GetWealthyClients(TopAmount)
	if err != nil {
		return nil, err
	}
	err = ts.red.PutTopWealthyClients(clients)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (ts *Task9Storage) GetTopWealthyClientsPostgres() ([]WealthyClient, error) {
	clients, err := ts.pgs.GetWealthyClients(TopAmount)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (ts *Task9Storage) CreateAccount(account *Account) error {
	err := ts.pgs.CreateAccount(account)
	if err != nil {
		return err
	}
	err = ts.red.DropTopWealthyClients()
	if err != nil {
		return err
	}
	return err
}

func (ts *Task9Storage) DeleteRandomAccount() error {
	err := ts.pgs.DeleteRandomAccount()
	if err != nil {
		return err
	}
	err = ts.red.DropTopWealthyClients()
	if err != nil {
		return err
	}
	return err
}

func (ts *Task9Storage) GetRandomActiveAccount() (*Account, error) {
	account, err := ts.pgs.GetRandomActiveAccount()
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (ts *Task9Storage) UpdateAccount(acc *Account) error {
	err := ts.pgs.UpdateAccount(acc)
	if err != nil {
		return err
	}
	err = ts.red.DropTopWealthyClients()
	if err != nil {
		return err
	}
	return nil
}

func (ts *Task9Storage) GetRandomClientId() (int64, error) {
	clientId, err := ts.pgs.GetRandomClientId()
	if err != nil {
		return 0, err
	}
	return clientId, nil
}
