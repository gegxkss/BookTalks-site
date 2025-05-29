<?php
// backend/check_auth.php
session_start();
header('Content-Type: application/json');

if (isset($_SESSION['user_id']) && isset($_SESSION['user'])) {
    echo json_encode([
        'authenticated' => true,
        'user_id' => $_SESSION['user_id'],
        'username' => $_SESSION['user']['nickname'] ?? $_SESSION['user']['email']
    ]);
} else {
    echo json_encode(['authenticated' => false]);
}
