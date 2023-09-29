-- +goose Up
-- +goose StatementBegin
CREATE TABLE dish_ingredients (
    dish_id INT UNSIGNED NOT NULL ,
    ingredient_id INT UNSIGNED NOT NULL ,
    quantity TINYINT UNSIGNED NOT NULL,
    FOREIGN KEY (dish_id) REFERENCES dishes (id)
        ON UPDATE RESTRICT
        ON DELETE CASCADE,
    FOREIGN KEY (ingredient_id) REFERENCES ingredients (id)
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dish_ingredients;
-- +goose StatementEnd
