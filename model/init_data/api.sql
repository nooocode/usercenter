SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
BEGIN;

INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (1000, '/api/auth/login', 'auth', 'POST', '用户登录', 1, 0);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (1001, '/api/auth/logout', 'auth', 'POST', '退出登录', 1, 0);

INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (2000, '/api/add', 'auth', 'POST', '新增API', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (2001, '/api/delete', 'auth', 'DELETE', '删除API', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (2002, '/api/update', 'auth', 'PUT', '更新API', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (2003, '/api/detail', 'auth', 'GET', '查询API明细', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (2004, '/api/query', 'auth', 'GET', '查询API(分页)', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (2005, '/api/all', 'auth', 'GET', '查询所有API', 1, 1);

INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (3001, '/api/menu/add', 'auth', 'POST', '新增菜单', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (3002, '/api/menu/delete', 'auth', 'DELETE', '删除菜单', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (3003, '/api/menu/update', 'auth', 'PUT', '更新菜单', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (3004, '/api/menu/detail', 'auth', 'GET', '查询菜单明细', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (3005, '/api/menu/query', 'auth', 'GET', '查询菜单(分页)', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (3006, '/api/menu/all', 'auth', 'GET', '查询所有菜单', 1, 1);

INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (4001, '/api/role/add', 'auth', 'POST', '新增角色', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (4002, '/api/role/delete', 'auth', 'DELETE', '删除角色', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (4003, '/api/role/update', 'auth', 'PUT', '更新角色', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (4004, '/api/role/detail', 'auth', 'GET', '查询角色明细', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (4005, '/api/role/query', 'auth', 'GET', '查询角色(分页)', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (4006, '/api/role/all', 'auth', 'GET', '查询所有角色', 1, 1);

INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (5001, '/api/user/add', 'auth', 'POST', '新增用户', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (5002, '/api/user/delete', 'auth', 'DELETE', '删除用户', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (5003, '/api/user/update', 'auth', 'PUT', '更新用户', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (5004, '/api/user/detail', 'auth', 'GET', '查询用户明细', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (5005, '/api/user/query', 'auth', 'GET', '查询用户(分页)', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (5006, '/api/user/all', 'auth', 'GET', '查询所有用户', 1, 1);


INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (6001, '/api/tenant/add', 'auth', 'POST', '新增租户', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (6002, '/api/tenant/delete', 'auth', 'DELETE', '删除租户', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (6003, '/api/tenant/update', 'auth', 'PUT', '更新租户', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (6004, '/api/tenant/detail', 'auth', 'GET', '查询租户明细', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (6005, '/api/tenant/query', 'auth', 'GET', '查询租户(分页)', 1, 1);
INSERT INTO `api` (`id`,  `uri`, `group`, `method`, `description`, `enable`, `check_auth`) VALUES (6006, '/api/tenant/all', 'auth', 'GET', '查询所有租户', 1, 1);

COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

alter table api AUTO_INCREMENT = 10000;