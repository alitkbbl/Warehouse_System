# 📦 Warehouse Management System

[![Go Version](https://img.shields.io/badge/Go-1.18%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/alitkbbl/Warehouse_System)

A robust Command Line Interface (CLI) Warehouse Management System built with **Go**. This application manages inventory, customer accounts, and sales transactions efficiently using a custom file-based database system. It features role-based access control (RBAC) to ensure secure operations.

## 📑 Table of Contents
- [Features](#-features)
- [System Architecture](#-system-architecture)
- [User Roles & Permissions](#-user-roles--permissions)
- [Getting Started](#-getting-started)
- [Data Storage Format](#-data-storage-format)
- [Business Logic](#-business-logic)
- [Project Structure](#-project-structure)

## 🚀 Features

- **🔐 Role-Based Authentication:** Secure login system differentiating between Admins, Employees, and Warehouse Staff.
- **📦 Inventory Management:** Real-time tracking of product names, quantities, and unit prices.
- **busts_in_silhouette: Customer Management:** Manage customer IDs, names, and credit balances.
- **💳 Transaction Processing:** Automated handling of sales, updating both inventory and customer balances instantly.
- **📊 Reporting:** Generate insights on top-selling products and high-value customers.
- **💾 Persistence:** All data is automatically saved to local text files, ensuring no data loss upon exit.

## 🏗 System Architecture

The system uses a flat-file database approach. It loads data from `.txt` files into memory slices upon startup (structs), performs operations in memory for speed, and writes changes back to the files to ensure persistence.

## 👥 User Roles & Permissions

Access is determined by the first digit of the username provided in `staff.txt`.

| Role | Username Prefix | Capabilities |
| :--- | :---: | :--- |
| **Admin** | `1` | Full access: View inventory, manage customers, view sales reports. |
| **Employee** | `2` | Sales focus: View inventory/customers, **process transactions**. |
| **Warehouse Staff** | `3` | Inventory focus: View inventory, **add new stock**. |

## 🛠 Getting Started

### Prerequisites
- **Go** (Golang) installed on your machine (version 1.18 or higher recommended).

### Installation

1.  **Clone the repository:**
```bash
git clone https://github.com/alitkbbl/Warehouse_System.git
cd Warehouse_System
```
2. **Verify Data Files**

   Ensure the following files exist in the root directory (or create them if missing):

- `staff.txt`
- `warehouse.txt`
- `customer.txt`
- `transaction.txt`

## 3. Run the Application
```bash
go run main.go
```

---

## 💾 Data Storage Format

The application uses a custom flat-file database system. Data is parsed line-by-line from specific text files located in the root directory.

### `staff.txt` (Credentials)
Stores user credentials in pairs (2 lines per user). The first digit of the username determines the role.
```text
username
password
username
password
```
### `warehouse.txt` (Inventory)
Stores product details in blocks of 3 lines.
text
product_name
quantity (integer)
unit_price (float)

### `customer.txt` (Clients)
Stores customer information in blocks of 4 lines.
text
first_name
last_name
customer_id (unique integer)
account_balance (float)

### `transaction.txt` (History)
Stores transaction records in blocks of 4 lines.
text
customer_id
product_name
quantity
status (e.g., completed)

---

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

