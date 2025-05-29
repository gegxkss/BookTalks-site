<?php
// backend/genres_list.php
require_once 'db.php';
header('Content-Type: application/json');

$sql = 'SELECT id, name FROM genre ORDER BY name';
$stmt = $pdo->query($sql);
$genres = $stmt->fetchAll();
echo json_encode(['genres' => $genres]);
