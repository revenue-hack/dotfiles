-- TODO: ローカル用の仮データなので後で削除します
set @user_id = 1;

INSERT INTO `categories` (`id`, `name`, `display_order`, `created_by`, `updated_by`)
VALUES
    (1, '新人研修', 1, @user_id, @user_id),
    (2, 'セキュリティ', 1, @user_id, @user_id),
    (3, 'キャリア支援', 1, @user_id, @user_id),
    (4, '全社員必須', 1, @user_id, @user_id);

INSERT INTO `courses` (`id`, `name`, `description`, `content_type`, `is_required`, `category_id`, `created_by`, `updated_by`)
VALUES
    (1, 'ビジネスマナーの基本', '', 1, 1, 1, @user_id, @user_id),
    (2, 'ロジカルシンキング研修', '', 1, 1, 1, @user_id, @user_id),
    (3, '情報セキュリティ研修', '', 1, 0, 2, @user_id, @user_id),
    (4, 'インサイダー取引防止研修', '', 1, 0, 2, @user_id, @user_id);
