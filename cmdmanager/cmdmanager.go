package cmdmanager

import "fmt"

type CMDManager struct {

}

func (cmd CMDManager) ReadLine() ([]string, error) {
	fmt.Println("Your prices (confirm with enter): ")

	var prices []string

	for{
		var price string
		fmt.Println("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}


func (cmd CMDManager) WriteResult(data interface{} /*any*/) error {
	fmt.Println(data)
	return nil
}

func New() *CMDManager{
	return &CMDManager{}
}