-- +mig Up
CREATE TABLE tasks (
  id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  data VARCHAR(255) NOT NULL,
  current_state ENUM('todo', 'done', 'block'),
  todo_since DATETIME,
  done_since DATETIME,
  blocked_since DATETIME,
  archived_at DATETIME,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME,
  project_id INT UNSIGNED NOT NULL,
  CONSTRAINT fk_project_tasks FOREIGN KEY (project_id) REFERENCES projects(id)
);
CREATE INDEX id_tasks_by_project_by_state
  ON tasks (project_id, archived_at, current_state);

-- +mig Down
ALTER TABLE tasks DROP FOREIGN KEY fk_project_tasks;
DROP INDEX id_tasks_by_project_by_state ON tasks;
DROP TABLE tasks;