<?php
// backend/login.php
require_once 'db.php';
session_start();

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $email = trim($_POST['email'] ?? '');
    $password = $_POST['password'] ?? '';

    if (!$email || !$password) {
        http_response_code(400);
        echo json_encode(['error' => 'Почта и пароль обязательны']);
        exit;
    }

    $stmt = $pdo->prepare('SELECT * FROM users WHERE email = ?');
    $stmt->execute([$email]);
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
        echo json_encode(['error' => 'Неверная почта или пароль']);
    }
} else {
    http_response_code(405);
    echo json_encode(['error' => 'Метод не разрешён']);
}
