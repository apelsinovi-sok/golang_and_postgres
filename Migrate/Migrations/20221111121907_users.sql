-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id UUID NOT NULL,
  Firstname VARCHAR(250) NOT NULL,
  Age INT NOT NULL,
  Created   TIMESTAMP,
  PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
