-- +goose Up
CREATE TABLE ingredient_units (
    id INT UNSIGNED primary key NOT NULL AUTO_INCREMENT,
    unit VARCHAR(50) NOT NULL
);

CREATE TABLE ingredients (
    id INT UNSIGNED primary key NOT NULL AUTO_INCREMENT,
    name VARCHAR(150) NOT NULL,
    unit INT UNSIGNED NOT NULL,
    FOREIGN KEY (unit) REFERENCES ingredient_units (id)
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
);
-- +goose Down
DROP TABLE ingredients;
DROP TABLE ingredient_units;