<?php
// backend/books_list.php
require_once 'db.php';
header('Content-Type: application/json');

$sql = 'SELECT b.id, b.name, a.first_name, a.last_name, g.name AS genre, b.coverimage_filename
        FROM book b
        LEFT JOIN author a ON b.author_id = a.id
        LEFT JOIN genre g ON b.genre_id = g.id
        ORDER BY b.name';
$stmt = $pdo->query($sql);
$books = $stmt->fetchAll();
echo json_encode(['books' => $books]);
