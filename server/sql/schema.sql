CREATE TABLE Accounts
(
    Account_ID BIGSERIAL PRIMARY KEY,
    Name  VARCHAR(50)        NOT NULL,
    Level smallint DEFAULT 0 NOT NULL,
    email VARCHAR(60) UNIQUE NOT NULL
);