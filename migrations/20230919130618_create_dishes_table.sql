-- +goose Up
-- +goose StatementBegin
CREATE TABLE dishes (
    id int unsigned primary key NOT NULL AUTO_INCREMENT,
    name varchar(150) NOT NULL,
    recipe varchar(2000) DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dishes;
-- +goose StatementEnd
