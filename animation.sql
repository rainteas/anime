-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2024-01-18 15:07:08
-- 服务器版本： 5.7.44-log
-- PHP 版本： 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `animation`
--

-- --------------------------------------------------------

--
-- 表的结构 `items`
--

CREATE TABLE `items` (
  `id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `link` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `pub_date` varchar(255) DEFAULT NULL,
  `download` int(11) DEFAULT NULL,
  `torrent_hash_string` varchar(255) NOT NULL,
  `torrent_id` int(11) NOT NULL,
  `torrent_name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `rss_meta`
--

CREATE TABLE `rss_meta` (
  `id` int(11) NOT NULL,
  `anime_name` varchar(255) DEFAULT NULL,
  `season` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `rss_meta`
--

INSERT INTO `rss_meta` (`id`, `anime_name`, `season`, `url`) VALUES
(1, '盾之勇者成名录 第三季', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3173&subgroupid=583'),
(2, '家里蹲吸血姬的苦闷', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3153&subgroupid=583'),
(3, '某大叔的VRMMO活动记', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3182&subgroupid=583'),
(4, '葬送的芙莉莲', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3141&subgroupid=583'),
(5, '哥布林杀手 第二季', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3164&subgroupid=583'),
(6, '因想当冒险者而前往大都市的女儿已经升到了S级', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3140&subgroupid=583'),
(7, '圣剑学院的魔剑使', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3172&subgroupid=583'),
(8, '16bit的感动 ANOTHER LAYER', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3178&subgroupid=583'),
(9, '狩龙人拉格纳', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3193&subgroupid=583'),
(10, '归还者的魔法要特别', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3176&subgroupid=583'),
(11, '尸体如山的死亡游戏', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=2988&subgroupid=583'),
(14, '捡走被人悔婚的千金，教会她坏坏的幸福生活', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3199&subgroupid=583'),
(15, '女友成双 第二季', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3206&subgroupid=583'),
(16, '圣女的魔力是万能的 第二季', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3181&subgroupid=583'),
(17, '东京复仇者 天竺篇', '2023 年秋季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3145&subgroupid=597'),
(18, '欢迎来到实力至上主义教室 第三季', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3260&subgroupid=583'),
(19, '秒杀外挂太强了，异世界的家伙们根本就不是对手。', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3253&subgroupid=583'),
(20, '魔都精兵的奴隶', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3231&subgroupid=583'),
(21, '弱角友崎同学 第二季', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3223&subgroupid=583'),
(22, '为了在异世界也能抚摸毛茸茸而努力', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3266&subgroupid=615'),
(23, '治愈魔法的错误使用方法', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3222&subgroupid=583'),
(24, '佐佐木与文鸟小哔', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3235&subgroupid=583'),
(25, '我独自升级', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3247&subgroupid=583'),
(26, '异修罗 ', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3261&subgroupid=370'),
(27, '百千家的妖怪王子', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3252&subgroupid=583'),
(28, '事与愿违的不死冒险者', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3221&subgroupid=583'),
(29, '至高之牌 第二季', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3224&subgroupid=583'),
(32, '反派大小姐等级99～我是隐藏BOSS但不是魔王～', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3255&subgroupid=583'),
(33, '轮回七次的反派大小姐，在前敌国享受随心所欲的新婚生活 ', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3225&subgroupid=583'),
(34, '最强肉盾的迷宫攻略～拥有稀少技能体力9999的肉盾，被勇者队伍辞退了～', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3227&subgroupid=583'),
(36, '最弱的驯养师开启的捡垃圾的旅途。', '2024 年冬季番组', 'https://mikanani.me/RSS/Bangumi?bangumiId=3251&subgroupid=583');

--
-- 转储表的索引
--

--
-- 表的索引 `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_title` (`title`);

--
-- 表的索引 `rss_meta`
--
ALTER TABLE `rss_meta`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_name` (`anime_name`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `items`
--
ALTER TABLE `items`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `rss_meta`
--
ALTER TABLE `rss_meta`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
