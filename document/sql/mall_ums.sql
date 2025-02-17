-- ----------------------------
-- Table structure for ums_admin
-- ----------------------------
DROP TABLE IF EXISTS `ums_admin`;
CREATE TABLE `ums_admin`  (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                              `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                              `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                              `icon` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
                              `email` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
                              `nick_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
                              `note` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注信息',
                              `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
                              `login_time` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
                              `status` int(1) NULL DEFAULT 1 COMMENT '帐号启用状态：0->禁用；1->启用',
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_admin_login_log
-- ----------------------------
DROP TABLE IF EXISTS `ums_admin_login_log`;
CREATE TABLE `ums_admin_login_log`  (
                                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                        `admin_id` bigint(20) NULL DEFAULT NULL,
                                        `create_time` datetime NULL DEFAULT NULL,
                                        `ip` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                        `address` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                        `user_agent` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '浏览器登录类型',
                                        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 413 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台用户登录日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_admin_permission_relation
-- ----------------------------
DROP TABLE IF EXISTS `ums_admin_permission_relation`;
CREATE TABLE `ums_admin_permission_relation`  (
                                                  `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                                  `admin_id` bigint(20) NULL DEFAULT NULL,
                                                  `permission_id` bigint(20) NULL DEFAULT NULL,
                                                  `type` int(1) NULL DEFAULT NULL,
                                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台用户和权限关系表(除角色中定义的权限以外的加减权限)' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_admin_role_relation
-- ----------------------------
DROP TABLE IF EXISTS `ums_admin_role_relation`;
CREATE TABLE `ums_admin_role_relation`  (
                                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                            `admin_id` bigint(20) NULL DEFAULT NULL,
                                            `role_id` bigint(20) NULL DEFAULT NULL,
                                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台用户和角色关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_growth_change_history
-- ----------------------------
DROP TABLE IF EXISTS `ums_growth_change_history`;
CREATE TABLE `ums_growth_change_history`  (
                                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                              `member_id` bigint(20) NULL DEFAULT NULL,
                                              `create_time` datetime NULL DEFAULT NULL,
                                              `change_type` int(1) NULL DEFAULT NULL COMMENT '改变类型：0->增加；1->减少',
                                              `change_count` int(11) NULL DEFAULT NULL COMMENT '积分改变数量',
                                              `operate_man` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作人员',
                                              `operate_note` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作备注',
                                              `source_type` int(1) NULL DEFAULT NULL COMMENT '积分来源：0->购物；1->管理员修改',
                                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '成长值变化历史记录表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_integration_change_history
-- ----------------------------
DROP TABLE IF EXISTS `ums_integration_change_history`;
CREATE TABLE `ums_integration_change_history`  (
                                                   `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                                   `member_id` bigint(20) NULL DEFAULT NULL,
                                                   `create_time` datetime NULL DEFAULT NULL,
                                                   `change_type` int(1) NULL DEFAULT NULL COMMENT '改变类型：0->增加；1->减少',
                                                   `change_count` int(11) NULL DEFAULT NULL COMMENT '积分改变数量',
                                                   `operate_man` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作人员',
                                                   `operate_note` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作备注',
                                                   `source_type` int(1) NULL DEFAULT NULL COMMENT '积分来源：0->购物；1->管理员修改',
                                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '积分变化历史记录表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_integration_consume_setting
-- ----------------------------
DROP TABLE IF EXISTS `ums_integration_consume_setting`;
CREATE TABLE `ums_integration_consume_setting`  (
                                                    `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                                    `deduction_per_amount` int(11) NULL DEFAULT NULL COMMENT '每一元需要抵扣的积分数量',
                                                    `max_percent_per_order` int(11) NULL DEFAULT NULL COMMENT '每笔订单最高抵用百分比',
                                                    `use_unit` int(11) NULL DEFAULT NULL COMMENT '每次使用积分最小单位100',
                                                    `coupon_status` int(1) NULL DEFAULT NULL COMMENT '是否可以和优惠券同用；0->不可以；1->可以',
                                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '积分消费设置' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member
-- ----------------------------
DROP TABLE IF EXISTS `ums_member`;
CREATE TABLE `ums_member`  (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT,
                               `member_level_id` bigint(20) NULL DEFAULT NULL,
                               `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
                               `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
                               `nickname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
                               `phone` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号码',
                               `status` int(1) NULL DEFAULT NULL COMMENT '帐号启用状态:0->禁用；1->启用',
                               `create_time` datetime NULL DEFAULT NULL COMMENT '注册时间',
                               `icon` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
                               `gender` int(1) NULL DEFAULT NULL COMMENT '性别：0->未知；1->男；2->女',
                               `birthday` date NULL DEFAULT NULL COMMENT '生日',
                               `city` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '所做城市',
                               `job` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '职业',
                               `personalized_signature` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '个性签名',
                               `source_type` int(1) NULL DEFAULT NULL COMMENT '用户来源',
                               `integration` int(11) NULL DEFAULT NULL COMMENT '积分',
                               `growth` int(11) NULL DEFAULT NULL COMMENT '成长值',
                               `luckey_count` int(11) NULL DEFAULT NULL COMMENT '剩余抽奖次数',
                               `history_integration` int(11) NULL DEFAULT NULL COMMENT '历史积分数量',
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `idx_username`(`username`) USING BTREE,
                               UNIQUE INDEX `idx_phone`(`phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member_level
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_level`;
CREATE TABLE `ums_member_level`  (
                                     `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                     `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                     `growth_point` int(11) NULL DEFAULT NULL,
                                     `default_status` int(1) NULL DEFAULT NULL COMMENT '是否为默认等级：0->不是；1->是',
                                     `free_freight_point` decimal(10, 2) NULL DEFAULT NULL COMMENT '免运费标准',
                                     `comment_growth_point` int(11) NULL DEFAULT NULL COMMENT '每次评价获取的成长值',
                                     `priviledge_free_freight` int(1) NULL DEFAULT NULL COMMENT '是否有免邮特权',
                                     `priviledge_sign_in` int(1) NULL DEFAULT NULL COMMENT '是否有签到特权',
                                     `priviledge_comment` int(1) NULL DEFAULT NULL COMMENT '是否有评论获奖励特权',
                                     `priviledge_promotion` int(1) NULL DEFAULT NULL COMMENT '是否有专享活动特权',
                                     `priviledge_member_price` int(1) NULL DEFAULT NULL COMMENT '是否有会员价格特权',
                                     `priviledge_birthday` int(1) NULL DEFAULT NULL COMMENT '是否有生日特权',
                                     `note` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会员等级表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member_login_log
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_login_log`;
CREATE TABLE `ums_member_login_log`  (
                                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                         `member_id` bigint(20) NULL DEFAULT NULL,
                                         `create_time` datetime NULL DEFAULT NULL,
                                         `ip` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                         `city` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                         `login_type` int(1) NULL DEFAULT NULL COMMENT '登录类型：0->PC；1->android;2->ios;3->小程序',
                                         `province` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会员登录记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member_member_tag_relation
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_member_tag_relation`;
CREATE TABLE `ums_member_member_tag_relation`  (
                                                   `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                                   `member_id` bigint(20) NULL DEFAULT NULL,
                                                   `tag_id` bigint(20) NULL DEFAULT NULL,
                                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户和标签关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member_product_category_relation
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_product_category_relation`;
CREATE TABLE `ums_member_product_category_relation`  (
                                                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                                         `member_id` bigint(20) NULL DEFAULT NULL,
                                                         `product_category_id` bigint(20) NULL DEFAULT NULL,
                                                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会员与产品分类关系表（用户喜欢的分类）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member_receive_address
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_receive_address`;
CREATE TABLE `ums_member_receive_address`  (
                                               `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                               `member_id` bigint(20) NULL DEFAULT NULL,
                                               `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '收货人名称',
                                               `phone_number` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                               `default_status` int(1) NULL DEFAULT NULL COMMENT '是否为默认',
                                               `post_code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮政编码',
                                               `province` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '省份/直辖市',
                                               `city` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '城市',
                                               `region` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '区',
                                               `detail_address` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '详细地址(街道)',
                                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会员收货地址表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member_rule_setting
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_rule_setting`;
CREATE TABLE `ums_member_rule_setting`  (
                                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                            `continue_sign_day` int(11) NULL DEFAULT NULL COMMENT '连续签到天数',
                                            `continue_sign_point` int(11) NULL DEFAULT NULL COMMENT '连续签到赠送数量',
                                            `consume_per_point` decimal(10, 2) NULL DEFAULT NULL COMMENT '每消费多少元获取1个点',
                                            `low_order_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '最低获取点数的订单金额',
                                            `max_point_per_order` int(11) NULL DEFAULT NULL COMMENT '每笔订单最高获取点数',
                                            `type` int(1) NULL DEFAULT NULL COMMENT '类型：0->积分规则；1->成长值规则',
                                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会员积分成长规则表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member_statistics_info
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_statistics_info`;
CREATE TABLE `ums_member_statistics_info`  (
                                               `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                               `member_id` bigint(20) NULL DEFAULT NULL,
                                               `consume_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '累计消费金额',
                                               `order_count` int(11) NULL DEFAULT NULL COMMENT '订单数量',
                                               `coupon_count` int(11) NULL DEFAULT NULL COMMENT '优惠券数量',
                                               `comment_count` int(11) NULL DEFAULT NULL COMMENT '评价数',
                                               `return_order_count` int(11) NULL DEFAULT NULL COMMENT '退货数量',
                                               `login_count` int(11) NULL DEFAULT NULL COMMENT '登录次数',
                                               `attend_count` int(11) NULL DEFAULT NULL COMMENT '关注数量',
                                               `fans_count` int(11) NULL DEFAULT NULL COMMENT '粉丝数量',
                                               `collect_product_count` int(11) NULL DEFAULT NULL,
                                               `collect_subject_count` int(11) NULL DEFAULT NULL,
                                               `collect_topic_count` int(11) NULL DEFAULT NULL,
                                               `collect_comment_count` int(11) NULL DEFAULT NULL,
                                               `invite_friend_count` int(11) NULL DEFAULT NULL,
                                               `recent_order_time` datetime NULL DEFAULT NULL COMMENT '最后一次下订单时间',
                                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会员统计信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member_tag
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_tag`;
CREATE TABLE `ums_member_tag`  (
                                   `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                   `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                   `finish_order_count` int(11) NULL DEFAULT NULL COMMENT '自动打标签完成订单数量',
                                   `finish_order_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '自动打标签完成订单金额',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户标签表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_member_task
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_task`;
CREATE TABLE `ums_member_task`  (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                    `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                    `growth` int(11) NULL DEFAULT NULL COMMENT '赠送成长值',
                                    `intergration` int(11) NULL DEFAULT NULL COMMENT '赠送积分',
                                    `type` int(1) NULL DEFAULT NULL COMMENT '任务类型：0->新手任务；1->日常任务',
                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会员任务表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_menu
-- ----------------------------
DROP TABLE IF EXISTS `ums_menu`;
CREATE TABLE `ums_menu`  (
                             `id` bigint(20) NOT NULL AUTO_INCREMENT,
                             `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父级ID',
                             `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
                             `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单名称',
                             `level` int(4) NULL DEFAULT NULL COMMENT '菜单级数',
                             `sort` int(4) NULL DEFAULT NULL COMMENT '菜单排序',
                             `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '前端名称',
                             `icon` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '前端图标',
                             `hidden` int(1) NULL DEFAULT NULL COMMENT '前端隐藏',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_permission
-- ----------------------------
DROP TABLE IF EXISTS `ums_permission`;
CREATE TABLE `ums_permission`  (
                                   `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                   `pid` bigint(20) NULL DEFAULT NULL COMMENT '父级权限id',
                                   `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
                                   `value` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限值',
                                   `icon` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图标',
                                   `type` int(1) NULL DEFAULT NULL COMMENT '权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）',
                                   `uri` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '前端资源路径',
                                   `status` int(1) NULL DEFAULT NULL COMMENT '启用状态；0->禁用；1->启用',
                                   `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
                                   `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台用户权限表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_resource
-- ----------------------------
DROP TABLE IF EXISTS `ums_resource`;
CREATE TABLE `ums_resource`  (
                                 `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                 `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
                                 `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '资源名称',
                                 `url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '资源URL',
                                 `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
                                 `category_id` bigint(20) NULL DEFAULT NULL COMMENT '资源分类ID',
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台资源表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_resource_category
-- ----------------------------
DROP TABLE IF EXISTS `ums_resource_category`;
CREATE TABLE `ums_resource_category`  (
                                          `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                          `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
                                          `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '分类名称',
                                          `sort` int(4) NULL DEFAULT NULL COMMENT '排序',
                                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '资源分类表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_role
-- ----------------------------
DROP TABLE IF EXISTS `ums_role`;
CREATE TABLE `ums_role`  (
                             `id` bigint(20) NOT NULL AUTO_INCREMENT,
                             `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
                             `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
                             `admin_count` int(11) NULL DEFAULT NULL COMMENT '后台用户数量',
                             `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
                             `status` int(1) NULL DEFAULT 1 COMMENT '启用状态：0->禁用；1->启用',
                             `sort` int(11) NULL DEFAULT 0,
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台用户角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_role_menu_relation
-- ----------------------------
DROP TABLE IF EXISTS `ums_role_menu_relation`;
CREATE TABLE `ums_role_menu_relation`  (
                                           `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                           `role_id` bigint(20) NULL DEFAULT NULL COMMENT '角色ID',
                                           `menu_id` bigint(20) NULL DEFAULT NULL COMMENT '菜单ID',
                                           PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 127 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台角色菜单关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_role_permission_relation
-- ----------------------------
DROP TABLE IF EXISTS `ums_role_permission_relation`;
CREATE TABLE `ums_role_permission_relation`  (
                                                 `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                                 `role_id` bigint(20) NULL DEFAULT NULL,
                                                 `permission_id` bigint(20) NULL DEFAULT NULL,
                                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台用户角色和权限关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ums_role_resource_relation
-- ----------------------------
DROP TABLE IF EXISTS `ums_role_resource_relation`;
CREATE TABLE `ums_role_resource_relation`  (
                                               `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                               `role_id` bigint(20) NULL DEFAULT NULL COMMENT '角色ID',
                                               `resource_id` bigint(20) NULL DEFAULT NULL COMMENT '资源ID',
                                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 249 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台角色资源关系表' ROW_FORMAT = DYNAMIC;

