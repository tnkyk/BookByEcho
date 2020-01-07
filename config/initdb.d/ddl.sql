/*
外部キーの命名規則：fk_*****_****

FOREIGN KEY ('A')
REFERENCES `mydb`.`table` (`B`)

名前はA=Bである。

また、A・Bはデータ型が一致してなくてはならない。


CONSTRAINT `fk_category_book_category`
    FOREIGN KEY (`category_id`)
    REFERENCES `mydb`.`category` (`category_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,

*/

-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `mydb` DEFAULT CHARACTER SET utf8 ;
USE `mydb` ;

-- -----------------------------------------------------
-- Table `mydb`.`user_prof`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`user_prof` (
  `user_prof_id` VARCHAR(50) CHARACTER SET 'utf8' NOT NULL,
  `name` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL,
  `favorite` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL,
  `user_image` VARCHAR(100) CHARACTER SET 'utf8' NOT NULL,
  `crearted_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`user_prof_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`user` (
  `user_id` VARCHAR(50) CHARACTER SET 'utf8' NOT NULL,
  `email` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL,
  `password` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL,
  `user_prof_id` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL,
  INDEX `user_prof_id_idx` (`user_prof_id` ASC) VISIBLE,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_prof`
    FOREIGN KEY (`user_prof_id`)
    REFERENCES `mydb`.`user_prof` (`user_prof_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`category`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`category` (
  `category_id` VARCHAR(50) CHARACTER SET 'utf8' NOT NULL,
  `name` VARCHAR(255) CHARACTER SET 'utf8' NOT NULL,
  PRIMARY KEY (`category_id`));


-- -----------------------------------------------------
-- Table `mydb`.`book`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`book` (
  `book_id` VARCHAR(50) CHARACTER SET 'utf8' NOT NULL,
  `title` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL,
  `author` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL,
  `content` VARCHAR(200) CHARACTER SET 'utf8' NOT NULL,
  `evaluation` INT NOT NULL,
  `category_id` VARCHAR(50) CHARACTER SET 'utf8' NULL,
  `user_id` VARCHAR(50) CHARACTER SET 'utf8' NOT NULL,
  INDEX `category_id_idx` (`category_id` ASC) VISIBLE,
  PRIMARY KEY (`book_id`),
  INDEX `provide_user_id_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_category_book_category`
    FOREIGN KEY (`category_id`)
    REFERENCES `mydb`.`category` (`category_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_provide_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `mydb`.`user` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`user_book`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`user_book` (
  `user_book_id` VARCHAR(50) CHARACTER SET 'utf8' NOT NULL,
  `reservation_user_id` VARCHAR(50) CHARACTER SET 'utf8' NOT NULL,
  `comment` TEXT CHARACTER SET 'utf8' NOT NULL,
  `evaluation` INT NOT NULL,
  `state` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL,
  `reservation_at` DATETIME NOT NULL,
  `finish_at` DATETIME NULL,
  `book_id` VARCHAR(50) CHARACTER SET 'utf8' NOT NULL,
  PRIMARY KEY (`user_book_id`),
  INDEX `book_id_idx` (`book_id` ASC) VISIBLE,
  INDEX `user_id_idx` (`reservation_user_id` ASC) VISIBLE,
  CONSTRAINT `fk_book_user_book`
    FOREIGN KEY (`book_id`)
    REFERENCES `mydb`.`book` (`book_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_user_book_user`
    FOREIGN KEY (`reservation_user_id`)
    REFERENCES `mydb`.`user` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
