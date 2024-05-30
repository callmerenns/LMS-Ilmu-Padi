INSERT INTO users (Name, Email, Password, Role, SubscriptionStatus) VALUES
('Al Tsaqif Nugraha Ahmad', 'altsaqifnugraha19@gmail.com', '2wsx1qaz', 'admin', 'active'),
('Angga Lesmana', 'anggalesmana131@gmail.com', '2wsx1qaz', 'normal-user', 'inactive'),
('Sien Khumaen Damarwendha', 'sienkhumaen@gmail.com', '2wsx1qaz', 'premium-user', 'active');

INSERT INTO courses (Title, Description, ContentURL, AuthorID, IsFree) VALUES
('Introduction to PostgreSQL', 'A beginner course on PostgreSQL', 'http://example.com/course1', 
    (SELECT UserID FROM users WHERE Email = 'altsaqifnugraha19@gmail.com'), TRUE),
('Advanced PostgreSQL', 'An advanced course on PostgreSQL', 'http://example.com/course2', 
    (SELECT UserID FROM users WHERE Email = 'anggalesmana131@gmail.com'), FALSE),
('PostgreSQL Performance Tuning', 'Learn how to tune PostgreSQL for better performance', 'http://example.com/course3', 
    (SELECT UserID FROM users WHERE Email = 'sienkhumaen@gmail.com'), TRUE);

INSERT INTO subscriptions (UserID, StartDate, EndDate, IsActive) VALUES
((SELECT UserID FROM users WHERE Email = 'altsaqifnugraha19@gmail.com'), '2024-01-01', '2025-01-01', TRUE),
((SELECT UserID FROM users WHERE Email = 'anggalesmana131@gmail.com'), '2024-01-01', '2024-06-01', FALSE),
((SELECT UserID FROM users WHERE Email = 'sienkhumaen@gmail.com'), '2024-01-01', '2025-01-01', TRUE);

INSERT INTO ads (Content, CourseID) VALUES
('Check out our new course on PostgreSQL!', 
    (SELECT CourseID FROM courses WHERE Title = 'Introduction to PostgreSQL')),
('Join our advanced PostgreSQL course now!', 
    (SELECT CourseID FROM courses WHERE Title = 'Advanced PostgreSQL')),
('Learn how to tune PostgreSQL with our new course!', 
    (SELECT CourseID FROM courses WHERE Title = 'PostgreSQL Performance Tuning'));

INSERT INTO password_resets (UserID, Token, Expiration) VALUES
((SELECT UserID FROM users WHERE Email = 'altsaqifnugraha19@gmail.com'), 'token123', '2024-06-01 00:00:00'),
((SELECT UserID FROM users WHERE Email = 'anggalesmana131@gmail.com'), 'token456', '2024-06-01 00:00:00'),
((SELECT UserID FROM users WHERE Email = 'sienkhumaen@gmail.com'), 'token789', '2024-06-01 00:00:00');
