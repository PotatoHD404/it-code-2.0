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
-- Dumping data for table `items`
--

/*!40000 ALTER TABLE `items` DISABLE KEYS */;
INSERT INTO `items` VALUES (1,'Пицца Пепперони',500.00),(2,'Пицца По-домащнему',520.00),(3,'Пицца 4 сыра',500.00),(4,'Пицца Баварская',550.00),(5,'Пицца Гавайская',530.00),(6,'Пицца Ветчина и грибы',550.00),(7,'Пицца Итальянцы в России',540.00),(8,'Пицца Всё и сразу',500.00),(9,'Пицца Деревенская',520.00),(10,'Пицца Чесночный цыпа',530.00),(11,'Роллы Филадельфия',350.00),(12,'Роллы Калифорния',320.00),(13,'Роллы Дракон',300.00),(14,'Роллы Фудзияма',280.00),(15,'Роллы Шен',300.00),(16,'Роллы Якудза',320.00),(17,'Роллы Тамаго',330.00),(18,'Роллы Киото',290.00),(19,'Роллы Император',320.00),(20,'Картофель по-деревенски',200.00),(21,'Картофель фри',150.00),(22,'Наггетсы',280.00),(23,'Удон с курицей в соусе терияки',300.00),(24,'Удон с креветками и сладким соусом чили',330.00),(25,'Соба с грибами в соусе терияки',280.00),(26,'Харусаме с курицей и имбирем',290.00),(27,'Фунчоза с курицей в сливочном соусе',270.00),(28,'Паста Карбонара',320.00),(29,'Паста Болоньезе',350.00),(30,'Паста с ветчиной и грибами',310.00),(31,'Паста с семгой',400.00),(32,'Салат Цезарь с индейкой',180.00),(33,'Салат Цезарь с креветками',200.00),(34,'Салат Олиьве',160.00),(35,'Салат Греческий',180.00),(36,'Суп Сборная солянка',220.00),(37,'Суп Украинский борш',230.00),(38,'Суп Уха по-фински',200.00),(39,'Чизкейк Клубничный',150.00),(40,'Чизкейк Карамельный с арахисом',160.00),(41,'Чизкейк Кокосовый',150.00),(42,'Компот из черной смородины 1л',200.00),(43,'Компот из черной смородины 0.5л',100.00),(44,'Компот клюквенный 1л',200.00),(45,'Компот клюквенный 0.5л',100.00),(46,'7АП 1л',110.00),(47,'Пепси 1л',100.00),(48,'Миринда 1л',100.00),(49,'Набор салфеток и зубочисток',10.00),(50,'Палочки',10.00),(51,'Вилка одноразовая',5.00),(52,'Ложка одноразовая',5.00),(53,'Соус сырный',25.00),(54,'Соус чесночный',25.00),(55,'Соус барбекю',25.00),(56,'Соус кисло-сладкий',25.00),(57,'Имбирь',25.00),(58,'Васаби',25.00);
/*!40000 ALTER TABLE `items` ENABLE KEYS */;
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
