CREATE TABLE `app`
(
    `id`          bigint       NOT NULL AUTO_INCREMENT,
    `name`        varchar(32)  NOT NULL,
    `description` varchar(256),
    `create_at`   datetime,
    `update_at`   datetime,
    `delete_at`   datetime,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `cluster`
(
    `id`           bigint       NOT NULL AUTO_INCREMENT,
    `app_name`     varchar(32)  NOT NULL,
    `cluster_name` varchar(32)  NOT NULL,
    `description`  varchar(256),
    `create_at`    datetime,
    `update_at`    datetime,
    `delete_at`    datetime,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `namespace`
(
    `id`             bigint       NOT NULL AUTO_INCREMENT,
    `app_name`       varchar(32)  NOT NULL,
    `cluster_name`   varchar(32)  NOT NULL,
    `namespace_name` varchar(32)  NOT NULL,
    `format`         varchar(16)  NOT NULL,
    `value`          text         NOT NULL,
    `Released`       tinyint,
    `description`    varchar(256),
    `create_at`      datetime,
    `update_at`      datetime,
    `delete_at`      datetime,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `release`
(
    `id`             bigint       NOT NULL AUTO_INCREMENT,
    `app_name`       varchar(32)  NOT NULL,
    `cluster_name`   varchar(32)  NOT NULL,
    `namespace_name` varchar(32)  NOT NULL,
    `format`         varchar(16)  NOT NULL,
    `value`          text         NOT NULL,
    `version`        varchar(16)  NOT NULL,
    `description`    varchar(256),
    `create_at`      datetime,
    `update_at`      datetime,
    `delete_at`      datetime,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;