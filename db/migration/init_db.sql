CREATE TABLE companies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL,
    representative_name VARCHAR(255),
    phone_number VARCHAR(20),
    postal_code VARCHAR(10),
    address VARCHAR(255)
);

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    company_id INT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    session_id VARCHAR(255) UNIQUE,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE clients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    company_id INT,
    client_name VARCHAR(255) NOT NULL,
    representative_name VARCHAR(255),
    phone_number VARCHAR(20),
    postal_code VARCHAR(10),
    address VARCHAR(255),
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE client_bank_accounts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    client_id INT,
    bank_name VARCHAR(255) NOT NULL,
    branch_name VARCHAR(255) NOT NULL,
    account_number VARCHAR(20) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE invoices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    company_id INT,
    client_id INT,
    issued_date DATE NOT NULL,
    amount_due DECIMAL(20, 2) NOT NULL,
    fee DECIMAL(20, 2),
    fee_rate DECIMAL(10, 2),
    consumption_tax DECIMAL(10, 2),
    tax_rate DECIMAL(10, 2),
    total_amount DECIMAL(10, 2) NOT NULL,
    due_date DATE,
    status ENUM('pending', 'processing', 'paid', 'error'),
    FOREIGN KEY (company_id) REFERENCES companies(id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

insert into companies values (1, "test_company", "shachou taro", null, null, null);

insert into clients values(1, 1, "test_client", "shachou jirou", null, null, null);

insert into users values (1, 1, "test user", "test@example.com", "test", "test");
