SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
BEGIN;
-- web管理后台
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (100000000,  0, 0, 'web', 'web', 0, '', 1, 0, 0, '管理后台菜单', 'setting', 0);

-- 首页
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (101000000,  1, 100000000, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, 0, 0, '首页', 'setting', 0);

-- 系统配置
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (108000000,  1, 100000000, 'system', 'system', 0, 'view/system/index.vue', 99, 0, 0, '系统配置', 's-cooperation', 0);
-- 用户管理
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (108010000,  2, 108000000, 'user', 'user', 0, 'view/system/user/user.vue', 1, 0, 0, '用户管理', 'coordinate', 0);
INSERT INTO `menu_funcs` (`id`, `menu_id`, `name`, `title`, `hidden`) VALUES (108010000, 108010000, 'AddUser', '新增', 0);
-- 新增用户API
INSERT INTO `menu_func_apis` (`id`,`menu_func_id`, `api_id`) VALUES (108010100,108010000, 5001);
-- 查询所有角色API
INSERT INTO `menu_func_apis` (`id`,`menu_func_id`, `api_id`) VALUES (108010200,108010000, 4006);
-- 查询用户明细API
INSERT INTO `menu_func_apis` (`id`,`menu_func_id`, `api_id`) VALUES (108010300,108010000, 5004);
-- 查询用户(分页)API
INSERT INTO `menu_func_apis` (`id`,`menu_func_id`, `api_id`) VALUES (108010400,108010000, 5005);

INSERT INTO `menu_funcs` (`id`, `menu_id`, `name`, `title`, `hidden`) VALUES (108010100, 108010000, 'QueryUser', '查询', 0);
INSERT INTO `menu_funcs` (`id`, `menu_id`, `name`, `title`, `hidden`) VALUES (108010200, 108010000, 'DeleteUser', '删除', 0);
INSERT INTO `menu_funcs` (`id`, `menu_id`, `name`, `title`, `hidden`) VALUES (108010300, 108010000, 'UpdateUser', '更新', 0);

INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (108020000,  2, 108000000, 'role', 'role', 0, 'view/system/role/role.vue', 2, 0, 0, '角色管理', 's-custom', 0);
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (108030000,  2, 108000000, 'customer', 'customer', 0, 'view/system/customer/customer.vue', 3, 0, 0, '客户列表（资源示例）', 's-custom', 0);
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (108050000,  2, 108000000, 'operation', 'operation', 0, 'view/system/operation/sysOperationRecord.vue', 4, 0, 0, '操作历史', 'time', 0);
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (108040000,  2, 108000000, 'state', 'state', 0, 'view/system/state.vue', 5, 0, 0, '服务器状态', 'cloudy', 0);

INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (109000000,  1, 100000000, 'develop', 'develop', 0, 'view/develop/system/system.vue', 100, 0, 0, '开发工具', 's-operation', 0);
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (109010000,  2, 109000000, 'menu', 'menu', 0, 'view/develop/menu/menu.vue', 1, 1, 0, '菜单管理', 's-order', 0);
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (109020000,  2, 109000000, 'api', 'api', 0, 'view/develop/api/api.vue', 2, 1, 0, 'api管理', 's-platform', 0);
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (109030000,  2, 109000000, 'autoCode', 'autoCode', 0, 'view/develop/autoCode/index.vue', 3, 1, 0, '代码生成器', 'cpu', 0);
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (109040000,  2, 109000000, 'formCreate', 'formCreate', 0, 'view/develop/formCreate/index.vue', 4, 1, 0, '表单生成器', 'magic-stick', 0);
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (109050000,  2, 109000000, 'dictionary', 'dictionary', 0, 'view/develop/dictionary/sysDictionary.vue', 5, 0, 0, '字典管理', 'notebook-2', 0);
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (109060000,  2, 109000000, 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/develop/dictionary/sysDictionaryDetail.vue', 6, 0, 0, '字典详情', 's-order', 0);

-- APP
INSERT INTO `menus` (`id`, `level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `cache`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (200000000,  0, 0, 'app', 'app', 0, '', 2, 0, 0, 'APP菜单', 'setting', 0);

COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

alter table menus AUTO_INCREMENT = 1000000000;

alter table menus_parameters AUTO_INCREMENT = 1000000000;

alter table menus_funcs AUTO_INCREMENT = 1000000000;

alter table menus_func_apis AUTO_INCREMENT = 1000000000;