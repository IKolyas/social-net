CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    second_name VARCHAR(255) NOT NULL,
    birthdate DATE NOT NULL,
    biography TEXT,
    city VARCHAR(255),
    password VARCHAR(255) NOT NULL
);

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    text TEXT NOT NULL,
    author_user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE dialogs (
    id UUID PRIMARY KEY,
    name VARCHAR(255),  -- Название диалога (опционально для групповых диалогов)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE dialog_participants (
    dialog_id UUID REFERENCES dialogs(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (dialog_id, user_id)
);

CREATE TABLE messages (
    id UUID PRIMARY KEY,
    dialog_id UUID REFERENCES dialogs(id) ON DELETE CASCADE,
    from_user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE friends (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    friend_id UUID REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, friend_id)
);

CREATE INDEX idx_users_city ON users(city);
CREATE INDEX idx_posts_author_user_id ON posts(author_user_id);
CREATE INDEX idx_messages_dialog_id ON messages(dialog_id);
CREATE INDEX idx_dialog_participants_dialog_id ON dialog_participants(dialog_id);
CREATE INDEX idx_dialog_participants_user_id ON dialog_participants(user_id);
CREATE INDEX idx_friends_user_id ON friends(user_id);
CREATE INDEX idx_friends_friend_id ON friends(friend_id);