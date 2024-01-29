-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 30, 2024 at 12:00 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `campaignweb`
--

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `short_description` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `perks` text DEFAULT NULL,
  `backer_count` int(11) DEFAULT NULL,
  `goal_amount` int(11) DEFAULT NULL,
  `current_amount` int(11) DEFAULT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'A Campaign for Adventist Church', 'Adventist Church', 'Adventist Church is Adventist Protestant Christian', 'this is for new church, and for the renew church', 0, 1000000000, 0, 'campaign-satu', '2024-01-24 11:49:01', '2024-01-25 16:42:47'),
(2, 1, 'Penggalangan', 'short', 'Long', 'hadiah satu, dua, dan tiga', 0, 10101010, 0, 'penggalangan-s-int-1', '2024-01-24 13:24:30', '2024-01-24 13:24:30'),
(3, 1, 'A Campaign for Adventist Church', 'Adventist Church', 'Adventist Church ist afsfefwdnni', 'this is for new church, and for the renew church', 0, 1000000000, 0, 'a-campaign-for-adventist-church-1', '2024-01-24 13:44:27', '2024-01-24 13:44:27'),
(4, 1, 'A Campaign for Adventist Church', 'Adventist Church', 'Adventist Church ist afsfefwdnni', 'this is for new church, and for the renew church', 0, 1000000000, 0, 'a-campaign-for-adventist-church-1', '2024-01-30 05:51:51', '2024-01-30 05:51:51');

-- --------------------------------------------------------

--
-- Table structure for table `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `is_primary` tinyint(4) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(1, 1, 'campaign/images-satu', 0, '2024-01-24 11:51:43', '2024-01-26 14:34:19'),
(2, 1, 'campaign/images-dua', 0, '2024-01-24 11:51:43', '2024-01-26 14:34:19'),
(4, 1, 'images/1-3.2.PNG', 0, '2024-01-26 14:31:55', '2024-01-26 14:34:19'),
(5, 1, 'images/1-3.2.PNG', 0, '2024-01-26 14:32:14', '2024-01-26 14:34:19'),
(6, 1, 'images/1-3.1.PNG', 1, '2024-01-26 14:34:19', '2024-01-26 14:34:19');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `payment_url` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `payment_url`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 500000, 'Hai', NULL, '', '2024-01-28 14:28:15', '2024-01-28 14:28:15'),
(3, 1, 1, 50000, 'Hai Campaign 2', NULL, '', '2024-01-28 14:28:15', '2024-01-28 14:28:15'),
(4, 9, 1, 5000000, 'pending', '', '', '2024-01-29 10:04:15', '2024-01-29 10:04:15'),
(5, 9, 1, 5000000, 'pending', '', '', '2024-01-29 10:12:19', '2024-01-29 10:12:19'),
(6, 9, 1, 5000000, 'pending', '', '', '2024-01-29 10:16:47', '2024-01-29 10:16:47'),
(7, 10, 1, 5000000, 'pending', '', '', '2024-01-29 10:17:35', '2024-01-29 10:17:35'),
(8, 1, 1, 5000000, 'pending', '', 'https://app.sandbox.midtrans.com/snap/v3/redirection/c84dec7c-06f4-4026-ae12-6974dfda2264', '2024-01-29 13:59:40', '2024-01-29 13:59:41'),
(9, 1, 1, 5000000, 'pending', '', 'https://app.sandbox.midtrans.com/snap/v3/redirection/69db7f0c-3cb2-4d17-ac62-d1292fdc38b3', '2024-01-29 14:08:46', '2024-01-29 14:08:47'),
(10, 1, 1, 5000000, 'pending', '', 'https://app.sandbox.midtrans.com/snap/v3/redirection/aab63aff-d22c-46e0-bf91-17e2ffdb229d', '2024-01-30 05:42:47', '2024-01-30 05:42:49');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `occupation` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password_hash` varchar(255) DEFAULT NULL,
  `avatar_file_name` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `token`, `created_at`, `updated_at`) VALUES
(1, 'Paul', 'Machine Learning Engineer', 'paul1@gmail.com', '$2a$04$I0lF6PBBmtMS47pltoBQi.p9R6UWf1xu3WEdCvHO9HQFi27pKjORC', 'images/1-3.2.PNG', 'user', NULL, '0000-00-00 00:00:00', '0000-00-00 00:00:00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
