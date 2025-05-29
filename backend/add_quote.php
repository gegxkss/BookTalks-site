<?php
// backend/add_quote.php
require_once 'db.php';
session_start();

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    if (!isset($_SESSION['user_id'])) {
        http_response_code(401);
        echo json_encode(['error' => 'Требуется авторизация']);
        exit;
    }
    $book_id = intval($_POST['book_id'] ?? 0);
    $text = trim($_POST['text'] ?? '');
    if (!$book_id || !$text) {
        http_response_code(400);
        echo json_encode(['error' => 'Все поля обязательны']);
        exit;
    }
    $stmt = $pdo->prepare('INSERT INTO quote (book_id, user_id, text) VALUES (?, ?, ?)');
    $stmt->execute([$book_id, $_SESSION['user_id'], $text]);
    echo json_encode(['success' => true]);
} else {
    http_response_code(405);
    echo json_encode(['error' => 'Метод не разрешён']);
}
