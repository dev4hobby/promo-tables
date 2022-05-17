CREATE TABLE `categories` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `orders` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `product_id` bigint(20) DEFAULT NULL,
  `count` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_orders_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `products` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `price` longtext COLLATE utf8mb4_unicode_ci,
  `image_url_list` json DEFAULT NULL,
  `desc_image_url_list` json DEFAULT NULL,
  `option_image_url_list` json DEFAULT NULL,
  `category_id_list` json DEFAULT NULL,
  `promo_category_id` bigint(20) DEFAULT NULL,
  `status` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT 'NOT_ON_SALE',
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `promo_categories` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `ordering` bigint(20) DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `promo_controllers` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `promo_category_id` bigint(20) DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `buy_limit_count` bigint(20) DEFAULT '1',
  `start_at` datetime(3) DEFAULT NULL,
  `end_at` datetime(3) DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT '0',
  `subscribe_condition` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `promo_purchased_logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `promo_category_id` bigint(20) DEFAULT NULL,
  `promo_controller_id` bigint(20) DEFAULT NULL,
  `count` bigint(20) DEFAULT NULL,
  `is_cancelled` tinyint(1) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(120) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nickname` longtext COLLATE utf8mb4_unicode_ci,
  `email` longtext COLLATE utf8mb4_unicode_ci,
  `password` longtext COLLATE utf8mb4_unicode_ci,
  `sex` longtext COLLATE utf8mb4_unicode_ci,
  `birth` datetime(3) DEFAULT NULL,
  `phone` longtext COLLATE utf8mb4_unicode_ci,
  `profile_image_url` longtext COLLATE utf8mb4_unicode_ci,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `is_banned` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `user_store_subscriptions` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `promo_controller_id_list` json DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;