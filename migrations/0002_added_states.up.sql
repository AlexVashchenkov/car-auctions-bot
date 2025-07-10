CREATE TABLE IF NOT EXISTS user_dialog_states -- состояния пользователей в диалогах
(
    id               SERIAL PRIMARY KEY,
    telegram_user_id BIGINT NOT NULL,
    role             TEXT   NOT NULL CHECK (role IN ('user', 'admin')),
    state            TEXT   NOT NULL,
    context          JSONB     DEFAULT '{}'::JSONB,
    updated_at       TIMESTAMP DEFAULT NOW()
);
