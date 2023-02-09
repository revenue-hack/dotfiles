-- 講習の検索APIの検証用のテストデータ

INSERT INTO `courses` (`id`, `course_type`, `title`, `description`, `is_required`, `category_id`, `status`, `created_by`, `updated_by`)
VALUES
    (1, 1, 'e-Learning 1', '', 0, 1,    1, 1, 1),
    (2, 1, 'e-Learning 2', '', 0, 1,    2, 1, 1),
    (3, 1, 'e-Learning 3', '', 0, null, 1, 1, 1),
    (4, 1, 'e-Learning 4', '', 0, null, 2, 1, 1),
    (5, 1, 'e-Learning 5', '', 1, 2,    1, 1, 1),
    (6, 1, 'e-Learning 6', '', 1, 2,    2, 1, 1),
    (7, 1, 'e-Learning 7', '', 1, null, 1, 1, 1),
    (8, 1, 'e-Learning 8', '', 1, null, 2, 1, 1);

INSERT INTO `course_schedules` (`id`, `course_id`, `created_by`, `updated_by`)
VALUES
    (10, 1, 1, 1),
    (20, 2, 1, 1),
    (30, 3, 1, 1),
    (40, 4, 1, 1),
    (50, 5, 1, 1),
    (60, 6, 1, 1),
    (70, 7, 1, 1),
    (80, 8, 1, 1);

INSERT INTO `e_learning_schedules` (`id`, `course_schedule_id`, `from`, `to`, `created_by`, `updated_by`)
VALUES
    (100, 10, '2023-02-01 10:00:00', '2023-03-01 18:30:00', 1, 1),
    (200, 20, '2023-02-02 10:00:00', '2023-03-02 18:30:00', 1, 1),
    (300, 30, null, null, 1, 1),
    (400, 40, null, null, 1, 1),
    (500, 50, '2023-02-05 10:00:00', '2023-03-05 18:30:00', 1, 1),
    (600, 60, '2023-02-06 10:00:00', '2023-03-06 18:30:00', 1, 1),
    (700, 70, null, null, 1, 1),
    (800, 80, null, null, 1, 1);
