CREATE TABLE IF NOT EXISTS roles(
    id VARCHAR(36),
    name VARCHAR(50),
    CONSTRAINT pk_roles_id PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(36),
    role_id VARCHAR(36),
    email VARCHAR(50),
    password VARCHAR(150),
    name VARCHAR(100),
    birth_date DATE,
    cpf VARCHAR(11),
    phone VARCHAR(30),
    status BOOLEAN,
    created_at DATE,
    CONSTRAINT pk_users_id PRIMARY KEY (id),
    CONSTRAINT fk_users_role_id FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE IF NOT EXISTS addresses(
    id VARCHAR(36),
    user_id VARCHAR(36),
    zip_code VARCHAR(30),
    state VARCHAR(50),
    city VARCHAR(50),
    neighborhood VARCHAR(50),
    street VARCHAR(50),
    number VARCHAR(10),
    complement VARCHAR(50),
    CONSTRAINT pk_addresses_id PRIMARY KEY(id),
    CONSTRAINT fk_addresses_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);
