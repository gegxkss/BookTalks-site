<?php
// backend/user_profile_data.php

require_once 'db.php';
session_start();

// Отключаем вывод ошибок в браузер, только лог
ini_set('display_errors', 0);
ini_set('display_startup_errors', 0);
error_reporting(E_ALL);

header('Content-Type: application/json; charset=utf-8');

if (!isset($_SESSION['user_id'])) {
    http_response_code(401);
    echo json_encode(['error' => 'Требуется авторизация'], JSON_UNESCAPED_UNICODE);
    exit;
}

$user_id = $_SESSION['user_id'];

try {
    // Получение основных данных пользователя
    $stmt = $pdo->prepare('SELECT id, nickname, email, first_name, last_name, sex, birth_date, profile_image FROM users WHERE id = ?');
    $stmt->execute([$user_id]);
    $user = $stmt->fetch();

    if (!$user) {
        http_response_code(404);
        echo json_encode(['error' => 'Пользователь не найден'], JSON_UNESCAPED_UNICODE);
        exit;
    }

    $books = [];
    $reviews = [];
    $quotes = [];

    $tables = $pdo->query("SHOW TABLES")->fetchAll(PDO::FETCH_COLUMN);
    if (in_array('books', $tables) && in_array('user_book', $tables)) {
        // Отключаем работу с таблицей rating, чтобы не было ошибок
        $stmt = $pdo->prepare('SELECT b.id, b.name, b.coverimage_filename FROM user_book ub
                              JOIN books b ON ub.book_id = b.id
                              WHERE ub.user_id = ?');
        $stmt->execute([$user_id]);
        $books = $stmt->fetchAll();
    }
    if (in_array('review', $tables)) {
        $stmt = $pdo->prepare('SELECT r.id, r.text, b.name AS book_name, b.coverimage_filename FROM review r
                              JOIN books b ON r.book_id = b.id
                              WHERE r.user_id = ?');
        $stmt->execute([$user_id]);
        $reviews = $stmt->fetchAll();
    }
    if (in_array('quote', $tables)) {
        $stmt = $pdo->prepare('SELECT q.id, q.text, b.name AS book_name, b.coverimage_filename FROM quote q
                              JOIN books b ON q.book_id = b.id
                              WHERE q.user_id = ?');
        $stmt->execute([$user_id]);
        $quotes = $stmt->fetchAll();
    }

    echo json_encode([
        'user' => $user,
        'books' => $books,
        'reviews' => $reviews,
        'quotes' => $quotes
    ], JSON_UNESCAPED_UNICODE);
    exit;

} catch (Throwable $e) {
    error_log('user_profile_data.php error: ' . $e->getMessage());
    http_response_code(500);
    echo json_encode(['error' => 'Внутренняя ошибка сервера'], JSON_UNESCAPED_UNICODE);
    exit;
}