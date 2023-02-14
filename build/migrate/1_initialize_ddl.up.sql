CREATE TABLE `categories` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `name` VARCHAR(255) NOT NULL COMMENT 'カテゴリ名',
  `display_order` SMALLINT UNSIGNED NOT NULL COMMENT '表示順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = 'カテゴリ管理';


CREATE TABLE `courses` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `course_type` TINYINT(3) UNSIGNED NOT NULL COMMENT '研修の種別\n  1: e-Learning\n  2: 集合研修',
  `title` VARCHAR(255) NOT NULL COMMENT '名称',
  `description` TEXT NULL COMMENT '説明文',
  `thumbnail_delivery_file_name` VARCHAR(255) NULL DEFAULT NULL COMMENT 'サムネイル画像名（配信用）',
  `thumbnail_original_file_name` VARCHAR(255) NULL DEFAULT NULL COMMENT 'サムネイル画像名（元ファイル名）',
  `is_required` TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '受講必須フラグ\n  0: 任意\n  1: 必須',
  `category_id` INT UNSIGNED NULL DEFAULT NULL COMMENT 'categories.id',
  `status` TINYINT(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '講習のステータス\n  1: 非公開\n  2: 公開',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `contents_category_id_fk_idx` (`category_id` ASC) VISIBLE,
  CONSTRAINT `contents_category_id_fk`
    FOREIGN KEY (`category_id`)
    REFERENCES `categories` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '講習の基礎情報を管理';


CREATE TABLE `contents` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `course_id` INT UNSIGNED NOT NULL COMMENT 'courses.id',
  `display_order` SMALLINT UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `e_learnings_course_id_fk_idx` (`course_id` ASC) VISIBLE,
  CONSTRAINT `e_learnings_course_id_fk`
    FOREIGN KEY (`course_id`)
    REFERENCES `courses` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '講習に紐づく動画・ファイル・外部URLの枠を管理';


CREATE TABLE `movies` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `content_id` INT UNSIGNED NOT NULL COMMENT 'contents.id',
  `delivery_file_name` VARCHAR(255) NOT NULL COMMENT '配信用ファイルのファイル名',
  `original_file_name` VARCHAR(255) NOT NULL COMMENT 'アップロード時の元ファイル名',
  `thumbnail_delivery_file_name` VARCHAR(255) NOT NULL COMMENT 'サムネイル画像名（配信用）',
  `duration` INT UNSIGNED NOT NULL COMMENT '動画の再生時間（秒）',
  `convert_status` TINYINT(3) UNSIGNED NOT NULL COMMENT '動画の変換ステータス\n  1: ファイルアップロード前\n  2: 動画ファイルアップロード完了\n  3: 動画変換中\n  4: 動画変換完了\n  9: 動画変換エラー',
  `convert_error_detail` TEXT NULL COMMENT '動画変換エラーの詳細情報',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `movies_content_id_fk_idx` (`content_id` ASC) VISIBLE,
  CONSTRAINT `movies_content_id_fk`
    FOREIGN KEY (`content_id`)
    REFERENCES `contents` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '動画コンテンツ管理';


CREATE TABLE `files` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `content_id` INT UNSIGNED NOT NULL COMMENT 'contents.id',
  `delivery_file_name` VARCHAR(255) NOT NULL COMMENT '配信用ファイルのファイル名',
  `original_file_name` VARCHAR(255) NOT NULL COMMENT 'アップロード時の元ファイル名',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `files_content_id_fk_idx` (`content_id` ASC) VISIBLE,
  CONSTRAINT `files_content_id_fk`
    FOREIGN KEY (`content_id`)
    REFERENCES `contents` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = 'ファイルコンテンツ管理';


CREATE TABLE `urls` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `content_id` INT UNSIGNED NOT NULL COMMENT 'contents.id',
  `title` VARCHAR(255) NOT NULL COMMENT '動画タイトル',
  `url` TEXT NOT NULL COMMENT 'URL',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `urls_content_id_fk_idx` (`content_id` ASC) VISIBLE,
  CONSTRAINT `urls_content_id_fk`
    FOREIGN KEY (`content_id`)
    REFERENCES `contents` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '外部URLコンテンツ管理';


CREATE TABLE `course_schedules` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `course_id` INT UNSIGNED NOT NULL COMMENT 'group_trainings.id',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `course_schedules_course_id_fk_idx` (`course_id` ASC) VISIBLE,
  CONSTRAINT `course_schedules_course_id_fk`
    FOREIGN KEY (`course_id`)
    REFERENCES `courses` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '研修の日程情報を管理';


CREATE TABLE `target_members` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `course_schedule_id` INT UNSIGNED NOT NULL COMMENT 'courses.id',
  `member_id` INT UNSIGNED NOT NULL COMMENT 'kaonaviのmembers.id',
  `status` TINYINT(3) UNSIGNED NOT NULL COMMENT '受講ステータス\n  1: 未実施\n  2: 実施済み',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `target_members_course_schedule_id_fk_idx` (`course_schedule_id` ASC) VISIBLE,
  CONSTRAINT `target_members_course_schedule_id_fk`
    FOREIGN KEY (`course_schedule_id`)
    REFERENCES `course_schedules` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '講習の日程別の受講対象メンバー管理';


CREATE TABLE `target_member_movie_watch_statuses` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `movie_id` INT UNSIGNED NOT NULL COMMENT 'movies.id',
  `target_member_id` INT UNSIGNED NOT NULL COMMENT 'target_members.id',
  `status` TINYINT(3) UNSIGNED NOT NULL COMMENT '動画の視聴ステータス\n  1: 未視聴\n  2: 視聴中\n  3: 視聴完了',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `target_member_movie_watch_statuses_target_member_id_fk_idx` (`target_member_id` ASC) VISIBLE,
  INDEX `target_member_movie_watch_statuses_content_id_fk_idx` (`movie_id` ASC) VISIBLE,
  CONSTRAINT `target_member_movie_watch_statuses_movie_id_fk`
    FOREIGN KEY (`movie_id`)
    REFERENCES `movies` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `target_member_movie_watch_statuses_target_member_id_fk`
    FOREIGN KEY (`target_member_id`)
    REFERENCES `target_members` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '受講対象メンバーの動画コンテンツの視聴ステータス管理';


CREATE TABLE `target_member_file_watch_statuses` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `file_id` INT UNSIGNED NOT NULL COMMENT 'files.id',
  `target_member_id` INT UNSIGNED NOT NULL COMMENT 'target_members.id',
  `status` TINYINT(3) UNSIGNED NOT NULL COMMENT 'ファイルの視聴ステータス\n  1: 未視聴\n  2: 視聴中\n  3: 視聴完了',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `target_member_file_watch_statuses_target_member_id_fk_idx` (`target_member_id` ASC) VISIBLE,
  INDEX `target_member_file_watch_statuses_file_id_fk_idx` (`file_id` ASC) VISIBLE,
  CONSTRAINT `target_member_file_watch_statuses_file_id_fk`
    FOREIGN KEY (`file_id`)
    REFERENCES `files` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `target_member_file_watch_statuses_target_member_id_fk`
    FOREIGN KEY (`target_member_id`)
    REFERENCES `target_members` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '受講対象メンバーのファイルコンテンツの視聴ステータス管理';


CREATE TABLE `target_member_url_watch_statuses` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `url_id` INT UNSIGNED NOT NULL COMMENT 'files.id',
  `target_member_id` INT UNSIGNED NOT NULL COMMENT 'target_members.id',
  `status` TINYINT(3) UNSIGNED NOT NULL COMMENT '外部URLの視聴ステータス\n  1: 未視聴\n  2: 視聴中\n  3: 視聴完了',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `target_member_url_watch_statuses_target_member_id_fk_idx` (`target_member_id` ASC) VISIBLE,
  INDEX `target_member_url_watch_statuses_url_id_fk_idx` (`url_id` ASC) VISIBLE,
  CONSTRAINT `target_member_url_watch_statuses_url_id_fk`
    FOREIGN KEY (`url_id`)
    REFERENCES `urls` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `target_member_url_watch_statuses_target_member_id_fk`
    FOREIGN KEY (`target_member_id`)
    REFERENCES `target_members` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '受講対象メンバーの外部URLコンテンツの視聴ステータス管理';


CREATE TABLE `forms` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `course_id` INT UNSIGNED NOT NULL COMMENT 'courses.id',
  `title` VARCHAR(255) NOT NULL COMMENT 'タイトル',
  `answer_limit` TINYINT(3) UNSIGNED NULL COMMENT 'テストの回答回数上限（NULL: 制限無し）',
  `answer_time_limit` SMALLINT UNSIGNED NULL COMMENT 'テストの回答時間（NULL: 制限無し）',
  `pass_point` SMALLINT UNSIGNED NULL COMMENT 'テスト合格点（NULL: 無し）',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `forms_course_id_fk_idx` (`course_id` ASC) VISIBLE,
  CONSTRAINT `forms_course_id_fk`
    FOREIGN KEY (`course_id`)
    REFERENCES `courses` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '講習のテストフォームを管理';


CREATE TABLE `form_parts` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `form_id` INT UNSIGNED NOT NULL COMMENT 'forms.id',
  `category` VARCHAR(255) NULL COMMENT '問題のカテゴリ分けに使用する文字列（NULL: カテゴリ無し）',
  `part_type` TINYINT(3) UNSIGNED NOT NULL COMMENT 'パーツ種別\n  1: ラジオボタン\n  2: チェックボックス',
  `question` TEXT NOT NULL COMMENT '問題文',
  `description` TEXT NULL COMMENT '問題の解説文',
  `point` SMALLINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '得点',
  `is_required` TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '必須フラグ\n  0: 任意\n  1: 必須',
  `is_random` TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '選択肢の並び順をランダムにするかを管理するフラグ\n  0: 設定順\n  1: ランダム',
  `display_order` SMALLINT NOT NULL COMMENT '並び順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `form_parts_form_id_fk_idx` (`form_id` ASC) VISIBLE,
  CONSTRAINT `form_parts_form_id_fk`
    FOREIGN KEY (`form_id`)
    REFERENCES `forms` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '講習のテストフォームに配置するパーツを管理';


CREATE TABLE `form_part_choices` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `form_part_id` INT UNSIGNED NOT NULL COMMENT 'form_parts.id',
  `name` VARCHAR(255) NOT NULL COMMENT '選択肢名',
  `is_correct` TINYINT(1) UNSIGNED NOT NULL COMMENT '正解の設問かを表すフラグ\n  0: 不正解\n  1: 正解',
  `display_order` SMALLINT NOT NULL COMMENT '並び順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `form_part_choices_form_part_id_fk_idx` (`form_part_id` ASC) VISIBLE,
  CONSTRAINT `form_part_choices_form_part_id_fk`
    FOREIGN KEY (`form_part_id`)
    REFERENCES `form_parts` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '講習のテストフォームに配置する選択肢パーツの選択肢を管理';


CREATE TABLE `form_answers` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `form_id` INT UNSIGNED NOT NULL COMMENT 'forms.id',
  `target_member_id` INT UNSIGNED NOT NULL COMMENT 'target_members.id',
  `start_at` DATETIME NOT NULL COMMENT '回答開始日時',
  `end_at` DATETIME NULL COMMENT '回答終了日時（NULL: 回答中',
  `status` TINYINT(3) UNSIGNED NOT NULL COMMENT 'ステータス\n  1: 回答中\n  2: 回答終了（時間内に完了）\n  3: 回答終了（時間切れ）',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `form_answers_form_id_fk_idx` (`form_id` ASC) VISIBLE,
  INDEX `form_answers_target_member_id_fk_idx` (`target_member_id` ASC) VISIBLE,
  CONSTRAINT `form_answers_form_id_fk`
    FOREIGN KEY (`form_id`)
    REFERENCES `forms` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `form_answers_target_member_id_fk`
    FOREIGN KEY (`target_member_id`)
    REFERENCES `target_members` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '講習のテストフォームの回答状況を管理';


CREATE TABLE `form_answer_choices` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `form_answer_id` INT UNSIGNED NOT NULL COMMENT 'form_answers.id',
  `form_part_choice_id` INT UNSIGNED NOT NULL COMMENT 'form_part_choices.id',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `form_answer_choices_form_answer_id_fk_idx` (`form_answer_id` ASC) VISIBLE,
  INDEX `form_answer_choices_form_part_choice_id_fk_idx` (`form_part_choice_id` ASC) VISIBLE,
  CONSTRAINT `form_answer_choices_form_answer_id_fk`
    FOREIGN KEY (`form_answer_id`)
    REFERENCES `form_answers` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `form_answer_choices_form_part_choice_id_fk`
    FOREIGN KEY (`form_part_choice_id`)
    REFERENCES `form_part_choices` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '講習のテストフォームの選択肢パーツの選択値を管理';


CREATE TABLE `reminders` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `course_id` INT UNSIGNED NOT NULL COMMENT 'courses.id',
  `from_login_user_id` INT UNSIGNED NULL COMMENT 'kaonaviのlogin_users.id（NULLは固定の送信元アドレスを使用',
  `subject` VARCHAR(200) NOT NULL COMMENT '件名',
  `body` TEXT NOT NULL COMMENT '本文',
  `extra_address` TEXT NULL COMMENT '追加で送信したいメールアドレスをカンマ区切りで複数保存',
  `target_type` TINYINT(3) UNSIGNED NOT NULL DEFAULT 2 COMMENT '送信対象\n  1: 未実施のみ\n  2: 対象者全員',
  `is_available` TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '設定が有効かを判定するフラグ\n  0: 無効\n  1: 有効',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `reminders_course_id_fk_idx` (`course_id` ASC) VISIBLE,
  CONSTRAINT `reminders_course_id_fk`
    FOREIGN KEY (`course_id`)
    REFERENCES `courses` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = 'リマインダーの設定を管理';


CREATE TABLE `reminder_schedules` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `reminder_id` INT UNSIGNED NOT NULL COMMENT 'reminders.id',
  `send_date` DATE NOT NULL COMMENT '送信日',
  `send_time` TIME NOT NULL COMMENT '送信時間',
  `is_available` TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '設定が有効かを判定するフラグ\n  0: 無効\n  1: 有効',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `reminder_schedules_reminder_id_fk_idx` (`reminder_id` ASC) VISIBLE,
  CONSTRAINT `reminder_schedules_reminder_id_fk`
    FOREIGN KEY (`reminder_id`)
    REFERENCES `reminders` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = 'リマインダーの送信スケジュールを管理';


CREATE TABLE `reminder_schedule_send_histories` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `reminder_schedule_id` INT UNSIGNED NOT NULL COMMENT 'reminders.id',
  `mail_group_id` INT UNSIGNED NOT NULL COMMENT 'kaonaviのmail_groups.id',
  `send_at` DATETIME NOT NULL COMMENT '送信日時',
  `status` TINYINT(3) UNSIGNED NOT NULL COMMENT 'ステータス\n  1: 送信完了\n  2: 送信エラー\n',
  `error_detail` TEXT NULL COMMENT '送信エラー時のエラー詳細',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `reminder_schedule_send_histories_reminder_schedule_id_fk_idx` (`reminder_schedule_id` ASC) VISIBLE,
  CONSTRAINT `reminder_schedule_send_histories_reminder_schedule_id_fk`
    FOREIGN KEY (`reminder_schedule_id`)
    REFERENCES `reminder_schedules` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = 'リマインダーメールの送信履歴を管理';


CREATE TABLE `e_learning_schedules` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `course_schedule_id` INT UNSIGNED NOT NULL COMMENT 'group_trainings.id',
  `from` DATETIME NULL DEFAULT NULL COMMENT '受講期間（開始）',
  `to` DATETIME NULL DEFAULT NULL COMMENT '受講期間（終了）',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `e_learning_schedules_course_schedule_id_fk_idx` (`course_schedule_id` ASC) VISIBLE,
  CONSTRAINT `e_learning_schedules_course_schedule_id_fk`
    FOREIGN KEY (`course_schedule_id`)
    REFERENCES `course_schedules` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = 'e-Learningの日程情報を管理';


CREATE TABLE `group_training_schedules` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `course_schedule_id` INT UNSIGNED NOT NULL COMMENT 'group_trainings.id',
  `from` DATETIME NOT NULL COMMENT '受講期間（開始）',
  `to` DATETIME NOT NULL COMMENT '受講期間（終了）',
  `capacity` SMALLINT UNSIGNED NULL COMMENT '定員（NULL: 定員なし）',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `group_training_schedules_course_schedule_id_fk_idx` (`course_schedule_id` ASC) VISIBLE,
  CONSTRAINT `group_training_schedules_course_schedule_id_fk`
    FOREIGN KEY (`course_schedule_id`)
    REFERENCES `course_schedules` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '集合研修の日程情報を管理';


CREATE TABLE `entry_target_members` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `course_id` INT UNSIGNED NOT NULL COMMENT 'group_trainings.id',
  `member_id` INT UNSIGNED NOT NULL COMMENT 'kaonaviのmembers.id',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `created_by` INT UNSIGNED NOT NULL COMMENT '作成者ID（login_users.id）',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `updated_by` INT UNSIGNED NOT NULL COMMENT '更新者ID（login_users.id）',
  PRIMARY KEY (`id`),
  INDEX `entry_target_members_course_id_fk_idx` (`course_id` ASC) VISIBLE,
  CONSTRAINT `entry_target_members_course_id_fk`
    FOREIGN KEY (`course_id`)
    REFERENCES `courses` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_ja_0900_as_cs_ks
COMMENT = '講習に申込可能なメンバーを管理（集合研修でのみ使用）';
