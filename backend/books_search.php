<?php
// backend/books_search.php
require_once 'db.php';
header('Content-Type: application/json');

$query = trim($_GET['q'] ?? '');
$sql = 'SELECT b.id, b.name, a.first_name, a.last_name, g.name AS b.coverimage_filename
        FROM book b
        LEFT JOIN author a ON b.author_id = a.id
        WHERE b.name LIKE ? OR a.first_name LIKE ? OR a.last_name LIKE ? OR g.name LIKE ?
        ORDER BY b.name';
$stmt = $pdo->prepare($sql);
$like = "%$query%";
$stmt->execute([$like, $like, $like, $like]);
$books = $stmt->fetchAll();
echo json_encode(['books' => $books]);
