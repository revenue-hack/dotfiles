-- 講習の検索APIの検証用のテストデータ
SET @createdAt = '2023-02-09 10:00:00';

INSERT INTO `courses` (`id`, `course_type`, `title`, `description`, `thumbnail_delivery_file_name`, `thumbnail_original_file_name`, `is_required`, `category_id`, `status`, `created_by`, `created_at`, `updated_by`)
VALUES
    (1, 1, 'e-Learning 1', 'e-Learningの説明1', 'delivery1.png', 'original1.png', 0, 1, 1, 1, @createdAt, 1),
    (2, 1, 'e-Learning 2', '', null, null, 1, 2, 2, 1, @createdAt, 1),
    (3, 2, '集合研修 1', '集合研修の説明1', 'delivery2.png', 'original2.png', 0, 1, 1, 1, @createdAt, 1),
    (4, 2, '集合研修 2', '', null, null, 1, 2, 2, 1, @createdAt, 1);

INSERT INTO `course_schedules` (`id`, `course_id`, `created_by`, `created_at`, `updated_by`)
VALUES
    (10, 1, 1, @createdAt, 1),
    (30, 3, 1, @createdAt, 1);

INSERT INTO `e_learning_schedules` (`id`, `course_schedule_id`, `from`, `to`, `created_by`, `created_at`, `updated_by`)
VALUES
    (100, 10, '2023-02-01 10:00:00', '2023-03-01 18:30:00', 1, @createdAt, 1),
    (300, 30, '2023-02-03 10:00:00', '2023-03-03 18:30:00', 1, @createdAt, 1);
