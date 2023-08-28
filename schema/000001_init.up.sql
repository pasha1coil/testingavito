CREATE DOMAIN slug_name as VARCHAR(100) NOT NULL CONSTRAINT non_empty CHECK(length(value)>0);
CREATE TABLE Users
(
    User_number INTEGER PRIMARY KEY,
    CONSTRAINT User_number_unique UNIQUE (User_number)
);

CREATE TABLE slugs
(
    slug_name varchar(100) PRIMARY KEY,
    CONSTRAINT slug_name_unique UNIQUE (slug_name)

);

CREATE TABLE UsersSlug
(
    id   BIGSERIAL PRIMARY KEY,
    UserID INTEGER REFERENCES Users(User_number),
    name_slug  varchar(100) REFERENCES slugs(slug_name)
);

CREATE TABLE History
(
    id   BIGSERIAL PRIMARY KEY,
    UserID INTEGER,
    name_slug  varchar(100),
    mode varchar(50),
    created DATE DEFAULT CURRENT_DATE
);