-- +mig Up
CREATE TABLE projects (
  id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(30),
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME
);

-- +mig Down
DROP TABLE projects;
