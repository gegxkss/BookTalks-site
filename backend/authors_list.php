<?php
// backend/authors_list.php
require_once 'db.php';
header('Content-Type: application/json');

$sql = 'SELECT id, first_name, last_name, surname FROM author ORDER BY last_name, first_name';
$stmt = $pdo->query($sql);
$authors = $stmt->fetchAll();
echo json_encode(['authors' => $authors]);
