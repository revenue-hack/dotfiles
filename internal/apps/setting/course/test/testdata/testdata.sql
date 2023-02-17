-- 講習の検索APIの検証用のテストデータ
SET @at = '2023-02-09 10:00:00';

INSERT INTO `categories` (`id`, `name`, `display_order`, `created_by`, `updated_by`)
VALUES
    (1, '新人研修', 1, 1, 1),
    (2, 'セキュリティ', 2, 1, 1),
    (3, 'キャリア支援', 3, 1, 1),
    (4, '全社員必須', 4, 1, 1),
    (5, '勉強会', 5, 1, 1);

INSERT INTO `courses` (`id`, `course_type`, `title`, `description`, `thumbnail_delivery_file_name`, `thumbnail_original_file_name`, `is_required`, `category_id`, `status`, `created_by`, `created_at`, `updated_by`, `updated_at`)
VALUES
    (1, 1, 'e-Learning 1', 'e-Learningの説明1', 'delivery1.png', 'original1.png', 0, 1, 1, 1, @at, 1, @at),
    (2, 1, 'e-Learning 2', '', null, null, 1, 2, 2, 1, @at, 1, @at),
    (3, 2, '集合研修 1', '集合研修の説明1', 'delivery2.png', 'original2.png', 0, 1, 1, 1, @at, 1, @at),
    (4, 2, '集合研修 2', '', null, null, 1, 2, 2, 1, @at, 1, @at);

INSERT INTO `course_schedules` (`id`, `course_id`, `created_by`, `created_at`, `updated_by`, `updated_at`)
VALUES
    (10, 1, 1, @at, 1, @at),
    (30, 3, 1, @at, 1, @at);

INSERT INTO `e_learning_schedules` (`id`, `course_schedule_id`, `from`, `to`, `created_by`, `created_at`, `updated_by`, `updated_at`)
VALUES
    (100, 10, '2023-02-01 10:00:00', '2023-03-01 18:30:00', 1, @at, 1, @at),
    (300, 30, '2023-02-03 10:00:00', '2023-03-03 18:30:00', 1, @at, 1, @at);
