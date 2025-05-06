SET
    SEARCH_PATH TO  PUBLIC;

CREATE TABLE userprofile (
  id         text        NOT NULL,
  first_name text        NOT NULL,
  last_name  text        NOT NULL,
  email      text        NOT NULL,
  password   text        NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  deleted_at timestamptz NULL,
  PRIMARY KEY (id)
);

CREATE TABLE role (
  id           SERIAL       NOT NULL,
  name         text         NOT NULL,
  description  text         NULL,
  created_at   timestamptz  NOT NULL DEFAULT NOW(),
  updated_at   timestamptz  NOT NULL DEFAULT NOW(),
  deleted_at   timestamptz  NULL,
  PRIMARY KEY (id)
);

CREATE TABLE userrole (
  id           SERIAL       NOT NULL,
  user_id      text         NOT NULL,
  role_id      SERIAL       NOT NULL,
  created_at   timestamptz  NOT NULL DEFAULT NOW(),
  updated_at   timestamptz  NOT NULL DEFAULT NOW(),
  deleted_at   timestamptz  NULL,
  PRIMARY KEY (id)
);

CREATE TABLE permission (
  id            SERIAL        NOT NULL,
  name          text          NOT NULL,
  srv_name      varchar(20)   NOT NULL,
  func_name1    varchar(20)   NOT NULL,
  func_name2    varchar(20)   NULL,       
  func_name3    varchar(20)   NULL,       
  created_at    timestamptz   NOT NULL DEFAULT NOW(),
  updated_at    timestamptz   NOT NULL DEFAULT NOW(),
  deleted_at    timestamptz   NULL,
  PRIMARY KEY (id)
);

CREATE TABLE rbac (
  id            SERIAL       NOT NULL,
  role_id       SERIAL       NOT NULL,
  permission_id SERIAL       NOT NULL,
  created_at    timestamptz  NOT NULL DEFAULT NOW(),
  updated_at    timestamptz  NOT NULL DEFAULT NOW(),
  deleted_at    timestamptz  NULL,
  PRIMARY KEY (id)
);


INSERT INTO userprofile(id, email, password, first_name, last_name) values('776c8b50-27b2-4b4b-99b1-1965f7a23f38','quykim117@gmail.com','$2a$10$fvVwsdAeKae9IvbCMNXP6OyHWShndPbPkFeJAuQkWYvPIT0QNvtJ6','quy', 'kim');
INSERT INTO userprofile(id,email, password, first_name, last_name) values('7007d648-f1ba-43e6-9ec7-e6823f83ab58','admin111@gmail.com','$2a$10$fvVwsdAeKae9IvbCMNXP6OyHWShndPbPkFeJAuQkWYvPIT0QNvtJ6','admin', 'john');

INSERT INTO role(name) values('employee');
INSERT INTO role(name) values('employer');

INSERT INTO userrole(user_id, role_id) values('776c8b50-27b2-4b4b-99b1-1965f7a23f38',1);
INSERT INTO userrole(user_id, role_id) values('7007d648-f1ba-43e6-9ec7-e6823f83ab58',2);

INSERT INTO permission(name, srv_name, func_name) values('view_assisgned_tasks', 'task', 'read_tasks_assigned');
INSERT INTO permission(name, srv_name, func_name) values('update_assisgned_tasks', 'task', 'update_task_status');
INSERT INTO permission(name, srv_name, func_name) values('create_and_assisgn_tasks', 'task', 'create_task','assign_task');
INSERT INTO permission(name, srv_name, func_name) values('view_tasks', 'task', 'list_tasks');

INSERT INTO rbac(role_id, permission_id) values(1, 1);
INSERT INTO rbac(role_id, permission_id) values(1, 2);
INSERT INTO rbac(role_id, permission_id) values(2, 3);
INSERT INTO rbac(role_id, permission_id) values(2, 4);