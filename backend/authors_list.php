<?php
// backend/authors_list.php
require_once 'db.php';

header('Content-Type: application/json');

try {
    $stmt = $pdo->query("SELECT id, first_name, last_name FROM author ORDER BY last_name, first_name");
    $authors = $stmt->fetchAll(PDO::FETCH_ASSOC);

    echo json_encode(['success' => true, 'authors' => $authors]);
} catch (PDOException $e) {
    echo json_encode(['success' => false, 'message' => 'Ошибка загрузки авторов: ' . $e->getMessage()]);
}
