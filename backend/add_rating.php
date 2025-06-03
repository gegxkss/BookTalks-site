<?php
session_start();
header('Content-Type: application/json; charset=utf-8');
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

// Убедимся, что заголовок JSON установлен корректно
header('Content-Type: application/json; charset=utf-8');

// Добавим проверку на ошибки в запросе
if (!isset($_POST['book_id']) || !isset($_POST['amount'])) {
    echo json_encode(['error' => 'Некорректные параметры запроса']);
    exit;
}

// Логирование ошибок для отладки
try {
    error_log('Параметры: book_id=' . $book_id . ', amount=' . $amount . ', user_id=' . $user_id);
    $stmt = $pdo->prepare('SELECT id FROM rating WHERE book_id = ? AND user_id = ?');
    $stmt->execute([$book_id, $user_id]);
    if ($stmt->fetch()) {
        // Обновить существующую оценку
        $stmt = $pdo->prepare('UPDATE rating SET amount = ? WHERE book_id = ? AND user_id = ?');
        $stmt->execute([$amount, $book_id, $user_id]);
        error_log('Обновлена оценка: ' . $amount);
    } else {
        // Вставить новую оценку
        $stmt = $pdo->prepare('INSERT INTO rating (book_id, user_id, amount) VALUES (?, ?, ?)');
        $stmt->execute([$book_id, $user_id, $amount]);
        error_log('Добавлена новая оценка: ' . $amount);
    }
    echo json_encode(['success' => true]);
} catch (Exception $e) {
    error_log('Ошибка базы данных: ' . $e->getMessage());
    echo json_encode(['error' => 'Ошибка базы данных']);
    exit;
}
