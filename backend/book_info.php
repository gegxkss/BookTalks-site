<?php
// backend/book_info.php
require_once 'db.php';
header('Content-Type: application/json');

$id = intval($_GET['id'] ?? 0);
if (!$id) {
    http_response_code(400);
    echo json_encode(['error' => 'ID книги не указан']);
    exit;
}
$sql = 'SELECT b.id, b.name, a.first_name, a.last_name, g.name AS genre, b.coverimage_filename, b.created_at
        FROM book b
        LEFT JOIN author a ON b.author_id = a.id
        LEFT JOIN genre g ON b.genre_id = g.id
        WHERE b.id = ?';
$stmt = $pdo->prepare($sql);
$stmt->execute([$id]);
$book = $stmt->fetch();
if ($book) {
    echo json_encode(['book' => $book]);
} else {
    http_response_code(404);
    echo json_encode(['error' => 'Книга не найдена']);
}
