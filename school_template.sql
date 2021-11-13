-- MySQL dump 10.19  Distrib 10.3.31-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: school
-- ------------------------------------------------------
-- Server version	10.3.31-MariaDB-0ubuntu0.20.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `behavioral_point`
--

DROP TABLE IF EXISTS `behavioral_point`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `behavioral_point` (
  `ref` int(11) NOT NULL AUTO_INCREMENT,
  `groupno` int(11) DEFAULT NULL,
  `groupname` varchar(255) DEFAULT NULL,
  `topicno` int(11) DEFAULT NULL,
  `topicname` varchar(255) DEFAULT NULL,
  `point` int(11) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=126 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `checkin_daily`
--

DROP TABLE IF EXISTS `checkin_daily`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `checkin_daily` (
  `created` date NOT NULL,
  `code` varchar(50) NOT NULL,
  `fullname` varchar(255) DEFAULT NULL,
  `grade` varchar(50) DEFAULT NULL,
  `time_in` time DEFAULT NULL,
  `time_out` time DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`created`,`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `checkin_group`
--

DROP TABLE IF EXISTS `checkin_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `checkin_group` (
  `ref` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `late1` time DEFAULT NULL,
  `late1_point` int(11) DEFAULT NULL,
  `late2` time DEFAULT NULL,
  `late2_point` int(11) DEFAULT NULL,
  `late3` time DEFAULT NULL,
  `late3_point` int(11) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `checkinout`
--

DROP TABLE IF EXISTS `checkinout`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `checkinout` (
  `ref` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created` datetime DEFAULT NULL,
  `idcard` int(11) DEFAULT NULL,
  `image_file` varchar(255) DEFAULT NULL,
  `camno` int(11) DEFAULT NULL,
  `temperature` decimal(10,2) DEFAULT NULL,
  `cutpoint` int(11) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=441 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `checkinprofile`
--

DROP TABLE IF EXISTS `checkinprofile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `checkinprofile` (
  `grp` int(11) NOT NULL,
  `date_in` date NOT NULL,
  `yearedu` int(11) DEFAULT NULL,
  `name` varchar(25) DEFAULT NULL,
  `time1` time DEFAULT NULL,
  `time2` time DEFAULT NULL,
  `time3` time DEFAULT NULL,
  `point1` int(11) DEFAULT NULL,
  `point2` int(11) DEFAULT NULL,
  `point3` int(11) DEFAULT NULL,
  PRIMARY KEY (`grp`,`date_in`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `checkinsubject`
--

DROP TABLE IF EXISTS `checkinsubject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `checkinsubject` (
  `created` date NOT NULL,
  `period` int(11) NOT NULL,
  `room_ref` int(11) NOT NULL,
  `teacher_ref` int(11) DEFAULT NULL,
  `subject_ref` int(11) DEFAULT NULL,
  `note` text DEFAULT NULL,
  PRIMARY KEY (`created`,`period`,`room_ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `checkinsubject_student`
--

DROP TABLE IF EXISTS `checkinsubject_student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `checkinsubject_student` (
  `created` date NOT NULL,
  `period` int(11) NOT NULL,
  `room_ref` int(11) NOT NULL,
  `student_ref` int(11) NOT NULL,
  `student_code` varchar(10) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`created`,`period`,`room_ref`,`student_ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `classroom`
--

DROP TABLE IF EXISTS `classroom`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `classroom` (
  `ref` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `grade_ref` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `building_ref` int(11) DEFAULT NULL,
  `floor` int(11) DEFAULT NULL,
  `teacher_ref` int(11) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=122 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cutpoint`
--

DROP TABLE IF EXISTS `cutpoint`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cutpoint` (
  `ref` int(11) NOT NULL AUTO_INCREMENT,
  `createdate` date DEFAULT NULL,
  `behavior_ref` int(11) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cutpoint_students`
--

DROP TABLE IF EXISTS `cutpoint_students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cutpoint_students` (
  `cutpoint_ref` int(11) NOT NULL,
  `student_ref` int(11) NOT NULL,
  PRIMARY KEY (`cutpoint_ref`,`student_ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `department`
--

DROP TABLE IF EXISTS `department`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `department` (
  `ref` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(10) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `employee`
--

DROP TABLE IF EXISTS `employee`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `employee` (
  `ref` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(10) DEFAULT NULL,
  `title_ref` int(11) DEFAULT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `dept_ref` int(11) DEFAULT NULL,
  `room_ref` int(11) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `grade`
--

DROP TABLE IF EXISTS `grade`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `grade` (
  `ref` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(25) DEFAULT NULL,
  `rank` int(11) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `importstudent`
--

DROP TABLE IF EXISTS `importstudent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `importstudent` (
  `created` datetime NOT NULL,
  `room_ref` int(11) DEFAULT NULL,
  `code` varchar(25) DEFAULT NULL,
  `no` int(11) DEFAULT NULL,
  `firstname` varchar(25) DEFAULT NULL,
  `lastname` varchar(25) DEFAULT NULL,
  PRIMARY KEY (`created`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `occupation`
--

DROP TABLE IF EXISTS `occupation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `occupation` (
  `ref` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `parents`
--

DROP TABLE IF EXISTS `parents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `parents` (
  `ref` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `idcard` varchar(13) DEFAULT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `gender` int(11) DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `occupation_code` int(11) DEFAULT NULL,
  `occupation_desc` varchar(255) DEFAULT NULL,
  `income` decimal(10,0) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `line` varchar(255) DEFAULT NULL,
  `facebook` varchar(255) DEFAULT NULL,
  `office_name` varchar(255) DEFAULT NULL,
  `addr1` varchar(255) DEFAULT NULL,
  `addr2` varchar(255) DEFAULT NULL,
  `addr3` varchar(255) DEFAULT NULL,
  `province_code` varchar(255) DEFAULT NULL,
  `province_name` varchar(255) DEFAULT NULL,
  `zipcode` varchar(5) DEFAULT NULL,
  `title` int(11) DEFAULT NULL,
  `lineuid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=399 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `point_students`
--

DROP TABLE IF EXISTS `point_students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `point_students` (
  `created` datetime NOT NULL,
  `behavior_ref` int(11) NOT NULL,
  `student_code` varchar(10) NOT NULL,
  `yearedu` int(11) DEFAULT NULL,
  `point` int(11) DEFAULT 0,
  PRIMARY KEY (`created`,`behavior_ref`,`student_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `province`
--

DROP TABLE IF EXISTS `province`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `province` (
  `ref` int(11) NOT NULL AUTO_INCREMENT,
  `rank` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `question_ans`
--

DROP TABLE IF EXISTS `question_ans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `question_ans` (
  `student_code` varchar(50) NOT NULL,
  `question_code` int(11) NOT NULL,
  `question_no` int(11) NOT NULL,
  `answer` int(11) DEFAULT NULL,
  PRIMARY KEY (`student_code`,`question_code`,`question_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sarary_rate`
--

DROP TABLE IF EXISTS `sarary_rate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sarary_rate` (
  `ref` int(11) NOT NULL AUTO_INCREMENT,
  `rank` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sdq`
--

DROP TABLE IF EXISTS `sdq`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sdq` (
  `code` varchar(10) NOT NULL,
  `type` int(11) NOT NULL,
  `composedby` varchar(10) DEFAULT NULL,
  `topic1_1` int(11) DEFAULT NULL,
  `topic1_2` int(11) DEFAULT NULL,
  `topic1_3` int(11) DEFAULT NULL,
  `topic1_4` int(11) DEFAULT NULL,
  `topic1_5` int(11) DEFAULT NULL,
  `topic1_6` int(11) DEFAULT NULL,
  `topic1_7` int(11) DEFAULT NULL,
  `topic1_8` int(11) DEFAULT NULL,
  `topic1_9` int(11) DEFAULT NULL,
  `topic1_10` int(11) DEFAULT NULL,
  `topic1_11` int(11) DEFAULT NULL,
  `topic1_12` int(11) DEFAULT NULL,
  `topic1_13` int(11) DEFAULT NULL,
  `topic1_14` int(11) DEFAULT NULL,
  `topic1_15` int(11) DEFAULT NULL,
  `topic1_16` int(11) DEFAULT NULL,
  `topic1_17` int(11) DEFAULT NULL,
  `topic1_18` int(11) DEFAULT NULL,
  `topic1_19` int(11) DEFAULT NULL,
  `topic1_20` int(11) DEFAULT NULL,
  `topic1_21` int(11) DEFAULT NULL,
  `topic1_22` int(11) DEFAULT NULL,
  `topic1_23` int(11) DEFAULT NULL,
  `topic1_24` int(11) DEFAULT NULL,
  `topic1_25` int(11) DEFAULT NULL,
  `topic2` int(11) DEFAULT NULL,
  `topic2_1` int(11) DEFAULT NULL,
  `topic2_2` int(11) DEFAULT NULL,
  `topic2_3` int(11) DEFAULT NULL,
  `topic2_4` int(11) DEFAULT NULL,
  `topic2_5` int(11) DEFAULT NULL,
  `topic2_6` int(11) DEFAULT NULL,
  `topic2_7` int(11) DEFAULT NULL,
  `topic3_1` int(11) DEFAULT NULL,
  `topic3_2` int(11) DEFAULT NULL,
  `topic3_3` int(11) DEFAULT NULL,
  `topic3_4` int(11) DEFAULT NULL,
  `topic3_5` int(11) DEFAULT NULL,
  `topic3_6` int(11) DEFAULT NULL,
  `topic3_7` int(11) DEFAULT NULL,
  `topic3_8` int(11) DEFAULT NULL,
  `topic3_9` int(11) DEFAULT NULL,
  `topic3_10` int(11) DEFAULT NULL,
  `topic3_11` int(11) DEFAULT NULL,
  `topic3_12` int(11) DEFAULT NULL,
  `topic3_13` int(11) DEFAULT NULL,
  `topic3_14` int(11) DEFAULT NULL,
  `topic3_15` int(11) DEFAULT NULL,
  `topic3_16` int(11) DEFAULT NULL,
  `topic3_17` int(11) DEFAULT NULL,
  `topic3_18` int(11) DEFAULT NULL,
  `topic3_19` int(11) DEFAULT NULL,
  `topic3_20` int(11) DEFAULT NULL,
  `topic3_21` int(11) DEFAULT NULL,
  `topic3_22` int(11) DEFAULT NULL,
  `topic3_23` int(11) DEFAULT NULL,
  `topic3_24` int(11) DEFAULT NULL,
  `topic3_25` int(11) DEFAULT NULL,
  `topic3_26` int(11) DEFAULT NULL,
  `topic3_27` int(11) DEFAULT NULL,
  `topic3_28` int(11) DEFAULT NULL,
  `topic3_29` int(11) DEFAULT NULL,
  `topic3_30` int(11) DEFAULT NULL,
  `topic3_31` int(11) DEFAULT NULL,
  `topic3_32` int(11) DEFAULT NULL,
  `topic3_33` int(11) DEFAULT NULL,
  `topic3_34` int(11) DEFAULT NULL,
  `topic3_35` int(11) DEFAULT NULL,
  `topic3_36` int(11) DEFAULT NULL,
  `topic3_37` int(11) DEFAULT NULL,
  `topic3_38` int(11) DEFAULT NULL,
  `topic3_39` int(11) DEFAULT NULL,
  `topic3_40` int(11) DEFAULT NULL,
  `topic3_41` int(11) DEFAULT NULL,
  `topic3_42` int(11) DEFAULT NULL,
  `topic3_43` int(11) DEFAULT NULL,
  `topic3_44` int(11) DEFAULT NULL,
  `topic3_45` int(11) DEFAULT NULL,
  `topic3_46` int(11) DEFAULT NULL,
  `topic3_47` int(11) DEFAULT NULL,
  `topic3_48` int(11) DEFAULT NULL,
  `topic3_49` int(11) DEFAULT NULL,
  `topic3_50` int(11) DEFAULT NULL,
  `topic3_51` int(11) DEFAULT NULL,
  `topic3_52` int(11) DEFAULT NULL,
  PRIMARY KEY (`code`,`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `student_parents`
--

DROP TABLE IF EXISTS `student_parents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `student_parents` (
  `student_ref` int(11) NOT NULL,
  `parent_ref` int(11) NOT NULL,
  PRIMARY KEY (`student_ref`,`parent_ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `student_subject`
--

DROP TABLE IF EXISTS `student_subject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `student_subject` (
  `student_ref` int(11) NOT NULL,
  `subject_ref` int(11) NOT NULL,
  `year` int(11) DEFAULT NULL,
  `point` decimal(10,0) DEFAULT NULL,
  `grade` int(11) DEFAULT NULL,
  `term` int(11) DEFAULT NULL,
  PRIMARY KEY (`student_ref`,`subject_ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `students`
--

DROP TABLE IF EXISTS `students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `students` (
  `ref` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(50) DEFAULT NULL,
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `gender` int(11) DEFAULT NULL,
  `idcard` varchar(13) DEFAULT NULL,
  `phone` varchar(25) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `line` varchar(50) DEFAULT NULL,
  `facebook` varchar(50) DEFAULT NULL,
  `nickname` varchar(25) DEFAULT NULL,
  `grade` int(11) DEFAULT NULL,
  `room` int(11) DEFAULT NULL,
  `checkinprofile` varchar(25) DEFAULT NULL,
  `title` int(11) DEFAULT NULL,
  `rfid` varchar(25) DEFAULT NULL,
  `no` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=3087 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `study_day`
--

DROP TABLE IF EXISTS `study_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `study_day` (
  `checkin_group_ref` int(10) unsigned NOT NULL,
  `date` date NOT NULL,
  PRIMARY KEY (`checkin_group_ref`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subject`
--

DROP TABLE IF EXISTS `subject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subject` (
  `ref` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `group_ref` int(11) DEFAULT NULL,
  `code` varchar(25) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `grade_ref` int(11) DEFAULT NULL,
  `teacher_ref` int(11) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=249 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subject_group`
--

DROP TABLE IF EXISTS `subject_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subject_group` (
  `ref` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subjectresult_student`
--

DROP TABLE IF EXISTS `subjectresult_student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subjectresult_student` (
  `yearedu` int(11) NOT NULL,
  `semester` int(11) NOT NULL,
  `grade_ref` int(11) DEFAULT NULL,
  `room_ref` int(11) NOT NULL,
  `student_ref` int(11) NOT NULL,
  `subject_ref` int(11) NOT NULL,
  `point` int(11) DEFAULT 0,
  `gpa` int(11) DEFAULT 0,
  `can_exam` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`yearedu`,`semester`,`room_ref`,`student_ref`,`subject_ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `teacher`
--

DROP TABLE IF EXISTS `teacher`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teacher` (
  `ref` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `dept_ref` int(11) DEFAULT NULL,
  `room_ref` int(11) DEFAULT NULL,
  `code` varchar(10) DEFAULT NULL,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=410 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `title`
--

DROP TABLE IF EXISTS `title`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `title` (
  `ref` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `rank` int(11) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `titlestudent`
--

DROP TABLE IF EXISTS `titlestudent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `titlestudent` (
  `ref` int(11) NOT NULL AUTO_INCREMENT,
  `rank` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `ref` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(25) DEFAULT NULL,
  `line_userid` varchar(50) DEFAULT NULL,
  `idcard` varchar(20) DEFAULT NULL,
  `name` varchar(25) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`ref`)
) ENGINE=InnoDB AUTO_INCREMENT=8343 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-11-11 19:43:47
