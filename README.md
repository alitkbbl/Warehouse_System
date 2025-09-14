Warehouse Management System

A simple warehouse management system built with Go that handles products, customers, and transactions with role-based access control.
Features

    Role-based authentication (admin, employee, warehouse staff)

    Product inventory management

    Customer account management

    Transaction processing

    Sales reporting (top products and customers)

User Roles
Role	Permissions
Admin	View inventory, view customers, view reports
Employee	View inventory, view customers, process transactions
Warehouse Staff	View inventory, add new products
Project Structure
text

.
├── main.go                 # Main application file
├── staff.txt               # User credentials and roles
├── warehouse.txt           # Product inventory data
├── customer.txt            # Customer information
└── transaction.txt         # Transaction records

Data File Formats
staff.txt
text

username
password
username
password
...

warehouse.txt
text

product_name
inventory_quantity
unit_price
...

customer.txt
text

first_name
last_name
customer_id
account_balance
...

transaction.txt
text

customer_id
product_name
quantity
completed_status
...

Installation & Usage

    Ensure Go is installed on your system

    Place all data files in the same directory as the executable

    Run the application:

bash

go run main.go

    Login with credentials from staff.txt

        Usernames starting with 1: Admin

        Usernames starting with 2: Employee

        Usernames starting with 3: Warehouse staff

Notes

    Customer accounts can have a negative balance up to -200,000

    Transactions are only processed when sufficient inventory and customer credit are available

    All data is persisted to text files after modifications

