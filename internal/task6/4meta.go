package task6

import "github.com/jedib0t/go-pretty/table"

// 4. Выполнить запрос к метаданным

// Получить имена и типы всех атрибутов таблицы

type GetTableAttributesResponse struct {
	Attributes []Attribute
}

func (r *GetTableAttributesResponse) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Name", "Type"})
	t.AppendFooter(table.Row{"Name", "Type"})
	for _, a := range r.Attributes {
		s := a.Row()
		row := make([]interface{}, 0, len(s))
		for _, v := range s {
			row = append(row, v)
		}
		t.AppendRow(row)
	}
	return t.Render()
}

type Attribute struct {
	Name string
	Type string
}

func (a *Attribute) Row() []string {
	return []string{a.Name, a.Type}
}

const getTableAttributes = `
SELECT column_name, data_type
FROM information_schema.columns
WHERE table_name = $1
`

func (t6s *Task6Storage) GetTableAttributes(tableName string) (*GetTableAttributesResponse, error) {
	rows, err := t6s.conn.Query(getTableAttributes, tableName)
	if err != nil {
		return nil, err
	}

	attrs := make([]Attribute, 0)
	for rows.Next() {
		var attr Attribute
		err := rows.Scan(&attr.Name, &attr.Type)
		if err != nil {
			return nil, err
		}
		attrs = append(attrs, attr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &GetTableAttributesResponse{Attributes: attrs}, nil
}
