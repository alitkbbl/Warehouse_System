package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Staff struct {
	Username string
	Password string
	Role     string // User role (admin, employee, warehouse)
}

type Product struct {
	Name      string
	Inventory int
	Price     int
}

type Customer struct {
	FirstName string
	LastName  string
	ID        int
	Balance   int
}

type Transaction struct {
	CustomerID  int
	ProductName string
	Quantity    int
	Completed   bool
}

// Main system structure
type WarehouseSystem struct {
	Staff         []Staff
	Products      []Product
	Customers     []Customer
	Transactions  []Transaction
	ProductStats  map[string]int // Product sales statistics
	CustomerStats map[int]int    // Customer purchase statistics
}

// File reading functions
func (ws *WarehouseSystem) LoadStaff(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username := strings.TrimSpace(scanner.Text())
		if !scanner.Scan() {
			break
		}
		password := strings.TrimSpace(scanner.Text())

		// Determine role based on username prefix
		var role string
		if strings.HasPrefix(username, "1") {
			role = "admin"
		} else if strings.HasPrefix(username, "2") {
			role = "employee"
		} else if strings.HasPrefix(username, "3") {
			role = "warehouse"
		} else {
			role = "unknown"
		}

		ws.Staff = append(ws.Staff, Staff{
			Username: username,
			Password: password,
			Role:     role,
		})
	}
	return scanner.Err()
}

func (ws *WarehouseSystem) LoadProducts(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())
		if !scanner.Scan() {
			break
		}
		inventory, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if !scanner.Scan() {
			break
		}
		price, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		ws.Products = append(ws.Products, Product{
			Name:      name,
			Inventory: inventory,
			Price:     price,
		})
	}
	return scanner.Err()
}

func (ws *WarehouseSystem) LoadCustomers(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		firstName := strings.TrimSpace(scanner.Text())
		if !scanner.Scan() {
			break
		}
		lastName := strings.TrimSpace(scanner.Text())
		if !scanner.Scan() {
			break
		}
		id, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if !scanner.Scan() {
			break
		}
		balance, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		ws.Customers = append(ws.Customers, Customer{
			FirstName: firstName,
			LastName:  lastName,
			ID:        id,
			Balance:   balance,
		})
	}
	return scanner.Err()
}

func (ws *WarehouseSystem) LoadTransactions(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		customerID, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if !scanner.Scan() {
			break
		}
		productName := strings.TrimSpace(scanner.Text())
		if !scanner.Scan() {
			break
		}
		quantity, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if !scanner.Scan() {
			break
		}
		completed, _ := strconv.ParseBool(strings.TrimSpace(scanner.Text()))

		ws.Transactions = append(ws.Transactions, Transaction{
			CustomerID:  customerID,
			ProductName: productName,
			Quantity:    quantity,
			Completed:   completed,
		})
	}
	return scanner.Err()
}

// File saving functions
func (ws *WarehouseSystem) SaveProducts(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, product := range ws.Products {
		file.WriteString(product.Name + "\n")
		file.WriteString(strconv.Itoa(product.Inventory) + "\n")
		file.WriteString(strconv.Itoa(product.Price) + "\n")
	}
	return nil
}

func (ws *WarehouseSystem) SaveCustomers(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, customer := range ws.Customers {
		file.WriteString(customer.FirstName + "\n")
		file.WriteString(customer.LastName + "\n")
		file.WriteString(strconv.Itoa(customer.ID) + "\n")
		file.WriteString(strconv.Itoa(customer.Balance) + "\n")
	}
	return nil
}

func (ws *WarehouseSystem) SaveTransactions(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, transaction := range ws.Transactions {
		file.WriteString(strconv.Itoa(transaction.CustomerID) + "\n")
		file.WriteString(transaction.ProductName + "\n")
		file.WriteString(strconv.Itoa(transaction.Quantity) + "\n")
		file.WriteString(strconv.FormatBool(transaction.Completed) + "\n")
	}
	return nil
}

// System functions
func (ws *WarehouseSystem) Login(username, password string) *Staff {
	for _, staff := range ws.Staff {
		if staff.Username == username && staff.Password == password {
			return &staff
		}
	}
	return nil
}

func (ws *WarehouseSystem) ShowProducts() {
	fmt.Println("\n=== Product List ===")
	for i, product := range ws.Products {
		if product.Name != "" {
			fmt.Printf("%d. Name: %s, Inventory: %d, Price: %d\n",
				i+1, product.Name, product.Inventory, product.Price)
		}
	}
	fmt.Println("====================")
}

func (ws *WarehouseSystem) ShowCustomers() {
	fmt.Println("\n=== Customer List ===")
	for i, customer := range ws.Customers {
		if customer.FirstName != "" {
			fmt.Printf("%d. Name: %s %s, ID: %d, Balance: %d\n",
				i+1, customer.FirstName, customer.LastName, customer.ID, customer.Balance)
		}
	}
	fmt.Println("====================")
}

func (ws *WarehouseSystem) AddProduct(name string, inventory, price int) {
	ws.Products = append(ws.Products, Product{
		Name:      name,
		Inventory: inventory,
		Price:     price,
	})
	ws.SaveProducts("warehouse.txt")
	fmt.Println("Product added successfully.")
}

func (ws *WarehouseSystem) ProcessTransactions() {
	fmt.Println("\n=== Processing Transactions ===")
	for i, transaction := range ws.Transactions {
		if !transaction.Completed {
			// Find customer
			var customer *Customer
			for j := range ws.Customers {
				if ws.Customers[j].ID == transaction.CustomerID {
					customer = &ws.Customers[j]
					break
				}
			}

			// Find product
			var product *Product
			for j := range ws.Products {
				if ws.Products[j].Name == transaction.ProductName {
					product = &ws.Products[j]
					break
				}
			}

			if customer == nil {
				fmt.Printf("Transaction %d: Customer with ID %d not found\n", i+1, transaction.CustomerID)
				continue
			}

			if product == nil {
				fmt.Printf("Transaction %d: Product with name %s not found\n", i+1, transaction.ProductName)
				continue
			}

			totalCost := transaction.Quantity * product.Price

			if product.Inventory < transaction.Quantity {
				fmt.Printf("Transaction %d: Insufficient inventory for product %s\n", i+1, transaction.ProductName)
				continue
			}

			if customer.Balance-totalCost < -200000 {
				fmt.Printf("Transaction %d: Insufficient customer balance\n", i+1)
				continue
			}

			// Process transaction
			product.Inventory -= transaction.Quantity
			customer.Balance -= totalCost
			ws.Transactions[i].Completed = true

			// Update statistics
			ws.ProductStats[product.Name] += totalCost
			ws.CustomerStats[customer.ID] += totalCost

			fmt.Printf("Transaction %d: Completed successfully\n", i+1)
		}
	}

	// Save changes
	ws.SaveProducts("warehouse.txt")
	ws.SaveCustomers("customer.txt")
	ws.SaveTransactions("transaction.txt")
	fmt.Println("Transaction processing completed.")
}

func (ws *WarehouseSystem) ShowTopProduct() {
	fmt.Println("\n=== Best Selling Product ===")
	var topProduct string
	maxSales := 0

	for product, sales := range ws.ProductStats {
		if sales > maxSales {
			maxSales = sales
			topProduct = product
		}
	}

	if topProduct != "" {
		fmt.Printf("Best selling product: %s with %d in sales\n", topProduct, maxSales)
	} else {
		fmt.Println("No transactions recorded yet.")
	}
	fmt.Println("========================")
}

func (ws *WarehouseSystem) ShowTopCustomer() {
	fmt.Println("\n=== Best Customer ===")
	var topCustomerID int
	maxPurchases := 0

	for customerID, purchases := range ws.CustomerStats {
		if purchases > maxPurchases {
			maxPurchases = purchases
			topCustomerID = customerID
		}
	}

	if topCustomerID != 0 {
		// Find customer information
		for _, customer := range ws.Customers {
			if customer.ID == topCustomerID {
				fmt.Printf("Best customer: %s %s with %d in purchases\n",
					customer.FirstName, customer.LastName, maxPurchases)
				break
			}
		}
	} else {
		fmt.Println("No transactions recorded yet.")
	}
	fmt.Println("=====================")
}

// Menu functions
func (ws *WarehouseSystem) AdminMenu() {
	for {
		fmt.Println("\n=== Admin Menu ===")
		fmt.Println("1. View inventory")
		fmt.Println("2. View customers")
		fmt.Println("3. View best selling product")
		fmt.Println("4. View best customer")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Please select an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ws.ShowProducts()
		case 2:
			ws.ShowCustomers()
		case 3:
			ws.ShowTopProduct()
		case 4:
			ws.ShowTopCustomer()
		case 5:
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func (ws *WarehouseSystem) EmployeeMenu() {
	for {
		fmt.Println("\n=== Employee Menu ===")
		fmt.Println("1. View inventory")
		fmt.Println("2. View customers")
		fmt.Println("3. Process transactions")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Please select an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ws.ShowProducts()
		case 2:
			ws.ShowCustomers()
		case 3:
			ws.ProcessTransactions()
		case 4:
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func (ws *WarehouseSystem) WarehouseMenu() {
	for {
		fmt.Println("\n=== Warehouse Menu ===")
		fmt.Println("1. View inventory")
		fmt.Println("2. Add new product")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Please select an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ws.ShowProducts()
		case 2:
			var name string
			var inventory, price int

			fmt.Print("Product name: ")
			fmt.Scan(&name)
			fmt.Print("Inventory quantity: ")
			fmt.Scan(&inventory)
			fmt.Print("Unit price: ")
			fmt.Scan(&price)

			ws.AddProduct(name, inventory, price)
		case 3:
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func main() {
	fmt.Println("=== Warehouse Management System ===")

	// Initialize system and load data
	system := WarehouseSystem{
		ProductStats:  make(map[string]int),
		CustomerStats: make(map[int]int),
	}

	// Load data
	if err := system.LoadStaff("staff.txt"); err != nil {
		fmt.Printf("Error loading staff: %v\n", err)
		return
	}

	if err := system.LoadProducts("warehouse.txt"); err != nil {
		fmt.Printf("Error loading products: %v\n", err)
		return
	}

	if err := system.LoadCustomers("customer.txt"); err != nil {
		fmt.Printf("Error loading customers: %v\n", err)
		return
	}

	if err := system.LoadTransactions("transaction.txt"); err != nil {
		fmt.Printf("Error loading transactions: %v\n", err)
		return
	}

	// Login screen
	for {
		var username, password string
		fmt.Print("\nUsername: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		staff := system.Login(username, password)
		if staff != nil {
			fmt.Printf("Welcome, %s!\n", staff.Username)

			switch staff.Role {
			case "admin":
				system.AdminMenu()
			case "employee":
				system.EmployeeMenu()
			case "warehouse":
				system.WarehouseMenu()
			default:
				fmt.Println("Unknown user role.")
			}
		} else {
			fmt.Println("Invalid username or password. Please try again.")
		}
	}
}
