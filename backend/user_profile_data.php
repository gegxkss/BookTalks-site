<?php
// backend/user_profile_data.php

require_once 'db.php';
session_start();

// Временный вывод ошибок PHP
ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);
error_reporting(E_ALL);

if (!isset($_SESSION['user_id'])) {
    http_response_code(401);
    echo json_encode(['error' => 'Требуется авторизация']);
    exit;
}

$user_id = $_SESSION['user_id'];

// Получение основных данных пользователя
$stmt = $pdo->prepare('SELECT id, nickname, email, first_name, last_name, sex, birth_date, profile_image FROM users WHERE id = ?');
$stmt->execute([$user_id]);
$user = $stmt->fetch();

// Если пользователь не найден, возвращаем ошибку JSON и завершаем выполнение
if (!$user) {
    http_response_code(404);
    header('Content-Type: application/json');
    echo json_encode(['error' => 'Пользователь не найден']);
    exit;
}

// Временный вывод для отладки
error_log(json_encode($user));

// Если нет таблиц books, user_book, review, quote — не выполняем эти запросы
$books = [];
$reviews = [];
$quotes = [];

// Проверяем наличие таблицы books
$tables = $pdo->query("SHOW TABLES")->fetchAll(PDO::FETCH_COLUMN);
if (in_array('books', $tables) && in_array('user_book', $tables)) {
    if (in_array('rating', $tables)) {
        $stmt = $pdo->prepare('SELECT b.id, b.name, b.coverimage_filename, r.amount AS rating FROM user_book ub
                              JOIN books b ON ub.book_id = b.id
                              LEFT JOIN rating r ON b.id = r.book_id AND r.user_id = ?
                              WHERE ub.user_id = ?');
        $stmt->execute([$user_id, $user_id]);
    } else {
        $stmt = $pdo->prepare('SELECT b.id, b.name, b.coverimage_filename FROM user_book ub
                              JOIN books b ON ub.book_id = b.id
                              WHERE ub.user_id = ?');
        $stmt->execute([$user_id]);
    }
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

// Временный вывод для отладки
error_log(json_encode(['user' => $user, 'books' => $books, 'reviews' => $reviews, 'quotes' => $quotes]));

// Временный вывод для отладки
if (json_last_error() !== JSON_ERROR_NONE) {
    error_log('Ошибка JSON: ' . json_last_error_msg());
}

// Временный вывод для отладки
if (!$user) {
    error_log('Пользователь с ID ' . $user_id . ' не найден в таблице users.');
} else {
    error_log('Данные пользователя: ' . json_encode($user));
}

header('Content-Type: application/json');
echo json_encode([
    'user' => $user,
    'books' => $books,
    'reviews' => $reviews,
    'quotes' => $quotes
]);
exit;