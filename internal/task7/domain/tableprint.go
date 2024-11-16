package domain

import "github.com/jedib0t/go-pretty/table"

type DomainType interface {
	Row() []string
	Head() []string
}

func PrintDomainRow(d DomainType) string {
	t := table.NewWriter()
	head := make([]interface{}, 0, len(d.Head()))
	for _, h := range d.Head() {
		head = append(head, h)
	}
	t.AppendHeader(head)
	s := d.Row()
	row := make([]interface{}, 0, len(s))
	for _, v := range s {
		row = append(row, v)
	}
	t.AppendRow(row)
	return t.Render()
}

func PrintDomainTable(d []DomainType) string {
	t := table.NewWriter()
	head := make([]interface{}, 0, len(d[0].Head()))
	for _, h := range d[0].Head() {
		head = append(head, h)
	}
	t.AppendHeader(head)
	t.AppendFooter(head)
	for _, domain := range d {
		s := domain.Row()
		row := make([]interface{}, 0, len(s))
		for _, v := range s {
			row = append(row, v)
		}
		t.AppendRow(row)
	}
	return t.Render()
}
