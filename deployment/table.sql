create table user
(
    id         bigint auto_increment
        primary key,
    name       varchar(20) default '' null comment '名称',
    account    varchar(20) default '' null comment '账号',
    password   varchar(50) default '' null comment '密码',
    createTime timestamp              null on update CURRENT_TIMESTAMP comment '创建时间',
    constraint user_account_uindex
        unique (account)
)
    comment '用户表';

