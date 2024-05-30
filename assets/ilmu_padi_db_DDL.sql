CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    UserID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) UNIQUE NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Role VARCHAR(50) NOT NULL,
    SubscriptionStatus VARCHAR(50),
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE courses (
    CourseID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Title VARCHAR(255) NOT NULL,
    Description TEXT,
    ContentURL VARCHAR(255),
    AuthorID UUID REFERENCES users(UserID),
    IsFree BOOLEAN DEFAULT FALSE,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE subscriptions (
    SubscriptionID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    UserID UUID REFERENCES users(UserID),
    StartDate TIMESTAMP NOT NULL,
    EndDate TIMESTAMP NOT NULL,
    IsActive BOOLEAN DEFAULT TRUE,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ads (
    AdID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Content TEXT NOT NULL,
    CourseID UUID REFERENCES courses(CourseID),
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE password_resets (
    ResetID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    UserID UUID REFERENCES users(UserID),
    Token VARCHAR(255) NOT NULL,
    Expiration TIMESTAMP NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);