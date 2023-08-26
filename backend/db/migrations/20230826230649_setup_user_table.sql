-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    display_name varchar(255) NOT NULL,
    about_me varchar(255),
    avatar varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;

-- +goose StatementEnd
