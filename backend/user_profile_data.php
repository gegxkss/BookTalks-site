<?php
// backend/user_profile_data.php

require_once 'db.php';
session_start();

if (!isset($_SESSION['user_id'])) {
    http_response_code(401);
    echo json_encode(['error' => 'Требуется авторизация']);
    exit;
}

$user_id = $_SESSION['user_id'];

// Получение основных данных пользователя
$stmt = $pdo->prepare('SELECT id, nickname, email, first_name, last_name, sex, birth_date, profile_image FROM user WHERE id = ?');
$stmt->execute([$user_id]);
$user = $stmt->fetch();

// Получение книг пользователя
$stmt = $pdo->prepare('SELECT b.id, b.name, b.coverimage_filename, r.amount AS rating FROM user_book ub
                      JOIN book b ON ub.book_id = b.id
                      LEFT JOIN rating r ON b.id = r.book_id AND r.user_id = ?
                      WHERE ub.user_id = ?');
$stmt->execute([$user_id, $user_id]);
$books = $stmt->fetchAll();

// Получение рецензий пользователя
$stmt = $pdo->prepare('SELECT r.id, r.text, b.name AS book_name, b.coverimage_filename FROM review r
                      JOIN book b ON r.book_id = b.id
                      WHERE r.user_id = ?');
$stmt->execute([$user_id]);
$reviews = $stmt->fetchAll();

// Получение цитат пользователя
$stmt = $pdo->prepare('SELECT q.id, q.text, b.name AS book_name, b.coverimage_filename FROM quote q
                      JOIN book b ON q.book_id = b.id
                      WHERE q.user_id = ?');
$stmt->execute([$user_id]);
$quotes = $stmt->fetchAll();

header('Content-Type: application/json');
echo json_encode([
    'user' => $user,
    'books' => $books,
    'reviews' => $reviews,
    'quotes' => $quotes
]);