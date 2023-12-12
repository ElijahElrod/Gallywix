CREATE TABLE  IF NOT EXISTS `tbl_ticks` (
    `timestamp` bigint unsigned NOT NULL,
    `symbol` varchar(8) NOT NULL,
    `bid` float NOT NULL,
    `ask` float NOT NULL,
    CONSTRAINT tick_pk
        PRIMARY KEY (`timestamp`, `symbol`)
) ENGINE = InnoDB;