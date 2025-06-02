<?php
require_once 'db.php';
header('Content-Type: application/json');

$book_id = isset($_GET['book_id']) ? intval($_GET['book_id']) : 0;

if ($book_id > 0) {
    $stmt = $pdo->prepare('SELECT q.*, u.nickname, u.id as user_id FROM quote q LEFT JOIN users u ON q.user_id = u.id WHERE q.book_id = ? ORDER BY q.id DESC');
    $stmt->execute([$book_id]);
} else {
    $stmt = $pdo->query('SELECT q.*, u.nickname, u.id as user_id FROM quote q LEFT JOIN users u ON q.user_id = u.id ORDER BY q.id DESC');
}

$quotes = $stmt->fetchAll();
echo json_encode(['quotes' => $quotes]);
