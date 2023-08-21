USE bia_db;

CREATE TABLE IF NOT EXISTS `bia_db`.`energy_consumption` (
  `id` VARCHAR(50) NOT NULL,
  `meter_id` INT NOT NULL,
  `active_energy`  DECIMAL(65,20) NOT NULL,
  `reactive_energy`  DECIMAL(65,20) NOT NULL,
  `capacitive_reactive`  DECIMAL(65,20) NOT NULL,
  `solar`  DECIMAL(65,20) NOT NULL,
  `date`  TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

ALTER TABLE energy_consumption ADD INDEX date_consumption_index (date asc);