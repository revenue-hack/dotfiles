-- TODO: ローカル用の仮データなので後で削除します
set @user_id = 1;

INSERT INTO `categories` (`id`, `name`, `display_order`, `created_by`, `updated_by`)
VALUES
    (1, '新人研修', 1, @user_id, @user_id),
    (2, 'セキュリティ', 2, @user_id, @user_id),
    (3, 'キャリア支援', 3, @user_id, @user_id),
    (4, '全社員必須', 4, @user_id, @user_id);

INSERT INTO `courses` (`id`, `course_type`, `title`, `description`, `is_required`, `category_id`, `status`, `created_by`, `updated_by`)
VALUES
    (1, 1, 'ビジネスマナーの基本', '', 1, 1, 2, @user_id, @user_id),
    (2, 1, 'ロジカルシンキング研修', '', 1, 1, 2, @user_id, @user_id),
    (3, 1, '情報セキュリティ研修', '', 0, 2, 2, @user_id, @user_id),
    (4, 1, 'インサイダー取引防止研修', '', 0, 2, 1, @user_id, @user_id);
