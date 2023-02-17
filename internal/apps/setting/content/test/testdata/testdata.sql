-- 講習の検索APIの検証用のテストデータ
SET @at = '2023-02-09 10:00:00';

INSERT INTO `courses` (`id`, `course_type`, `title`, `description`, `thumbnail_delivery_file_name`, `thumbnail_original_file_name`, `is_required`, `category_id`, `status`, `created_at`, `created_by`, `updated_at`, `updated_by`)
VALUES
    (1, 1, 'e-Learning1', null, null, null, 1, null, 1, @at, 1, @at, 1),
    (2, 1, 'e-Learning2', null, null, null, 1, null, 1, @at, 1, @at, 1);

INSERT INTO `contents` (`id`, `course_id`, `content_type`, `display_order`, `created_at`, `created_by`, `updated_at`, `updated_by`)
VALUES
    (11, 1, 1, 1, @at, 1, @at, 1),
    (12, 1, 1, 3, @at, 1, @at, 1),
    (13, 1, 1, 2, @at, 1, @at, 1),
    (21, 1, 2, 4, @at, 1, @at, 1),
    (22, 1, 2, 6, @at, 1, @at, 1),
    (23, 1, 2, 5, @at, 1, @at, 1),
    (31, 1, 3, 7, @at, 1, @at, 1),
    (32, 1, 3, 9, @at, 1, @at, 1),
    (33, 1, 3, 8, @at, 1, @at, 1);

INSERT INTO `movies` (`id`, `content_id`, `delivery_file_name`, `original_file_name`, `thumbnail_delivery_file_name`, `duration`, `convert_status`, `convert_error_detail`, `created_at`, `created_by`, `updated_at`, `updated_by`)
VALUES
    (111, 11, 'movie_delivery1.m3u8', 'movie_original1.mp4', 'movie_thumbnail1.png', 30, 1, null, @at, 1, @at, 1),
    (112, 12, 'movie_delivery2.m3u8', 'movie_original2.mp4', 'movie_thumbnail2.png', 60, 2, null, @at, 1, @at, 1),
    (113, 13, 'movie_delivery3.m3u8', 'movie_original3.mp4', 'movie_thumbnail3.png', 15, 3, null, @at, 1, @at, 1);

INSERT INTO `files` (`id`, `content_id`, `delivery_file_name`, `original_file_name`, `created_at`, `created_by`, `updated_at`, `updated_by`)
VALUES
    (221, 21, 'file_delivery1.pptx', 'file_original1.pptx', @at, 1, @at, 1),
    (222, 22, 'file_delivery2.pptx', 'file_original2.pptx', @at, 1, @at, 1),
    (223, 23, 'file_delivery3.pptx', 'file_original3.pptx', @at, 1, @at, 1);

INSERT INTO `urls` (`id`, `content_id`, `title`, `url`, `created_at`, `created_by`, `updated_at`, `updated_by`)
VALUES
    (331, 31, 'kaonavi Tech Talk #1', 'https://www.youtube.com/watch?v=HIC4AynFDdw', @at, 1, @at, 1),
    (332, 32, 'kaonavi Tech Talk #2', 'https://www.youtube.com/watch?v=3Cs-PVZXsyU', @at, 1, @at, 1),
    (333, 33, 'kaonavi Tech Talk #11', 'https://www.youtube.com/watch?v=HPgx7r_I-Ko', @at, 1, @at, 1);
