-- 講習の検索APIの検証用のテストデータ

INSERT INTO `courses` (`id`, `course_type`, `title`, `description`, `is_required`, `category_id`, `status`, `created_by`, `updated_by`)
VALUES
    (1, 1, 'e-Learning 1', '', 0, 1, 1, 1, 1),
    (2, 1, 'e-Learning 2', '', 1, 2, 2, 1, 1),
    (3, 2, '集合研修 1', '', 0, 1, 1, 1, 1),
    (4, 2, '集合研修 2', '', 1, 2, 2, 1, 1);

INSERT INTO `course_schedules` (`id`, `course_id`, `created_by`, `updated_by`)
VALUES
    (10, 1, 1, 1),
    (30, 3, 1, 1);

INSERT INTO `e_learning_schedules` (`id`, `course_schedule_id`, `from`, `to`, `created_by`, `updated_by`)
VALUES
    (100, 10, '2023-02-01 10:00:00', '2023-03-01 18:30:00', 1, 1),
    (300, 30, '2023-02-03 10:00:00', '2023-03-03 18:30:00', 1, 1);
