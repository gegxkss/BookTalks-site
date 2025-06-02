<?php
session_start();
header('Content-Type: application/json');
require_once 'db.php';

if ($_SERVER['REQUEST_METHOD'] !== 'POST') {
    echo json_encode(['error' => 'Метод не разрешён']);
    exit;
}
if (!isset($_SESSION['user_id'])) {
    echo json_encode(['error' => 'Требуется авторизация']);
    exit;
}
$user_id = $_SESSION['user_id'];
$book_id = isset($_POST['book_id']) ? intval($_POST['book_id']) : 0;
$amount = isset($_POST['amount']) ? intval($_POST['amount']) : 0;
if (!$book_id || !$amount) {
    echo json_encode(['error' => 'Все поля обязательны']);
    exit;
}
try {
    $stmt = $pdo->prepare('SELECT id FROM rating WHERE book_id = ? AND user_id = ?');
    $stmt->execute([$book_id, $user_id]);
    if ($stmt->fetch()) {
        // Обновить существующую оценку
        $stmt = $pdo->prepare('UPDATE rating SET rating = ? WHERE book_id = ? AND user_id = ?');
        $stmt->execute([$amount, $book_id, $user_id]);
    } else {
        // Вставить новую оценку
        $stmt = $pdo->prepare('INSERT INTO rating (book_id, user_id, rating) VALUES (?, ?, ?)');
        $stmt->execute([$book_id, $user_id, $amount]);
    }
    echo json_encode(['success' => true]);
} catch (Exception $e) {
    echo json_encode(['error' => 'Ошибка базы данных: ' . $e->getMessage()]);
}
