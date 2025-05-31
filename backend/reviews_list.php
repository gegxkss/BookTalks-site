<?php
require_once 'db.php';
header('Content-Type: application/json');

$sql = 'SELECT r.id, r.text, r.book_id, r.user_id, r.created_at, b.name AS book_name, a.first_name AS author_first_name, a.last_name AS author_last_name, u.nickname AS user_nickname
        FROM review r
        JOIN book b ON r.book_id = b.id
        JOIN author a ON b.author_id = a.id
        JOIN users u ON r.user_id = u.id
        ORDER BY r.created_at DESC';
$stmt = $pdo->query($sql);
$reviews = $stmt->fetchAll();
echo json_encode(['reviews' => $reviews]);
