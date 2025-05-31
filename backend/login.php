<?php
// backend/login.php
require_once 'db.php';
session_start();
header('Content-Type: application/json');

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $login = trim($_POST['email'] ?? ''); // email или nickname
    $password = $_POST['password'] ?? '';

    if (!$login || !$password) {
        http_response_code(400);
        echo json_encode(['error' => 'Почта/ник и пароль обязательны']);
        exit;
    }

    // Поиск по email или nickname
    $stmt = $pdo->prepare('SELECT * FROM user WHERE email = ? OR nickname = ?');
    $stmt->execute([$login, $login]);
    $user = $stmt->fetch();

    if ($user && password_verify($password, $user['password'])) {
        $_SESSION['user_id'] = $user['id'];
        $_SESSION['user'] = [
            'nickname' => $user['nickname'],
            'email' => $user['email']
        ];
        echo json_encode(['success' => true, 'user' => [
            'id' => $user['id'],
            'nickname' => $user['nickname'],
            'email' => $user['email'],
            'first_name' => $user['first_name'],
            'last_name' => $user['last_name']
        ]]);
    } else {
        http_response_code(401);
        echo json_encode(['error' => 'Неверная почта/ник или пароль']);
    }
} else {
    http_response_code(405);
    echo json_encode(['error' => 'Метод не разрешён']);
}
