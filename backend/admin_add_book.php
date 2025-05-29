<?php
// backend/admin_add_book.php
require_once 'db.php';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $name = trim($_POST['name'] ?? '');
    $genre_id = intval($_POST['genre_id'] ?? 0);
    $author_id = intval($_POST['author_id'] ?? 0);
    $coverimage_filename = null;

    if (!$name || !$genre_id || !$author_id) {
        http_response_code(400);
        echo json_encode(['error' => 'Все поля обязательны']);
        exit;
    }

    // Обработка загрузки обложки
    if (isset($_FILES['cover_image']) && $_FILES['cover_image']['error'] === UPLOAD_ERR_OK) {
        $uploadDir = __DIR__ . '/../frontend/uploads/';
        if (!is_dir($uploadDir)) {
            mkdir($uploadDir, 0777, true);
        }
        $fileName = uniqid('cover_', true) . '.' . pathinfo($_FILES['cover_image']['name'], PATHINFO_EXTENSION);
        $targetPath = $uploadDir . $fileName;
        if (move_uploaded_file($_FILES['cover_image']['tmp_name'], $targetPath)) {
            $coverimage_filename = 'uploads/' . $fileName;
        }
    }

    $stmt = $pdo->prepare('INSERT INTO book (name, genre_id, author_id, coverimage_filename) VALUES (?, ?, ?, ?)');
    try {
        $stmt->execute([$name, $genre_id, $author_id, $coverimage_filename]);
        echo json_encode(['success' => true]);
    } catch (PDOException $e) {
        http_response_code(500);
        echo json_encode(['error' => 'Ошибка при добавлении книги']);
    }
} else {
    http_response_code(405);
    echo json_encode(['error' => 'Метод не разрешён']);
}
