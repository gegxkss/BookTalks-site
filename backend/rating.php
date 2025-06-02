<?php
require_once 'db.php';
header('Content-Type: application/json');

$book_id = intval($_GET['book_id'] ?? 0);
if (!$book_id) {
    http_response_code(400);
    echo json_encode(['error' => 'ID книги не указан']);
    exit;
}
$sql = 'SELECT ROUND(AVG(rating), 1) as rating FROM rating WHERE book_id = ?';
$stmt = $pdo->prepare($sql);
$stmt->execute([$book_id]);
$row = $stmt->fetch();
echo json_encode(['rating' => $row && $row['rating'] ? $row['rating'] : null]);
