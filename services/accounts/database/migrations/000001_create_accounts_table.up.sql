CREATE TABLE IF NOT EXISTS accounts(
   account_id UUID PRIMARY KEY,
   email VARCHAR (255) UNIQUE NOT NULL,
   name VARCHAR (255) NOT NULL,
   password_hash VARCHAR (255) NOT NULL
);
