--  创建数据库
create database if not exists rice_oj;

use rice_oj;

-- 题目表
create table if not exists t_question (
     id bigint primary key auto_increment comment '题目唯一标识',
     title varchar(64) not null comment '题目标题',
     content text  null comment '题目内容',
     tags	varchar(64)	null comment '题目标签JSON字符串（比如简单、二叉树）',
     cases	varchar(128)	null comment '用例JSON字符串（比如输入用例、输出用例）',
     limit_config	varchar(128)	null comment '限制配置JSON字符串（比如时间限制，内存限制）',
     answer	text	null comment 	'题解',
     user_id	datetime	not null comment '创建用户id',
     create_at datetime default CURRENT_TIMESTAMP comment  '创建时间',
     update_at datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment '更新时间',
     is_delete	tinyint default 0 comment 	'是否被删除（逻辑删除）0-未删除，1-删除'
);


-- 题目提交表
create table if not exists t_question_submit (
     id           bigint      primary key auto_increment                    comment '题目提交唯一标识',
     question_id  bigint      not null                                      comment '题目id',
     code         text        not null                                      comment '提交代码',
     language     varchar(64) not null                                      comment '语言（JAVA、Go等）',
     status       tinyint     default 0                                     comment '判题结果 1-等待中 2-失败 3-成功 4-系统错误',
     judge_info   varchar(128) null                                         comment '判题信息JSON字符串（比如消耗时间，占用内存）',
     user_id      bigint      not null                                      comment '提交用户id',
     create_at    datetime    default CURRENT_TIMESTAMP                     comment '创建时间',
     update_at    datetime    default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment '更新时间',
     is_delete    tinyint     default 0                                     comment '是否被删除（逻辑删除）0-未删除，1-删除',

     index idx_question_id (question_id),
     index idx_user_id (user_id)
) engine=InnoDB default charset=utf8mb4 comment '题目提交表';


