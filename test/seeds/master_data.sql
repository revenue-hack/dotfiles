-- 共通で使用するマスターデータです

INSERT INTO `categories` (`id`, `name`, `display_order`, `created_by`, `updated_by`)
VALUES
    (1, '新人研修', 1, 1, 1),
    (2, 'セキュリティ', 2, 1, 1),
    (3, 'キャリア支援', 3, 1, 1),
    (4, '全社員必須', 4, 1, 1),
    (5, '勉強会', 5, 1, 1);
