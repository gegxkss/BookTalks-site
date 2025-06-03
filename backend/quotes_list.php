<?php
require_once 'db.php';
header('Content-Type: application/json');

$book_id = isset($_GET['book_id']) ? intval($_GET['book_id']) : 0;

if ($book_id > 0) {
    $stmt = $pdo->prepare('SELECT q.*, u.nickname as user_nickname, u.id as user_id, b.name as book_name, b.coverimage_filename, a.first_name as author_first_name, a.last_name as author_last_name FROM quote q LEFT JOIN users u ON q.user_id = u.id LEFT JOIN book b ON q.book_id = b.id LEFT JOIN author a ON b.author_id = a.id WHERE q.book_id = ? ORDER BY q.id DESC');
    $stmt->execute([$book_id]);
} else {
    $stmt = $pdo->query('SELECT q.*, u.nickname as user_nickname, u.id as user_id, b.name as book_name, b.coverimage_filename, a.first_name as author_first_name, a.last_name as author_last_name FROM quote q LEFT JOIN users u ON q.user_id = u.id LEFT JOIN book b ON q.book_id = b.id LEFT JOIN author a ON b.author_id = a.id ORDER BY q.id DESC');
}

// Добавим обработку ошибок и логирование
if (!$stmt) {
    http_response_code(500);
    echo json_encode(['error' => 'Ошибка выполнения SQL-запроса']);
    exit;
}

$quotes = $stmt->fetchAll();

// Проверим, есть ли данные
if (empty($quotes)) {
    echo json_encode(['quotes' => [], 'message' => 'Цитаты не найдены']);
    exit;
}

echo json_encode(['quotes' => $quotes]);
