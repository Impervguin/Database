package menu

import (
	"fmt"
	"sort"
)

type Menu struct {
	Items map[int]MenuItem
}

type MenuItem interface {
	GetNum() int
	GetName() string
	DoBiz() (string, error)
}

func NewMenu() *Menu {
	return &Menu{
		Items: make(map[int]MenuItem, 10),
	}
}

func (m *Menu) AddItem(item MenuItem) {
	m.Items[item.GetNum()] = item
}

func (m *Menu) GetItem(num int) (MenuItem, bool) {
	item, ok := m.Items[num]
	return item, ok
}

func (m *Menu) RemoveItem(num int) {
	delete(m.Items, num)
}

func (m *Menu) Serve() {
	nums := make([]int, 0, len(m.Items))
	for num := range m.Items {
		nums = append(nums, num)
	}
	sort.Ints(nums)
	exitNum := nums[len(nums)-1] + 1

	for true {
		fmt.Println("\nChoose a menu item:")
		for _, num := range nums {
			item, _ := m.GetItem(num)
			fmt.Printf("%d. %s\n", num, item.GetName())
		}
		fmt.Printf("%d. %s\n", exitNum, "Exit")

		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		item, ok := m.GetItem(choice)
		if !ok {
			if choice == exitNum {
				fmt.Println("Thank you for using the menu. Goodbye!")
				return
			}
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		result, err := item.DoBiz()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Println(result)
	}
}
