CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE loans (
    loan_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id),
    nik VARCHAR(16) NOT NULL, 
    loan_limit DECIMAL(10, 2) CHECK(loan_limit >= 100000)
);

