<?php
require_once 'db.php';
header('Content-Type: application/json');

$sql = 'SELECT q.id, q.text, q.book_id, q.user_id, b.name AS book_name, a.first_name AS author_first_name, a.last_name AS author_last_name, u.nickname AS user_nickname
        FROM quote q
        JOIN book b ON q.book_id = b.id
        JOIN author a ON b.author_id = a.id
        JOIN users u ON q.user_id = u.id
        ORDER BY q.id DESC';
$stmt = $pdo->query($sql);
$quotes = $stmt->fetchAll();
echo json_encode(['quotes' => $quotes]);
