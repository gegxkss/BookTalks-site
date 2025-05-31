<?php
// backend/register.php

require_once 'db.php';
session_start();
header('Content-Type: application/json');

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $nickname = trim($_POST['nickname'] ?? '');
    $firstName = trim($_POST['first_name'] ?? '');
    $lastName = trim($_POST['last_name'] ?? '');
    $sex = $_POST['sex'] ?? null;
    $birthDate = $_POST['birth_date'] ?? null;
    $email = trim($_POST['email'] ?? '');
    $password = $_POST['password'] ?? '';

    // Валидация
    if (!$nickname || !$email || !$password) {
        http_response_code(400);
        echo json_encode(['error' => 'Никнейм, почта и пароль обязательны']);
        exit;
    }
    if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
        http_response_code(400);
        echo json_encode(['error' => 'Некорректный email']);
        exit;
    }
    if (strlen($password) < 6) {
        http_response_code(400);
        echo json_encode(['error' => 'Пароль должен быть не менее 6 символов']);
        exit;
    }

    // Проверка уникальности
    $stmt = $pdo->prepare('SELECT COUNT(*) FROM user WHERE email = ? OR nickname = ?');
    $stmt->execute([$email, $nickname]);
    if ($stmt->fetchColumn() > 0) {
        http_response_code(409);
        echo json_encode(['error' => 'Пользователь с таким никнеймом или email уже существует']);
        exit;
    }

    $hashedPassword = password_hash($password, PASSWORD_DEFAULT);

    // Обработка загрузки изображения
    $profileImagePath = null;
    if (isset($_FILES['profile_image']) && $_FILES['profile_image']['error'] === UPLOAD_ERR_OK) {
        $uploadDir = __DIR__ . '/../frontend/uploads/';
        if (!is_dir($uploadDir)) {
            mkdir($uploadDir, 0777, true);
        }
        $fileName = uniqid('profile_', true) . '.' . pathinfo($_FILES['profile_image']['name'], PATHINFO_EXTENSION);
        $targetPath = $uploadDir . $fileName;
        if (move_uploaded_file($_FILES['profile_image']['tmp_name'], $targetPath)) {
            $profileImagePath = 'uploads/' . $fileName;
        }
    }

    $stmt = $pdo->prepare('INSERT INTO user (nickname, email, password, first_name, last_name, sex, birth_date, profile_image) VALUES (?, ?, ?, ?, ?, ?, ?, ?)');
    try {
        $stmt->execute([
            $nickname, $email, $hashedPassword, $firstName, $lastName, $sex, $birthDate, $profileImagePath
        ]);
        $user_id = $pdo->lastInsertId();
        // Автоматический вход
        $_SESSION['user_id'] = $user_id;
        $_SESSION['user'] = [
            'nickname' => $nickname,
            'email' => $email
        ];
        echo json_encode(['success' => true, 'profile_image_url' => $profileImagePath]);
    } catch (PDOException $e) {
        http_response_code(500);
        echo json_encode(['error' => 'Ошибка регистрации']);
    }
} else {
    http_response_code(405);
    echo json_encode(['error' => 'Метод не разрешён']);
}
