-- phpMyAdmin SQL Dump
-- version 4.6.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 2016-10-26 19:57:06
-- 服务器版本： 10.0.25-MariaDB
-- PHP Version: 7.0.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sfs`
--

-- --------------------------------------------------------

--
-- 表的结构 `point`
--

CREATE TABLE `point` (
	  `id` int(11) NOT NULL,
	  `name` varchar(45) DEFAULT NULL,
	  `status` int(11) DEFAULT NULL,
	  `type` int(11) DEFAULT NULL,
	  `pre_hours` float DEFAULT NULL,
	  `pre_stars` int(11) DEFAULT NULL,
	  `hours` int(11) DEFAULT NULL,
	  `stars` int(11) DEFAULT NULL,
	  `points` float NOT NULL DEFAULT '0',
	  `start_time` varchar(45) DEFAULT NULL,
	  `end_time` varchar(45) DEFAULT NULL,
	  `create_user_id` int(11) DEFAULT NULL,
	  `create_time` varchar(45) DEFAULT NULL,
	  `update_user_id` int(11) DEFAULT NULL,
	  `update_time` varchar(45) DEFAULT NULL,
	  `remark` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `point`
--

INSERT INTO `point` (`id`, `name`, `status`, `type`, `pre_hours`, `pre_stars`, `hours`, `stars`, `points`, `start_time`, `end_time`, `create_user_id`, `create_time`, `update_user_id`, `update_time`, `remark`) VALUES
(6, '111', 10, 10, 0, 0, 1, 2, 2, '2016-10-17', '2016-10-23 23:59:59', 0, '2016-10-26 07:52:58', 0, '2016-10-26 07:53:25', ''),
(7, '333', 10, 10, 0, 0, 1, 3, 3, '2016-10-17', '2016-10-23 23:59:59', 0, '2016-10-26 07:52:58', 0, '2016-10-26 07:53:25', ''),
(8, '222', 10, 10, 0, 0, 23, 1, 23, '2016-10-24', '2016-10-30 23:59:59', 0, '2016-10-26 07:52:58', 0, '2016-10-26 07:53:25', ''),
(9, '44', 10, 20, 0, 0, 23, 3, 69, '2016-10-24', '2016-10-30 23:59:59', 0, '2016-10-26 07:52:58', 0, '2016-10-26 07:53:25', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `point`
--
ALTER TABLE `point`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `point`
--
ALTER TABLE `point`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

