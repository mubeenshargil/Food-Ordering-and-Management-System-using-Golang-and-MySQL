-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3307
-- Generation Time: Jun 23, 2026 at 01:41 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `food_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `menu`
--

CREATE TABLE `menu` (
  `ID` int(11) NOT NULL,
  `Name` varchar(255) NOT NULL,
  `Description` varchar(255) NOT NULL,
  `Price` float NOT NULL,
  `Image` varchar(255) NOT NULL,
  `Status` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `menu`
--

INSERT INTO `menu` (`ID`, `Name`, `Description`, `Price`, `Image`, `Status`) VALUES
(1, 'Margherita Pizza', 'Classic Italian pizza with tomato sauce, mozzarella cheese, and fresh basil.', 8.99, 'https://images.unsplash.com/photo-1604382355076-af4b0eb60143', 1),
(2, 'Pepperoni Pizza', 'Wood-fired pizza topped with spicy pepperoni and mozzarella cheese.', 10.5, 'https://images.unsplash.com/photo-1513104890138-7c749659a591', 1),
(3, 'Quattro Formaggi', 'Traditional Italian four cheese pizza with mozzarella, parmesan, gorgonzola, and ricotta.', 11.99, 'https://images.unsplash.com/photo-1594007654729-407eedc4be65', 1),
(4, 'Pasta Carbonara', 'Creamy Roman pasta with egg, parmesan cheese, pancetta, and black pepper.', 12.5, 'https://images.unsplash.com/photo-1612874742237-6526221588e3', 1),
(5, 'Lasagna', 'Layers of pasta with beef ragu, béchamel sauce, and mozzarella.', 13.99, 'https://images.unsplash.com/photo-1619895092538-128341789043', 1),
(7, 'Spaghetti Bolognese', 'Italian spaghetti served with rich meat sauce and parmesan cheese.', 10.99, 'https://images.unsplash.com/photo-1621996346565-e3dbc646d9a9', 1),
(8, 'Fettuccine Alfredo', 'Creamy Alfredo pasta with butter, parmesan, and Italian herbs.', 11.25, 'https://images.unsplash.com/photo-1645112411341-6c4fd023714a', 1),
(9, 'Bruschetta', 'Toasted Italian bread topped with tomatoes, garlic, basil, and olive oil.', 6.99, 'https://images.unsplash.com/photo-1506280754576-f6fa8a873550', 1),
(10, 'Caprese Salad', 'Fresh mozzarella, tomatoes, basil, and balsamic glaze.', 7.5, 'https://images.unsplash.com/photo-1566843972142-a7fcb70de55a', 1),
(11, 'Tiramisu', 'Classic Italian dessert made with mascarpone, coffee, and cocoa.', 5.99, 'https://images.unsplash.com/photo-1571877227200-a0d98ea607e9', 1),
(12, 'Panna Cotta', 'Silky Italian vanilla cream dessert with berry sauce.', 5.5, 'https://images.unsplash.com/photo-1488477181946-6428a0291777', 1),
(13, 'Italian Seafood Pasta', 'Pasta with shrimp, mussels, calamari, and tomato garlic sauce.', 15.99, 'https://images.unsplash.com/photo-1551183053-bf91a1d81141', 1),
(14, 'Gnocchi', 'Soft potato dumplings served with creamy tomato basil sauce.', 10.25, 'https://images.unsplash.com/photo-1516100882582-96c3a05fe590', 1),
(15, 'Calzone', 'Folded Italian pizza stuffed with cheese, ham, and tomato sauce.', 11.75, 'https://images.unsplash.com/photo-1541745537411-b8046dc6d66c', 1),
(16, 'Minestrone Soup', 'Traditional Italian vegetable soup with beans and pasta.', 6.5, 'https://images.unsplash.com/photo-1547592166-23ac45744acd', 1),
(17, 'Ravioli Ricotta', 'Fresh ravioli stuffed with ricotta cheese and spinach.', 12.25, 'https://images.unsplash.com/photo-1608219992759-8d74ed8d76eb', 1),
(18, 'Italian Meatballs', 'Beef meatballs in rich tomato sauce with parmesan.', 9.99, 'https://images.unsplash.com/photo-1529042410759-befb1204b468', 1),
(19, 'Pesto Pasta', 'Fresh basil pesto pasta with parmesan and pine nuts.', 10.75, 'https://images.unsplash.com/photo-1473093295043-cdd812d0e601', 1),
(20, 'Cannoli', 'Crispy Sicilian pastry filled with sweet ricotta cream.', 4.99, 'https://images.unsplash.com/photo-1551024601-bec78aea704b', 1),
(28, 'Lasagna new one', 'testing purpose - updated', 13.99, '/static/uploads/images.jfif', 1);

-- --------------------------------------------------------

--
-- Table structure for table `orderhistory`
--

CREATE TABLE `orderhistory` (
  `id` int(11) NOT NULL,
  `orderId` int(23) NOT NULL,
  `updatedBy` int(23) DEFAULT NULL,
  `updatedAt` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
  `status` varchar(255) NOT NULL,
  `reason` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `orderhistory`
--

INSERT INTO `orderhistory` (`id`, `orderId`, `updatedBy`, `updatedAt`, `status`, `reason`) VALUES
(1, 1, NULL, '2026-06-21 21:56:18.410859', 'PENDING', NULL),
(2, 2, NULL, '2026-06-21 21:56:22.853073', 'PENDING', NULL),
(3, 3, NULL, '2026-06-21 21:56:25.440109', 'PENDING', NULL),
(4, 4, NULL, '2026-06-21 21:56:34.860057', 'PENDING', NULL),
(5, 1, 2, '2026-06-21 21:56:28.464868', 'ACCEPTED', NULL),
(6, 1, 5, '2026-06-21 22:51:11.889799', 'PREPARING', NULL),
(8, 1, 5, '2026-06-21 22:51:45.320552', 'ON THE WAY', NULL),
(9, 1, 5, '2026-06-21 21:56:37.611235', 'DELIVERED', NULL),
(10, 1, 5, '2026-06-21 21:56:42.604928', 'CLOSED', NULL),
(12, 4, 5, '2026-06-21 18:54:23.175753', 'CANCELLED', 'Order was cancelled due to an incorrect address.'),
(15, 2, 5, '2026-06-21 19:02:09.854958', 'CANCELLED', '          Order was cancelled due to an incorrect Phone Number.                          '),
(16, 5, NULL, '2026-06-22 00:59:47.871397', 'PENDING', NULL),
(17, 6, NULL, '2026-06-22 01:10:28.123456', 'PENDING', NULL),
(18, 7, NULL, '2026-06-22 01:11:01.732363', 'PENDING', NULL),
(19, 8, NULL, '2026-06-22 01:12:15.571224', 'PENDING', NULL),
(20, 9, NULL, '2026-06-22 01:13:24.914158', 'PENDING', NULL),
(21, 3, 2, '2026-06-22 14:54:05.414803', 'ACCEPTED', ''),
(22, 3, 2, '2026-06-22 14:54:21.200762', 'PREPARING', ''),
(23, 3, 2, '2026-06-22 14:54:26.422840', 'READY', ''),
(24, 3, 2, '2026-06-22 14:54:29.239958', 'ON THE WAY', ''),
(25, 3, 2, '2026-06-22 14:54:32.690722', 'DELIVERED', ''),
(26, 3, 2, '2026-06-22 14:54:36.121722', 'CLOSED', ''),
(27, 10, NULL, '2026-06-22 19:20:15.693952', 'PENDING', NULL),
(28, 5, 2, '2026-06-22 19:22:30.381571', 'ACCEPTED', '');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `ID` int(11) NOT NULL,
  `UserID` int(24) NOT NULL,
  `MenuID` int(24) NOT NULL,
  `Quantity` int(24) NOT NULL,
  `price` float NOT NULL,
  `total` float NOT NULL,
  `Status` varchar(255) NOT NULL,
  `customer_name` varchar(255) NOT NULL,
  `customer_surname` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `orderTime` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
  `reason` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`ID`, `UserID`, `MenuID`, `Quantity`, `price`, `total`, `Status`, `customer_name`, `customer_surname`, `phone`, `address`, `email`, `orderTime`, `reason`) VALUES
(1, 0, 2, 2, 10.5, 21, 'CLOSED', 'Hiroko Hart', 'Carson', '+1 (546) 315-7652', 'Nobis perferendis de', 'ryses@mailinator.com', '2026-06-21 18:29:09.072762', NULL),
(2, 0, 3, 1, 11.99, 11.99, 'CANCELLED', 'Malcolm Holder', 'Cooke', '+1 (763) 672-7313', 'Temporibus est tenet', 'syxopaveko@mailinator.com', '2026-06-21 19:02:09.847410', '          Order was cancelled due to an incorrect Phone Number.                          '),
(3, 0, 1, 1, 8.99, 8.99, 'CLOSED', 'MacKenzie Obrien', 'Potter', '+1 (475) 816-7804', 'Officia eu sunt opt', 'xitiv@mailinator.com', '2026-06-22 14:54:36.119907', ''),
(4, 0, 5, 2, 13.99, 27.98, 'CANCELLED', 'Delilah Whitaker', 'Christian', '+1 (114) 787-2035', 'Quia sunt dolorem as', 'mavyj@mailinator.com', '2026-06-21 18:55:59.561598', 'Order was cancelled due to an incorrect address.'),
(5, 0, 2, 1, 10.5, 10.5, 'ACCEPTED', 'Portia Day', 'Rodgers', '+1 (986) 188-4902', 'Dolor vero reprehend', 'fygihe@mailinator.com', '2026-06-22 19:22:30.371623', ''),
(8, 0, 2, 3, 10.5, 31.5, 'PENDING', 'Amanda Edwards', 'Trujillo', '+1 (568) 193-9121', 'Nulla commodi enim i', 'pylykate@mailinator.com', '2026-06-22 01:12:15.568458', NULL),
(9, 0, 10, 1, 7.5, 7.5, 'PENDING', 'Quinn Larson', 'Nichols', '+1 (531) 368-5208', 'Mollitia sed dolor t', 'qesow@mailinator.com', '2026-06-22 01:13:24.904609', NULL),
(10, 0, 15, 1, 11.75, 11.75, 'PENDING', 'Kadeem Kline', 'Joyner', '+1 (445) 665-4859', 'Vero blanditiis temp', 'xizakupovy@mailinator.com', '2026-06-22 19:20:15.691156', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `ID` int(11) NOT NULL,
  `Name` varchar(255) NOT NULL,
  `Surname` varchar(255) NOT NULL,
  `Phone` varchar(255) NOT NULL,
  `Email` varchar(255) NOT NULL,
  `PlainPassword` varchar(255) NOT NULL,
  `Password` varchar(300) NOT NULL,
  `Role` varchar(255) NOT NULL,
  `status` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`ID`, `Name`, `Surname`, `Phone`, `Email`, `PlainPassword`, `Password`, `Role`, `status`) VALUES
(2, 'Mubeen', 'GIl', '+3999999999999', 'mubeen.shargil@gmail.com', '123456789', '$2a$10$frctKZaXrvHtq7q.lgYgeeWK7zH1rCPw8hDJr/aFq8YmAlQ.TaEnS', 'admin', 1),
(5, 'Cashier', 'David', '3929531094', 'cashier@gmail.com', '123456789', '$2a$10$mKCMGXBvhen3VIrrw5qXXuom2bbKBSFwaoNRcYHf51NIrvh.oVvli', 'cashier', 1),
(8, 'xxx', 'mick', '3929531098', 'xxx.mick@gmail.com', '123456789', '$2a$10$vPNuQGQodW2nuidxnESyx.P4IZCjdNFIbNssBwn3PgBKYHTfb6k8y', 'customer', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `menu`
--
ALTER TABLE `menu`
  ADD PRIMARY KEY (`ID`);

--
-- Indexes for table `orderhistory`
--
ALTER TABLE `orderhistory`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`ID`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`ID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `menu`
--
ALTER TABLE `menu`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT for table `orderhistory`
--
ALTER TABLE `orderhistory`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
