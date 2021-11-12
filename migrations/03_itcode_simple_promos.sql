-- MySQL dump 10.13  Distrib 8.0.26-16, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: itcode
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*!50717 SELECT COUNT(*) INTO @rocksdb_has_p_s_session_variables FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'performance_schema' AND TABLE_NAME = 'session_variables' */;
/*!50717 SET @rocksdb_get_is_supported = IF (@rocksdb_has_p_s_session_variables, 'SELECT COUNT(*) INTO @rocksdb_is_supported FROM performance_schema.session_variables WHERE VARIABLE_NAME=\'rocksdb_bulk_load\'', 'SELECT 0') */;
/*!50717 PREPARE s FROM @rocksdb_get_is_supported */;
/*!50717 EXECUTE s */;
/*!50717 DEALLOCATE PREPARE s */;
/*!50717 SET @rocksdb_enable_bulk_load = IF (@rocksdb_is_supported, 'SET SESSION rocksdb_bulk_load = 1', 'SET @rocksdb_dummy_bulk_load = 0') */;
/*!50717 PREPARE s FROM @rocksdb_enable_bulk_load */;
/*!50717 EXECUTE s */;
/*!50717 DEALLOCATE PREPARE s */;

--
-- Dumping data for table `promos`
--

/*!40000 ALTER TABLE `promos` DISABLE KEYS */;
INSERT INTO `promos` VALUES (1,NULL,'10percent',1,'percent_discount',10.00,'Скидка 10% за заказ по промокоду','order'),(2,NULL,'noaction',2,'percent_discount',0.00,'Скидка 0 %','order'),(3,NULL,'discount1000',3,'price_discount',1000.00,'Скидка 1000 за заказ по промокоду','order'),(4,NULL,'1peperoni',4,'gift',NULL,'Подарок 1 пеперони за заказ с промокодом','order'),(5,NULL,'2peperoni',5,'gift',NULL,'Подарок 2 пеперони за заказ с промокодом','order'),(99,400.00,'test',99,'gift',30.00,'Сложная скидка','item');
/*!40000 ALTER TABLE `promos` ENABLE KEYS */;

--
-- Dumping data for table `promo_condition_item`
--

/*!40000 ALTER TABLE `promo_condition_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `promo_condition_item` ENABLE KEYS */;

--
-- Dumping data for table `promo_exclusions`
--

/*!40000 ALTER TABLE `promo_exclusions` DISABLE KEYS */;
INSERT INTO `promo_exclusions` VALUES (1,99,1),(2,99,2);
/*!40000 ALTER TABLE `promo_exclusions` ENABLE KEYS */;

--
-- Dumping data for table `promo_gift_items`
--

/*!40000 ALTER TABLE `promo_gift_items` DISABLE KEYS */;
INSERT INTO `promo_gift_items` VALUES (1,4,1),(2,5,1),(3,5,1);
/*!40000 ALTER TABLE `promo_gift_items` ENABLE KEYS */;

--
-- Dumping data for table `promo_item_selector`
--

/*!40000 ALTER TABLE `promo_item_selector` DISABLE KEYS */;
INSERT INTO `promo_item_selector` VALUES (1,1,1),(2,10,99),(3,15,99),(4,20,99);
/*!40000 ALTER TABLE `promo_item_selector` ENABLE KEYS */;
/*!50112 SET @disable_bulk_load = IF (@is_rocksdb_supported, 'SET SESSION rocksdb_bulk_load = @old_rocksdb_bulk_load', 'SET @dummy_rocksdb_bulk_load = 0') */;
/*!50112 PREPARE s FROM @disable_bulk_load */;
/*!50112 EXECUTE s */;
/*!50112 DEALLOCATE PREPARE s */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-11-04 14:17:30
