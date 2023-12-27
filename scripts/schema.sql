CREATE TABLE  IF NOT EXISTS `tbl_ticks` (
    `timestamp` bigint unsigned NOT NULL,
    `symbol` varchar(8) NOT NULL,
    `bid` float NOT NULL,
    `ask` float NOT NULL,
    CONSTRAINT tick_pk
        PRIMARY KEY (`timestamp`, `symbol`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `tbl_orders` (
    `timestamp` bigint unsigned NOT NULL,
    `productId` varchar(10) NOT NULL,
    `side`      varchar(4) NOT NULL,
    `size`      varchar(30) NOT NULL,
    `price`     varchar(30) NOT NULL,
    `status`    varchar(30) NOT NULL,
    CONSTRAINT  order_pk
        PRIMARY KEY (`timestamp`, productId)
) ENGINE = InnoDB;