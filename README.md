# Warehouse Management System

A simple warehouse management system built with Go that handles products, customers, and transactions with role-based access control.

## Features

*   Role-based authentication (admin, employee, warehouse staff)
*   Product inventory management
*   Customer account management
*   Transaction processing
*   Sales reporting (top products and customers)

## User Roles

| Role | Permissions |
| :--- | :--- |
| Admin | View inventory, view customers, view reports |
| Employee | View inventory, view customers, process transactions |
| Warehouse Staff | View inventory, add new products |

## Project Structure
.
├── main.go # Main application file
├── staff.txt # User credentials and roles
├── warehouse.txt # Product inventory data
├── customer.txt # Customer information
└── transaction.txt # Transaction records


## Data File Formats

### staff.txt

username
password
username
password
...
text


### warehouse.txt

product_name
inventory_quantity
unit_price
...
text


### customer.txt

first_name
last_name
customer_id
account_balance
...
text


### transaction.txt

customer_id
product_name
quantity
completed_status
...
text


## Installation & Usage

1.  Ensure Go is installed on your system.
2.  Place all data files (`staff.txt`, `warehouse.txt`, `customer.txt`, `transaction.txt`) in the same directory as the source code or the executable.
3.  Run the application:
    ```bash
    go run main.go
    ```
4.  Login with credentials from `staff.txt`:
    *   Usernames starting with `1`: Admin
    *   Usernames starting with `2`: Employee
    *   Usernames starting with `3`: Warehouse staff

## Notes

*   Customer accounts can have a negative balance up to -200,000.
*   Transactions are only processed when sufficient inventory and customer credit are available.
*   All data is persisted to text files after modifications.

